// Code generated by go-swagger; DO NOT EDIT.

package records

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

// NewDeleteZonesUUIDRecordsParams creates a new DeleteZonesUUIDRecordsParams object
// with the default values initialized.
func NewDeleteZonesUUIDRecordsParams() *DeleteZonesUUIDRecordsParams {
	var ()
	return &DeleteZonesUUIDRecordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteZonesUUIDRecordsParamsWithTimeout creates a new DeleteZonesUUIDRecordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteZonesUUIDRecordsParamsWithTimeout(timeout time.Duration) *DeleteZonesUUIDRecordsParams {
	var ()
	return &DeleteZonesUUIDRecordsParams{

		timeout: timeout,
	}
}

// NewDeleteZonesUUIDRecordsParamsWithContext creates a new DeleteZonesUUIDRecordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteZonesUUIDRecordsParamsWithContext(ctx context.Context) *DeleteZonesUUIDRecordsParams {
	var ()
	return &DeleteZonesUUIDRecordsParams{

		Context: ctx,
	}
}

// NewDeleteZonesUUIDRecordsParamsWithHTTPClient creates a new DeleteZonesUUIDRecordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteZonesUUIDRecordsParamsWithHTTPClient(client *http.Client) *DeleteZonesUUIDRecordsParams {
	var ()
	return &DeleteZonesUUIDRecordsParams{
		HTTPClient: client,
	}
}

/*DeleteZonesUUIDRecordsParams contains all the parameters to send to the API endpoint
for the delete zones UUID records operation typically these are written to a http.Request
*/
type DeleteZonesUUIDRecordsParams struct {

	/*UUID
	  Zone uuid

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) WithTimeout(timeout time.Duration) *DeleteZonesUUIDRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) WithContext(ctx context.Context) *DeleteZonesUUIDRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) WithHTTPClient(client *http.Client) *DeleteZonesUUIDRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) WithUUID(uuid string) *DeleteZonesUUIDRecordsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete zones UUID records params
func (o *DeleteZonesUUIDRecordsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteZonesUUIDRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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