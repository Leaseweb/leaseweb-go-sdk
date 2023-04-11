package leaseweb

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const SERVICES_API_VERSION = "v1"

type ServicesApi struct{}

type Services struct {
	Services []Service `json:"services"`
	Metadata Metadata  `json:"metadata"`
}

type Service struct {
	// TODO;
	// "attributes": {}
	BillingCycle        string      `json:"billingCycle"`
	Cancellable         bool        `json:"cancellable"`
	ContractId          string      `json:"contractId"`
	ContractTerm        string      `json:"contractTerm"`
	ContractTermEndDate string      `json:"contractTermEndDate"`
	Currency            string      `json:"currency"`
	DeliveryDate        string      `json:"deliveryDate"`
	DeliveryEstimate    string      `json:"deliveryEstimate"`
	EndDate             string      `json:"endDate"`
	EquipmentId         string      `json:"equipmentId"`
	Id                  string      `json:"id"`
	OrderDate           string      `json:"orderDate"`
	PricePerFrequency   json.Number `json:"pricePerFrequency"`
	ProductId           string      `json:"productId"`
	Reference           string      `json:"reference"`
	StartDate           string      `json:"startDate"`
	Status              string      `json:"status"`
	Uncancellable       bool        `json:"uncancellable"`
}

type ServicesCancellationReasons struct {
	CancellationReasons []ServicesCancellationReason `json:"cancellationReasons"`
}

type ServicesCancellationReason struct {
	Reason     string `json:"reason"`
	ReasonCode string `json:"reasonCode"`
}

type ServicesListOptions struct {
	PaginationOptions
	ProductId   *string `param:"productId"`
	Reference   *string `param:"reference"`
	EquipmentId *string `param:"equipmentId"`
}

func (sa ServicesApi) getPath(endpoint string) string {
	return "/services/" + SERVICES_API_VERSION + endpoint
}

func (sa ServicesApi) List(ctx context.Context, opts ServicesListOptions) (*Services, error) {
	path := sa.getPath("/services")
	query := options.Encode(opts)
	result := &Services{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (sa ServicesApi) ListCancellationReasons(ctx context.Context) (*ServicesCancellationReasons, error) {
	path := sa.getPath("/services/cancellationReasons")
	result := &ServicesCancellationReasons{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (sa ServicesApi) Get(ctx context.Context, id string) (*Service, error) {
	path := sa.getPath("/services/" + id)
	result := &Service{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (sa ServicesApi) Cancel(ctx context.Context, id, reason, reasonCode string) error {
	payload := map[string]string{
		"reason":     reason,
		"reasonCode": reasonCode,
	}
	path := sa.getPath("/services/" + id + "/cancel")
	return doRequest(ctx, http.MethodPost, path, "", nil, payload)
}

func (sa ServicesApi) Uncancel(ctx context.Context, id string) error {
	path := sa.getPath("/services/" + id + "/uncancel")
	return doRequest(ctx, http.MethodPost, path, "")
}
