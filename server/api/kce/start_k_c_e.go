// Code generated by go-swagger; DO NOT EDIT.

package kce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// StartKCEHandlerFunc turns a function with the right signature into a start k c e handler
type StartKCEHandlerFunc func(StartKCEParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn StartKCEHandlerFunc) Handle(params StartKCEParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// StartKCEHandler interface for that can handle valid start k c e params
type StartKCEHandler interface {
	Handle(StartKCEParams, interface{}) middleware.Responder
}

// NewStartKCE creates a new http.Handler for the start k c e operation
func NewStartKCE(ctx *middleware.Context, handler StartKCEHandler) *StartKCE {
	return &StartKCE{Context: ctx, Handler: handler}
}

/*
	StartKCE swagger:route POST /kce/{kceId}/start kce startKCE

Boot up a KCE virtual machine.
*/
type StartKCE struct {
	Context *middleware.Context
	Handler StartKCEHandler
}

func (o *StartKCE) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewStartKCEParams()
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
