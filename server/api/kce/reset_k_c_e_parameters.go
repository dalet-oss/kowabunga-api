// Code generated by go-swagger; DO NOT EDIT.

package kce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewResetKCEParams creates a new ResetKCEParams object
//
// There are no default values defined in the spec.
func NewResetKCEParams() ResetKCEParams {

	return ResetKCEParams{}
}

// ResetKCEParams contains all the bound params for the reset k c e operation
// typically these are obtained from a http.Request
//
// swagger:parameters ResetKCE
type ResetKCEParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the KCE virtual machine to query.
	  Required: true
	  In: path
	*/
	KceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewResetKCEParams() beforehand.
func (o *ResetKCEParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rKceID, rhkKceID, _ := route.Params.GetOK("kceId")
	if err := o.bindKceID(rKceID, rhkKceID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindKceID binds and validates parameter KceID from path.
func (o *ResetKCEParams) bindKceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.KceID = raw

	return nil
}
