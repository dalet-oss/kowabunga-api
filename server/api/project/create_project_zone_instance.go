// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateProjectZoneInstanceHandlerFunc turns a function with the right signature into a create project zone instance handler
type CreateProjectZoneInstanceHandlerFunc func(CreateProjectZoneInstanceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateProjectZoneInstanceHandlerFunc) Handle(params CreateProjectZoneInstanceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateProjectZoneInstanceHandler interface for that can handle valid create project zone instance params
type CreateProjectZoneInstanceHandler interface {
	Handle(CreateProjectZoneInstanceParams, interface{}) middleware.Responder
}

// NewCreateProjectZoneInstance creates a new http.Handler for the create project zone instance operation
func NewCreateProjectZoneInstance(ctx *middleware.Context, handler CreateProjectZoneInstanceHandler) *CreateProjectZoneInstance {
	return &CreateProjectZoneInstance{Context: ctx, Handler: handler}
}

/*
	CreateProjectZoneInstance swagger:route POST /project/{projectId}/zone/{zoneId}/instance project zone instance createProjectZoneInstance

Creates a new virtual machine instance in specified zone.
*/
type CreateProjectZoneInstance struct {
	Context *middleware.Context
	Handler CreateProjectZoneInstanceHandler
}

func (o *CreateProjectZoneInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateProjectZoneInstanceParams()
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
