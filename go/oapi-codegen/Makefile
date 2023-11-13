.PHONY: setup
setup:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: gen
gen:
	oapi-codegen -generate gin swagger.yaml > handler/handler.gen.go && \
	oapi-codegen -generate types swagger.yaml > model/model.gen.go

.PHONY: gentmp
gentmp:
	oapi-codegen -templates templates/ -generate gin swagger.yaml > handler/handler.gen.go && \
	oapi-codegen -templates templates/ -generate types swagger.yaml > model/model.gen.go
