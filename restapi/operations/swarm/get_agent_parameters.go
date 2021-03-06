// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAgentParams creates a new GetAgentParams object
//
// There are no default values defined in the spec.
func NewGetAgentParams() GetAgentParams {

	return GetAgentParams{}
}

// GetAgentParams contains all the bound params for the get agent operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAgent
type GetAgentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	AgentID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAgentParams() beforehand.
func (o *GetAgentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAgentID, rhkAgentID, _ := route.Params.GetOK("agent_id")
	if err := o.bindAgentID(rAgentID, rhkAgentID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAgentID binds and validates parameter AgentID from path.
func (o *GetAgentParams) bindAgentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("agent_id", "path", "int64", raw)
	}
	o.AgentID = value

	return nil
}
