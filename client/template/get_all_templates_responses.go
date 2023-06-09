// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllTemplatesReader is a Reader for the GetAllTemplates structure.
type GetAllTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllTemplatesOK creates a GetAllTemplatesOK with default headers values
func NewGetAllTemplatesOK() *GetAllTemplatesOK {
	return &GetAllTemplatesOK{}
}

/*
GetAllTemplatesOK describes a response with status code 200, with default header values.

Returns the an array of volume template IDs.
*/
type GetAllTemplatesOK struct {
	Payload []string
}

// IsSuccess returns true when this get all templates o k response has a 2xx status code
func (o *GetAllTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all templates o k response has a 3xx status code
func (o *GetAllTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all templates o k response has a 4xx status code
func (o *GetAllTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all templates o k response has a 5xx status code
func (o *GetAllTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all templates o k response a status code equal to that given
func (o *GetAllTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all templates o k response
func (o *GetAllTemplatesOK) Code() int {
	return 200
}

func (o *GetAllTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /template][%d] getAllTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetAllTemplatesOK) String() string {
	return fmt.Sprintf("[GET /template][%d] getAllTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetAllTemplatesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAllTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
