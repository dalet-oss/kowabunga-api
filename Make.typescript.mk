NODE_DIR = node_modules
YARN = $(NODE_DIR)/.bin/yarn
OPENAPI_GENERATOR = $(NODE_DIR)/.bin/openapi-generator-cli

SDK_TYPESCRIPT_DIR = $(SDK_DIR)/typescript
SDK_ANGULAR_DIR = $(SDK_TYPESCRIPT_DIR)/angular
SDK_ANGULAR_CLIENT_NAME = kowabunga
SDK_ANGULAR_VERSION = 14.2.12

.PHONY: get-yarn
get-yarn: ;$(info $(M) [Npm] installing yarn…) @
	$Q test -x $(YARN) || npm install --silent yarn

.PHONY: get-openapi-generator
get-openapi-generator: get-yarn ;$(info $(M) [Yarn] installing openapi-generator-cli…) @
	$Q test -x $(OPENAPI_GENERATOR) || $(YARN) add --silent @openapitools/openapi-generator-cli
	$Q chmod a+x $(OPENAPI_GENERATOR)

.PHONY: typescript-sdk
typescript-sdk: typescript-sdk-client ; @

.PHONY: typescript-sdk-client
typescript-sdk-client: typescript-sdk-client-angular ; @

.PHONY: typescript-sdk-client-angular
typescript-sdk-client-angular: get-openapi-generator ; $(info $(M) [OpenAPIv3] generate TypeScript Angular SDK client code…) @
	$Q $(OPENAPI_GENERATOR) generate \
	  -g typescript-angular \
	  --package-name $(SDK_ANGULAR_CLIENT_NAME) \
	  -i $(OPENAPI_DEFINITION) \
	  -o $(SDK_ANGULAR_DIR) \
	  --additional-properties=ngVersion=$(SDK_ANGULAR_VERSION) \
	  $(OUT)

.PHONY: clean-typescript
clean-typescript: ; @
	$Q rm -rf $(NODE_DIR)
	$Q rm -rf package.json
	$Q rm -rf package-lock.json
	$Q rm -rf yarn.lock
	$Q rm -rf yarn-error.log
