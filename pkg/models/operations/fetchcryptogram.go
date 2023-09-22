// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/utils"
	"net/http"
)

type FetchCryptogramRequest struct {
	CustomerID    string  `pathParam:"style=simple,explode=false,name=customer_id"`
	InstrumentID  string  `pathParam:"style=simple,explode=false,name=instrument_id"`
	XAPIVersion   *string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID     string  `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret string  `header:"style=simple,explode=false,name=x-client-secret"`
}

func (f FetchCryptogramRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FetchCryptogramRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FetchCryptogramRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *FetchCryptogramRequest) GetInstrumentID() string {
	if o == nil {
		return ""
	}
	return o.InstrumentID
}

func (o *FetchCryptogramRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *FetchCryptogramRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *FetchCryptogramRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type FetchCryptogramResponse struct {
	ContentType string
	// OK
	Cryptogram *shared.Cryptogram
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	StatusCode    int
	RawResponse   *http.Response
}

func (o *FetchCryptogramResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FetchCryptogramResponse) GetCryptogram() *shared.Cryptogram {
	if o == nil {
		return nil
	}
	return o.Cryptogram
}

func (o *FetchCryptogramResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *FetchCryptogramResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *FetchCryptogramResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FetchCryptogramResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
