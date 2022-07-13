package leaseweb

import (
	"fmt"
	"net/url"
)

const PRIVATE_CLOUD_API_DEFAULT_VERSION = "v2"

type PrivateCloudApi struct {
	version string
}

func (pca *PrivateCloudApi) SetVersion(version string) {
	pca.version = version
}

func (pca PrivateCloudApi) GetVersion() string {
	if pca.version == "" {
		return PRIVATE_CLOUD_API_DEFAULT_VERSION
	}
	return pca.version
}

func (pca PrivateCloudApi) GetPath() string {
	return "/cloud"
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

type NetworkTraffic struct {
	Type             string `json:"type"`
	TrafficType      string `json:"trafficType"`
	DataTrafficUnit  string `json:"datatrafficUnit"`
	DataTrafficLimit int    `json:"datatrafficLimit"`
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

type TimeValuePair struct {
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
}

type BasicMetric struct {
	Unit   string          `json:"unit"`
	Values []TimeValuePair `json:"values"`
}

type MetricMetadata struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Granularity string `json:"granularity"`
	Aggregation string `json:"aggregation"`
}

type DataTrafficMetrics struct {
	Metric   DataTrafficMetric `json:"metrics"`
	Metadata MetricMetadata    `json:"_metadata"`
}

type DataTrafficMetric struct {
	DataTrafficUp   BasicMetric `json:"DATATRAFFIC_UP"`
	DataTrafficDown BasicMetric `json:"DATATRAFFIC_DOWN"`
}

type BandWidthMetrics struct {
	Metric   BandWidthMetric `json:"metrics"`
	Metadata MetricMetadata  `json:"_metadata"`
}

type BandWidthMetric struct {
	UpPublic   BasicMetric `json:"UP_PUBLIC"`
	DownPublic BasicMetric `json:"DOWN_PUBLIC"`
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

func (pca PrivateCloudApi) ListPrivateClouds(args ...interface{}) (PrivateClouds, error) {

	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	privateClouds := &PrivateClouds{}
	r := leasewebRequest{
		response: privateClouds,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return PrivateClouds{}, err
	}

	return *privateClouds, nil
}

func (pca PrivateCloudApi) GetPrivateCloud(privateCloudId string) (PrivateCloud, error) {

	privateCloud := &PrivateCloud{}
	r := leasewebRequest{
		response: privateCloud,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId,
	}
	err := doRequest(r)
	if err != nil {
		return PrivateCloud{}, err
	}

	return *privateCloud, nil
}

func (pca PrivateCloudApi) ListCredentials(privateCloudId string, credentialType string, args ...int) (Credentials, error) {

	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	credentials := &Credentials{}
	r := leasewebRequest{
		response: credentials,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId + "/credentials/" + credentialType + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return Credentials{}, err
	}

	return *credentials, nil
}

func (pca PrivateCloudApi) GetCredentials(privateCloudId string, credentialType string, username string) (Credential, error) {

	credential := &Credential{}
	r := leasewebRequest{
		response: credential,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId + "/credentials/" + credentialType + "/" + username,
	}
	err := doRequest(r)
	if err != nil {
		return Credential{}, err
	}

	return *credential, nil
}

func (pca PrivateCloudApi) GetDataTrafficMetrics(privateCloudId string, args ...interface{}) (DataTrafficMetrics, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	dataTrafficMetrics := &DataTrafficMetrics{}
	r := leasewebRequest{
		response: dataTrafficMetrics,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId + "/metrics/datatraffic" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return DataTrafficMetrics{}, err
	}

	return *dataTrafficMetrics, nil
}

func (pca PrivateCloudApi) GetBandWidthMetrics(privateCloudId string, args ...interface{}) (BandWidthMetrics, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	bandWidthMetrics := &BandWidthMetrics{}
	r := leasewebRequest{
		response: bandWidthMetrics,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId + "/metrics/bandwidth" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return BandWidthMetrics{}, err
	}

	return *bandWidthMetrics, nil
}

func (pca PrivateCloudApi) GetCpuMetrics(privateCloudId string, args ...interface{}) (CpuMetrics, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	cpuMetrics := &CpuMetrics{}
	r := leasewebRequest{
		response: cpuMetrics,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId + "/metrics/cpu" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return CpuMetrics{}, err
	}

	return *cpuMetrics, nil
}

func (pca PrivateCloudApi) GetMemoryMetrics(privateCloudId string, args ...interface{}) (MemoryMetrics, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	memoryMetrics := &MemoryMetrics{}
	r := leasewebRequest{
		response: memoryMetrics,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId + "/metrics/memory" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return MemoryMetrics{}, err
	}

	return *memoryMetrics, nil
}

func (pca PrivateCloudApi) GetStorageMetrics(privateCloudId string, args ...interface{}) (StorageMetrics, error) {

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

	queryString := ""
	if v.Encode() != "" {
		queryString = "?" + v.Encode()
	}

	storageMetrics := &StorageMetrics{}
	r := leasewebRequest{
		response: storageMetrics,
		method:   GET,
		endpoint: pca.GetPath() + "/" + pca.GetVersion() + "/privateClouds/" + privateCloudId + "/metrics/storage" + queryString,
	}
	err := doRequest(r)
	if err != nil {
		return StorageMetrics{}, err
	}

	return *storageMetrics, nil
}
