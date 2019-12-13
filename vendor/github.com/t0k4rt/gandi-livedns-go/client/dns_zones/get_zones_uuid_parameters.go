// Code generated by go-swagger; DO NOT EDIT.

package dns_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewGetZonesUUIDParams creates a new GetZonesUUIDParams object
// with the default values initialized.
func NewGetZonesUUIDParams() *GetZonesUUIDParams {
	var ()
	return &GetZonesUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetZonesUUIDParamsWithTimeout creates a new GetZonesUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetZonesUUIDParamsWithTimeout(timeout time.Duration) *GetZonesUUIDParams {
	var ()
	return &GetZonesUUIDParams{

		timeout: timeout,
	}
}

// NewGetZonesUUIDParamsWithContext creates a new GetZonesUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetZonesUUIDParamsWithContext(ctx context.Context) *GetZonesUUIDParams {
	var ()
	return &GetZonesUUIDParams{

		Context: ctx,
	}
}

// NewGetZonesUUIDParamsWithHTTPClient creates a new GetZonesUUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetZonesUUIDParamsWithHTTPClient(client *http.Client) *GetZonesUUIDParams {
	var ()
	return &GetZonesUUIDParams{
		HTTPClient: client,
	}
}

/*GetZonesUUIDParams contains all the parameters to send to the API endpoint
for the get zones UUID operation typically these are written to a http.Request
*/
type GetZonesUUIDParams struct {

	/*UUID
	  Zone uuid

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get zones UUID params
func (o *GetZonesUUIDParams) WithTimeout(timeout time.Duration) *GetZonesUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get zones UUID params
func (o *GetZonesUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get zones UUID params
func (o *GetZonesUUIDParams) WithContext(ctx context.Context) *GetZonesUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get zones UUID params
func (o *GetZonesUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get zones UUID params
func (o *GetZonesUUIDParams) WithHTTPClient(client *http.Client) *GetZonesUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get zones UUID params
func (o *GetZonesUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get zones UUID params
func (o *GetZonesUUIDParams) WithUUID(uuid string) *GetZonesUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get zones UUID params
func (o *GetZonesUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetZonesUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
