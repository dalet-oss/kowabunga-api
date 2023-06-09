// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RebootInstanceHandlerFunc turns a function with the right signature into a reboot instance handler
type RebootInstanceHandlerFunc func(RebootInstanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RebootInstanceHandlerFunc) Handle(params RebootInstanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RebootInstanceHandler interface for that can handle valid reboot instance params
type RebootInstanceHandler interface {
	Handle(RebootInstanceParams, interface{}) middleware.Responder
}

// NewRebootInstance creates a new http.Handler for the reboot instance operation
func NewRebootInstance(ctx *middleware.Context, handler RebootInstanceHandler) *RebootInstance {
	return &RebootInstance{Context: ctx, Handler: handler}
}

/*
	RebootInstance swagger:route POST /instance/{instanceId}/reboot instance rebootInstance

Perform a virtual machine software reboot.
*/
type RebootInstance struct {
	Context *middleware.Context
	Handler RebootInstanceHandler
}

func (o *RebootInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRebootInstanceParams()
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
