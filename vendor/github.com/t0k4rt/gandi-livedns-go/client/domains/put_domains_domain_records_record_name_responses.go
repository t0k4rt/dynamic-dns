// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// PutDomainsDomainRecordsRecordNameReader is a Reader for the PutDomainsDomainRecordsRecordName structure.
type PutDomainsDomainRecordsRecordNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDomainsDomainRecordsRecordNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutDomainsDomainRecordsRecordNameCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutDomainsDomainRecordsRecordNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutDomainsDomainRecordsRecordNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutDomainsDomainRecordsRecordNameCreated creates a PutDomainsDomainRecordsRecordNameCreated with default headers values
func NewPutDomainsDomainRecordsRecordNameCreated() *PutDomainsDomainRecordsRecordNameCreated {
	return &PutDomainsDomainRecordsRecordNameCreated{}
}

/*PutDomainsDomainRecordsRecordNameCreated handles this case with default header values.

OK
*/
type PutDomainsDomainRecordsRecordNameCreated struct {
	Payload *models.Return200
}

func (o *PutDomainsDomainRecordsRecordNameCreated) Error() string {
	return fmt.Sprintf("[PUT /domains/{domain}/records/{record_name}][%d] putDomainsDomainRecordsRecordNameCreated  %+v", 201, o.Payload)
}

func (o *PutDomainsDomainRecordsRecordNameCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDomainsDomainRecordsRecordNameBadRequest creates a PutDomainsDomainRecordsRecordNameBadRequest with default headers values
func NewPutDomainsDomainRecordsRecordNameBadRequest() *PutDomainsDomainRecordsRecordNameBadRequest {
	return &PutDomainsDomainRecordsRecordNameBadRequest{}
}

/*PutDomainsDomainRecordsRecordNameBadRequest handles this case with default header values.

Not OK
*/
type PutDomainsDomainRecordsRecordNameBadRequest struct {
	Payload *models.Return400
}

func (o *PutDomainsDomainRecordsRecordNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /domains/{domain}/records/{record_name}][%d] putDomainsDomainRecordsRecordNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutDomainsDomainRecordsRecordNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDomainsDomainRecordsRecordNameDefault creates a PutDomainsDomainRecordsRecordNameDefault with default headers values
func NewPutDomainsDomainRecordsRecordNameDefault(code int) *PutDomainsDomainRecordsRecordNameDefault {
	return &PutDomainsDomainRecordsRecordNameDefault{
		_statusCode: code,
	}
}

/*PutDomainsDomainRecordsRecordNameDefault handles this case with default header values.

Unexpected error
*/
type PutDomainsDomainRecordsRecordNameDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the put domains domain records record name default response
func (o *PutDomainsDomainRecordsRecordNameDefault) Code() int {
	return o._statusCode
}

func (o *PutDomainsDomainRecordsRecordNameDefault) Error() string {
	return fmt.Sprintf("[PUT /domains/{domain}/records/{record_name}][%d] PutDomainsDomainRecordsRecordName default  %+v", o._statusCode, o.Payload)
}

func (o *PutDomainsDomainRecordsRecordNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutDomainsDomainRecordsRecordNameBody put domains domain records record name body
swagger:model PutDomainsDomainRecordsRecordNameBody
*/
type PutDomainsDomainRecordsRecordNameBody struct {

	// items
	Items []*models.Record `json:"items"`
}

// Validate validates this put domains domain records record name body
func (o *PutDomainsDomainRecordsRecordNameBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutDomainsDomainRecordsRecordNameBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("record" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutDomainsDomainRecordsRecordNameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutDomainsDomainRecordsRecordNameBody) UnmarshalBinary(b []byte) error {
	var res PutDomainsDomainRecordsRecordNameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
