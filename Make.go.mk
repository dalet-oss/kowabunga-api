SDK_GO_DIR = $(SDK_DIR)/go
SDK_GO_MODELS = $(SDK_GO_DIR)/models
SDK_GO_SERVER_DIR = $(SDK_GO_DIR)/server
SDK_GO_SERVER_OPERATIONS = api
SDK_GO_SERVER_NAME = kowabunga
SDK_GO_CLIENT_DIR = $(SDK_GO_DIR)/client
SDK_GO_CLIENT_NAME = kowabunga

.PHONY: go-sdk
go-sdk: go-sdk-server go-sdk-client ; @

.PHONY: go-sdk-server
go-sdk-server: get-goswagger ; $(info $(M) [OpenAPIv2] generate Golang SDK Server code…) @
	$Q $(SWAGGER) generate server -q \
	  -f $(OPENAPI_DEFINITION) \
	  -s $(SDK_GO_SERVER_DIR) \
	  -a $(SDK_GO_SERVER_OPERATIONS) \
	  -m $(SDK_GO_MODELS) \
	  --exclude-main \
	  --name=$(SDK_GO_SERVER_NAME)

.PHONY: go-sdk-client
go-sdk-client: get-goswagger ; $(info $(M) [OpenAPIv2] generate Golang SDK client code…) @
	$Q $(SWAGGER) generate client -q \
	  -f $(OPENAPI_DEFINITION) \
	  -c $(SDK_GO_CLIENT_DIR) \
	  -m $(SDK_GO_MODELS) \
	  --name=$(SDK_GO_CLIENT_NAME)

.PHONY: clean-go
clean-go: ; @
