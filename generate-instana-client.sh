#!/usr/bin/env bash

if [[ ! -f "resources/openapi.yaml" ]]; then
	wget https://instana.github.io/openapi/openapi.yaml -O resources/openapi.yaml
fi

if [[ ! -f openapi-generator-cli.jar ]]; then
	wget http://central.maven.org/maven2/org/openapitools/openapi-generator-cli/4.1.0/openapi-generator-cli-4.1.0.jar -O openapi-generator-cli.jar
fi

GO_POST_PROCESS_FILE="gofmt -s -w" java -jar openapi-generator-cli.jar generate -i resources/openapi.yaml -g go \
    -o pkg/instana \
    --additional-properties packageName=instana \
    --additional-properties isGoSubmodule=true \
    --additional-properties packageVersion=162 \
    --type-mappings=object=interface{} \
    --enable-post-process-file \
    --skip-validate-spec
