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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// PoolAPIController binds http requests to an api service and writes the service results to the http response
type PoolAPIController struct {
	service PoolAPIServicer
	errorHandler ErrorHandler
}

// PoolAPIOption for how the controller is set up.
type PoolAPIOption func(*PoolAPIController)

// WithPoolAPIErrorHandler inject ErrorHandler into controller
func WithPoolAPIErrorHandler(h ErrorHandler) PoolAPIOption {
	return func(c *PoolAPIController) {
		c.errorHandler = h
	}
}

// NewPoolAPIController creates a default api controller
func NewPoolAPIController(s PoolAPIServicer, opts ...PoolAPIOption) Router {
	controller := &PoolAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PoolAPIController
func (c *PoolAPIController) Routes() Routes {
	return Routes{
		"CreatePool": Route{
			strings.ToUpper("Post"),
			"/api/v1/zone/{zoneId}/pool",
			c.CreatePool,
		},
		"CreateTemplate": Route{
			strings.ToUpper("Post"),
			"/api/v1/pool/{poolId}/template",
			c.CreateTemplate,
		},
		"DeletePool": Route{
			strings.ToUpper("Delete"),
			"/api/v1/pool/{poolId}",
			c.DeletePool,
		},
		"GetAllPools": Route{
			strings.ToUpper("Get"),
			"/api/v1/pool",
			c.GetAllPools,
		},
		"GetPool": Route{
			strings.ToUpper("Get"),
			"/api/v1/pool/{poolId}",
			c.GetPool,
		},
		"GetPoolTemplates": Route{
			strings.ToUpper("Get"),
			"/api/v1/pool/{poolId}/templates",
			c.GetPoolTemplates,
		},
		"GetPoolVolumes": Route{
			strings.ToUpper("Get"),
			"/api/v1/pool/{poolId}/volumes",
			c.GetPoolVolumes,
		},
		"GetZonePools": Route{
			strings.ToUpper("Get"),
			"/api/v1/zone/{zoneId}/pools",
			c.GetZonePools,
		},
		"UpdatePool": Route{
			strings.ToUpper("Put"),
			"/api/v1/pool/{poolId}",
			c.UpdatePool,
		},
		"UpdatePoolDefaultTemplate": Route{
			strings.ToUpper("Put"),
			"/api/v1/pool/{poolId}/template/{templateId}/default",
			c.UpdatePoolDefaultTemplate,
		},
		"UpdateZoneDefaultPool": Route{
			strings.ToUpper("Put"),
			"/api/v1/zone/{zoneId}/pool/{poolId}/default",
			c.UpdateZoneDefaultPool,
		},
	}
}

// CreatePool - 
func (c *PoolAPIController) CreatePool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	storagePoolParam := StoragePool{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&storagePoolParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertStoragePoolRequired(storagePoolParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertStoragePoolConstraints(storagePoolParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var hostIdParam string
	if query.Has("hostId") {
		param := query.Get("hostId")

		hostIdParam = param
	} else {
	}
	result, err := c.service.CreatePool(r.Context(), zoneIdParam, storagePoolParam, hostIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateTemplate - 
func (c *PoolAPIController) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	templateParam := Template{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&templateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTemplateRequired(templateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertTemplateConstraints(templateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateTemplate(r.Context(), poolIdParam, templateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeletePool - 
func (c *PoolAPIController) DeletePool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	result, err := c.service.DeletePool(r.Context(), poolIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAllPools - 
func (c *PoolAPIController) GetAllPools(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAllPools(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPool - 
func (c *PoolAPIController) GetPool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	result, err := c.service.GetPool(r.Context(), poolIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPoolTemplates - 
func (c *PoolAPIController) GetPoolTemplates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	result, err := c.service.GetPoolTemplates(r.Context(), poolIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPoolVolumes - 
func (c *PoolAPIController) GetPoolVolumes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	result, err := c.service.GetPoolVolumes(r.Context(), poolIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetZonePools - 
func (c *PoolAPIController) GetZonePools(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	result, err := c.service.GetZonePools(r.Context(), zoneIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdatePool - 
func (c *PoolAPIController) UpdatePool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	storagePoolParam := StoragePool{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&storagePoolParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertStoragePoolRequired(storagePoolParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertStoragePoolConstraints(storagePoolParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdatePool(r.Context(), poolIdParam, storagePoolParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdatePoolDefaultTemplate - 
func (c *PoolAPIController) UpdatePoolDefaultTemplate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	templateIdParam := params["templateId"]
	if templateIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"templateId"}, nil)
		return
	}
	result, err := c.service.UpdatePoolDefaultTemplate(r.Context(), poolIdParam, templateIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateZoneDefaultPool - 
func (c *PoolAPIController) UpdateZoneDefaultPool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	poolIdParam := params["poolId"]
	if poolIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"poolId"}, nil)
		return
	}
	result, err := c.service.UpdateZoneDefaultPool(r.Context(), zoneIdParam, poolIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}