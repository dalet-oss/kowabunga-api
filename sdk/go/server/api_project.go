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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ProjectAPIController binds http requests to an api service and writes the service results to the http response
type ProjectAPIController struct {
	service ProjectAPIServicer
	errorHandler ErrorHandler
}

// ProjectAPIOption for how the controller is set up.
type ProjectAPIOption func(*ProjectAPIController)

// WithProjectAPIErrorHandler inject ErrorHandler into controller
func WithProjectAPIErrorHandler(h ErrorHandler) ProjectAPIOption {
	return func(c *ProjectAPIController) {
		c.errorHandler = h
	}
}

// NewProjectAPIController creates a default api controller
func NewProjectAPIController(s ProjectAPIServicer, opts ...ProjectAPIOption) Router {
	controller := &ProjectAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ProjectAPIController
func (c *ProjectAPIController) Routes() Routes {
	return Routes{
		"CreateProject": Route{
			strings.ToUpper("Post"),
			"/api/v1/project",
			c.CreateProject,
		},
		"CreateProjectDnsRecord": Route{
			strings.ToUpper("Post"),
			"/api/v1/project/{projectId}/record",
			c.CreateProjectDnsRecord,
		},
		"CreateProjectRegionKFS": Route{
			strings.ToUpper("Post"),
			"/api/v1/project/{projectId}/region/{regionId}/kfs",
			c.CreateProjectRegionKFS,
		},
		"CreateProjectRegionKGW": Route{
			strings.ToUpper("Post"),
			"/api/v1/project/{projectId}/region/{regionId}/kgw",
			c.CreateProjectRegionKGW,
		},
		"CreateProjectRegionVolume": Route{
			strings.ToUpper("Post"),
			"/api/v1/project/{projectId}/region/{regionId}/volume",
			c.CreateProjectRegionVolume,
		},
		"CreateProjectZoneInstance": Route{
			strings.ToUpper("Post"),
			"/api/v1/project/{projectId}/zone/{zoneId}/instance",
			c.CreateProjectZoneInstance,
		},
		"CreateProjectZoneKCE": Route{
			strings.ToUpper("Post"),
			"/api/v1/project/{projectId}/zone/{zoneId}/kce",
			c.CreateProjectZoneKCE,
		},
		"DeleteProject": Route{
			strings.ToUpper("Delete"),
			"/api/v1/project/{projectId}",
			c.DeleteProject,
		},
		"ListProjectDnsRecords": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/records",
			c.ListProjectDnsRecords,
		},
		"ListProjectRegionKFSs": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/region/{regionId}/kfs",
			c.ListProjectRegionKFSs,
		},
		"ListProjectRegionKGWs": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/region/{regionId}/kgws",
			c.ListProjectRegionKGWs,
		},
		"ListProjectRegionVolumes": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/region/{regionId}/volumes",
			c.ListProjectRegionVolumes,
		},
		"ListProjectZoneInstances": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/zone/{zoneId}/instances",
			c.ListProjectZoneInstances,
		},
		"ListProjectZoneKCEs": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/zone/{zoneId}/kces",
			c.ListProjectZoneKCEs,
		},
		"ListProjects": Route{
			strings.ToUpper("Get"),
			"/api/v1/project",
			c.ListProjects,
		},
		"ReadProject": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}",
			c.ReadProject,
		},
		"ReadProjectCost": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/cost",
			c.ReadProjectCost,
		},
		"ReadProjectUsage": Route{
			strings.ToUpper("Get"),
			"/api/v1/project/{projectId}/usage",
			c.ReadProjectUsage,
		},
		"UpdateProject": Route{
			strings.ToUpper("Put"),
			"/api/v1/project/{projectId}",
			c.UpdateProject,
		},
	}
}

// CreateProject - 
func (c *ProjectAPIController) CreateProject(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	projectParam := Project{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&projectParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertProjectRequired(projectParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertProjectConstraints(projectParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var subnetSizeParam int32
	if query.Has("subnetSize") {
		param, err := parseNumericParameter[int32](
			query.Get("subnetSize"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		subnetSizeParam = param
	} else {
	}
	result, err := c.service.CreateProject(r.Context(), projectParam, subnetSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateProjectDnsRecord - 
func (c *ProjectAPIController) CreateProjectDnsRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	dnsRecordParam := DnsRecord{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&dnsRecordParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDnsRecordRequired(dnsRecordParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDnsRecordConstraints(dnsRecordParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateProjectDnsRecord(r.Context(), projectIdParam, dnsRecordParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateProjectRegionKFS - 
func (c *ProjectAPIController) CreateProjectRegionKFS(w http.ResponseWriter, r *http.Request) {
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
	regionIdParam := params["regionId"]
	if regionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"regionId"}, nil)
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
	result, err := c.service.CreateProjectRegionKFS(r.Context(), projectIdParam, regionIdParam, kfsParam, nfsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateProjectRegionKGW - 
func (c *ProjectAPIController) CreateProjectRegionKGW(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	regionIdParam := params["regionId"]
	if regionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"regionId"}, nil)
		return
	}
	kgwParam := Kgw{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&kgwParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertKgwRequired(kgwParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertKgwConstraints(kgwParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateProjectRegionKGW(r.Context(), projectIdParam, regionIdParam, kgwParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateProjectRegionVolume - 
func (c *ProjectAPIController) CreateProjectRegionVolume(w http.ResponseWriter, r *http.Request) {
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
	regionIdParam := params["regionId"]
	if regionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"regionId"}, nil)
		return
	}
	volumeParam := Volume{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&volumeParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertVolumeRequired(volumeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertVolumeConstraints(volumeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var poolIdParam string
	if query.Has("poolId") {
		param := query.Get("poolId")

		poolIdParam = param
	} else {
	}
	var templateIdParam string
	if query.Has("templateId") {
		param := query.Get("templateId")

		templateIdParam = param
	} else {
	}
	result, err := c.service.CreateProjectRegionVolume(r.Context(), projectIdParam, regionIdParam, volumeParam, poolIdParam, templateIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateProjectZoneInstance - 
func (c *ProjectAPIController) CreateProjectZoneInstance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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
	instanceParam := Instance{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&instanceParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInstanceRequired(instanceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertInstanceConstraints(instanceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateProjectZoneInstance(r.Context(), projectIdParam, zoneIdParam, instanceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateProjectZoneKCE - 
func (c *ProjectAPIController) CreateProjectZoneKCE(w http.ResponseWriter, r *http.Request) {
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
	kceParam := Kce{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&kceParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertKceRequired(kceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertKceConstraints(kceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	var poolIdParam string
	if query.Has("poolId") {
		param := query.Get("poolId")

		poolIdParam = param
	} else {
	}
	var templateIdParam string
	if query.Has("templateId") {
		param := query.Get("templateId")

		templateIdParam = param
	} else {
	}
	var publicParam bool
	if query.Has("public") {
		param, err := parseBoolParameter(
			query.Get("public"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		publicParam = param
	} else {
	}
	result, err := c.service.CreateProjectZoneKCE(r.Context(), projectIdParam, zoneIdParam, kceParam, poolIdParam, templateIdParam, publicParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteProject - 
func (c *ProjectAPIController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	result, err := c.service.DeleteProject(r.Context(), projectIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListProjectDnsRecords - 
func (c *ProjectAPIController) ListProjectDnsRecords(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	result, err := c.service.ListProjectDnsRecords(r.Context(), projectIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListProjectRegionKFSs - 
func (c *ProjectAPIController) ListProjectRegionKFSs(w http.ResponseWriter, r *http.Request) {
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
	regionIdParam := params["regionId"]
	if regionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"regionId"}, nil)
		return
	}
	var nfsIdParam string
	if query.Has("nfsId") {
		param := query.Get("nfsId")

		nfsIdParam = param
	} else {
	}
	result, err := c.service.ListProjectRegionKFSs(r.Context(), projectIdParam, regionIdParam, nfsIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListProjectRegionKGWs - 
func (c *ProjectAPIController) ListProjectRegionKGWs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	regionIdParam := params["regionId"]
	if regionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"regionId"}, nil)
		return
	}
	result, err := c.service.ListProjectRegionKGWs(r.Context(), projectIdParam, regionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListProjectRegionVolumes - 
func (c *ProjectAPIController) ListProjectRegionVolumes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	regionIdParam := params["regionId"]
	if regionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"regionId"}, nil)
		return
	}
	result, err := c.service.ListProjectRegionVolumes(r.Context(), projectIdParam, regionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListProjectZoneInstances - 
func (c *ProjectAPIController) ListProjectZoneInstances(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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
	result, err := c.service.ListProjectZoneInstances(r.Context(), projectIdParam, zoneIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListProjectZoneKCEs - 
func (c *ProjectAPIController) ListProjectZoneKCEs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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
	result, err := c.service.ListProjectZoneKCEs(r.Context(), projectIdParam, zoneIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListProjects - 
func (c *ProjectAPIController) ListProjects(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var subnetSizeParam int32
	if query.Has("subnetSize") {
		param, err := parseNumericParameter[int32](
			query.Get("subnetSize"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		subnetSizeParam = param
	} else {
	}
	result, err := c.service.ListProjects(r.Context(), subnetSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReadProject - 
func (c *ProjectAPIController) ReadProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	result, err := c.service.ReadProject(r.Context(), projectIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReadProjectCost - 
func (c *ProjectAPIController) ReadProjectCost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	result, err := c.service.ReadProjectCost(r.Context(), projectIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ReadProjectUsage - 
func (c *ProjectAPIController) ReadProjectUsage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	result, err := c.service.ReadProjectUsage(r.Context(), projectIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateProject - 
func (c *ProjectAPIController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdParam := params["projectId"]
	if projectIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"projectId"}, nil)
		return
	}
	projectParam := Project{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&projectParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertProjectRequired(projectParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertProjectConstraints(projectParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateProject(r.Context(), projectIdParam, projectParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
