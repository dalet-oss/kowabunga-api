// Code generated by go-swagger; DO NOT EDIT.

package adapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllAdaptersHandlerFunc turns a function with the right signature into a get all adapters handler
type GetAllAdaptersHandlerFunc func(GetAllAdaptersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllAdaptersHandlerFunc) Handle(params GetAllAdaptersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllAdaptersHandler interface for that can handle valid get all adapters params
type GetAllAdaptersHandler interface {
	Handle(GetAllAdaptersParams, interface{}) middleware.Responder
}

// NewGetAllAdapters creates a new http.Handler for the get all adapters operation
func NewGetAllAdapters(ctx *middleware.Context, handler GetAllAdaptersHandler) *GetAllAdapters {
	return &GetAllAdapters{Context: ctx, Handler: handler}
}

/*
	GetAllAdapters swagger:route GET /adapter adapter getAllAdapters

Returns the IDs of network adapters.
*/
type GetAllAdapters struct {
	Context *middleware.Context
	Handler GetAllAdaptersHandler
}

func (o *GetAllAdapters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllAdaptersParams()
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