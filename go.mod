module github.com/instana/instana-openapi-golang-example

require (
	github.com/antihax/optional v1.0.0
	github.com/instana/instana-openapi-golang-example/pkg/instana v1.185.824
)

replace github.com/instana/instana-openapi-golang-example/pkg/instana v1.185.824 => ./pkg/instana

go 1.13
