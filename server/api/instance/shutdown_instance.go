// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ShutdownInstanceHandlerFunc turns a function with the right signature into a shutdown instance handler
type ShutdownInstanceHandlerFunc func(ShutdownInstanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ShutdownInstanceHandlerFunc) Handle(params ShutdownInstanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ShutdownInstanceHandler interface for that can handle valid shutdown instance params
type ShutdownInstanceHandler interface {
	Handle(ShutdownInstanceParams, interface{}) middleware.Responder
}

// NewShutdownInstance creates a new http.Handler for the shutdown instance operation
func NewShutdownInstance(ctx *middleware.Context, handler ShutdownInstanceHandler) *ShutdownInstance {
	return &ShutdownInstance{Context: ctx, Handler: handler}
}

/*
	ShutdownInstance swagger:route POST /instance/{instanceId}/shutdown instance shutdownInstance

Initiate a software shutdown of a virtual machine.
*/
type ShutdownInstance struct {
	Context *middleware.Context
	Handler ShutdownInstanceHandler
}

func (o *ShutdownInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewShutdownInstanceParams()
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