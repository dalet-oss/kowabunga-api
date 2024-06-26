/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.40.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
	"errors"
)

// KgwAPIService is a service that implements the logic for the KgwAPIServicer
// This service should implement the business logic for every endpoint for the KgwAPI API.
// Include any external packages or services that will be required by this service.
type KgwAPIService struct {
}

// NewKgwAPIService creates a default api service
func NewKgwAPIService() KgwAPIServicer {
	return &KgwAPIService{}
}

// DeleteKGW - 
func (s *KgwAPIService) DeleteKGW(ctx context.Context, kgwId string) (ImplResponse, error) {
	// TODO - update DeleteKGW with the required logic for this service method.
	// Add api_kgw_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiErrorNotFound{}) or use other options such as http.Ok ...
	// return Response(404, ApiErrorNotFound{}), nil

	// TODO: Uncomment the next line to return response Response(409, ApiErrorConflict{}) or use other options such as http.Ok ...
	// return Response(409, ApiErrorConflict{}), nil

	// TODO: Uncomment the next line to return response Response(422, ApiErrorUnprocessableEntity{}) or use other options such as http.Ok ...
	// return Response(422, ApiErrorUnprocessableEntity{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteKGW method not implemented")
}

// ListKGWs - 
func (s *KgwAPIService) ListKGWs(ctx context.Context) (ImplResponse, error) {
	// TODO - update ListKGWs with the required logic for this service method.
	// Add api_kgw_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []string{}) or use other options such as http.Ok ...
	// return Response(200, []string{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListKGWs method not implemented")
}

// ReadKGW - 
func (s *KgwAPIService) ReadKGW(ctx context.Context, kgwId string) (ImplResponse, error) {
	// TODO - update ReadKGW with the required logic for this service method.
	// Add api_kgw_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Kgw{}) or use other options such as http.Ok ...
	// return Response(200, Kgw{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiErrorNotFound{}) or use other options such as http.Ok ...
	// return Response(404, ApiErrorNotFound{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ReadKGW method not implemented")
}

// UpdateKGW - 
func (s *KgwAPIService) UpdateKGW(ctx context.Context, kgwId string, kgw Kgw) (ImplResponse, error) {
	// TODO - update UpdateKGW with the required logic for this service method.
	// Add api_kgw_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Kgw{}) or use other options such as http.Ok ...
	// return Response(200, Kgw{}), nil

	// TODO: Uncomment the next line to return response Response(400, ApiErrorBadRequest{}) or use other options such as http.Ok ...
	// return Response(400, ApiErrorBadRequest{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiErrorNotFound{}) or use other options such as http.Ok ...
	// return Response(404, ApiErrorNotFound{}), nil

	// TODO: Uncomment the next line to return response Response(422, ApiErrorUnprocessableEntity{}) or use other options such as http.Ok ...
	// return Response(422, ApiErrorUnprocessableEntity{}), nil

	// TODO: Uncomment the next line to return response Response(507, ApiErrorInsufficientStorage{}) or use other options such as http.Ok ...
	// return Response(507, ApiErrorInsufficientStorage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateKGW method not implemented")
}
