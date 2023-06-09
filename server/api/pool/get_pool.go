// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPoolHandlerFunc turns a function with the right signature into a get pool handler
type GetPoolHandlerFunc func(GetPoolParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPoolHandlerFunc) Handle(params GetPoolParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetPoolHandler interface for that can handle valid get pool params
type GetPoolHandler interface {
	Handle(GetPoolParams, interface{}) middleware.Responder
}

// NewGetPool creates a new http.Handler for the get pool operation
func NewGetPool(ctx *middleware.Context, handler GetPoolHandler) *GetPool {
	return &GetPool{Context: ctx, Handler: handler}
}

/*
	GetPool swagger:route GET /pool/{poolId} pool getPool

Returns a description of the pool
*/
type GetPool struct {
	Context *middleware.Context
	Handler GetPoolHandler
}

func (o *GetPool) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPoolParams()
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
