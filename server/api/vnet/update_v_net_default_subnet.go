// Code generated by go-swagger; DO NOT EDIT.

package vnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateVNetDefaultSubnetHandlerFunc turns a function with the right signature into a update v net default subnet handler
type UpdateVNetDefaultSubnetHandlerFunc func(UpdateVNetDefaultSubnetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateVNetDefaultSubnetHandlerFunc) Handle(params UpdateVNetDefaultSubnetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateVNetDefaultSubnetHandler interface for that can handle valid update v net default subnet params
type UpdateVNetDefaultSubnetHandler interface {
	Handle(UpdateVNetDefaultSubnetParams, interface{}) middleware.Responder
}

// NewUpdateVNetDefaultSubnet creates a new http.Handler for the update v net default subnet operation
func NewUpdateVNetDefaultSubnet(ctx *middleware.Context, handler UpdateVNetDefaultSubnetHandler) *UpdateVNetDefaultSubnet {
	return &UpdateVNetDefaultSubnet{Context: ctx, Handler: handler}
}

/*
	UpdateVNetDefaultSubnet swagger:route PUT /vnet/{vnetId}/subnet/{subnetId}/default vnet subnet updateVNetDefaultSubnet

Set a virtual network default subnet.
*/
type UpdateVNetDefaultSubnet struct {
	Context *middleware.Context
	Handler UpdateVNetDefaultSubnetHandler
}

func (o *UpdateVNetDefaultSubnet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateVNetDefaultSubnetParams()
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
