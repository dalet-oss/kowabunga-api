// Code generated by go-swagger; DO NOT EDIT.

package kce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllKCEsReader is a Reader for the GetAllKCEs structure.
type GetAllKCEsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllKCEsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllKCEsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllKCEsOK creates a GetAllKCEsOK with default headers values
func NewGetAllKCEsOK() *GetAllKCEsOK {
	return &GetAllKCEsOK{}
}

/*
GetAllKCEsOK describes a response with status code 200, with default header values.

Returns the an array of KCE virtual machines IDs.
*/
type GetAllKCEsOK struct {
	Payload []string
}

// IsSuccess returns true when this get all k c es o k response has a 2xx status code
func (o *GetAllKCEsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all k c es o k response has a 3xx status code
func (o *GetAllKCEsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all k c es o k response has a 4xx status code
func (o *GetAllKCEsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all k c es o k response has a 5xx status code
func (o *GetAllKCEsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all k c es o k response a status code equal to that given
func (o *GetAllKCEsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all k c es o k response
func (o *GetAllKCEsOK) Code() int {
	return 200
}

func (o *GetAllKCEsOK) Error() string {
	return fmt.Sprintf("[GET /kce][%d] getAllKCEsOK  %+v", 200, o.Payload)
}

func (o *GetAllKCEsOK) String() string {
	return fmt.Sprintf("[GET /kce][%d] getAllKCEsOK  %+v", 200, o.Payload)
}

func (o *GetAllKCEsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAllKCEsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
