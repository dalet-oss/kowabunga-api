/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.10.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
	"errors"
)

// KfsAPIService is a service that implements the logic for the KfsAPIServicer
// This service should implement the business logic for every endpoint for the KfsAPI API.
// Include any external packages or services that will be required by this service.
type KfsAPIService struct {
}

// NewKfsAPIService creates a default api service
func NewKfsAPIService() KfsAPIServicer {
	return &KfsAPIService{}
}

// CreateProjectZoneKfs - 
func (s *KfsAPIService) CreateProjectZoneKfs(ctx context.Context, projectId string, zoneId string, kfs Kfs, nfsId string, notify bool) (ImplResponse, error) {
	// TODO - update CreateProjectZoneKfs with the required logic for this service method.
	// Add api_kfs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Kfs{}) or use other options such as http.Ok ...
	// return Response(201, Kfs{}), nil

	// TODO: Uncomment the next line to return response Response(400, ApiErrorBadRequest{}) or use other options such as http.Ok ...
	// return Response(400, ApiErrorBadRequest{}), nil

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

	return Response(http.StatusNotImplemented, nil), errors.New("CreateProjectZoneKfs method not implemented")
}

// DeleteKFS - 
func (s *KfsAPIService) DeleteKFS(ctx context.Context, kfsId string) (ImplResponse, error) {
	// TODO - update DeleteKFS with the required logic for this service method.
	// Add api_kfs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

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

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteKFS method not implemented")
}

// GetAllKFSs - 
func (s *KfsAPIService) GetAllKFSs(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetAllKFSs with the required logic for this service method.
	// Add api_kfs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []string{}) or use other options such as http.Ok ...
	// return Response(200, []string{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAllKFSs method not implemented")
}

// GetKFS - 
func (s *KfsAPIService) GetKFS(ctx context.Context, kfsId string) (ImplResponse, error) {
	// TODO - update GetKFS with the required logic for this service method.
	// Add api_kfs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Kfs{}) or use other options such as http.Ok ...
	// return Response(200, Kfs{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiErrorNotFound{}) or use other options such as http.Ok ...
	// return Response(404, ApiErrorNotFound{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetKFS method not implemented")
}

// GetNfsKfs - 
func (s *KfsAPIService) GetNfsKfs(ctx context.Context, nfsId string) (ImplResponse, error) {
	// TODO - update GetNfsKfs with the required logic for this service method.
	// Add api_kfs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []string{}) or use other options such as http.Ok ...
	// return Response(200, []string{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiErrorNotFound{}) or use other options such as http.Ok ...
	// return Response(404, ApiErrorNotFound{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNfsKfs method not implemented")
}

// GetProjectZoneKfs - 
func (s *KfsAPIService) GetProjectZoneKfs(ctx context.Context, projectId string, zoneId string, nfsId string, notify bool) (ImplResponse, error) {
	// TODO - update GetProjectZoneKfs with the required logic for this service method.
	// Add api_kfs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []string{}) or use other options such as http.Ok ...
	// return Response(200, []string{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiErrorNotFound{}) or use other options such as http.Ok ...
	// return Response(404, ApiErrorNotFound{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetProjectZoneKfs method not implemented")
}

// UpdateKFS - 
func (s *KfsAPIService) UpdateKFS(ctx context.Context, kfsId string, kfs Kfs) (ImplResponse, error) {
	// TODO - update UpdateKFS with the required logic for this service method.
	// Add api_kfs_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Kfs{}) or use other options such as http.Ok ...
	// return Response(200, Kfs{}), nil

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

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateKFS method not implemented")
}