module github.com/instana/instana-openapi-golang-example

require (
	github.com/antihax/optional v1.0.0
	github.com/instana/instana-openapi-golang-example/pkg/instana v1.185.824
	github.com/kr/pretty v0.2.1 // indirect
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e
)

replace github.com/instana/instana-openapi-golang-example/pkg/instana v1.185.824 => ./pkg/instana

go 1.13
