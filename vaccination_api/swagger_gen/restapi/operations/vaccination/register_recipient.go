// Code generated by go-swagger; DO NOT EDIT.

package vaccination

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterRecipientHandlerFunc turns a function with the right signature into a register recipient handler
type RegisterRecipientHandlerFunc func(RegisterRecipientParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterRecipientHandlerFunc) Handle(params RegisterRecipientParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RegisterRecipientHandler interface for that can handle valid register recipient params
type RegisterRecipientHandler interface {
	Handle(RegisterRecipientParams, interface{}) middleware.Responder
}

// NewRegisterRecipient creates a new http.Handler for the register recipient operation
func NewRegisterRecipient(ctx *middleware.Context, handler RegisterRecipientHandler) *RegisterRecipient {
	return &RegisterRecipient{Context: ctx, Handler: handler}
}

/*RegisterRecipient swagger:route POST /recipients vaccination registerRecipient

Vaccination recipient registration can be based on pre enrollment or ad-hoc

Vaccination registration captures information like name, dob, nationality etc for the vaccination certificate.

*/
type RegisterRecipient struct {
	Context *middleware.Context
	Handler RegisterRecipientHandler
}

func (o *RegisterRecipient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegisterRecipientParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
