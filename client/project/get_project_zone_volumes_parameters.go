// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetProjectZoneVolumesParams creates a new GetProjectZoneVolumesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectZoneVolumesParams() *GetProjectZoneVolumesParams {
	return &GetProjectZoneVolumesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectZoneVolumesParamsWithTimeout creates a new GetProjectZoneVolumesParams object
// with the ability to set a timeout on a request.
func NewGetProjectZoneVolumesParamsWithTimeout(timeout time.Duration) *GetProjectZoneVolumesParams {
	return &GetProjectZoneVolumesParams{
		timeout: timeout,
	}
}

// NewGetProjectZoneVolumesParamsWithContext creates a new GetProjectZoneVolumesParams object
// with the ability to set a context for a request.
func NewGetProjectZoneVolumesParamsWithContext(ctx context.Context) *GetProjectZoneVolumesParams {
	return &GetProjectZoneVolumesParams{
		Context: ctx,
	}
}

// NewGetProjectZoneVolumesParamsWithHTTPClient creates a new GetProjectZoneVolumesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectZoneVolumesParamsWithHTTPClient(client *http.Client) *GetProjectZoneVolumesParams {
	return &GetProjectZoneVolumesParams{
		HTTPClient: client,
	}
}

/*
GetProjectZoneVolumesParams contains all the parameters to send to the API endpoint

	for the get project zone volumes operation.

	Typically these are written to a http.Request.
*/
type GetProjectZoneVolumesParams struct {

	/* ProjectID.

	   The ID of the project to query.
	*/
	ProjectID string

	/* ZoneID.

	   The ID of the zone to query.
	*/
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project zone volumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectZoneVolumesParams) WithDefaults() *GetProjectZoneVolumesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project zone volumes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectZoneVolumesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) WithTimeout(timeout time.Duration) *GetProjectZoneVolumesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) WithContext(ctx context.Context) *GetProjectZoneVolumesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) WithHTTPClient(client *http.Client) *GetProjectZoneVolumesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) WithProjectID(projectID string) *GetProjectZoneVolumesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithZoneID adds the zoneID to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) WithZoneID(zoneID string) *GetProjectZoneVolumesParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the get project zone volumes params
func (o *GetProjectZoneVolumesParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectZoneVolumesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	// path param zoneId
	if err := r.SetPathParam("zoneId", o.ZoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}