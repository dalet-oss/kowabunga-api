// Code generated by go-swagger; DO NOT EDIT.

package kfs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteKFSHandlerFunc turns a function with the right signature into a delete k f s handler
type DeleteKFSHandlerFunc func(DeleteKFSParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteKFSHandlerFunc) Handle(params DeleteKFSParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteKFSHandler interface for that can handle valid delete k f s params
type DeleteKFSHandler interface {
	Handle(DeleteKFSParams, interface{}) middleware.Responder
}

// NewDeleteKFS creates a new http.Handler for the delete k f s operation
func NewDeleteKFS(ctx *middleware.Context, handler DeleteKFSHandler) *DeleteKFS {
	return &DeleteKFS{Context: ctx, Handler: handler}
}

/*
	DeleteKFS swagger:route DELETE /kfs/{kfsId} kfs deleteKFS

Deletes an existing KFS storage volume.
*/
type DeleteKFS struct {
	Context *middleware.Context
	Handler DeleteKFSHandler
}

func (o *DeleteKFS) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteKFSParams()
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