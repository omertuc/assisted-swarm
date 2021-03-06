// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HealthReader is a Reader for the Health structure.
type HealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewHealthNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHealthNoContent creates a HealthNoContent with default headers values
func NewHealthNoContent() *HealthNoContent {
	return &HealthNoContent{}
}

/* HealthNoContent describes a response with status code 204, with default header values.

Success.
*/
type HealthNoContent struct {
}

func (o *HealthNoContent) Error() string {
	return fmt.Sprintf("[GET /health][%d] healthNoContent ", 204)
}

func (o *HealthNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
