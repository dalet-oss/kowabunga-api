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

// KfsAPIController binds http requests to an api service and writes the service results to the http response
type KfsAPIController struct {
	service KfsAPIServicer
	errorHandler ErrorHandler
}

// KfsAPIOption for how the controller is set up.
type KfsAPIOption func(*KfsAPIController)

// WithKfsAPIErrorHandler inject ErrorHandler into controller
func WithKfsAPIErrorHandler(h ErrorHandler) KfsAPIOption {
	return func(c *KfsAPIController) {
		c.errorHandler = h
	}
}

// NewKfsAPIController creates a default api controller
func NewKfsAPIController(s KfsAPIServicer, opts ...KfsAPIOption) Router {
	controller := &KfsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the KfsAPIController
func (c *KfsAPIController) Routes() Routes {
	return Routes{
		"CreateProjectZoneKfs": Route{
			strings.ToUpper("Post"),
			"/api/v1/project/{projectId}/zone/{zoneId}/kfs",
			c.CreateProjectZoneKfs,
		},
		"DeleteKFS": Route{
			strings.ToUpper("Delete"),
			"/api/v1/kfs/{kfsId}",
			c.DeleteKFS,
		},
		"GetAllKFSs": Route{
			strings.ToUpper("Get"),
			"/api/v1/kfs",
			c.GetAllKFSs,
		},
		"GetKFS": Route{
			strings.ToUpper("Get"),
			"/api/v1/kfs/{kfsId}",
			c.GetKFS,
		},
		"GetNfsKfs": Route{
			strings.ToUpper("Get"),
			"/api/v1/nfs/{nfsId}/kfs",
			c.GetNfsKfs,
		},
		"GetProjectZoneKfs": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/zone/{zoneId}/kfs",
			c.GetProjectZoneKfs,
		},
		"UpdateKFS": Route{
			strings.ToUpper("Put"),
			"/api/v1/kfs/{kfsId}",
			c.UpdateKFS,
		},
	}
}

// CreateProjectZoneKfs - 
func (c *KfsAPIController) CreateProjectZoneKfs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	kfsParam := Kfs{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&kfsParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertKfsRequired(kfsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertKfsConstraints(kfsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var nfsIdParam string
	if query.Has("nfsId") {
		param := query.Get("nfsId")

		nfsIdParam = param
	} else {
	}
	var notifyParam bool
	if query.Has("notify") {
		param, err := parseBoolParameter(
			query.Get("notify"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		notifyParam = param
	} else {
	}
	result, err := c.service.CreateProjectZoneKfs(r.Context(), projectIdParam, zoneIdParam, kfsParam, nfsIdParam, notifyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteKFS - 
func (c *KfsAPIController) DeleteKFS(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kfsIdParam := params["kfsId"]
	if kfsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"kfsId"}, nil)
		return
	}
	result, err := c.service.DeleteKFS(r.Context(), kfsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAllKFSs - 
func (c *KfsAPIController) GetAllKFSs(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAllKFSs(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetKFS - 
func (c *KfsAPIController) GetKFS(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kfsIdParam := params["kfsId"]
	if kfsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"kfsId"}, nil)
		return
	}
	result, err := c.service.GetKFS(r.Context(), kfsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetNfsKfs - 
func (c *KfsAPIController) GetNfsKfs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nfsIdParam := params["nfsId"]
	if nfsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"nfsId"}, nil)
		return
	}
	result, err := c.service.GetNfsKfs(r.Context(), nfsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetProjectZoneKfs - 
func (c *KfsAPIController) GetProjectZoneKfs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	var nfsIdParam string
	if query.Has("nfsId") {
		param := query.Get("nfsId")

		nfsIdParam = param
	} else {
	}
	var notifyParam bool
	if query.Has("notify") {
		param, err := parseBoolParameter(
			query.Get("notify"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		notifyParam = param
	} else {
	}
	result, err := c.service.GetProjectZoneKfs(r.Context(), projectIdParam, zoneIdParam, nfsIdParam, notifyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateKFS - 
func (c *KfsAPIController) UpdateKFS(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kfsIdParam := params["kfsId"]
	if kfsIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"kfsId"}, nil)
		return
	}
	kfsParam := Kfs{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&kfsParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertKfsRequired(kfsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertKfsConstraints(kfsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateKFS(r.Context(), kfsIdParam, kfsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}