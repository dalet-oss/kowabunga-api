// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/sdk/go/models"
)

// GetHostReader is a Reader for the GetHost structure.
type GetHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /host/{hostId}] GetHost", response, response.Code())
	}
}

// NewGetHostOK creates a GetHostOK with default headers values
func NewGetHostOK() *GetHostOK {
	return &GetHostOK{}
}

/*
GetHostOK describes a response with status code 200, with default header values.

Returns the host object.
*/
type GetHostOK struct {
	Payload *models.Host
}

// IsSuccess returns true when this get host o k response has a 2xx status code
func (o *GetHostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get host o k response has a 3xx status code
func (o *GetHostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host o k response has a 4xx status code
func (o *GetHostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host o k response has a 5xx status code
func (o *GetHostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get host o k response a status code equal to that given
func (o *GetHostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get host o k response
func (o *GetHostOK) Code() int {
	return 200
}

func (o *GetHostOK) Error() string {
	return fmt.Sprintf("[GET /host/{hostId}][%d] getHostOK  %+v", 200, o.Payload)
}

func (o *GetHostOK) String() string {
	return fmt.Sprintf("[GET /host/{hostId}][%d] getHostOK  %+v", 200, o.Payload)
}

func (o *GetHostOK) GetPayload() *models.Host {
	return o.Payload
}

func (o *GetHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostNotFound creates a GetHostNotFound with default headers values
func NewGetHostNotFound() *GetHostNotFound {
	return &GetHostNotFound{}
}

/*
GetHostNotFound describes a response with status code 404, with default header values.

Invalid host ID was provided.
*/
type GetHostNotFound struct {
}

// IsSuccess returns true when this get host not found response has a 2xx status code
func (o *GetHostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host not found response has a 3xx status code
func (o *GetHostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host not found response has a 4xx status code
func (o *GetHostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host not found response has a 5xx status code
func (o *GetHostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get host not found response a status code equal to that given
func (o *GetHostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get host not found response
func (o *GetHostNotFound) Code() int {
	return 404
}

func (o *GetHostNotFound) Error() string {
	return fmt.Sprintf("[GET /host/{hostId}][%d] getHostNotFound ", 404)
}

func (o *GetHostNotFound) String() string {
	return fmt.Sprintf("[GET /host/{hostId}][%d] getHostNotFound ", 404)
}

func (o *GetHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}