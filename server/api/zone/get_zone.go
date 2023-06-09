// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetZoneHandlerFunc turns a function with the right signature into a get zone handler
type GetZoneHandlerFunc func(GetZoneParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetZoneHandlerFunc) Handle(params GetZoneParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetZoneHandler interface for that can handle valid get zone params
type GetZoneHandler interface {
	Handle(GetZoneParams, interface{}) middleware.Responder
}

// NewGetZone creates a new http.Handler for the get zone operation
func NewGetZone(ctx *middleware.Context, handler GetZoneHandler) *GetZone {
	return &GetZone{Context: ctx, Handler: handler}
}

/*
	GetZone swagger:route GET /zone/{zoneId} zone getZone

Returns a description of the zone
*/
type GetZone struct {
	Context *middleware.Context
	Handler GetZoneHandler
}

func (o *GetZone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetZoneParams()
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
