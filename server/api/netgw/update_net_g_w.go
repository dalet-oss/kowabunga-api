// Code generated by go-swagger; DO NOT EDIT.

package netgw

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateNetGWHandlerFunc turns a function with the right signature into a update net g w handler
type UpdateNetGWHandlerFunc func(UpdateNetGWParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateNetGWHandlerFunc) Handle(params UpdateNetGWParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateNetGWHandler interface for that can handle valid update net g w params
type UpdateNetGWHandler interface {
	Handle(UpdateNetGWParams, interface{}) middleware.Responder
}

// NewUpdateNetGW creates a new http.Handler for the update net g w operation
func NewUpdateNetGW(ctx *middleware.Context, handler UpdateNetGWHandler) *UpdateNetGW {
	return &UpdateNetGW{Context: ctx, Handler: handler}
}

/*
	UpdateNetGW swagger:route PUT /netgw/{netgwId} netgw updateNetGW

Updates a network gateway configuration.
*/
type UpdateNetGW struct {
	Context *middleware.Context
	Handler UpdateNetGWHandler
}

func (o *UpdateNetGW) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateNetGWParams()
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
