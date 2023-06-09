// Code generated by go-swagger; DO NOT EDIT.

package zone

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

	"github.com/dalet-oss/kowabunga-api/models"
)

// NewCreatePoolParams creates a new CreatePoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePoolParams() *CreatePoolParams {
	return &CreatePoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePoolParamsWithTimeout creates a new CreatePoolParams object
// with the ability to set a timeout on a request.
func NewCreatePoolParamsWithTimeout(timeout time.Duration) *CreatePoolParams {
	return &CreatePoolParams{
		timeout: timeout,
	}
}

// NewCreatePoolParamsWithContext creates a new CreatePoolParams object
// with the ability to set a context for a request.
func NewCreatePoolParamsWithContext(ctx context.Context) *CreatePoolParams {
	return &CreatePoolParams{
		Context: ctx,
	}
}

// NewCreatePoolParamsWithHTTPClient creates a new CreatePoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePoolParamsWithHTTPClient(client *http.Client) *CreatePoolParams {
	return &CreatePoolParams{
		HTTPClient: client,
	}
}

/*
CreatePoolParams contains all the parameters to send to the API endpoint

	for the create pool operation.

	Typically these are written to a http.Request.
*/
type CreatePoolParams struct {

	// Body.
	Body *models.StoragePool

	/* HostID.

	   the ID of the associated host (useless for RBD pools, mandatory for local ones).
	*/
	HostID *string

	/* ZoneID.

	   the ID of the associated zone.
	*/
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePoolParams) WithDefaults() *CreatePoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create pool params
func (o *CreatePoolParams) WithTimeout(timeout time.Duration) *CreatePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pool params
func (o *CreatePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pool params
func (o *CreatePoolParams) WithContext(ctx context.Context) *CreatePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pool params
func (o *CreatePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pool params
func (o *CreatePoolParams) WithHTTPClient(client *http.Client) *CreatePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pool params
func (o *CreatePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pool params
func (o *CreatePoolParams) WithBody(body *models.StoragePool) *CreatePoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pool params
func (o *CreatePoolParams) SetBody(body *models.StoragePool) {
	o.Body = body
}

// WithHostID adds the hostID to the create pool params
func (o *CreatePoolParams) WithHostID(hostID *string) *CreatePoolParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the create pool params
func (o *CreatePoolParams) SetHostID(hostID *string) {
	o.HostID = hostID
}

// WithZoneID adds the zoneID to the create pool params
func (o *CreatePoolParams) WithZoneID(zoneID string) *CreatePoolParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the create pool params
func (o *CreatePoolParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.HostID != nil {

		// query param hostId
		var qrHostID string

		if o.HostID != nil {
			qrHostID = *o.HostID
		}
		qHostID := qrHostID
		if qHostID != "" {

			if err := r.SetQueryParam("hostId", qHostID); err != nil {
				return err
			}
		}
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
