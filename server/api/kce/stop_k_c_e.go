// Code generated by go-swagger; DO NOT EDIT.

package kce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// StopKCEHandlerFunc turns a function with the right signature into a stop k c e handler
type StopKCEHandlerFunc func(StopKCEParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn StopKCEHandlerFunc) Handle(params StopKCEParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// StopKCEHandler interface for that can handle valid stop k c e params
type StopKCEHandler interface {
	Handle(StopKCEParams, interface{}) middleware.Responder
}

// NewStopKCE creates a new http.Handler for the stop k c e operation
func NewStopKCE(ctx *middleware.Context, handler StopKCEHandler) *StopKCE {
	return &StopKCE{Context: ctx, Handler: handler}
}

/*
	StopKCE swagger:route POST /kce/{kceId}/stop kce stopKCE

Initiate a hardware stop of a KCE virtual machine.
*/
type StopKCE struct {
	Context *middleware.Context
	Handler StopKCEHandler
}

func (o *StopKCE) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewStopKCEParams()
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