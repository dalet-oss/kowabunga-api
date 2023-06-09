// Code generated by go-swagger; DO NOT EDIT.

package pool

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

// NewUpdatePoolDefaultTemplateParams creates a new UpdatePoolDefaultTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePoolDefaultTemplateParams() *UpdatePoolDefaultTemplateParams {
	return &UpdatePoolDefaultTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePoolDefaultTemplateParamsWithTimeout creates a new UpdatePoolDefaultTemplateParams object
// with the ability to set a timeout on a request.
func NewUpdatePoolDefaultTemplateParamsWithTimeout(timeout time.Duration) *UpdatePoolDefaultTemplateParams {
	return &UpdatePoolDefaultTemplateParams{
		timeout: timeout,
	}
}

// NewUpdatePoolDefaultTemplateParamsWithContext creates a new UpdatePoolDefaultTemplateParams object
// with the ability to set a context for a request.
func NewUpdatePoolDefaultTemplateParamsWithContext(ctx context.Context) *UpdatePoolDefaultTemplateParams {
	return &UpdatePoolDefaultTemplateParams{
		Context: ctx,
	}
}

// NewUpdatePoolDefaultTemplateParamsWithHTTPClient creates a new UpdatePoolDefaultTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePoolDefaultTemplateParamsWithHTTPClient(client *http.Client) *UpdatePoolDefaultTemplateParams {
	return &UpdatePoolDefaultTemplateParams{
		HTTPClient: client,
	}
}

/*
UpdatePoolDefaultTemplateParams contains all the parameters to send to the API endpoint

	for the update pool default template operation.

	Typically these are written to a http.Request.
*/
type UpdatePoolDefaultTemplateParams struct {

	/* PoolID.

	   The ID of the storage pool to update.
	*/
	PoolID string

	/* TemplateID.

	   The ID of the volume template to set as default.
	*/
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update pool default template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePoolDefaultTemplateParams) WithDefaults() *UpdatePoolDefaultTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update pool default template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePoolDefaultTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) WithTimeout(timeout time.Duration) *UpdatePoolDefaultTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) WithContext(ctx context.Context) *UpdatePoolDefaultTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) WithHTTPClient(client *http.Client) *UpdatePoolDefaultTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPoolID adds the poolID to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) WithPoolID(poolID string) *UpdatePoolDefaultTemplateParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithTemplateID adds the templateID to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) WithTemplateID(templateID string) *UpdatePoolDefaultTemplateParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the update pool default template params
func (o *UpdatePoolDefaultTemplateParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePoolDefaultTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param poolId
	if err := r.SetPathParam("poolId", o.PoolID); err != nil {
		return err
	}

	// path param templateId
	if err := r.SetPathParam("templateId", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
