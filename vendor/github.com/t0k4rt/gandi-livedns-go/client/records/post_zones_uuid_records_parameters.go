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
	models "github.com/t0k4rt/gandi-livedns-go/models"
	"golang.org/x/net/context"
)

// NewPostZonesUUIDRecordsParams creates a new PostZonesUUIDRecordsParams object
// with the default values initialized.
func NewPostZonesUUIDRecordsParams() *PostZonesUUIDRecordsParams {
	var ()
	return &PostZonesUUIDRecordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostZonesUUIDRecordsParamsWithTimeout creates a new PostZonesUUIDRecordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostZonesUUIDRecordsParamsWithTimeout(timeout time.Duration) *PostZonesUUIDRecordsParams {
	var ()
	return &PostZonesUUIDRecordsParams{

		timeout: timeout,
	}
}

// NewPostZonesUUIDRecordsParamsWithContext creates a new PostZonesUUIDRecordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostZonesUUIDRecordsParamsWithContext(ctx context.Context) *PostZonesUUIDRecordsParams {
	var ()
	return &PostZonesUUIDRecordsParams{

		Context: ctx,
	}
}

// NewPostZonesUUIDRecordsParamsWithHTTPClient creates a new PostZonesUUIDRecordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostZonesUUIDRecordsParamsWithHTTPClient(client *http.Client) *PostZonesUUIDRecordsParams {
	var ()
	return &PostZonesUUIDRecordsParams{
		HTTPClient: client,
	}
}

/*PostZonesUUIDRecordsParams contains all the parameters to send to the API endpoint
for the post zones UUID records operation typically these are written to a http.Request
*/
type PostZonesUUIDRecordsParams struct {

	/*Record
	  Zone name

	*/
	Record *models.Record
	/*UUID
	  Zone uuid

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) WithTimeout(timeout time.Duration) *PostZonesUUIDRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) WithContext(ctx context.Context) *PostZonesUUIDRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) WithHTTPClient(client *http.Client) *PostZonesUUIDRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecord adds the record to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) WithRecord(record *models.Record) *PostZonesUUIDRecordsParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) SetRecord(record *models.Record) {
	o.Record = record
}

// WithUUID adds the uuid to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) WithUUID(uuid string) *PostZonesUUIDRecordsParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the post zones UUID records params
func (o *PostZonesUUIDRecordsParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PostZonesUUIDRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}