// Code generated by go-swagger; DO NOT EDIT.

package vnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetVNetSubnetsReader is a Reader for the GetVNetSubnets structure.
type GetVNetSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVNetSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVNetSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetVNetSubnetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVNetSubnetsOK creates a GetVNetSubnetsOK with default headers values
func NewGetVNetSubnetsOK() *GetVNetSubnetsOK {
	return &GetVNetSubnetsOK{}
}

/*
GetVNetSubnetsOK describes a response with status code 200, with default header values.

Returns an array of subnet IDs.
*/
type GetVNetSubnetsOK struct {
	Payload []string
}

// IsSuccess returns true when this get v net subnets o k response has a 2xx status code
func (o *GetVNetSubnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v net subnets o k response has a 3xx status code
func (o *GetVNetSubnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v net subnets o k response has a 4xx status code
func (o *GetVNetSubnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v net subnets o k response has a 5xx status code
func (o *GetVNetSubnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v net subnets o k response a status code equal to that given
func (o *GetVNetSubnetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v net subnets o k response
func (o *GetVNetSubnetsOK) Code() int {
	return 200
}

func (o *GetVNetSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}/subnets][%d] getVNetSubnetsOK  %+v", 200, o.Payload)
}

func (o *GetVNetSubnetsOK) String() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}/subnets][%d] getVNetSubnetsOK  %+v", 200, o.Payload)
}

func (o *GetVNetSubnetsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetVNetSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVNetSubnetsNotFound creates a GetVNetSubnetsNotFound with default headers values
func NewGetVNetSubnetsNotFound() *GetVNetSubnetsNotFound {
	return &GetVNetSubnetsNotFound{}
}

/*
GetVNetSubnetsNotFound describes a response with status code 404, with default header values.

Invalid virtual network ID was provided.
*/
type GetVNetSubnetsNotFound struct {
}

// IsSuccess returns true when this get v net subnets not found response has a 2xx status code
func (o *GetVNetSubnetsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v net subnets not found response has a 3xx status code
func (o *GetVNetSubnetsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v net subnets not found response has a 4xx status code
func (o *GetVNetSubnetsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v net subnets not found response has a 5xx status code
func (o *GetVNetSubnetsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v net subnets not found response a status code equal to that given
func (o *GetVNetSubnetsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v net subnets not found response
func (o *GetVNetSubnetsNotFound) Code() int {
	return 404
}

func (o *GetVNetSubnetsNotFound) Error() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}/subnets][%d] getVNetSubnetsNotFound ", 404)
}

func (o *GetVNetSubnetsNotFound) String() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}/subnets][%d] getVNetSubnetsNotFound ", 404)
}

func (o *GetVNetSubnetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
