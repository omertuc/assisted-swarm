// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExitParams creates a new ExitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExitParams() *ExitParams {
	return &ExitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExitParamsWithTimeout creates a new ExitParams object
// with the ability to set a timeout on a request.
func NewExitParamsWithTimeout(timeout time.Duration) *ExitParams {
	return &ExitParams{
		timeout: timeout,
	}
}

// NewExitParamsWithContext creates a new ExitParams object
// with the ability to set a context for a request.
func NewExitParamsWithContext(ctx context.Context) *ExitParams {
	return &ExitParams{
		Context: ctx,
	}
}

// NewExitParamsWithHTTPClient creates a new ExitParams object
// with the ability to set a custom HTTPClient for a request.
func NewExitParamsWithHTTPClient(client *http.Client) *ExitParams {
	return &ExitParams{
		HTTPClient: client,
	}
}

/* ExitParams contains all the parameters to send to the API endpoint
   for the exit operation.

   Typically these are written to a http.Request.
*/
type ExitParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the exit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExitParams) WithDefaults() *ExitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the exit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the exit params
func (o *ExitParams) WithTimeout(timeout time.Duration) *ExitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exit params
func (o *ExitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exit params
func (o *ExitParams) WithContext(ctx context.Context) *ExitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exit params
func (o *ExitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exit params
func (o *ExitParams) WithHTTPClient(client *http.Client) *ExitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exit params
func (o *ExitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
