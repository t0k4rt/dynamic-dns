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

// DeleteZonesUUIDRecordsRecordNameReader is a Reader for the DeleteZonesUUIDRecordsRecordName structure.
type DeleteZonesUUIDRecordsRecordNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteZonesUUIDRecordsRecordNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteZonesUUIDRecordsRecordNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteZonesUUIDRecordsRecordNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteZonesUUIDRecordsRecordNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteZonesUUIDRecordsRecordNameOK creates a DeleteZonesUUIDRecordsRecordNameOK with default headers values
func NewDeleteZonesUUIDRecordsRecordNameOK() *DeleteZonesUUIDRecordsRecordNameOK {
	return &DeleteZonesUUIDRecordsRecordNameOK{}
}

/*DeleteZonesUUIDRecordsRecordNameOK handles this case with default header values.

OK
*/
type DeleteZonesUUIDRecordsRecordNameOK struct {
	Payload *models.Return200
}

func (o *DeleteZonesUUIDRecordsRecordNameOK) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}/records/{record_name}][%d] deleteZonesUuidRecordsRecordNameOK  %+v", 200, o.Payload)
}

func (o *DeleteZonesUUIDRecordsRecordNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZonesUUIDRecordsRecordNameBadRequest creates a DeleteZonesUUIDRecordsRecordNameBadRequest with default headers values
func NewDeleteZonesUUIDRecordsRecordNameBadRequest() *DeleteZonesUUIDRecordsRecordNameBadRequest {
	return &DeleteZonesUUIDRecordsRecordNameBadRequest{}
}

/*DeleteZonesUUIDRecordsRecordNameBadRequest handles this case with default header values.

Not OK
*/
type DeleteZonesUUIDRecordsRecordNameBadRequest struct {
	Payload *models.Return400
}

func (o *DeleteZonesUUIDRecordsRecordNameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}/records/{record_name}][%d] deleteZonesUuidRecordsRecordNameBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteZonesUUIDRecordsRecordNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteZonesUUIDRecordsRecordNameDefault creates a DeleteZonesUUIDRecordsRecordNameDefault with default headers values
func NewDeleteZonesUUIDRecordsRecordNameDefault(code int) *DeleteZonesUUIDRecordsRecordNameDefault {
	return &DeleteZonesUUIDRecordsRecordNameDefault{
		_statusCode: code,
	}
}

/*DeleteZonesUUIDRecordsRecordNameDefault handles this case with default header values.

Unexpected error
*/
type DeleteZonesUUIDRecordsRecordNameDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the delete zones UUID records record name default response
func (o *DeleteZonesUUIDRecordsRecordNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteZonesUUIDRecordsRecordNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /zones/{uuid}/records/{record_name}][%d] DeleteZonesUUIDRecordsRecordName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteZonesUUIDRecordsRecordNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
