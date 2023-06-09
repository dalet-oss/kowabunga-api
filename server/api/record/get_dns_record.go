// Code generated by go-swagger; DO NOT EDIT.

package record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDNSRecordHandlerFunc turns a function with the right signature into a get Dns record handler
type GetDNSRecordHandlerFunc func(GetDNSRecordParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDNSRecordHandlerFunc) Handle(params GetDNSRecordParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDNSRecordHandler interface for that can handle valid get Dns record params
type GetDNSRecordHandler interface {
	Handle(GetDNSRecordParams, interface{}) middleware.Responder
}

// NewGetDNSRecord creates a new http.Handler for the get Dns record operation
func NewGetDNSRecord(ctx *middleware.Context, handler GetDNSRecordHandler) *GetDNSRecord {
	return &GetDNSRecord{Context: ctx, Handler: handler}
}

/*
	GetDNSRecord swagger:route GET /record/{recordId} record getDnsRecord

Returns a DNS record object
*/
type GetDNSRecord struct {
	Context *middleware.Context
	Handler GetDNSRecordHandler
}

func (o *GetDNSRecord) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDNSRecordParams()
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
