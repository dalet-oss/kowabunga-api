SDK_GO_DIR = $(SDK_DIR)/go

SDK_GO_CLIENT_DIR = $(SDK_GO_DIR)
SDK_GO_CLIENT_GENERATOR = go
SDK_GO_CLIENT_NAME = client

.PHONY: go-sdk
go-sdk: go-sdk-client ; @

.PHONY: go-sdk-client
go-sdk-client: get-openapi-generator ; $(info $(M) [OpenAPIv3] generate Golang SDK client code…) @
	$Q $(OPENAPI_GENERATOR) generate \
	  -g $(SDK_GO_CLIENT_GENERATOR) \
	  --package-name $(SDK_GO_CLIENT_NAME) \
	  --openapi-normalizer KEEP_ONLY_FIRST_TAG_IN_OPERATION=true \
	  -p withGoMod=false \
	  -i $(OPENAPI_DEFINITION) \
	  -o $(SDK_GO_CLIENT_DIR) \
	  $(OUT)
	  $Q rm -f $(SDK_GO_CLIENT_DIR)/.gitignore
	  $Q rm -rf $(SDK_GO_CLIENT_DIR)/.openapi-generator
	  $Q rm -f $(SDK_GO_CLIENT_DIR)/.openapi-generator-ignore
	  $Q rm -f $(SDK_GO_CLIENT_DIR)/.travis.yml
	  $Q rm -rf $(SDK_GO_CLIENT_DIR)/api
	  $Q rm -f $(SDK_GO_CLIENT_DIR)/git_push.sh
	  $Q rm -rf $(SDK_GO_CLIENT_DIR)/docs
	  $Q rm -rf $(SDK_GO_CLIENT_DIR)/README.md
	  $Q rm -rf $(SDK_GO_CLIENT_DIR)/test

.PHONY: clean-go
clean-go: ; @
