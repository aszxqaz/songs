#!/bin/bash
set -e

readonly service="$1"

oapi-codegen -generate types -o "internal/common/client/$service/openapi_types.gen.go" -package "$service" "internal/openapi/spec/$service.yml"
oapi-codegen -generate client -o "internal/common/client/$service/openapi_client_gen.go" -package "$service" "internal/openapi/spec/$service.yml"
