// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateVolumeHandlerFunc turns a function with the right signature into a update volume handler
type UpdateVolumeHandlerFunc func(UpdateVolumeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateVolumeHandlerFunc) Handle(params UpdateVolumeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateVolumeHandler interface for that can handle valid update volume params
type UpdateVolumeHandler interface {
	Handle(UpdateVolumeParams, interface{}) middleware.Responder
}

// NewUpdateVolume creates a new http.Handler for the update volume operation
func NewUpdateVolume(ctx *middleware.Context, handler UpdateVolumeHandler) *UpdateVolume {
	return &UpdateVolume{Context: ctx, Handler: handler}
}

/*
	UpdateVolume swagger:route PUT /volume/{volumeId} volume updateVolume

Updates/resizes a storage volume configuration.
*/
type UpdateVolume struct {
	Context *middleware.Context
	Handler UpdateVolumeHandler
}

func (o *UpdateVolume) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateVolumeParams()
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
