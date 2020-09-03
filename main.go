package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/instana/instana-openapi-golang-example/pkg/instana"
)

func main() {
	// grab settings from environment
	apiKey := os.Getenv("INSTANA_KEY")
	tenant := os.Getenv("INSTANA_TENANT")
	unit := os.Getenv("INSTANA_UNIT")

	// create golang client specific configuration
	configuration := instana.NewConfiguration()
	host := fmt.Sprintf("%s-%s.instana.io", unit, tenant)
	configuration.Host = host
	configuration.BasePath = fmt.Sprintf("https://%s", host)

	client := instana.NewAPIClient(configuration)

	// Instana uses the `apiToken` prefix in the `Authorization` header
	auth := context.WithValue(context.Background(), instana.ContextAPIKey, instana.APIKey{
		Key:    apiKey,
		Prefix: "apiToken",
	})

	searchFieldResults, _, err := client.InfrastructureCatalogApi.GetInfrastructureCatalogSearchFields(auth)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	fmt.Println("Available search fields supported by Instana:")
	for _, field := range searchFieldResults {
		fmt.Println(fmt.Sprintf("%s (%s) - %s", field.Keyword, field.Context, field.Description))
	}

	// search for the kubernetes pod plugin
	// https://instana.github.io/openapi/#operation/getInfrastructureCatalogPlugins
	plugins, _, err := client.InfrastructureCatalogApi.GetInfrastructureCatalogPlugins(auth)
	if err != nil {
		println(fmt.Errorf("Error reading plugins: %s", err))

		os.Exit(1)
	}

	println("\n\nLet's find ourselves some kubernetes pods!")

	// search for kubernetes plugins:
	for _, plugin := range plugins {
		if strings.Contains(plugin.Plugin, "kubernetes") {
			fmt.Printf("Found Kubernetes plugin %s with ID %s\n", plugin.Label, plugin.Plugin)
		}
	}

	println("\n\nThe plugin id for Kubernets Pods is `kubernetesPod`, let's find out what metrics are collected:")

	kubernetesPodPluginID := "kubernetesPod"
	metrics, _, err := client.InfrastructureCatalogApi.GetInfrastructureCatalogMetrics(auth, kubernetesPodPluginID, &instana.GetInfrastructureCatalogMetricsOpts{})
	if err != nil {
		println(fmt.Errorf("Error reading available metrics: %s", err))

		os.Exit(1)
	}

	for _, metric := range metrics {
		fmt.Printf("Metric for plugin `%s`: %s (ID: %s)\n", kubernetesPodPluginID, metric.Label, metric.MetricId)
	}

	println("\n\nLet's find a Kubernetes Pod that contains an Instana Agent")

	// let's find all the Snapshots that involve Kubernetes pods!
	// https://instana.github.io/openapi/#operation/getSnapshots
	snapshotsWithKubernetesPods, _, err := client.InfrastructureMetricsApi.GetSnapshots(auth, &instana.GetSnapshotsOpts{
		Plugin: optional.NewString(kubernetesPodPluginID),

		// We know that clusters, when monitored with Instana usually have pods with a name of `instana-agent-*`
		Query: optional.NewString("entity.kubernetes.pod.name:instana-agent*"),

		// We can travel through time and query data from entities that are no more!
		Offline: optional.NewBool(true),

		// Instana uses Milliseconds accross the board
		WindowSize: optional.NewInt64(7 * 86400 * 1000),
		To:         optional.NewInt64(time.Now().Unix() * int64(time.Millisecond)),
	})
	if err != nil {
		println(fmt.Errorf("Error reading snapshots: %s", err))

		os.Exit(1)
	}

	for _, snapshot := range snapshotsWithKubernetesPods.Items {
		fmt.Printf("Kubernetes Pod %s on Host %s with snapshot ID %s \n", snapshot.Label, snapshot.Host, snapshot.SnapshotId)
	}

	println("\n\nLet's put everything together: Querying the cpuRequests pod metrics from a specific snapshot")

	metricID := "cpuRequests"
	metricsQuery := instana.GetCombinedMetrics{
		Plugin: kubernetesPodPluginID,
		Metrics: []string{
			metricID,
		},

		SnapshotIds: []string{
			snapshotsWithKubernetesPods.Items[0].SnapshotId,
		},

		TimeFrame: instana.TimeFrame{
			To:         time.Now().Unix() * 1000,
			WindowSize: 300000,
		},

		// 5 Minutes
		Rollup: 60,
	}

	metricsResult, _, err := client.InfrastructureMetricsApi.GetInfrastructureMetrics(auth, &instana.GetInfrastructureMetricsOpts{
		GetCombinedMetrics: optional.NewInterface(metricsQuery),
	})

	if err != nil {
		println(fmt.Errorf("Error reading metrics: %s", err.(instana.GenericOpenAPIError)))

		os.Exit(1)
	}

	for _, metric := range metricsResult.Items {
		for _, bracket := range metric.Metrics[metricID] {
			parsedTime := time.Unix(0, int64(bracket[0])*int64(time.Millisecond))
			fmt.Printf("CPU requests of Kubernetes Pod %s at %s: %f\n", snapshotsWithKubernetesPods.Items[0].Label, parsedTime, bracket[1])
		}
	}
}
