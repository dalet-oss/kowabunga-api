// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateProjectZoneKceHandlerFunc turns a function with the right signature into a create project zone kce handler
type CreateProjectZoneKceHandlerFunc func(CreateProjectZoneKceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateProjectZoneKceHandlerFunc) Handle(params CreateProjectZoneKceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateProjectZoneKceHandler interface for that can handle valid create project zone kce params
type CreateProjectZoneKceHandler interface {
	Handle(CreateProjectZoneKceParams, interface{}) middleware.Responder
}

// NewCreateProjectZoneKce creates a new http.Handler for the create project zone kce operation
func NewCreateProjectZoneKce(ctx *middleware.Context, handler CreateProjectZoneKceHandler) *CreateProjectZoneKce {
	return &CreateProjectZoneKce{Context: ctx, Handler: handler}
}

/*
	CreateProjectZoneKce swagger:route POST /project/{projectId}/zone/{zoneId}/kce project zone kce createProjectZoneKce

Creates a new KCE virtual machine in specified zone.
*/
type CreateProjectZoneKce struct {
	Context *middleware.Context
	Handler CreateProjectZoneKceHandler
}

func (o *CreateProjectZoneKce) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateProjectZoneKceParams()
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