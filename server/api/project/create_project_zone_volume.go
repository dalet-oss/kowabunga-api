// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateProjectZoneVolumeHandlerFunc turns a function with the right signature into a create project zone volume handler
type CreateProjectZoneVolumeHandlerFunc func(CreateProjectZoneVolumeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateProjectZoneVolumeHandlerFunc) Handle(params CreateProjectZoneVolumeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateProjectZoneVolumeHandler interface for that can handle valid create project zone volume params
type CreateProjectZoneVolumeHandler interface {
	Handle(CreateProjectZoneVolumeParams, interface{}) middleware.Responder
}

// NewCreateProjectZoneVolume creates a new http.Handler for the create project zone volume operation
func NewCreateProjectZoneVolume(ctx *middleware.Context, handler CreateProjectZoneVolumeHandler) *CreateProjectZoneVolume {
	return &CreateProjectZoneVolume{Context: ctx, Handler: handler}
}

/*
	CreateProjectZoneVolume swagger:route POST /project/{projectId}/zone/{zoneId}/volume project zone volume createProjectZoneVolume

Creates a new storage volume in specified zone.
*/
type CreateProjectZoneVolume struct {
	Context *middleware.Context
	Handler CreateProjectZoneVolumeHandler
}

func (o *CreateProjectZoneVolume) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateProjectZoneVolumeParams()
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