YQ = $(BIN_DIR)/yq
YQ_VERSION = v4.34.1

JINJA = jinjanate

SWAGGER_YAML_TO_HTML = $(BIN_DIR)/swagger-yaml-to-html

OPENAPI_DIR = openapi
OPENAPI_TEMPLATIZE = "./templatize.sh"
OPENAPI_TEMPLATES_DIR = $(OPENAPI_DIR)/templates
OPENAPI_DEFINITION = $(OPENAPI_DIR)/openapi.generated.yml
OPENAPI_DOC = $(DOCS_DIR)/index.html

.PHONY: get-yq
get-yq: bin; $(info $(M) [Misc] downloading yq…) @
	$Q test -x $(YQ) || curl -sL https://github.com/mikefarah/yq/releases/download/$(YQ_VERSION)/yq_$(shell uname -s | tr '[:upper:]' '[:lower:]')_$(shell uname -m | sed 's%x86_64%amd64%') --output $(YQ)
	$Q chmod a+x $(YQ)

.PHONY: get-swagger-yaml-to-html
get-swagger-yaml-to-html: bin; $(info $(M) [Misc] downloading swagger-yaml-to-html…) @
	$Q test -x $(SWAGGER_YAML_TO_HTML) || curl -sL https://raw.githubusercontent.com/yousan/swagger-yaml-to-html/master/swagger-yaml-to-html.py --output $(SWAGGER_YAML_TO_HTML)
	$Q chmod a+x $(SWAGGER_YAML_TO_HTML)

.PHONY: get-jinjanator
get-jinjanator: ; $(info $(M) [Misc] installing jinjanator…) @
	$Q which -s $(JINJA) || pip3 install jinjanator

.PHONY: api
api: specs validate doc ; @

.PHONY: specs
specs: get-jinjanator get-yq ; $(info $(M) [OpenAPIv3] merge fragments…) @
	$Q $(OPENAPI_TEMPLATIZE)
	$Q $(YQ) ea '. as $$item ireduce ({}; . * $$item )' $(OPENAPI_TEMPLATES_DIR)/*.yml > $(OPENAPI_DEFINITION)

.PHONY: validate
validate: get-openapi-generator ; $(info $(M) [OpenAPIv3] valid API syntax…) @
	$Q $(OPENAPI_GENERATOR) validate \
	  --recommend \
	  -i $(OPENAPI_DEFINITION)

.PHONY: doc
doc: get-swagger-yaml-to-html ; $(info $(M) [OpenAPIv3] generate HTML documentation…) @
	$Q mkdir -p $(shell dirname $(OPENAPI_DOC))
	$Q python3 $(SWAGGER_YAML_TO_HTML) < $(OPENAPI_DEFINITION) > $(OPENAPI_DOC)

.PHONY: clean-api
clean-api: ; @
	$Q rm -rf $(OPENAPI_DOC)
	$Q rm -rf $(OPENAPI_DEFINITION)
	$Q rm -rf $(OPENAPI_TEMPLATES_DIR)
	$Q rm -rf $(DOCS_DIR)
