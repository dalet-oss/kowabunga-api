// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdatePoolDefaultTemplateHandlerFunc turns a function with the right signature into a update pool default template handler
type UpdatePoolDefaultTemplateHandlerFunc func(UpdatePoolDefaultTemplateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePoolDefaultTemplateHandlerFunc) Handle(params UpdatePoolDefaultTemplateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdatePoolDefaultTemplateHandler interface for that can handle valid update pool default template params
type UpdatePoolDefaultTemplateHandler interface {
	Handle(UpdatePoolDefaultTemplateParams, interface{}) middleware.Responder
}

// NewUpdatePoolDefaultTemplate creates a new http.Handler for the update pool default template operation
func NewUpdatePoolDefaultTemplate(ctx *middleware.Context, handler UpdatePoolDefaultTemplateHandler) *UpdatePoolDefaultTemplate {
	return &UpdatePoolDefaultTemplate{Context: ctx, Handler: handler}
}

/*
	UpdatePoolDefaultTemplate swagger:route PUT /pool/{poolId}/template/{templateId}/default pool template updatePoolDefaultTemplate

Set a storage pool default volume template.
*/
type UpdatePoolDefaultTemplate struct {
	Context *middleware.Context
	Handler UpdatePoolDefaultTemplateHandler
}

func (o *UpdatePoolDefaultTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdatePoolDefaultTemplateParams()
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