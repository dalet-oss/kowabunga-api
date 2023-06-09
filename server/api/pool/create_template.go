// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateTemplateHandlerFunc turns a function with the right signature into a create template handler
type CreateTemplateHandlerFunc func(CreateTemplateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTemplateHandlerFunc) Handle(params CreateTemplateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateTemplateHandler interface for that can handle valid create template params
type CreateTemplateHandler interface {
	Handle(CreateTemplateParams, interface{}) middleware.Responder
}

// NewCreateTemplate creates a new http.Handler for the create template operation
func NewCreateTemplate(ctx *middleware.Context, handler CreateTemplateHandler) *CreateTemplate {
	return &CreateTemplate{Context: ctx, Handler: handler}
}

/*
	CreateTemplate swagger:route POST /pool/{poolId}/template pool template createTemplate

Creates a new volume template.
*/
type CreateTemplate struct {
	Context *middleware.Context
	Handler CreateTemplateHandler
}

func (o *CreateTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateTemplateParams()
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
