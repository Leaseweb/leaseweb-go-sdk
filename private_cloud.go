package leaseweb

import (
	"fmt"
	"net/url"
)

const PRIVATE_CLOUD_API_DEFAULT_VERSION = "v2"

type PrivateCloudApi struct {
	version string
}

type PrivateClouds struct {
	PrivateClouds []PrivateCloud `json:"privateClouds"`
	Metadata      Metadata       `json:"_metadata"`
}

type PrivateCloud struct {
	Id              string               `json:"id"`
	CustomerId      string               `json:"customerId"`
	DataCenter      string               `json:"dataCenter"`
	ServiceOffering string               `json:"serviceOffering"`
	Sla             string               `json:"sla"`
	Contract        PrivateCloudContract `json:"contract"`
	NetworkTraffic  NetworkTraffic       `json:"networkTraffic"`
	Ips             []PrivateCloudIp     `json:"ips"`
	Hardware        PrivateCloudHardware `json:"hardware"`
}

type PrivateCloudContract struct {
	Id                string  `json:"id"`
	StartsAt          string  `json:"startsAt"`
	EndsAt            string  `json:"endsAt"`
	BillingCycle      int     `json:"billingCycle"`
	BillingFrequency  string  `json:"billingFrequency"`
	PricePerFrequency float32 `json:"pricePerFrequency"`
	Currency          string  `json:"currency"`
}

type PrivateCloudIp struct {
	Ip      string `json:"ip"`
	Version int    `json:"version"`
	Type    string `json:"type"`
}

type PrivateCloudHardware struct {
	Cpu     Cpu            `json:"cpu"`
	Memory  UnitAmountPair `json:"memory"`
	Storage UnitAmountPair `json:"storage"`
}

type Cpu struct {
	Cores int `json:"cores"`
}

type Credential struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

type Credentials struct {
	Credentials []Credential `json:"credentials"`
	Metadata    Metadata     `json:"_metadata"`
}

type UnitAmountPair struct {
	Unit   string `json:"unit"`
	Amount int    `json:"amount"`
}

type DataTrafficMetrics struct {
	Metric   DataTrafficMetric `json:"metrics"`
	Metadata MetricMetadata    `json:"_metadata"`
}

type DataTrafficMetric struct {
	DataTrafficUp   BasicMetric `json:"DATATRAFFIC_UP"`
	DataTrafficDown BasicMetric `json:"DATATRAFFIC_DOWN"`
}

type CpuMetrics struct {
	Metric   CpuMetric      `json:"metrics"`
	Metadata MetricMetadata `json:"_metadata"`
}

type CpuMetric struct {
	Cpu BasicMetric `json:"CPU"`
}

type MemoryMetrics struct {
	Metric   MemoryMetric   `json:"metrics"`
	Metadata MetricMetadata `json:"_metadata"`
}

type MemoryMetric struct {
	Memory BasicMetric `json:"MEMORY"`
}

type StorageMetrics struct {
	Metric   StorageMetric  `json:"metrics"`
	Metadata MetricMetadata `json:"_metadata"`
}

type StorageMetric struct {
	Storage BasicMetric `json:"STORAGE"`
}

func (pca *PrivateCloudApi) SetVersion(version string) {
	pca.version = version
}

func (pca PrivateCloudApi) getPath(endpoint string) string {
	version := pca.version
	if version == "" {
		version = PRIVATE_CLOUD_API_DEFAULT_VERSION
	}
	return "/cloud/" + version + endpoint
}

func (pca PrivateCloudApi) ListPrivateClouds(args ...interface{}) (*PrivateClouds, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := pca.getPath("/privateClouds?" + v.Encode())
	result := &PrivateClouds{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetPrivateCloud(privateCloudId string) (*PrivateCloud, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId)
	result := &PrivateCloud{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) ListCredentials(privateCloudId string, credentialType string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := pca.getPath("/privateClouds/" + privateCloudId + "/credentials/" + credentialType + v.Encode())
	result := &Credentials{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetCredentials(privateCloudId string, credentialType string, username string) (*Credential, error) {
	path := pca.getPath("/privateClouds/" + privateCloudId + "/credentials/" + credentialType + "/" + username)
	result := &Credential{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetDataTrafficMetrics(privateCloudId string, args ...interface{}) (*DataTrafficMetrics, error) {
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

	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/datatraffic?" + v.Encode())
	result := &DataTrafficMetrics{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetBandWidthMetrics(privateCloudId string, args ...interface{}) (*BandWidthMetrics, error) {
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

	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/bandwidth?" + v.Encode())
	result := &BandWidthMetrics{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetCpuMetrics(privateCloudId string, args ...interface{}) (*CpuMetrics, error) {
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
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/cpu?" + v.Encode())
	result := &CpuMetrics{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetMemoryMetrics(privateCloudId string, args ...interface{}) (*MemoryMetrics, error) {
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
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/memory?" + v.Encode())
	result := &MemoryMetrics{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (pca PrivateCloudApi) GetStorageMetrics(privateCloudId string, args ...interface{}) (*StorageMetrics, error) {
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
	path := pca.getPath("/privateClouds/" + privateCloudId + "/metrics/storage?" + v.Encode())
	result := &StorageMetrics{}
	if err := doRequest(GET, path, result); err != nil {
		return nil, err
	}
	return result, nil
}
