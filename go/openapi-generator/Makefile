.PHONY: gen
gen:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    	-i /local/swagger.yaml \
    	-g go-gin-server \
    	-o /local/gen
