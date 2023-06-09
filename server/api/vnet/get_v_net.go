// Code generated by go-swagger; DO NOT EDIT.

package vnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetVNetHandlerFunc turns a function with the right signature into a get v net handler
type GetVNetHandlerFunc func(GetVNetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVNetHandlerFunc) Handle(params GetVNetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetVNetHandler interface for that can handle valid get v net params
type GetVNetHandler interface {
	Handle(GetVNetParams, interface{}) middleware.Responder
}

// NewGetVNet creates a new http.Handler for the get v net operation
func NewGetVNet(ctx *middleware.Context, handler GetVNetHandler) *GetVNet {
	return &GetVNet{Context: ctx, Handler: handler}
}

/*
	GetVNet swagger:route GET /vnet/{vnetId} vnet getVNet

Returns a description of the virtual network
*/
type GetVNet struct {
	Context *middleware.Context
	Handler GetVNetHandler
}

func (o *GetVNet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetVNetParams()
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
