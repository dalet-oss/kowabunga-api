// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllZonesReader is a Reader for the GetAllZones structure.
type GetAllZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllZonesOK creates a GetAllZonesOK with default headers values
func NewGetAllZonesOK() *GetAllZonesOK {
	return &GetAllZonesOK{}
}

/*
GetAllZonesOK describes a response with status code 200, with default header values.

Returns the an array of zone IDs.
*/
type GetAllZonesOK struct {
	Payload []string
}

// IsSuccess returns true when this get all zones o k response has a 2xx status code
func (o *GetAllZonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all zones o k response has a 3xx status code
func (o *GetAllZonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all zones o k response has a 4xx status code
func (o *GetAllZonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all zones o k response has a 5xx status code
func (o *GetAllZonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all zones o k response a status code equal to that given
func (o *GetAllZonesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all zones o k response
func (o *GetAllZonesOK) Code() int {
	return 200
}

func (o *GetAllZonesOK) Error() string {
	return fmt.Sprintf("[GET /zone][%d] getAllZonesOK  %+v", 200, o.Payload)
}

func (o *GetAllZonesOK) String() string {
	return fmt.Sprintf("[GET /zone][%d] getAllZonesOK  %+v", 200, o.Payload)
}

func (o *GetAllZonesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAllZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
