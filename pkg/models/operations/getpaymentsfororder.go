// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"net/http"
)

type GetPaymentsfororderRequest struct {
	OrderID       string  `pathParam:"style=simple,explode=false,name=order_id"`
	XAPIVersion   *string `header:"style=simple,explode=false,name=x-api-version"`
	XClientID     string  `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret string  `header:"style=simple,explode=false,name=x-client-secret"`
}

func (o *GetPaymentsfororderRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *GetPaymentsfororderRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *GetPaymentsfororderRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *GetPaymentsfororderRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type GetPaymentsfororder200ApplicationJSONType string

const (
	GetPaymentsfororder200ApplicationJSONTypePaymentsEntity GetPaymentsfororder200ApplicationJSONType = "PaymentsEntity"
)

type GetPaymentsfororder200ApplicationJSON struct {
	PaymentsEntity *shared.PaymentsEntity

	Type GetPaymentsfororder200ApplicationJSONType
}

func CreateGetPaymentsfororder200ApplicationJSONPaymentsEntity(paymentsEntity shared.PaymentsEntity) GetPaymentsfororder200ApplicationJSON {
	typ := GetPaymentsfororder200ApplicationJSONTypePaymentsEntity

	return GetPaymentsfororder200ApplicationJSON{
		PaymentsEntity: &paymentsEntity,
		Type:           typ,
	}
}

func (u *GetPaymentsfororder200ApplicationJSON) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	paymentsEntity := new(shared.PaymentsEntity)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&paymentsEntity); err == nil {
		u.PaymentsEntity = paymentsEntity
		u.Type = GetPaymentsfororder200ApplicationJSONTypePaymentsEntity
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetPaymentsfororder200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.PaymentsEntity != nil {
		return json.Marshal(u.PaymentsEntity)
	}

	return nil, nil
}

type GetPaymentsfororderResponse struct {
	ContentType string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	// OK
	GetPaymentsfororder200ApplicationJSONOneOf *GetPaymentsfororder200ApplicationJSON
	Headers                                    map[string][]string
	StatusCode                                 int
	RawResponse                                *http.Response
}

func (o *GetPaymentsfororderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentsfororderResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *GetPaymentsfororderResponse) GetGetPaymentsfororder200ApplicationJSONOneOf() *GetPaymentsfororder200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetPaymentsfororder200ApplicationJSONOneOf
}

func (o *GetPaymentsfororderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetPaymentsfororderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentsfororderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
