// Code generated by go-swagger; DO NOT EDIT.

package vnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetVNetSubnetsHandlerFunc turns a function with the right signature into a get v net subnets handler
type GetVNetSubnetsHandlerFunc func(GetVNetSubnetsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVNetSubnetsHandlerFunc) Handle(params GetVNetSubnetsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetVNetSubnetsHandler interface for that can handle valid get v net subnets params
type GetVNetSubnetsHandler interface {
	Handle(GetVNetSubnetsParams, interface{}) middleware.Responder
}

// NewGetVNetSubnets creates a new http.Handler for the get v net subnets operation
func NewGetVNetSubnets(ctx *middleware.Context, handler GetVNetSubnetsHandler) *GetVNetSubnets {
	return &GetVNetSubnets{Context: ctx, Handler: handler}
}

/*
	GetVNetSubnets swagger:route GET /vnet/{vnetId}/subnets vnet subnet getVNetSubnets

Returns the IDs of the subnets existing in the virtual network.
*/
type GetVNetSubnets struct {
	Context *middleware.Context
	Handler GetVNetSubnetsHandler
}

func (o *GetVNetSubnets) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetVNetSubnetsParams()
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
