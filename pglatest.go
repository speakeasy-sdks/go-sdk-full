// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package gosdkfull

import (
	"fmt"
	"github.com/speakeasy-sdks/go-sdk-full/v3/internal/hooks"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Sandbox server
	"https://sandbox.cashfree.com/pg",
	// Production server
	"https://api.cashfree.com/pg",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

type PGLatest struct {
	// Cashfree's token Vault helps you save cards and tokenize them in a PCI complaint manner. We support creation of network tokens which can be used across acquiring banks
	TokenVault      *TokenVault
	EligibilityAPIs *EligibilityAPIs
	PaymentLinks    *PaymentLinks
	Offers          *Offers
	Orders          *Orders
	// The Authentication API allows merchants to show a native screen and capture OTP on their own page and submit to Cashfree. This feature is only available on request.
	Authentication *Authentication
	Payments       *Payments
	Refunds        *Refunds
	Settlements    *Settlements
	Reconciliation *Reconciliation
	// softPOS' agent and order management system now supported by APIs
	SoftPOS *SoftPOS

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*PGLatest)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *PGLatest) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *PGLatest) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *PGLatest) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *PGLatest) {
		sdk.sdkConfiguration.Client = client
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *PGLatest) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *PGLatest {
	sdk := &PGLatest{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2022-09-01",
			SDKVersion:        "3.3.1",
			GenVersion:        "2.280.6",
			UserAgent:         "speakeasy-sdk/go 3.3.1 2.280.6 2022-09-01 github.com/speakeasy-sdks/go-sdk-full",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.TokenVault = newTokenVault(sdk.sdkConfiguration)

	sdk.EligibilityAPIs = newEligibilityAPIs(sdk.sdkConfiguration)

	sdk.PaymentLinks = newPaymentLinks(sdk.sdkConfiguration)

	sdk.Offers = newOffers(sdk.sdkConfiguration)

	sdk.Orders = newOrders(sdk.sdkConfiguration)

	sdk.Authentication = newAuthentication(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.Refunds = newRefunds(sdk.sdkConfiguration)

	sdk.Settlements = newSettlements(sdk.sdkConfiguration)

	sdk.Reconciliation = newReconciliation(sdk.sdkConfiguration)

	sdk.SoftPOS = newSoftPOS(sdk.sdkConfiguration)

	return sdk
}
