// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProjectUsageHandlerFunc turns a function with the right signature into a get project usage handler
type GetProjectUsageHandlerFunc func(GetProjectUsageParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectUsageHandlerFunc) Handle(params GetProjectUsageParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetProjectUsageHandler interface for that can handle valid get project usage params
type GetProjectUsageHandler interface {
	Handle(GetProjectUsageParams, interface{}) middleware.Responder
}

// NewGetProjectUsage creates a new http.Handler for the get project usage operation
func NewGetProjectUsage(ctx *middleware.Context, handler GetProjectUsageHandler) *GetProjectUsage {
	return &GetProjectUsage{Context: ctx, Handler: handler}
}

/*
	GetProjectUsage swagger:route GET /project/{projectId}/usage project getProjectUsage

Returns the current resources usage for the project.
*/
type GetProjectUsage struct {
	Context *middleware.Context
	Handler GetProjectUsageHandler
}

func (o *GetProjectUsage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProjectUsageParams()
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