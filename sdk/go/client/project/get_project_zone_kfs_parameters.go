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

// NewGetProjectZoneKfsParams creates a new GetProjectZoneKfsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectZoneKfsParams() *GetProjectZoneKfsParams {
	return &GetProjectZoneKfsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectZoneKfsParamsWithTimeout creates a new GetProjectZoneKfsParams object
// with the ability to set a timeout on a request.
func NewGetProjectZoneKfsParamsWithTimeout(timeout time.Duration) *GetProjectZoneKfsParams {
	return &GetProjectZoneKfsParams{
		timeout: timeout,
	}
}

// NewGetProjectZoneKfsParamsWithContext creates a new GetProjectZoneKfsParams object
// with the ability to set a context for a request.
func NewGetProjectZoneKfsParamsWithContext(ctx context.Context) *GetProjectZoneKfsParams {
	return &GetProjectZoneKfsParams{
		Context: ctx,
	}
}

// NewGetProjectZoneKfsParamsWithHTTPClient creates a new GetProjectZoneKfsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectZoneKfsParamsWithHTTPClient(client *http.Client) *GetProjectZoneKfsParams {
	return &GetProjectZoneKfsParams{
		HTTPClient: client,
	}
}

/*
GetProjectZoneKfsParams contains all the parameters to send to the API endpoint

	for the get project zone kfs operation.

	Typically these are written to a http.Request.
*/
type GetProjectZoneKfsParams struct {

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

// WithDefaults hydrates default values in the get project zone kfs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectZoneKfsParams) WithDefaults() *GetProjectZoneKfsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project zone kfs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectZoneKfsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project zone kfs params
func (o *GetProjectZoneKfsParams) WithTimeout(timeout time.Duration) *GetProjectZoneKfsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project zone kfs params
func (o *GetProjectZoneKfsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project zone kfs params
func (o *GetProjectZoneKfsParams) WithContext(ctx context.Context) *GetProjectZoneKfsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project zone kfs params
func (o *GetProjectZoneKfsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project zone kfs params
func (o *GetProjectZoneKfsParams) WithHTTPClient(client *http.Client) *GetProjectZoneKfsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project zone kfs params
func (o *GetProjectZoneKfsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get project zone kfs params
func (o *GetProjectZoneKfsParams) WithProjectID(projectID string) *GetProjectZoneKfsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project zone kfs params
func (o *GetProjectZoneKfsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithZoneID adds the zoneID to the get project zone kfs params
func (o *GetProjectZoneKfsParams) WithZoneID(zoneID string) *GetProjectZoneKfsParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the get project zone kfs params
func (o *GetProjectZoneKfsParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectZoneKfsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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