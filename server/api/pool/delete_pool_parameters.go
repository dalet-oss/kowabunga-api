// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeletePoolParams creates a new DeletePoolParams object
//
// There are no default values defined in the spec.
func NewDeletePoolParams() DeletePoolParams {

	return DeletePoolParams{}
}

// DeletePoolParams contains all the bound params for the delete pool operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeletePool
type DeletePoolParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the pool to get.
	  Required: true
	  In: path
	*/
	PoolID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeletePoolParams() beforehand.
func (o *DeletePoolParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPoolID, rhkPoolID, _ := route.Params.GetOK("poolId")
	if err := o.bindPoolID(rPoolID, rhkPoolID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPoolID binds and validates parameter PoolID from path.
func (o *DeletePoolParams) bindPoolID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.PoolID = raw

	return nil
}
