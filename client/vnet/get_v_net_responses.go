// Code generated by go-swagger; DO NOT EDIT.

package vnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/models"
)

// GetVNetReader is a Reader for the GetVNet structure.
type GetVNetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVNetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVNetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetVNetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVNetOK creates a GetVNetOK with default headers values
func NewGetVNetOK() *GetVNetOK {
	return &GetVNetOK{}
}

/*
GetVNetOK describes a response with status code 200, with default header values.

Returns the virtual network object.
*/
type GetVNetOK struct {
	Payload *models.VNet
}

// IsSuccess returns true when this get v net o k response has a 2xx status code
func (o *GetVNetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v net o k response has a 3xx status code
func (o *GetVNetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v net o k response has a 4xx status code
func (o *GetVNetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v net o k response has a 5xx status code
func (o *GetVNetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v net o k response a status code equal to that given
func (o *GetVNetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v net o k response
func (o *GetVNetOK) Code() int {
	return 200
}

func (o *GetVNetOK) Error() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}][%d] getVNetOK  %+v", 200, o.Payload)
}

func (o *GetVNetOK) String() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}][%d] getVNetOK  %+v", 200, o.Payload)
}

func (o *GetVNetOK) GetPayload() *models.VNet {
	return o.Payload
}

func (o *GetVNetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VNet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVNetNotFound creates a GetVNetNotFound with default headers values
func NewGetVNetNotFound() *GetVNetNotFound {
	return &GetVNetNotFound{}
}

/*
GetVNetNotFound describes a response with status code 404, with default header values.

Invalid virtual network ID was provided.
*/
type GetVNetNotFound struct {
}

// IsSuccess returns true when this get v net not found response has a 2xx status code
func (o *GetVNetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v net not found response has a 3xx status code
func (o *GetVNetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v net not found response has a 4xx status code
func (o *GetVNetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v net not found response has a 5xx status code
func (o *GetVNetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v net not found response a status code equal to that given
func (o *GetVNetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v net not found response
func (o *GetVNetNotFound) Code() int {
	return 404
}

func (o *GetVNetNotFound) Error() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}][%d] getVNetNotFound ", 404)
}

func (o *GetVNetNotFound) String() string {
	return fmt.Sprintf("[GET /vnet/{vnetId}][%d] getVNetNotFound ", 404)
}

func (o *GetVNetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
