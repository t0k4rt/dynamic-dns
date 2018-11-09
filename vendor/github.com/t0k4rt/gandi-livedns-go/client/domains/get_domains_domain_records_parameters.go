// Code generated by go-swagger; DO NOT EDIT.

package domains

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

// NewGetDomainsDomainRecordsParams creates a new GetDomainsDomainRecordsParams object
// with the default values initialized.
func NewGetDomainsDomainRecordsParams() *GetDomainsDomainRecordsParams {
	var ()
	return &GetDomainsDomainRecordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainsDomainRecordsParamsWithTimeout creates a new GetDomainsDomainRecordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDomainsDomainRecordsParamsWithTimeout(timeout time.Duration) *GetDomainsDomainRecordsParams {
	var ()
	return &GetDomainsDomainRecordsParams{

		timeout: timeout,
	}
}

// NewGetDomainsDomainRecordsParamsWithContext creates a new GetDomainsDomainRecordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDomainsDomainRecordsParamsWithContext(ctx context.Context) *GetDomainsDomainRecordsParams {
	var ()
	return &GetDomainsDomainRecordsParams{

		Context: ctx,
	}
}

// NewGetDomainsDomainRecordsParamsWithHTTPClient creates a new GetDomainsDomainRecordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDomainsDomainRecordsParamsWithHTTPClient(client *http.Client) *GetDomainsDomainRecordsParams {
	var ()
	return &GetDomainsDomainRecordsParams{
		HTTPClient: client,
	}
}

/*GetDomainsDomainRecordsParams contains all the parameters to send to the API endpoint
for the get domains domain records operation typically these are written to a http.Request
*/
type GetDomainsDomainRecordsParams struct {

	/*Domain
	  Domain to inspect

	*/
	Domain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) WithTimeout(timeout time.Duration) *GetDomainsDomainRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) WithContext(ctx context.Context) *GetDomainsDomainRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) WithHTTPClient(client *http.Client) *GetDomainsDomainRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) WithDomain(domain string) *GetDomainsDomainRecordsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the get domains domain records params
func (o *GetDomainsDomainRecordsParams) SetDomain(domain string) {
	o.Domain = domain
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainsDomainRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}