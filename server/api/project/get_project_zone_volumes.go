// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProjectZoneVolumesHandlerFunc turns a function with the right signature into a get project zone volumes handler
type GetProjectZoneVolumesHandlerFunc func(GetProjectZoneVolumesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectZoneVolumesHandlerFunc) Handle(params GetProjectZoneVolumesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetProjectZoneVolumesHandler interface for that can handle valid get project zone volumes params
type GetProjectZoneVolumesHandler interface {
	Handle(GetProjectZoneVolumesParams, interface{}) middleware.Responder
}

// NewGetProjectZoneVolumes creates a new http.Handler for the get project zone volumes operation
func NewGetProjectZoneVolumes(ctx *middleware.Context, handler GetProjectZoneVolumesHandler) *GetProjectZoneVolumes {
	return &GetProjectZoneVolumes{Context: ctx, Handler: handler}
}

/*
	GetProjectZoneVolumes swagger:route GET /project/{projectId}/zone/{zoneId}/volumes project zone volume getProjectZoneVolumes

Returns the IDs of the storage volumes existing in the project in the specified zone.
*/
type GetProjectZoneVolumes struct {
	Context *middleware.Context
	Handler GetProjectZoneVolumesHandler
}

func (o *GetProjectZoneVolumes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProjectZoneVolumesParams()
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