package leaseweb

import (
	"fmt"
	"net/http"
	"net/url"
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

func (dra DedicatedRackApi) getPath(endpoint string) string {
	return "/bareMetals/" + DEDICATED_RACK_API_VERSION + endpoint
}

func (dra DedicatedRackApi) List(args ...interface{}) (*DedicatedRacks, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("privateNetworkCapable", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("privateNetworkEnabled", fmt.Sprint(args[3]))
	}

	path := dra.getPath("/privateRacks?" + v.Encode())
	result := &DedicatedRacks{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) Get(privateRackId string) (*DedicatedRack, error) {
	path := dra.getPath("/privateRacks/" + privateRackId)
	result := &DedicatedRack{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) Update(privateRackId, reference string) error {
	payload := map[string]string{"reference": reference}
	path := dra.getPath("/privateRacks/" + privateRackId)
	return doRequest(http.MethodPut, path, nil, payload)
}

func (dra DedicatedRackApi) ListNullRoutes(privateRackId string, args ...interface{}) (*NullRoutes, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	path := dra.getPath("/privateRacks/" + privateRackId + "/nullRouteHistory?" + v.Encode())
	result := &NullRoutes{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListIps(privateRackId string, args ...interface{}) (*Ips, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}
	if len(args) >= 3 {
		v.Add("networkType", fmt.Sprint(args[2]))
	}
	if len(args) >= 4 {
		v.Add("version", fmt.Sprint(args[3]))
	}
	if len(args) >= 5 {
		v.Add("nullRouted", fmt.Sprint(args[4]))
	}
	if len(args) >= 6 {
		v.Add("ips", fmt.Sprint(args[5]))
	}

	path := dra.getPath("/privateRacks/" + privateRackId + "/ips" + v.Encode())
	result := &Ips{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetIp(privateRackId, ip string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateIp(privateRackId, ip string, payload map[string]string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) NullRouteAnIp(privateRackId, ip string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip + "/null")
	result := &Ip{}
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) RemoveNullRouteAnIp(privateRackId, ip string) (*Ip, error) {
	path := dra.getPath("/privateRacks/" + privateRackId + "/ips/" + ip + "/unnull")
	result := &Ip{}
	if err := doRequest(http.MethodPost, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListCredentials(privateRackId string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &Credentials{}
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials?" + v.Encode())
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) CreateCredential(privateRackId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := make(map[string]string)
	payload["type"] = credentialType
	payload["username"] = username
	payload["password"] = password
	path := dra.getPath("/privateRacks/" + privateRackId + "/credentials")
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListCredentialsByType(serverId, credentialType string, args ...int) (*Credentials, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &Credentials{}
	path := dra.getPath("/privateRacks/" + serverId + "/credentials/" + credentialType + "?" + v.Encode())
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetCredential(serverId, credentialType, username string) (*Credential, error) {
	result := &Credential{}
	path := dra.getPath("/privateRacks/" + serverId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) DeleteCredential(serverId, credentialType, username string) error {
	path := dra.getPath("/privateRacks/" + serverId + "/credentials/" + credentialType + "/" + username)
	return doRequest(http.MethodDelete, path)
}

func (dra DedicatedRackApi) UpdateCredential(serverId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := map[string]string{"password": password}
	path := dra.getPath("/privateRacks/" + serverId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetDataTrafficMetrics(serverId string, args ...interface{}) (*DataTrafficMetricsV1, error) {
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

	path := dra.getPath("/privateRacks/" + serverId + "/metrics/datatraffic?" + v.Encode())
	result := &DataTrafficMetricsV1{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetBandWidthMetrics(serverId string, args ...interface{}) (*BandWidthMetrics, error) {
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

	path := dra.getPath("/privateRacks/" + serverId + "/metrics/bandwidth?" + v.Encode())
	result := &BandWidthMetrics{}
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dra DedicatedRackApi) GetDdosNotificationSetting(serverId string) (*DdosNotificationSetting, error) {
	result := &DdosNotificationSetting{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/ddos")
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateDdosNotificationSetting(serverId string, payload map[string]string) error {
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/ddos")
	if err := doRequest(http.MethodPut, path, nil, payload); err != nil {
		return err
	}
	return nil
}

func (dra DedicatedRackApi) ListBandWidthNotificationSettings(serverId string, args ...int) (*BandWidthNotificationSettings, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &BandWidthNotificationSettings{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/bandwidth?" + v.Encode())
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) CreateBandWidthNotificationSetting(serverId, frequency, threshold, unit string) (*NotificationSetting, error) {
	payload := map[string]string{
		"frequency": frequency,
		"threshold": threshold,
		"unit":      unit,
	}
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/bandwidth")
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) DeleteBandWidthNotificationSetting(serverId, notificationId string) error {
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/bandwidth/" + notificationId)
	return doRequest(http.MethodDelete, path)
}

func (dra DedicatedRackApi) GetBandWidthNotificationSetting(serverId, notificationId string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/bandwidth/" + notificationId)
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateBandWidthNotificationSetting(serverId, notificationSettingId string, payload map[string]string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/bandwidth/" + notificationSettingId)
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) ListDataTrafficNotificationSettings(serverId string, args ...int) (*DataTrafficNotificationSettings, error) {
	v := url.Values{}
	if len(args) >= 1 {
		v.Add("offset", fmt.Sprint(args[0]))
	}
	if len(args) >= 2 {
		v.Add("limit", fmt.Sprint(args[1]))
	}

	result := &DataTrafficNotificationSettings{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/datatraffic?" + v.Encode())
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) CreateDataTrafficNotificationSetting(serverId, frequency, threshold, unit string) (*NotificationSetting, error) {
	payload := map[string]string{
		"frequency": frequency,
		"threshold": threshold,
		"unit":      unit,
	}
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/datatraffic")
	if err := doRequest(http.MethodPost, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) DeleteDataTrafficNotificationSetting(serverId, notificationId string) error {
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/datatraffic/" + notificationId)
	return doRequest(http.MethodDelete, path)
}

func (dra DedicatedRackApi) GetDataTrafficNotificationSetting(serverId, notificationId string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/datatraffic/" + notificationId)
	if err := doRequest(http.MethodGet, path, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dra DedicatedRackApi) UpdateDataTrafficNotificationSetting(serverId, notificationSettingId string, payload map[string]string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dra.getPath("/privateRacks/" + serverId + "/notificationSettings/datatraffic/" + notificationSettingId)
	if err := doRequest(http.MethodPut, path, result, payload); err != nil {
		return result, err
	}
	return result, nil
}
