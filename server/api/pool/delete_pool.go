// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeletePoolHandlerFunc turns a function with the right signature into a delete pool handler
type DeletePoolHandlerFunc func(DeletePoolParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePoolHandlerFunc) Handle(params DeletePoolParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePoolHandler interface for that can handle valid delete pool params
type DeletePoolHandler interface {
	Handle(DeletePoolParams, interface{}) middleware.Responder
}

// NewDeletePool creates a new http.Handler for the delete pool operation
func NewDeletePool(ctx *middleware.Context, handler DeletePoolHandler) *DeletePool {
	return &DeletePool{Context: ctx, Handler: handler}
}

/*
	DeletePool swagger:route DELETE /pool/{poolId} pool deletePool

Deletes an existing pool.
*/
type DeletePool struct {
	Context *middleware.Context
	Handler DeletePoolHandler
}

func (o *DeletePool) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePoolParams()
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