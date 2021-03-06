// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift-assisted/assisted-swarm/models"
)

// ListAgentsOKCode is the HTTP code returned for type ListAgentsOK
const ListAgentsOKCode int = 200

/*ListAgentsOK Success

swagger:response listAgentsOK
*/
type ListAgentsOK struct {

	/*
	  In: Body
	*/
	Payload models.AgentList `json:"body,omitempty"`
}

// NewListAgentsOK creates ListAgentsOK with default headers values
func NewListAgentsOK() *ListAgentsOK {

	return &ListAgentsOK{}
}

// WithPayload adds the payload to the list agents o k response
func (o *ListAgentsOK) WithPayload(payload models.AgentList) *ListAgentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list agents o k response
func (o *ListAgentsOK) SetPayload(payload models.AgentList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAgentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.AgentList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListAgentsBadRequestCode is the HTTP code returned for type ListAgentsBadRequest
const ListAgentsBadRequestCode int = 400

/*ListAgentsBadRequest Error.

swagger:response listAgentsBadRequest
*/
type ListAgentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListAgentsBadRequest creates ListAgentsBadRequest with default headers values
func NewListAgentsBadRequest() *ListAgentsBadRequest {

	return &ListAgentsBadRequest{}
}

// WithPayload adds the payload to the list agents bad request response
func (o *ListAgentsBadRequest) WithPayload(payload *models.Error) *ListAgentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list agents bad request response
func (o *ListAgentsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAgentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAgentsUnauthorizedCode is the HTTP code returned for type ListAgentsUnauthorized
const ListAgentsUnauthorizedCode int = 401

/*ListAgentsUnauthorized Unauthorized.

swagger:response listAgentsUnauthorized
*/
type ListAgentsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListAgentsUnauthorized creates ListAgentsUnauthorized with default headers values
func NewListAgentsUnauthorized() *ListAgentsUnauthorized {

	return &ListAgentsUnauthorized{}
}

// WithPayload adds the payload to the list agents unauthorized response
func (o *ListAgentsUnauthorized) WithPayload(payload *models.Error) *ListAgentsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list agents unauthorized response
func (o *ListAgentsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAgentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAgentsForbiddenCode is the HTTP code returned for type ListAgentsForbidden
const ListAgentsForbiddenCode int = 403

/*ListAgentsForbidden Forbidden.

swagger:response listAgentsForbidden
*/
type ListAgentsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListAgentsForbidden creates ListAgentsForbidden with default headers values
func NewListAgentsForbidden() *ListAgentsForbidden {

	return &ListAgentsForbidden{}
}

// WithPayload adds the payload to the list agents forbidden response
func (o *ListAgentsForbidden) WithPayload(payload *models.Error) *ListAgentsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list agents forbidden response
func (o *ListAgentsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAgentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAgentsInternalServerErrorCode is the HTTP code returned for type ListAgentsInternalServerError
const ListAgentsInternalServerErrorCode int = 500

/*ListAgentsInternalServerError Error.

swagger:response listAgentsInternalServerError
*/
type ListAgentsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListAgentsInternalServerError creates ListAgentsInternalServerError with default headers values
func NewListAgentsInternalServerError() *ListAgentsInternalServerError {

	return &ListAgentsInternalServerError{}
}

// WithPayload adds the payload to the list agents internal server error response
func (o *ListAgentsInternalServerError) WithPayload(payload *models.Error) *ListAgentsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list agents internal server error response
func (o *ListAgentsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAgentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
