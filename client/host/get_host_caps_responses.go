// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/models"
)

// GetHostCapsReader is a Reader for the GetHostCaps structure.
type GetHostCapsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostCapsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostCapsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHostCapsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHostCapsOK creates a GetHostCapsOK with default headers values
func NewGetHostCapsOK() *GetHostCapsOK {
	return &GetHostCapsOK{}
}

/*
GetHostCapsOK describes a response with status code 200, with default header values.

Returns the host capabilities.
*/
type GetHostCapsOK struct {
	Payload *models.HostCaps
}

// IsSuccess returns true when this get host caps o k response has a 2xx status code
func (o *GetHostCapsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get host caps o k response has a 3xx status code
func (o *GetHostCapsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host caps o k response has a 4xx status code
func (o *GetHostCapsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host caps o k response has a 5xx status code
func (o *GetHostCapsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get host caps o k response a status code equal to that given
func (o *GetHostCapsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get host caps o k response
func (o *GetHostCapsOK) Code() int {
	return 200
}

func (o *GetHostCapsOK) Error() string {
	return fmt.Sprintf("[GET /host/{hostId}/caps][%d] getHostCapsOK  %+v", 200, o.Payload)
}

func (o *GetHostCapsOK) String() string {
	return fmt.Sprintf("[GET /host/{hostId}/caps][%d] getHostCapsOK  %+v", 200, o.Payload)
}

func (o *GetHostCapsOK) GetPayload() *models.HostCaps {
	return o.Payload
}

func (o *GetHostCapsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostCaps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostCapsNotFound creates a GetHostCapsNotFound with default headers values
func NewGetHostCapsNotFound() *GetHostCapsNotFound {
	return &GetHostCapsNotFound{}
}

/*
GetHostCapsNotFound describes a response with status code 404, with default header values.

Invalid host ID was provided.
*/
type GetHostCapsNotFound struct {
}

// IsSuccess returns true when this get host caps not found response has a 2xx status code
func (o *GetHostCapsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host caps not found response has a 3xx status code
func (o *GetHostCapsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host caps not found response has a 4xx status code
func (o *GetHostCapsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host caps not found response has a 5xx status code
func (o *GetHostCapsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get host caps not found response a status code equal to that given
func (o *GetHostCapsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get host caps not found response
func (o *GetHostCapsNotFound) Code() int {
	return 404
}

func (o *GetHostCapsNotFound) Error() string {
	return fmt.Sprintf("[GET /host/{hostId}/caps][%d] getHostCapsNotFound ", 404)
}

func (o *GetHostCapsNotFound) String() string {
	return fmt.Sprintf("[GET /host/{hostId}/caps][%d] getHostCapsNotFound ", 404)
}

func (o *GetHostCapsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
