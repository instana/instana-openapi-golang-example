package main

import (
	"context"
	"fmt"
	"os"

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
}
