package leaseweb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
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

func (pca PrivateCloudApi) getPath(endpoint string) string {
	return "/cloud/" + PRIVATE_CLOUD_API_VERSION + endpoint
}

func (pca PrivateCloudApi) List(ctx context.Context, args ...interface{}) (*PrivateClouds, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := pca.getPath("/privateClouds")
	query := v.Encode()
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

func (pca PrivateCloudApi) ListCredentials(ctx context.Context, privateCloudId string, credentialType string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := pca.getPath("/privateClouds/" + privateCloudId + "/credentials/" + credentialType)
	query := v.Encode()
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

func (pca PrivateCloudApi) GetDataTrafficMetrics(ctx context.Context, privateCloudId string, args ...interface{}) (*DataTrafficMetricsV2, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("granularity", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("aggregation", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("from", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("to", fmt.Sprint(args[3]))
	}

	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/datatraffic")
	query := v.Encode()
	result := &DataTrafficMetricsV2{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetBandWidthMetrics(ctx context.Context, privateCloudId string, args ...interface{}) (*BandWidthMetrics, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("granularity", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("aggregation", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("from", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("to", fmt.Sprint(args[3]))
	}

	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/bandwidth")
	query := v.Encode()
	result := &BandWidthMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetCpuMetrics(ctx context.Context, privateCloudId string, args ...interface{}) (*PrivateCloudCpuMetrics, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("granularity", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("aggregation", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("from", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("to", fmt.Sprint(args[3]))
	}
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/cpu")
	query := v.Encode()
	result := &PrivateCloudCpuMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetMemoryMetrics(ctx context.Context, privateCloudId string, args ...interface{}) (*PrivateCloudMemoryMetrics, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("granularity", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("aggregation", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("from", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("to", fmt.Sprint(args[3]))
	}
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/memory")
	query := v.Encode()
	result := &PrivateCloudMemoryMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetStorageMetrics(ctx context.Context, privateCloudId string, args ...interface{}) (*PrivateCloudStorageMetrics, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("granularity", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("aggregation", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("from", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("to", fmt.Sprint(args[3]))
	}
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/storage")
	query := v.Encode()
	result := &PrivateCloudStorageMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}
