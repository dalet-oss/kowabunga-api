// Code generated by go-swagger; DO NOT EDIT.

package netgw

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/models"
)

// GetNetGWReader is a Reader for the GetNetGW structure.
type GetNetGWReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetGWReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetGWOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNetGWNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /netgw/{netgwId}] GetNetGW", response, response.Code())
	}
}

// NewGetNetGWOK creates a GetNetGWOK with default headers values
func NewGetNetGWOK() *GetNetGWOK {
	return &GetNetGWOK{}
}

/*
GetNetGWOK describes a response with status code 200, with default header values.

Returns the network gateway object.
*/
type GetNetGWOK struct {
	Payload *models.NetGW
}

// IsSuccess returns true when this get net g w o k response has a 2xx status code
func (o *GetNetGWOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get net g w o k response has a 3xx status code
func (o *GetNetGWOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get net g w o k response has a 4xx status code
func (o *GetNetGWOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get net g w o k response has a 5xx status code
func (o *GetNetGWOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get net g w o k response a status code equal to that given
func (o *GetNetGWOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get net g w o k response
func (o *GetNetGWOK) Code() int {
	return 200
}

func (o *GetNetGWOK) Error() string {
	return fmt.Sprintf("[GET /netgw/{netgwId}][%d] getNetGWOK  %+v", 200, o.Payload)
}

func (o *GetNetGWOK) String() string {
	return fmt.Sprintf("[GET /netgw/{netgwId}][%d] getNetGWOK  %+v", 200, o.Payload)
}

func (o *GetNetGWOK) GetPayload() *models.NetGW {
	return o.Payload
}

func (o *GetNetGWOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetGW)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetGWNotFound creates a GetNetGWNotFound with default headers values
func NewGetNetGWNotFound() *GetNetGWNotFound {
	return &GetNetGWNotFound{}
}

/*
GetNetGWNotFound describes a response with status code 404, with default header values.

Invalid network gateway ID was provided.
*/
type GetNetGWNotFound struct {
}

// IsSuccess returns true when this get net g w not found response has a 2xx status code
func (o *GetNetGWNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get net g w not found response has a 3xx status code
func (o *GetNetGWNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get net g w not found response has a 4xx status code
func (o *GetNetGWNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get net g w not found response has a 5xx status code
func (o *GetNetGWNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get net g w not found response a status code equal to that given
func (o *GetNetGWNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get net g w not found response
func (o *GetNetGWNotFound) Code() int {
	return 404
}

func (o *GetNetGWNotFound) Error() string {
	return fmt.Sprintf("[GET /netgw/{netgwId}][%d] getNetGWNotFound ", 404)
}

func (o *GetNetGWNotFound) String() string {
	return fmt.Sprintf("[GET /netgw/{netgwId}][%d] getNetGWNotFound ", 404)
}

func (o *GetNetGWNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}