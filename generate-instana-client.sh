#!/usr/bin/env bash

if [[ ! -f "resources/openapi.yaml" ]]; then
	curl -0 https://instana.github.io/openapi/openapi.yaml -o resources/openapi.yaml
fi

if [[ ! -f openapi-generator-cli.jar ]]; then
	curl -0 https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.0.0-beta/openapi-generator-cli-5.0.0-beta.jar -o openapi-generator-cli.jar
fi

GO_POST_PROCESS_FILE="gofmt -s -w" java -jar openapi-generator-cli.jar generate -i resources/openapi.yaml -g go \
    -o pkg/instana \
    --additional-properties packageName=instana \
    --additional-properties isGoSubmodule=true \
    --additional-properties packageVersion=1.185.824 \
    --type-mappings=object=interface{} \
    --enable-post-process-file \
    --skip-validate-spec
