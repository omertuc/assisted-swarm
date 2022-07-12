// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ExitHandlerFunc turns a function with the right signature into a exit handler
type ExitHandlerFunc func(ExitParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ExitHandlerFunc) Handle(params ExitParams) middleware.Responder {
	return fn(params)
}

// ExitHandler interface for that can handle valid exit params
type ExitHandler interface {
	Handle(ExitParams) middleware.Responder
}

// NewExit creates a new http.Handler for the exit operation
func NewExit(ctx *middleware.Context, handler ExitHandler) *Exit {
	return &Exit{Context: ctx, Handler: handler}
}

/* Exit swagger:route GET /exit swarm exit

Exit the process.

*/
type Exit struct {
	Context *middleware.Context
	Handler ExitHandler
}

func (o *Exit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewExitParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}