// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProjectZoneKGWsHandlerFunc turns a function with the right signature into a get project zone k g ws handler
type GetProjectZoneKGWsHandlerFunc func(GetProjectZoneKGWsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectZoneKGWsHandlerFunc) Handle(params GetProjectZoneKGWsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetProjectZoneKGWsHandler interface for that can handle valid get project zone k g ws params
type GetProjectZoneKGWsHandler interface {
	Handle(GetProjectZoneKGWsParams, interface{}) middleware.Responder
}

// NewGetProjectZoneKGWs creates a new http.Handler for the get project zone k g ws operation
func NewGetProjectZoneKGWs(ctx *middleware.Context, handler GetProjectZoneKGWsHandler) *GetProjectZoneKGWs {
	return &GetProjectZoneKGWs{Context: ctx, Handler: handler}
}

/*
	GetProjectZoneKGWs swagger:route GET /project/{projectId}/zone/{zoneId}/kgws project zone kgw getProjectZoneKGWs

Returns the IDs of the KGW existing in the project in the specified zone.
*/
type GetProjectZoneKGWs struct {
	Context *middleware.Context
	Handler GetProjectZoneKGWsHandler
}

func (o *GetProjectZoneKGWs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetProjectZoneKGWsParams()
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