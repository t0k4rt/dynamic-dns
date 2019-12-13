// Code generated by go-swagger; DO NOT EDIT.

package records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// DeleteZonesUUIDRecordsRecordNameRecordTypeReader is a Reader for the DeleteZonesUUIDRecordsRecordNameRecordType structure.
type DeleteZonesUUIDRecordsRecordNameRecordTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteZonesUUIDRecordsRecordNameRecordTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteZonesUUIDRecordsRecordNameRecordTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteZonesUUIDRecordsRecordNameRecordTypeOK creates a DeleteZonesUUIDRecordsRecordNameRecordTypeOK with default headers values
func NewDeleteZonesUUIDRecordsRecordNameRecordTypeOK() *DeleteZonesUUIDRecordsRecordNameRecordTypeOK {
	return &DeleteZonesUUIDRecordsRecordNameRecordTypeOK{}
}

/*DeleteZonesUUIDRecordsRecordNameRecordTypeOK handles this case with default header values.

OK
*/
type DeleteZonesUUIDRecordsRecordNameRecordTypeOK struct {
	Payload *models.Return200
}

func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeOK) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}/records/{record_name}/{record_type}][%d] deleteZonesUuidRecordsRecordNameRecordTypeOK  %+v", 200, o.Payload)
}

func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest creates a DeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest with default headers values
func NewDeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest() *DeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest {
	return &DeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest{}
}

/*DeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest handles this case with default header values.

Not OK
*/
type DeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest struct {
	Payload *models.Return400
}

func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}/records/{record_name}/{record_type}][%d] deleteZonesUuidRecordsRecordNameRecordTypeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZonesUUIDRecordsRecordNameRecordTypeDefault creates a DeleteZonesUUIDRecordsRecordNameRecordTypeDefault with default headers values
func NewDeleteZonesUUIDRecordsRecordNameRecordTypeDefault(code int) *DeleteZonesUUIDRecordsRecordNameRecordTypeDefault {
	return &DeleteZonesUUIDRecordsRecordNameRecordTypeDefault{
		_statusCode: code,
	}
}

/*DeleteZonesUUIDRecordsRecordNameRecordTypeDefault handles this case with default header values.

Unexpected error
*/
type DeleteZonesUUIDRecordsRecordNameRecordTypeDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the delete zones UUID records record name record type default response
func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeDefault) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}/records/{record_name}/{record_type}][%d] DeleteZonesUUIDRecordsRecordNameRecordType default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteZonesUUIDRecordsRecordNameRecordTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}