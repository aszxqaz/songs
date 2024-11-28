.PHONY: openapi
openapi:
	@./scripts/openapi-http-chi.sh info internal/info/ports/http http
	@./scripts/openapi-http-types.sh info internal/info/ports/http http
	@./scripts/openapi-http-client.sh info

	@./scripts/openapi-http-client.sh songs
	@./scripts/openapi-http-types.sh songs internal/songs/ports/http/contracts contracts
	@./scripts/openapi-http-chi.sh songs internal/songs/ports/http/contracts contracts