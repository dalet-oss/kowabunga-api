// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPoolTemplatesReader is a Reader for the GetPoolTemplates structure.
type GetPoolTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPoolTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPoolTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPoolTemplatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPoolTemplatesOK creates a GetPoolTemplatesOK with default headers values
func NewGetPoolTemplatesOK() *GetPoolTemplatesOK {
	return &GetPoolTemplatesOK{}
}

/*
GetPoolTemplatesOK describes a response with status code 200, with default header values.

Returns an array of volume template IDs.
*/
type GetPoolTemplatesOK struct {
	Payload []string
}

// IsSuccess returns true when this get pool templates o k response has a 2xx status code
func (o *GetPoolTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pool templates o k response has a 3xx status code
func (o *GetPoolTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pool templates o k response has a 4xx status code
func (o *GetPoolTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pool templates o k response has a 5xx status code
func (o *GetPoolTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pool templates o k response a status code equal to that given
func (o *GetPoolTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pool templates o k response
func (o *GetPoolTemplatesOK) Code() int {
	return 200
}

func (o *GetPoolTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /pool/{poolId}/templates][%d] getPoolTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetPoolTemplatesOK) String() string {
	return fmt.Sprintf("[GET /pool/{poolId}/templates][%d] getPoolTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetPoolTemplatesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetPoolTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPoolTemplatesNotFound creates a GetPoolTemplatesNotFound with default headers values
func NewGetPoolTemplatesNotFound() *GetPoolTemplatesNotFound {
	return &GetPoolTemplatesNotFound{}
}

/*
GetPoolTemplatesNotFound describes a response with status code 404, with default header values.

Invalid storage pool ID was provided.
*/
type GetPoolTemplatesNotFound struct {
}

// IsSuccess returns true when this get pool templates not found response has a 2xx status code
func (o *GetPoolTemplatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pool templates not found response has a 3xx status code
func (o *GetPoolTemplatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pool templates not found response has a 4xx status code
func (o *GetPoolTemplatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pool templates not found response has a 5xx status code
func (o *GetPoolTemplatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get pool templates not found response a status code equal to that given
func (o *GetPoolTemplatesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get pool templates not found response
func (o *GetPoolTemplatesNotFound) Code() int {
	return 404
}

func (o *GetPoolTemplatesNotFound) Error() string {
	return fmt.Sprintf("[GET /pool/{poolId}/templates][%d] getPoolTemplatesNotFound ", 404)
}

func (o *GetPoolTemplatesNotFound) String() string {
	return fmt.Sprintf("[GET /pool/{poolId}/templates][%d] getPoolTemplatesNotFound ", 404)
}

func (o *GetPoolTemplatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
