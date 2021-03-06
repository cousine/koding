package j_compute_stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJComputeStackDeleteParams creates a new JComputeStackDeleteParams object
// with the default values initialized.
func NewJComputeStackDeleteParams() *JComputeStackDeleteParams {
	var ()
	return &JComputeStackDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJComputeStackDeleteParamsWithTimeout creates a new JComputeStackDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJComputeStackDeleteParamsWithTimeout(timeout time.Duration) *JComputeStackDeleteParams {
	var ()
	return &JComputeStackDeleteParams{

		timeout: timeout,
	}
}

// NewJComputeStackDeleteParamsWithContext creates a new JComputeStackDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewJComputeStackDeleteParamsWithContext(ctx context.Context) *JComputeStackDeleteParams {
	var ()
	return &JComputeStackDeleteParams{

		Context: ctx,
	}
}

/*JComputeStackDeleteParams contains all the parameters to send to the API endpoint
for the j compute stack delete operation typically these are written to a http.Request
*/
type JComputeStackDeleteParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j compute stack delete params
func (o *JComputeStackDeleteParams) WithTimeout(timeout time.Duration) *JComputeStackDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j compute stack delete params
func (o *JComputeStackDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j compute stack delete params
func (o *JComputeStackDeleteParams) WithContext(ctx context.Context) *JComputeStackDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j compute stack delete params
func (o *JComputeStackDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j compute stack delete params
func (o *JComputeStackDeleteParams) WithBody(body models.DefaultSelector) *JComputeStackDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j compute stack delete params
func (o *JComputeStackDeleteParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j compute stack delete params
func (o *JComputeStackDeleteParams) WithID(id string) *JComputeStackDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j compute stack delete params
func (o *JComputeStackDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JComputeStackDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
