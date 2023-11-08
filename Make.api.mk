YQ = $(BIN_DIR)/yq
YQ_VERSION = v4.34.1

SWAGGER = $(BIN_DIR)/goswagger
SWAGGER_VERSION = v0.30.5
SWAGGER_YAML_TO_HTML = $(BIN_DIR)/swagger-yaml-to-html

OPENAPI_DIR = openapi
OPENAPI_DEFINITION = swagger.generated.yml
OPENAPI_DOC = $(DOCS_DIR)/index.html

.PHONY: get-yq
get-yq: bin; $(info $(M) [Misc] downloading yq…) @
	$Q test -x $(YQ) || curl -sL https://github.com/mikefarah/yq/releases/download/$(YQ_VERSION)/yq_$(shell uname -s | tr '[:upper:]' '[:lower:]')_$(shell uname -m | sed 's%x86_64%amd64%') --output $(YQ)
	$Q chmod a+x $(YQ)

.PHONY: get-goswagger
get-goswagger: bin; $(info $(M) [Misc] downloading go-swagger…) @
	$Q test -x $(SWAGGER) || curl -sL https://github.com/go-swagger/go-swagger/releases/download/$(SWAGGER_VERSION)/swagger_$(shell uname -s | tr '[:upper:]' '[:lower:]')_$(shell uname -m | sed 's%x86_64%amd64%') --output $(SWAGGER)
	$Q chmod a+x $(SWAGGER)

.PHONY: get-swagger-yaml-to-html
get-swagger-yaml-to-html: bin; $(info $(M) [Misc] downloading swagger-yaml-to-html…) @
	$Q test -x $(SWAGGER_YAML_TO_HTML) || curl -sL https://raw.githubusercontent.com/yousan/swagger-yaml-to-html/master/swagger-yaml-to-html.py --output $(SWAGGER_YAML_TO_HTML)
	$Q chmod a+x $(SWAGGER_YAML_TO_HTML)

.PHONY: api
api: specs validate doc ; @

.PHONY: specs
specs: get-yq ; $(info $(M) [OpenAPIv2] merge fragments…) @
	$Q $(YQ) ea '. as $$item ireduce ({}; . * $$item )' $(OPENAPI_DIR)/*.yml > $(OPENAPI_DEFINITION)

.PHONY: validate
validate: get-goswagger ; $(info $(M) [OpenAPIv2] validate API syntax…) @
	$Q $(SWAGGER) validate -q $(OPENAPI_DEFINITION)

.PHONY: doc
doc: get-swagger-yaml-to-html ; $(info $(M) [OpenAPIv2] generate HTML documentation…) @
	$Q mkdir -p $(shell dirname $(OPENAPI_DOC))
	$Q python3 $(SWAGGER_YAML_TO_HTML) < $(OPENAPI_DEFINITION) > $(OPENAPI_DOC)

.PHONY: clean-api
clean-api: ; @
	$Q rm -rf $(OPENAPI_DOC)
	$Q rm -rf $(OPENAPI_DEFINITION)
	$Q rm -rf $(DOCS_DIR)
