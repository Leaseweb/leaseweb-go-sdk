package leaseweb

import (
	"context"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const PRIVATE_CLOUD_API_VERSION = "v2"

type PrivateCloudApi struct{}

type PrivateClouds struct {
	PrivateClouds []PrivateCloud `json:"privateClouds"`
	Metadata      Metadata       `json:"_metadata"`
}

type PrivateCloud struct {
	Id              string         `json:"id"`
	CustomerId      string         `json:"customerId"`
	DataCenter      string         `json:"dataCenter"`
	ServiceOffering string         `json:"serviceOffering"`
	Sla             string         `json:"sla"`
	Contract        Contract       `json:"contract"`
	NetworkTraffic  NetworkTraffic `json:"networkTraffic"`
	Ips             []Ip           `json:"ips"`
	Hardware        Hardware       `json:"hardware"`
}

type PrivateCloudCpuMetrics struct {
	Metric   PrivateCloudCpuMetric `json:"metrics"`
	Metadata MetricMetadata        `json:"_metadata"`
}

type PrivateCloudCpuMetric struct {
	Cpu BasicMetric `json:"CPU"`
}

type PrivateCloudMemoryMetrics struct {
	Metric   PrivateCloudMemoryMetric `json:"metrics"`
	Metadata MetricMetadata           `json:"_metadata"`
}

type PrivateCloudMemoryMetric struct {
	Memory BasicMetric `json:"MEMORY"`
}

type PrivateCloudStorageMetrics struct {
	Metric   PrivateCloudStorageMetric `json:"metrics"`
	Metadata MetricMetadata            `json:"_metadata"`
}

type PrivateCloudStorageMetric struct {
	Storage BasicMetric `json:"STORAGE"`
}

type PrivateCloudListOptions struct {
	Offset *int `param:"offset"`
	Limit  *int `param:"limit"`
}

type PrivateCloudListCredentialsOptions struct {
	Offset *int `param:"offset"`
	Limit  *int `param:"limit"`
}

type PrivateCloudMetricsDataTrafficOptions struct {
	Granularity *string `param:"granularity"`
	Aggregation *string `param:"aggregation"`
	From        *string `param:"from"`
	To          *string `param:"to"`
}

type PrivateCloudMetricsBandWidthOptions struct {
	Granularity *string `param:"granularity"`
	Aggregation *string `param:"aggregation"`
	From        *string `param:"from"`
	To          *string `param:"to"`
}

type PrivateCloudMetricsCpuOptions struct {
	Granularity *string `param:"granularity"`
	Aggregation *string `param:"aggregation"`
	From        *string `param:"from"`
	To          *string `param:"to"`
}

type PrivateCloudMetricsMemoryOptions struct {
	Granularity *string `param:"granularity"`
	Aggregation *string `param:"aggregation"`
	From        *string `param:"from"`
	To          *string `param:"to"`
}

type PrivateCloudMetricsStorageOptions struct {
	Granularity *string `param:"granularity"`
	Aggregation *string `param:"aggregation"`
	From        *string `param:"from"`
	To          *string `param:"to"`
}

func (pca PrivateCloudApi) getPath(endpoint string) string {
	return "/cloud/" + PRIVATE_CLOUD_API_VERSION + endpoint
}

func (pca PrivateCloudApi) List(ctx context.Context, opts PrivateCloudListOptions) (*PrivateClouds, error) {
	path := pca.getPath("/privateClouds")
	query := options.Encode(opts)
	result := &PrivateClouds{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) Get(ctx context.Context, privateCloudId string) (*PrivateCloud, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId)
	result := &PrivateCloud{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) ListCredentials(ctx context.Context, privateCloudId string, credentialType string, opts PrivateCloudListCredentialsOptions) (*Credentials, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/credentials/" + credentialType)
	query := options.Encode(opts)
	result := &Credentials{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetCredential(ctx context.Context, privateCloudId string, credentialType string, username string) (*Credential, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/credentials/" + credentialType + "/" + username)
	result := &Credential{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetDataTrafficMetrics(ctx context.Context, privateCloudId string, opts PrivateCloudMetricsDataTrafficOptions) (*DataTrafficMetricsV2, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/datatraffic")
	query := options.Encode(opts)
	result := &DataTrafficMetricsV2{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetBandWidthMetrics(ctx context.Context, privateCloudId string, opts PrivateCloudMetricsBandWidthOptions) (*BandWidthMetrics, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/bandwidth")
	query := options.Encode(opts)
	result := &BandWidthMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetCpuMetrics(ctx context.Context, privateCloudId string, opts PrivateCloudMetricsCpuOptions) (*PrivateCloudCpuMetrics, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/cpu")
	query := options.Encode(opts)
	result := &PrivateCloudCpuMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetMemoryMetrics(ctx context.Context, privateCloudId string, opts PrivateCloudMetricsMemoryOptions) (*PrivateCloudMemoryMetrics, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/memory")
	query := options.Encode(opts)
	result := &PrivateCloudMemoryMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetStorageMetrics(ctx context.Context, privateCloudId string, opts PrivateCloudMetricsStorageOptions) (*PrivateCloudStorageMetrics, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/storage")
	query := options.Encode(opts)
	result := &PrivateCloudStorageMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}
