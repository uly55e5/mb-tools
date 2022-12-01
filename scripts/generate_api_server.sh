#!/bin/bash

JAVA_CMD=java
OPENAPI_JAR=~/opt/openapi-generator-cli.jar

MB_TOOLS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"/..

${JAVA_CMD} -jar ${OPENAPI_JAR}  generate -c ${MB_TOOLS_DIR}/data/openapi/config.yml -o ${MB_TOOLS_DIR}/cmd/openapi-server/