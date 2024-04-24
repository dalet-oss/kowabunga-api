YQ = $(BIN_DIR)/yq
YQ_VERSION = v4.34.1

JINJA = jinjanate

OPENAPI_DIR = openapi
OPENAPI_TEMPLATIZE = "./templatize.sh"
OPENAPI_TEMPLATES_DIR = $(OPENAPI_DIR)/templates
OPENAPI_DEFINITION = $(OPENAPI_DIR)/openapi.generated.yml

OPENAPI_HTML_GENERATOR = html2
OPENAPI_DOCS_DIR_HTML = $(DOCS_DIR)/html

OPENAPI_MARKDOWN_GENERATOR = markdown
OPENAPI_DOCS_DIR_MARKDOWN = $(DOCS_DIR)/markdown

OPENAPI_YAML_GENERATOR = openapi-yaml
OPENAPI_DOCS_DIR_YAML = $(DOCS_DIR)/yaml

.PHONY: get-yq
get-yq: bin; $(info $(M) [Misc] downloading yq…) @
	$Q test -x $(YQ) || curl -sL https://github.com/mikefarah/yq/releases/download/$(YQ_VERSION)/yq_$(shell uname -s | tr '[:upper:]' '[:lower:]')_$(shell uname -m | sed 's%x86_64%amd64%') --output $(YQ)
	$Q chmod a+x $(YQ)

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
doc: docs-yaml docs-html docs-markdown ; @

.PHONY: docs-yaml
docs-yaml: ; $(info $(M) [OpenAPIv3] generate YAML schema…) @
	$Q mkdir -p $(OPENAPI_DOCS_DIR_HTML)
	$Q $(OPENAPI_GENERATOR) generate \
	  -g $(OPENAPI_YAML_GENERATOR) \
	  --package-name $(PACKAGE_NAME) \
	  -i $(OPENAPI_DEFINITION) \
	  -o $(OPENAPI_DOCS_DIR_YAML) \
	  $(OUT)

.PHONY: docs-html
docs-html: ; $(info $(M) [OpenAPIv3] generate HTML documentation…) @
	$Q mkdir -p $(OPENAPI_DOCS_DIR_HTML)
	$Q $(OPENAPI_GENERATOR) generate \
	  -g $(OPENAPI_HTML_GENERATOR) \
	  --package-name $(PACKAGE_NAME) \
	  -i $(OPENAPI_DEFINITION) \
	  -o $(OPENAPI_DOCS_DIR_HTML) \
	  $(OUT)

.PHONY: docs-markdown
docs-markdown: ; $(info $(M) [OpenAPIv3] generate Markdown documentation…) @
	$Q mkdir -p $(OPENAPI_DOCS_DIR_HTML)
	$Q $(OPENAPI_GENERATOR) generate \
	  -g $(OPENAPI_MARKDOWN_GENERATOR) \
	  --package-name $(PACKAGE_NAME) \
	  -i $(OPENAPI_DEFINITION) \
	  -o $(OPENAPI_DOCS_DIR_MARKDOWN) \
	  $(OUT)

.PHONY: clean-api
clean-api: ; @
	$Q rm -rf $(OPENAPI_DEFINITION)
	$Q rm -rf $(OPENAPI_TEMPLATES_DIR)
	$Q rm -rf $(DOCS_DIR)
