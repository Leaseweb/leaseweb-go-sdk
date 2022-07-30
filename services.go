package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
)

const SERVICES_API_VERSION = "v1"

type ServicesApi struct {
}

type Services struct {
	Services []Service `json:"services"`
	Metadata Metadata  `json:"metadata"`
}

type Service struct {
	// TODO;
	// "attributes": {}
	BillingCycle        string  `json:"billingCycle"`
	Cancellable         bool    `json:"cancellable"`
	ContractId          string  `json:"contractId"`
	ContractTerm        string  `json:"contractTerm"`
	ContractTermEndDate string  `json:"contractTermEndDate"`
	Currency            string  `json:"currency"`
	DeliveryDate        string  `json:"deliveryDate"`
	DeliveryEstimate    string  `json:"deliveryEstimate"`
	EndDate             string  `json:"endDate"`
	EquipmentId         string  `json:"equipmentId"`
	Id                  string  `json:"id"`
	OrderDate           string  `json:"orderDate"`
	PricePerFrequency   float32 `json:"pricePerFrequency"`
	ProductId           string  `json:"productId"`
	Reference           string  `json:"reference"`
	StartDate           string  `json:"startDate"`
	Status              string  `json:"status"`
	Uncancellable       bool    `json:"uncancellable"`
}

type CancellationReasons struct {
	CancellationReasons []CancellationReason `json:"cancellationReasons"`
}

type CancellationReason struct {
	Reason     string `json:"reason"`
	ReasonCode string `json:"reasonCode"`
}

func (sa ServicesApi) getPath(endpoint string) string {
	return "/services/" + SERVICES_API_VERSION + endpoint
}

func (sa ServicesApi) ListServices(args ...int) (*Services, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := sa.getPath("/services?" + v.Encode())
	result := &Services{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (sa ServicesApi) ListCancellationReasons() (*CancellationReasons, error) {
	path := sa.getPath("/services/cancellationReasons")
	result := &CancellationReasons{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (sa ServicesApi) GetService(id string) (*Service, error) {
	path := sa.getPath("/services/" + id)
	result := &Service{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (sa ServicesApi) CancelService(id, reason, reasonCode string) error {
	payload := map[string]string{
		"reason":     reason,
		"reasonCode": reasonCode,
	}
	path := sa.getPath("/services/" + id + "/cancel")
	return doRequest(http.MethodPost, path, nil, payload)
}

func (sa ServicesApi) UncancelService(id string) error {
	path := sa.getPath("/services/" + id + "/uncancel")
	return doRequest(http.MethodPost, path)
}
