// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateZoneHandlerFunc turns a function with the right signature into a update zone handler
type UpdateZoneHandlerFunc func(UpdateZoneParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateZoneHandlerFunc) Handle(params UpdateZoneParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateZoneHandler interface for that can handle valid update zone params
type UpdateZoneHandler interface {
	Handle(UpdateZoneParams, interface{}) middleware.Responder
}

// NewUpdateZone creates a new http.Handler for the update zone operation
func NewUpdateZone(ctx *middleware.Context, handler UpdateZoneHandler) *UpdateZone {
	return &UpdateZone{Context: ctx, Handler: handler}
}

/*
	UpdateZone swagger:route PUT /zone/{zoneId} zone updateZone

Updates a zone configuration.
*/
type UpdateZone struct {
	Context *middleware.Context
	Handler UpdateZoneHandler
}

func (o *UpdateZone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateZoneParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}