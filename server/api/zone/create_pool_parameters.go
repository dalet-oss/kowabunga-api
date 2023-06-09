// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/dalet-oss/kowabunga-api/models"
)

// NewCreatePoolParams creates a new CreatePoolParams object
//
// There are no default values defined in the spec.
func NewCreatePoolParams() CreatePoolParams {

	return CreatePoolParams{}
}

// CreatePoolParams contains all the bound params for the create pool operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreatePool
type CreatePoolParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.StoragePool
	/*the ID of the associated host (useless for RBD pools, mandatory for local ones).
	  In: query
	*/
	HostID *string
	/*the ID of the associated zone.
	  Required: true
	  In: path
	*/
	ZoneID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreatePoolParams() beforehand.
func (o *CreatePoolParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.StoragePool
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
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
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	qHostID, qhkHostID, _ := qs.GetOK("hostId")
	if err := o.bindHostID(qHostID, qhkHostID, route.Formats); err != nil {
		res = append(res, err)
	}

	rZoneID, rhkZoneID, _ := route.Params.GetOK("zoneId")
	if err := o.bindZoneID(rZoneID, rhkZoneID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHostID binds and validates parameter HostID from query.
func (o *CreatePoolParams) bindHostID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.HostID = &raw

	return nil
}

// bindZoneID binds and validates parameter ZoneID from path.
func (o *CreatePoolParams) bindZoneID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ZoneID = raw

	return nil
}
