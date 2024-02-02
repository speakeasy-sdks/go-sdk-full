// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/utils"
	"net/http"
)

type DeleteSpecificSavedInstrumentRequest struct {
	CustomerID    string  `pathParam:"style=simple,explode=false,name=customer_id"`
	InstrumentID  string  `pathParam:"style=simple,explode=false,name=instrument_id"`
	XAPIVersion   *string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID     string  `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret string  `header:"style=simple,explode=false,name=x-client-secret"`
}

func (d DeleteSpecificSavedInstrumentRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeleteSpecificSavedInstrumentRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeleteSpecificSavedInstrumentRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *DeleteSpecificSavedInstrumentRequest) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *DeleteSpecificSavedInstrumentRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *DeleteSpecificSavedInstrumentRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *DeleteSpecificSavedInstrumentRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type DeleteSpecificSavedInstrumentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	// OK
	FetchAllSavedInstruments *shared.FetchAllSavedInstruments
	Headers                  map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteSpecificSavedInstrumentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSpecificSavedInstrumentResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *DeleteSpecificSavedInstrumentResponse) GetFetchAllSavedInstruments() *shared.FetchAllSavedInstruments {
	if o == nil {
		return nil
	}
	return o.FetchAllSavedInstruments
}

func (o *DeleteSpecificSavedInstrumentResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *DeleteSpecificSavedInstrumentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSpecificSavedInstrumentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
