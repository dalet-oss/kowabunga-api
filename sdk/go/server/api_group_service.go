/*
 * Kowabunga API documentation
 *
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * API version: 0.20.0
 * Contact: csops@dalet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
	"errors"
)

// GroupAPIService is a service that implements the logic for the GroupAPIServicer
// This service should implement the business logic for every endpoint for the GroupAPI API.
// Include any external packages or services that will be required by this service.
type GroupAPIService struct {
}

// NewGroupAPIService creates a default api service
func NewGroupAPIService() GroupAPIServicer {
	return &GroupAPIService{}
}

// CreateGroup - 
func (s *GroupAPIService) CreateGroup(ctx context.Context, group Group) (ImplResponse, error) {
	// TODO - update CreateGroup with the required logic for this service method.
	// Add api_group_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Group{}) or use other options such as http.Ok ...
	// return Response(201, Group{}), nil

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

	// TODO: Uncomment the next line to return response Response(507, ApiErrorInsufficientStorage{}) or use other options such as http.Ok ...
	// return Response(507, ApiErrorInsufficientStorage{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateGroup method not implemented")
}

// DeleteGroup - 
func (s *GroupAPIService) DeleteGroup(ctx context.Context, groupId string) (ImplResponse, error) {
	// TODO - update DeleteGroup with the required logic for this service method.
	// Add api_group_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

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

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteGroup method not implemented")
}

// ListGroups - 
func (s *GroupAPIService) ListGroups(ctx context.Context) (ImplResponse, error) {
	// TODO - update ListGroups with the required logic for this service method.
	// Add api_group_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []string{}) or use other options such as http.Ok ...
	// return Response(200, []string{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListGroups method not implemented")
}

// ReadGroup - 
func (s *GroupAPIService) ReadGroup(ctx context.Context, groupId string) (ImplResponse, error) {
	// TODO - update ReadGroup with the required logic for this service method.
	// Add api_group_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Group{}) or use other options such as http.Ok ...
	// return Response(200, Group{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiErrorUnauthorized{}) or use other options such as http.Ok ...
	// return Response(401, ApiErrorUnauthorized{}), nil

	// TODO: Uncomment the next line to return response Response(403, ApiErrorForbidden{}) or use other options such as http.Ok ...
	// return Response(403, ApiErrorForbidden{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiErrorNotFound{}) or use other options such as http.Ok ...
	// return Response(404, ApiErrorNotFound{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ReadGroup method not implemented")
}

// UpdateGroup - 
func (s *GroupAPIService) UpdateGroup(ctx context.Context, groupId string, group Group) (ImplResponse, error) {
	// TODO - update UpdateGroup with the required logic for this service method.
	// Add api_group_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Group{}) or use other options such as http.Ok ...
	// return Response(200, Group{}), nil

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

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateGroup method not implemented")
}