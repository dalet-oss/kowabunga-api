// Code generated by go-swagger; DO NOT EDIT.

package nfs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/dalet-oss/kowabunga-api/models"
)

// NewUpdateNfsStorageParams creates a new UpdateNfsStorageParams object
//
// There are no default values defined in the spec.
func NewUpdateNfsStorageParams() UpdateNfsStorageParams {

	return UpdateNfsStorageParams{}
}

// UpdateNfsStorageParams contains all the bound params for the update nfs storage operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateNfsStorage
type UpdateNfsStorageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.StorageNFS
	/*The ID of the NFS storage to get.
	  Required: true
	  In: path
	*/
	NfsID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateNfsStorageParams() beforehand.
func (o *UpdateNfsStorageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.StorageNFS
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}

	rNfsID, rhkNfsID, _ := route.Params.GetOK("nfsId")
	if err := o.bindNfsID(rNfsID, rhkNfsID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNfsID binds and validates parameter NfsID from path.
func (o *UpdateNfsStorageParams) bindNfsID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.NfsID = raw

	return nil
}