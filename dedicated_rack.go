package leaseweb

import (
	"context"
	"net/http"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const DEDICATED_RACK_API_VERSION = "v2"

type DedicatedRackApi struct{}

type DedicatedRacks struct {
	PrivateRacks []DedicatedRack `json:"privateRacks"`
	Metadata     Metadata        `json:"_metadata"`
}

type DedicatedRack struct {
	Contract            Contract            `json:"contract"`
	FeatureAvailability FeatureAvailability `json:"featureAvailability"`
	Id                  string              `json:"id"`
	Location            Location            `json:"location"`
	NetworkInterfaces   NetworkInterfaces   `json:"networkInterfaces"`
	PowerPorts          []Port              `json:"powerPorts"`
	Units               []DedicatedRackUnit `json:"units"`
}

type DedicatedRackUnit struct {
	Unit           string   `json:"unit"`
	Status         string   `json:"status"`
	ConnectedUnits []string `json:"connectedUnits"`
}

type DedicatedRackListOptions struct {
	PaginationOptions
	Reference             *string `param:"reference"`
	PrivateNetworkCapable *bool   `param:"privateNetworkCapable"`
	PrivateNetworkEnabled *bool   `param:"privateNetworkEnabled"`
}

type DedicatedRackListIpsOptions struct {
	PaginationOptions
	NetworkType *string  `param:"networkType"`
	Version     *string  `param:"version"`
	NullRouted  *string  `param:"nullRouted"`
	IPs         []string `param:"ips"`
}

func (dra DedicatedRackApi) getPath(endpoint string) string {
	return "/bareMetals/" + DEDICATED_RACK_API_VERSION + endpoint
}

func (dra DedicatedRackApi) List(ctx context.Context, opts DedicatedRackListOptions) (*DedicatedRacks, error) {
	path := dra.getPath("/privateRacks")
	query := options.Encode(opts)
	result := &DedicatedRacks{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) Get(ctx context.Context, privateRackId string) (*DedicatedRack, error) {
	path := dra.getPath("/privateRacks/" + privateRackId)
	result := &DedicatedRack{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) Update(ctx context.Context, privateRackId, reference string) error {
	payload := map[string]string{"reference": reference}
	path := dra.getPath("/privateRacks/" + privateRackId)
	return doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (dra DedicatedRackApi) ListNullRoutes(ctx context.Context, privateRackId string, opts PaginationOptions) (*NullRoutes, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/nullRouteHistory")
	query := options.Encode(opts)
	result := &NullRoutes{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListIps(ctx context.Context, privateRackId string, opts DedicatedRackListIpsOptions) (*Ips, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips")
	query := options.Encode(opts)
	result := &Ips{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetIp(ctx context.Context, privateRackId, ip string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateIp(ctx context.Context, privateRackId, ip string, payload map[string]string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) NullRouteAnIp(ctx context.Context, privateRackId, ip string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip + "/null")
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) RemoveNullRouteAnIp(ctx context.Context, privateRackId, ip string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip + "/unnull")
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListCredentials(ctx context.Context, privateRackId string, opts PaginationOptions) (*Credentials, error) {
	result := &Credentials{}
	query := options.Encode(opts)
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials")
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) CreateCredential(ctx context.Context, privateRackId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := make(map[string]string)
	payload["type"] = credentialType
	payload["username"] = username
	payload["password"] = password
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListCredentialsByType(ctx context.Context, privateRackId, credentialType string, opts PaginationOptions) (*Credentials, error) {
	result := &Credentials{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials/" + credentialType)
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetCredential(ctx context.Context, privateRackId, credentialType, username string) (*Credential, error) {
	result := &Credential{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) DeleteCredential(ctx context.Context, privateRackId, credentialType, username string) error {
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials/" + credentialType + "/" + username)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dra DedicatedRackApi) UpdateCredential(ctx context.Context, privateRackId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := map[string]string{"password": password}
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetDataTrafficMetrics(ctx context.Context, privateRackId string, opts MetricsOptions) (*DataTrafficMetricsV1, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/metrics/datatraffic")
	query := options.Encode(opts)
	result := &DataTrafficMetricsV1{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetBandWidthMetrics(ctx context.Context, privateRackId string, opts MetricsOptions) (*BandWidthMetrics, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/metrics/bandwidth")
	query := options.Encode(opts)
	result := &BandWidthMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetDdosNotificationSetting(ctx context.Context, privateRackId string) (*DdosNotificationSetting, error) {
	result := &DdosNotificationSetting{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/ddos")
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateDdosNotificationSetting(ctx context.Context, privateRackId string, payload map[string]string) error {
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/ddos")
	if err := doRequest(ctx, http.MethodPut, path, "", nil, payload); err != nil {
		return err
	}
	return nil
}

func (dra DedicatedRackApi) ListBandWidthNotificationSettings(ctx context.Context, privateRackId string, opts PaginationOptions) (*BandWidthNotificationSettings, error) {
	result := &BandWidthNotificationSettings{}
	query := options.Encode(opts)
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/bandwidth")
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) CreateBandWidthNotificationSetting(ctx context.Context, privateRackId, frequency, threshold, unit string) (*NotificationSetting, error) {
	payload := map[string]string{
		"frequency": frequency,
		"threshold": threshold,
		"unit":      unit,
	}
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/bandwidth")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) DeleteBandWidthNotificationSetting(ctx context.Context, privateRackId, notificationId string) error {
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/bandwidth/" + notificationId)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dra DedicatedRackApi) GetBandWidthNotificationSetting(ctx context.Context, privateRackId, notificationId string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/bandwidth/" + notificationId)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateBandWidthNotificationSetting(ctx context.Context, privateRackId, notificationSettingId string, payload map[string]string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/bandwidth/" + notificationSettingId)
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListDataTrafficNotificationSettings(ctx context.Context, privateRackId string, opts PaginationOptions) (*DataTrafficNotificationSettings, error) {
	result := &DataTrafficNotificationSettings{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/datatraffic")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) CreateDataTrafficNotificationSetting(ctx context.Context, privateRackId, frequency, threshold, unit string) (*NotificationSetting, error) {
	payload := map[string]string{
		"frequency": frequency,
		"threshold": threshold,
		"unit":      unit,
	}
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/datatraffic")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) DeleteDataTrafficNotificationSetting(ctx context.Context, privateRackId, notificationId string) error {
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/datatraffic/" + notificationId)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dra DedicatedRackApi) GetDataTrafficNotificationSetting(ctx context.Context, privateRackId, notificationId string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/datatraffic/" + notificationId)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateDataTrafficNotificationSetting(ctx context.Context, privateRackId, notificationSettingId string, payload map[string]string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/notificationSettings/datatraffic/" + notificationSettingId)
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}
