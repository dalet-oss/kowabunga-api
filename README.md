# kowabunga-api
Dalet's Kowabunga (Linux KVM Orchestrator) OpenAPI definition and SDKs

Kowabunga API is compliant with [OpenAPI v2](https://swagger.io/specification/v2/) specifications (also known as Swagger).

Any change to the API must be performed in the **openapi** directory files. Makefile and build scripts will take care of auto-generating Golang code for both Server and Client SDKs, ready-to-consumed by your favorite apps.
