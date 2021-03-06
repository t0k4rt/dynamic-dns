// Code generated by go-swagger; DO NOT EDIT.

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// GetDomainsDomainRecordsRecordNameRecordTypeReader is a Reader for the GetDomainsDomainRecordsRecordNameRecordType structure.
type GetDomainsDomainRecordsRecordNameRecordTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsDomainRecordsRecordNameRecordTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDomainsDomainRecordsRecordNameRecordTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDomainsDomainRecordsRecordNameRecordTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDomainsDomainRecordsRecordNameRecordTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainsDomainRecordsRecordNameRecordTypeOK creates a GetDomainsDomainRecordsRecordNameRecordTypeOK with default headers values
func NewGetDomainsDomainRecordsRecordNameRecordTypeOK() *GetDomainsDomainRecordsRecordNameRecordTypeOK {
	return &GetDomainsDomainRecordsRecordNameRecordTypeOK{}
}

/*GetDomainsDomainRecordsRecordNameRecordTypeOK handles this case with default header values.

OK
*/
type GetDomainsDomainRecordsRecordNameRecordTypeOK struct {
	Payload *models.Record
}

func (o *GetDomainsDomainRecordsRecordNameRecordTypeOK) Error() string {
	return fmt.Sprintf("[GET /domains/{domain}/records/{record_name}/{record_type}][%d] getDomainsDomainRecordsRecordNameRecordTypeOK  %+v", 200, o.Payload)
}

func (o *GetDomainsDomainRecordsRecordNameRecordTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Record)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsDomainRecordsRecordNameRecordTypeBadRequest creates a GetDomainsDomainRecordsRecordNameRecordTypeBadRequest with default headers values
func NewGetDomainsDomainRecordsRecordNameRecordTypeBadRequest() *GetDomainsDomainRecordsRecordNameRecordTypeBadRequest {
	return &GetDomainsDomainRecordsRecordNameRecordTypeBadRequest{}
}

/*GetDomainsDomainRecordsRecordNameRecordTypeBadRequest handles this case with default header values.

Not OK
*/
type GetDomainsDomainRecordsRecordNameRecordTypeBadRequest struct {
	Payload *models.Return400
}

func (o *GetDomainsDomainRecordsRecordNameRecordTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /domains/{domain}/records/{record_name}/{record_type}][%d] getDomainsDomainRecordsRecordNameRecordTypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetDomainsDomainRecordsRecordNameRecordTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsDomainRecordsRecordNameRecordTypeDefault creates a GetDomainsDomainRecordsRecordNameRecordTypeDefault with default headers values
func NewGetDomainsDomainRecordsRecordNameRecordTypeDefault(code int) *GetDomainsDomainRecordsRecordNameRecordTypeDefault {
	return &GetDomainsDomainRecordsRecordNameRecordTypeDefault{
		_statusCode: code,
	}
}

/*GetDomainsDomainRecordsRecordNameRecordTypeDefault handles this case with default header values.

Unexpected error
*/
type GetDomainsDomainRecordsRecordNameRecordTypeDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the get domains domain records record name record type default response
func (o *GetDomainsDomainRecordsRecordNameRecordTypeDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainsDomainRecordsRecordNameRecordTypeDefault) Error() string {
	return fmt.Sprintf("[GET /domains/{domain}/records/{record_name}/{record_type}][%d] GetDomainsDomainRecordsRecordNameRecordType default  %+v", o._statusCode, o.Payload)
}

func (o *GetDomainsDomainRecordsRecordNameRecordTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
