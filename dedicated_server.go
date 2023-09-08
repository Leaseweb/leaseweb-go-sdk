package leaseweb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/LeaseWeb/leaseweb-go-sdk/options"
)

const DEDICATED_SERVER_API_VERSION = "v2"

type DedicatedServerApi struct{}

type DedicatedServers struct {
	Servers  []DedicatedServer `json:"servers"`
	Metadata Metadata          `json:"_metadata"`
}

type DedicatedServer struct {
	AssetId             string               `json:"assetId"`
	Contract            Contract             `json:"contract"`
	FeatureAvailability FeatureAvailability  `json:"featureAvailability"`
	Id                  string               `json:"id"`
	Location            Location             `json:"location"`
	NetworkInterfaces   NetworkInterfaces    `json:"networkInterfaces"`
	PowerPorts          []Port               `json:"powerPorts"`
	PrivateNetworks     []PrivateNetwork     `json:"privateNetworks"`
	Rack                DedicatedServerRack  `json:"rack"`
	SerialNumber        string               `json:"serialNumber"`
	Specs               DedicatedServerSpecs `json:"specs"`
}

type DedicatedServerSpecs struct {
	Chassis             string                       `json:"chassis"`
	HardwareRaidCapable bool                         `json:"hardwareRaidCapable"`
	Cpu                 DedicatedServerSpecCpu       `json:"cpu"`
	Ram                 DedicatedServerSpecRam       `json:"ram"`
	Hdd                 []DedicatedServerSpecHdd     `json:"hdd"`
	PciCards            []DedicatedServerSpecPciCard `json:"pciCards"`
}

type DedicatedServerSpecCpu struct {
	Quantity json.Number `json:"quantity"`
	Type     string      `json:"type"`
}

type DedicatedServerSpecRam struct {
	Size json.Number `json:"size"`
	Unit string      `json:"unit"`
}

type DedicatedServerSpecHdd struct {
	Amount          json.Number `json:"amount"`
	Id              string      `json:"id"`
	PerformanceType string      `json:"performanceType"`
	Size            json.Number `json:"size"`
	Type            string      `json:"type"`
	Unit            string      `json:"unit"`
}

type DedicatedServerSpecPciCard struct {
	Description string `json:"description"`
}

type DedicatedServerRack struct {
	Type string `json:"type"`
}

type DedicatedServerHardware struct {
	Id            string                             `json:"id"`
	ParserVersion string                             `json:"parserVersion"`
	Result        DedicatedServerHardwareInformation `json:"result"`
	ScannedAt     string                             `json:"scannedAt"`
	ServerId      string                             `json:"serverId"`
}

type DedicatedServerHardwareInformation struct {
	Chassis  DedicatedServerChassis   `json:"chassis"`
	Cpu      []DedicatedServerCpu     `json:"cpu"`
	Ipmi     DedicatedServerIpmi      `json:"ipmi"`
	Disks    []DedicatedServerDisks   `json:"disks"`
	Memories []DedicatedServerMemory  `json:"memory"`
	Networks []DedicatedServerNetwork `json:"network"`
}

type DedicatedServerChassis struct {
	Description string                            `json:"description"`
	Firmware    DedicatedServerChassisFirmware    `json:"firmware"`
	Motherboard DedicatedServerChassisMotherboard `json:"motherboard"`
	Product     string                            `json:"product"`
	Serial      string                            `json:"serial"`
	Vendor      string                            `json:"vendor"`
}

type DedicatedServerChassisFirmware struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Vendor      string `json:"vendor"`
	Version     string `json:"version"`
}

type DedicatedServerChassisMotherboard struct {
	Product string `json:"product"`
	Serial  string `json:"serial"`
	Vendor  string `json:"vendor"`
}

type DedicatedServerCpu struct {
	Capabilities DedicatedServerCpuCapabilities `json:"capabilities"`
	Description  string                         `json:"description"`
	HZ           string                         `json:"hz"`
	SerialNumber string                         `json:"serial_number"`
	Settings     DedicatedServerCpuSettings     `json:"settings"`
	Slot         string                         `json:"slot"`
	Vendor       string                         `json:"vendor"`
}

type DedicatedServerCpuSettings struct {
	Cores        string `json:"cores"`
	EnabledCores string `json:"enabledcores"`
	Threads      string `json:"threads"`
}

type DedicatedServerCpuCapabilities struct {
	CpuFreq string `json:"cpufreq"`
	HT      string `json:"ht"`
	VMX     bool   `json:"vmx"`
	X8664   string `json:"x86-64"`
}

type DedicatedServerIpmi struct {
	Defgateway string `json:"defgateway"`
	Firmware   string `json:"firmware"`
	IpAddress  string `json:"ipaddress"`
	IpSource   string `json:"ipsource"`
	MacAddress string `json:"macaddress"`
	SubnetMask string `json:"subnetmask"`
	Vendor     string `json:"vendor"`
}

type DedicatedServerMemory struct {
	ClockHZ      string `json:"clock_hz"`
	Description  string `json:"description"`
	Id           string `json:"id"`
	SerialNumber string `json:"serial_number"`
	SizeBytes    string `json:"size_bytes"`
}

type DedicatedServerDisks struct {
	Description  string                       `json:"description"`
	Id           string                       `json:"id"`
	Product      string                       `json:"product"`
	SerialNumber string                       `json:"serial_number"`
	Size         string                       `json:"size"`
	SmartCTL     DedicatedServerDisksSmartCTL `json:"smartctl"`
	Vendor       string                       `json:"vendor"`
}

type DedicatedServerDisksSmartCTL struct {
	ATAVersion      string                                 `json:"ata_version"`
	Attributes      DedicatedServerDisksSmartCTLAttributes `json:"attributes"`
	DeviceModel     string                                 `json:"device_model"`
	ExecutionStatus string                                 `json:"execution_status"`
	FirmwareVersion string                                 `json:"firmware_version"`
	ISSAS           bool                                   `json:"is_sas"`
	OverallHealth   string                                 `json:"overall_health"`
	RPM             string                                 `json:"rpm"`
	SATAVersion     string                                 `json:"sata_version"`
	SectorSize      string                                 `json:"sector_size"`
	SerialNumber    string                                 `json:"serial_number"`
	SmartErrorLog   string                                 `json:"smart_error_log"`
	SmartSupport    DedicatedServerDisksSmartCTLSupport    `json:"smart_support"`
	SmartctlVersion string                                 `json:"smartctl_version"`
	UserCapacity    string                                 `json:"user_capacity"`
}

type DedicatedServerDisksSmartCTLSupport struct {
	Available bool `json:"available"`
	Enabled   bool `json:"enabled"`
}

type DedicatedServerDisksSmartCTLAttributes struct {
	PowerOnHours        DedicatedServerDisksSmartCTLAttributesType `json:"Power_On_Hours"`
	ReallocatedSectorCT DedicatedServerDisksSmartCTLAttributesType `json:"Reallocated_Sector_Ct"`
}

type DedicatedServerDisksSmartCTLAttributesType struct {
	Flag       string `json:"flag"`
	Id         string `json:"id"`
	RawValue   string `json:"raw_value"`
	Thresh     string `json:"thresh"`
	Type       string `json:"type"`
	Updated    string `json:"updated"`
	Value      string `json:"value"`
	WhenFailed string `json:"when_failed"`
	Worst      string `json:"worst"`
}

type DedicatedServerNetwork struct {
	Capabilities DedicatedServerNetworkCapabilities `json:"capabilities"`
	LLDP         DedicatedServerNetworkLLDP         `json:"lldp"`
	LogicalName  string                             `json:"logical_name"`
	MacAddress   string                             `json:"mac_address"`
	Product      string                             `json:"product"`
	Settings     DedicatedServerNetworkSettings     `json:"settings"`
	Vendor       string                             `json:"vendor"`
}

type DedicatedServerNetworkSettings struct {
	AutoNegotiation string `json:"autonegotiation"`
	Broadcast       string `json:"broadcast"`
	Driver          string `json:"driver"`
	DriverVersion   string `json:"driverversion"`
	Duplex          string `json:"duplex"`
	Firmware        string `json:"firmware"`
	Ip              string `json:"ip"`
	Latency         string `json:"latency"`
	Link            string `json:"link"`
	Multicast       string `json:"multicast"`
	Port            string `json:"port"`
	Speed           string `json:"speed"`
}

type DedicatedServerNetworkLLDP struct {
	Chassis DedicatedServerNetworkLLDPChassis `json:"chassis"`
	Port    DedicatedServerNetworkPort        `json:"port"`
	Vlan    DedicatedServerNetworkLLDPVlan    `json:"vlan"`
}

type DedicatedServerNetworkLLDPVlan struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Name  string `json:"name"`
}

type DedicatedServerNetworkLLDPChassis struct {
	Description string `json:"description"`
	MacAddress  string `json:"mac_address"`
	Name        string `json:"name"`
}

type DedicatedServerNetworkPort struct {
	Description     string                                    `json:"description"`
	AutoNegotiation DedicatedServerNetworkPortAutoNegotiation `json:"auto_negotiation"`
}

type DedicatedServerNetworkPortAutoNegotiation struct {
	Enabled   string `json:"enabled"`
	Supported string `json:"supported"`
}

type DedicatedServerNetworkCapabilities struct {
	AutoNegotiation string                                       `json:"autonegotiation"`
	BusMaster       string                                       `json:"bus_master"`
	CapList         string                                       `json:"cap_list"`
	Ethernet        string                                       `json:"ethernet"`
	LinkSpeeds      DedicatedServerNetworkCapabilitiesLinkSpeeds `json:"link_speeds"`
	MSI             string                                       `json:"msi"`
	MSIX            string                                       `json:"msix"`
	PciExpress      string                                       `json:"pciexpress"`
	Physical        string                                       `json:"physical"`
	PM              string                                       `json:"pm"`
	TP              string                                       `json:"tp"`
}

type DedicatedServerNetworkCapabilitiesLinkSpeeds struct {
	BTFD1000 string `json:"1000bt-fd"`
	BT100    string `json:"100bt"`
	BTFD100  string `json:"100bt-fd"`
	BT10     string `json:"10bt"`
	BTFD10   string `json:"10bt-fd"`
}

type DedicatedServerNetworkInterfaces struct {
	NetworkInterfaces []DedicatedServerNetworkInterface `json:"networkInterfaces"`
	Metadata          Metadata                          `json:"_metadata"`
}

type DedicatedServerNetworkInterface struct {
	LinkSpeed       string `json:"linkSpeed"`
	OperStatus      string `json:"operStatus"`
	Status          string `json:"status"`
	SwitchInterface string `json:"switchInterface"`
	SwitchName      string `json:"switchName"`
	Type            string `json:"type"`
}

type DedicatedServerDhcpReservations struct {
	Leases   []DedicatedServerDhcpReservation `json:"leases"`
	Metadata Metadata                         `json:"_metadata"`
}

type DedicatedServerDhcpReservation struct {
	BootFile          string                                          `json:"bootfile"`
	CreatedAt         string                                          `json:"createdAt"`
	Gateway           string                                          `json:"gateway"`
	Hostname          string                                          `json:"hostname"`
	Ip                string                                          `json:"ip"`
	LastClientRequest DedicatedServerDhcpReservationLastClientRequest `json:"lastClientRequest"`
	Mac               string                                          `json:"mac"`
	Netmask           string                                          `json:"netmask"`
	Site              string                                          `json:"site"`
	UpdatedAt         string                                          `json:"updatedAt"`
}

type DedicatedServerDhcpReservationLastClientRequest struct {
	RelayAgent string `json:"relayAgent"`
	Type       string `json:"type"`
	UserAgent  string `json:"userAgent"`
}

type DedicatedServerJobs struct {
	Jobs     []DedicatedServerJob `json:"jobs"`
	Metadata Metadata             `json:"_metadata"`
}

type DedicatedServerJob struct {
	CreatedAt string                     `json:"createdAt"`
	Flow      string                     `json:"flow"`
	IsRunning bool                       `json:"isRunning"`
	Node      string                     `json:"node"`
	Payload   DedicatedServerJobPayload  `json:"payload"`
	Progress  DedicatedServerJobProgress `json:"progress"`
	ServerId  string                     `json:"serverId"`
	Status    string                     `json:"status"`
	Tasks     []DedicatedServerJobTask   `json:"tasks"`
	Type      string                     `json:"type"`
	UpdatedAt string                     `json:"updatedAt"`
	Uuid      string                     `json:"uuid"`
	Metadata  struct {
		BatchId string `json:"BATCH_ID"`
	} `json:"metadata"`
}

type DedicatedServerJobPayload map[string]interface{}

type DedicatedServerJobProgress struct {
	Canceled   json.Number `json:"canceled"`
	Expired    json.Number `json:"expired"`
	Failed     json.Number `json:"failed"`
	Finished   json.Number `json:"finished"`
	InProgress json.Number `json:"inprogress"`
	Pending    json.Number `json:"pending"`
	Percentage json.Number `json:"percentage"`
	Total      json.Number `json:"total"`
	Waiting    json.Number `json:"waiting"`
}

type DedicatedServerJobTask struct {
	Description      string                                `json:"description"`
	Message          string                                `json:"errorMessage"`
	Flow             string                                `json:"flow"`
	OnError          string                                `json:"onError"`
	Status           string                                `json:"status"`
	StatusTimestamps DedicatedServerJobTaskStatusTimeStamp `json:"statusTimestamps"`
	Uuid             string                                `json:"uuid"`
}

type DedicatedServerJobTaskStatusTimeStamp struct {
	Canceled string `json:"CANCELED"`
	Pending  string `json:"PENDING"`
	Waiting  string `json:"WAITING"`
}

type DedicatedServerOperatingSystems struct {
	Metadata         Metadata                         `json:"_metadata"`
	OperatingSystems []DedicatedServerOperatingSystem `json:"operatingSystems"`
}

type DedicatedServerOperatingSystem struct {
	Id                   string                                 `json:"id"`
	Name                 string                                 `json:"name"`
	Architecture         string                                 `json:"architecture"`
	Configurable         bool                                   `json:"configurable"`
	Family               string                                 `json:"family"`
	Type                 string                                 `json:"type"`
	Version              string                                 `json:"version"`
	Features             []string                               `json:"features"`
	SupportedBootDevices []string                               `json:"supportedBootDevices"`
	SupportedFileSystems []string                               `json:"supportedFileSystems"`
	Defaults             DedicatedServerOperatingSystemDefaults `json:"defaults"`
}

type DedicatedServerOperatingSystemDefaults struct {
	Device     string                                    `json:"device"`
	Partitions []DedicatedServerOperatingSystemPartition `json:"partitions"`
}

type DedicatedServerOperatingSystemPartition struct {
	Bootable   bool   `json:"bootable"`
	Filesystem string `json:"filesystem"`
	Mountpoint string `json:"mountpoint"`
	Primary    bool   `json:"primary"`
	Size       string `json:"size"`
}

type DedicatedServerControlPanels struct {
	Metadata      Metadata                      `json:"_metadata"`
	ControlPanels []DedicatedServerControlPanel `json:"controlPanels"`
}

type DedicatedServerControlPanel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DedicatedServerRescueImages struct {
	Metadata     Metadata                     `json:"_metadata"`
	RescueImages []DedicatedServerRescueImage `json:"rescueImages"`
}

type DedicatedServerRescueImage struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DedicatedServerListOptions struct {
	PaginationOptions
	IP                    *string `param:"ip"`
	MacAddress            *string `param:"macAddress"`
	Site                  *string `param:"site"`
	PrivateRackID         *string `param:"privateRackId"`
	Reference             *string `param:"reference"`
	PrivateNetworkCapable *bool   `param:"privateNetworkCapable"`
	PrivateNetworkEnabled *bool   `param:"privateNetworkEnabled"`
}

type DedicatedServerListIpsOptions struct {
	PaginationOptions
	NetworkType *string  `param:"networkType"`
	Version     *string  `param:"version"`
	NullRouted  *string  `param:"nullRouted"`
	IPs         []string `param:"ips"`
}

type DedicatedServerListJobOptions struct {
	PaginationOptions
	Type      *string `param:"type"`
	Status    *string `param:"status"`
	IsRunning *bool   `param:"isRunning"`
}

type DedicatedServerListOperatingSystemsOptions struct {
	PaginationOptions
	ControlPanelId *string `param:"controlPanelId"`
}

type DedicatedServerListControlPanelsOptions struct {
	PaginationOptions
	OperatingSystemId *string `param:"operatingSystemId"`
}

func (dsa DedicatedServerApi) getPath(endpoint string) string {
	return "/bareMetals/" + DEDICATED_SERVER_API_VERSION + endpoint
}

func (dsa DedicatedServerApi) List(ctx context.Context, opts DedicatedServerListOptions) (*DedicatedServers, error) {
	path := dsa.getPath("/servers")
	query := options.Encode(opts)
	result := &DedicatedServers{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

// AllWithOpts will return all dedicated servers and allows filtering options be specified as an argument
func (dsa DedicatedServerApi) AllWithOpts(ctx context.Context, opts DedicatedServerListOptions) ([]DedicatedServer, error) {
	var allServers []DedicatedServer

	// handle default pagination if Offset or Limit are undefined
	if opts.PaginationOptions.Offset == nil || opts.PaginationOptions.Limit == nil {
		opts.PaginationOptions = PaginationOptions{Offset: Int(0), Limit: Int(20)}
	}

	for {
		// List all dedicated servers with defined opts
		result, err := DedicatedServerApi{}.List(ctx, opts)
		if err != nil {
			return nil, err
		}

		// Break if no more servers are found
		if len(result.Servers) == 0 {
			break
		}

		// Append servers to list
		allServers = append(allServers, result.Servers...)
		// Increment Offset with the value of Limit
		*opts.PaginationOptions.Offset += *opts.PaginationOptions.Limit
	}

	return allServers, nil
}

// All will return all dedicated servers without optional filters specified
func (dsa DedicatedServerApi) All(ctx context.Context) ([]DedicatedServer, error) {
	return dsa.AllWithOpts(ctx, DedicatedServerListOptions{})
}

func (dsa DedicatedServerApi) Get(ctx context.Context, serverId string) (*DedicatedServer, error) {
	path := dsa.getPath("/servers/" + serverId)
	result := &DedicatedServer{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) Update(ctx context.Context, serverId string, payload map[string]interface{}) error {
	path := dsa.getPath("/servers/" + serverId)
	return doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (dsa DedicatedServerApi) GetHardwareInformation(ctx context.Context, serverId string) (*DedicatedServerHardware, error) {
	path := dsa.getPath("/servers/" + serverId + "/hardwareInfo")
	result := &DedicatedServerHardware{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListIps(ctx context.Context, serverId string, opts DedicatedServerListIpsOptions) (*Ips, error) {
	path := dsa.getPath("/servers/" + serverId + "/ips")
	query := options.Encode(opts)
	result := &Ips{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) GetIp(ctx context.Context, serverId, ip string) (*Ip, error) {
	path := dsa.getPath("/servers/" + serverId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) UpdateIp(ctx context.Context, serverId, ip string, payload map[string]string) (*Ip, error) {
	path := dsa.getPath("/servers/" + serverId + "/ips/" + ip)
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) NullRouteAnIp(ctx context.Context, serverId, ip string) (*Ip, error) {
	path := dsa.getPath("/servers/" + serverId + "/ips/" + ip + "/null")
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) RemoveNullRouteAnIp(ctx context.Context, serverId, ip string) (*Ip, error) {
	path := dsa.getPath("/servers/" + serverId + "/ips/" + ip + "/unnull")
	result := &Ip{}
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListNullRoutes(ctx context.Context, serverId string, opts PaginationOptions) (*NullRoutes, error) {
	path := dsa.getPath("/servers/" + serverId + "/nullRouteHistory")
	query := options.Encode(opts)
	result := &NullRoutes{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListNetworkInterfaces(ctx context.Context, serverId string, opts PaginationOptions) (*DedicatedServerNetworkInterfaces, error) {
	path := dsa.getPath("/servers/" + serverId + "/networkInterfaces")
	result := &DedicatedServerNetworkInterfaces{}
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) CloseAllNetworkInterfaces(ctx context.Context, serverId string) error {
	path := dsa.getPath("/servers/" + serverId + "/networkInterfaces/close")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dsa DedicatedServerApi) OpenAllNetworkInterfaces(ctx context.Context, serverId string) error {
	path := dsa.getPath("/servers/" + serverId + "/networkInterfaces/open")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dsa DedicatedServerApi) GetNetworkInterface(ctx context.Context, serverId, networkType string) (*DedicatedServerNetworkInterface, error) {
	path := dsa.getPath("/servers/" + serverId + "/networkInterfaces/" + networkType)
	result := &DedicatedServerNetworkInterface{}
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) CloseNetworkInterface(ctx context.Context, serverId, networkType string) error {
	path := dsa.getPath("/servers/" + serverId + "/networkInterfaces/" + networkType + "/close")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dsa DedicatedServerApi) OpenNetworkInterface(ctx context.Context, serverId, networkType string) error {
	path := dsa.getPath("/servers/" + serverId + "/networkInterfaces/" + networkType + "/open")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dsa DedicatedServerApi) DeleteServerFromPrivateNetwork(ctx context.Context, serverId, privateNetworkId string) error {
	path := dsa.getPath("/servers/" + serverId + "/privateNetworks/" + privateNetworkId)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dsa DedicatedServerApi) AddServerToPrivateNetwork(ctx context.Context, serverId, privateNetworkId string, linkSpeed int) error {
	payload := map[string]int{"linkSpeed": linkSpeed}
	path := dsa.getPath("/servers/" + serverId + "/privateNetworks/" + privateNetworkId)
	return doRequest(ctx, http.MethodPut, path, "", nil, payload)
}

func (dsa DedicatedServerApi) DeleteDhcpReservation(ctx context.Context, serverId string) error {
	path := dsa.getPath("/servers/" + serverId + "/leases")
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dsa DedicatedServerApi) ListDhcpReservation(ctx context.Context, serverId string, opts PaginationOptions) (*DedicatedServerDhcpReservations, error) {
	path := dsa.getPath("/servers/" + serverId + "/leases")
	result := &DedicatedServerDhcpReservations{}
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) CreateDhcpReservation(ctx context.Context, serverId string, payload map[string]string) error {
	path := dsa.getPath("/servers/" + serverId + "/leases")
	return doRequest(ctx, http.MethodPost, path, "", nil, payload)
}

func (dsa DedicatedServerApi) CancelActiveJob(ctx context.Context, serverId string) (*DedicatedServerJob, error) {
	result := &DedicatedServerJob{}
	path := dsa.getPath("/servers/" + serverId + "/cancelActiveJob")
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ExpireActiveJob(ctx context.Context, serverId string) (*DedicatedServerJob, error) {
	result := &DedicatedServerJob{}
	path := dsa.getPath("/servers/" + serverId + "/expireActiveJob")
	if err := doRequest(ctx, http.MethodPost, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) LaunchHardwareScan(ctx context.Context, serverId string, payload map[string]interface{}) (*DedicatedServerJob, error) {
	result := &DedicatedServerJob{}
	path := dsa.getPath("/servers/" + serverId + "/hardwareScan")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) LaunchInstallation(ctx context.Context, serverId string, payload map[string]interface{}) (*DedicatedServerJob, error) {
	result := &DedicatedServerJob{}
	path := dsa.getPath("/servers/" + serverId + "/install")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) LaunchIpmiRest(ctx context.Context, serverId string, payload map[string]interface{}) (*DedicatedServerJob, error) {
	result := &DedicatedServerJob{}
	path := dsa.getPath("/servers/" + serverId + "/ipmiRest")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListJobs(ctx context.Context, serverId string, opts DedicatedServerListJobOptions) (*DedicatedServerJobs, error) {
	result := &DedicatedServerJobs{}
	path := dsa.getPath("/servers/" + serverId + "/jobs")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) GetJob(ctx context.Context, serverId, jobId string) (*DedicatedServerJob, error) {
	result := &DedicatedServerJob{}
	path := dsa.getPath("/servers/" + serverId + "/jobs/" + jobId)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) LaunchRescueMode(ctx context.Context, serverId string, payload map[string]interface{}) (*DedicatedServerJob, error) {
	result := &DedicatedServerJob{}
	path := dsa.getPath("/servers/" + serverId + "/rescueMode")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListCredentials(ctx context.Context, serverId string, opts PaginationOptions) (*Credentials, error) {
	result := &Credentials{}
	path := dsa.getPath("/servers/" + serverId + "/credentials")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) CreateCredential(ctx context.Context, serverId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := make(map[string]string)
	payload["type"] = credentialType
	payload["username"] = username
	payload["password"] = password
	path := dsa.getPath("/servers/" + serverId + "/credentials")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListCredentialsByType(ctx context.Context, serverId, credentialType string, opts PaginationOptions) (*Credentials, error) {
	result := &Credentials{}
	path := dsa.getPath("/servers/" + serverId + "/credentials/" + credentialType)
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) GetCredential(ctx context.Context, serverId, credentialType, username string) (*Credential, error) {
	result := &Credential{}
	path := dsa.getPath("/servers/" + serverId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) DeleteCredential(ctx context.Context, serverId, credentialType, username string) error {
	path := dsa.getPath("/servers/" + serverId + "/credentials/" + credentialType + "/" + username)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dsa DedicatedServerApi) UpdateCredential(ctx context.Context, serverId, credentialType, username, password string) (*Credential, error) {
	result := &Credential{}
	payload := map[string]string{"password": password}
	path := dsa.getPath("/servers/" + serverId + "/credentials/" + credentialType + "/" + username)
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) GetDataTrafficMetrics(ctx context.Context, serverId string, opts MetricsOptions) (*DataTrafficMetricsV1, error) {
	path := dsa.getPath("/servers/" + serverId + "/metrics/datatraffic")
	query := options.Encode(opts)
	result := &DataTrafficMetricsV1{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) GetBandWidthMetrics(ctx context.Context, serverId string, opts MetricsOptions) (*BandWidthMetrics, error) {
	path := dsa.getPath("/servers/" + serverId + "/metrics/bandwidth")
	query := options.Encode(opts)
	result := &BandWidthMetrics{}
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListBandWidthNotificationSettings(ctx context.Context, serverId string, opts PaginationOptions) (*BandWidthNotificationSettings, error) {
	result := &BandWidthNotificationSettings{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/bandwidth")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) CreateBandWidthNotificationSetting(ctx context.Context, serverId, frequency string, threshold float64, unit string) (*NotificationSetting, error) {

	payload := map[string]string{
		"frequency": frequency,
		"threshold": fmt.Sprintf("%g", threshold),
		"unit":      unit,
	}
	result := &NotificationSetting{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/bandwidth")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) DeleteBandWidthNotificationSetting(ctx context.Context, serverId, notificationId string) error {
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/bandwidth/" + notificationId)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dsa DedicatedServerApi) GetBandWidthNotificationSetting(ctx context.Context, serverId, notificationId string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/bandwidth/" + notificationId)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) UpdateBandWidthNotificationSetting(ctx context.Context, serverId, notificationSettingId string, payload map[string]string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/bandwidth/" + notificationSettingId)
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListDataTrafficNotificationSettings(ctx context.Context, serverId string, opts PaginationOptions) (*DataTrafficNotificationSettings, error) {
	result := &DataTrafficNotificationSettings{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/datatraffic")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) CreateDataTrafficNotificationSetting(ctx context.Context, serverId, frequency string, threshold float64, unit string) (*NotificationSetting, error) {
	payload := map[string]string{
		"frequency": frequency,
		"threshold": fmt.Sprintf("%g", threshold),
		"unit":      unit,
	}
	result := &NotificationSetting{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/datatraffic")
	if err := doRequest(ctx, http.MethodPost, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) DeleteDataTrafficNotificationSetting(ctx context.Context, serverId, notificationId string) error {
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/datatraffic/" + notificationId)
	return doRequest(ctx, http.MethodDelete, path, "")
}

func (dsa DedicatedServerApi) GetDataTrafficNotificationSetting(ctx context.Context, serverId, notificationId string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/datatraffic/" + notificationId)
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) UpdateDataTrafficNotificationSetting(ctx context.Context, serverId, notificationSettingId string, payload map[string]string) (*NotificationSetting, error) {
	result := &NotificationSetting{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/datatraffic/" + notificationSettingId)
	if err := doRequest(ctx, http.MethodPut, path, "", result, payload); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) GetDdosNotificationSetting(ctx context.Context, serverId string) (*DdosNotificationSetting, error) {
	result := &DdosNotificationSetting{}
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/ddos")
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) UpdateDdosNotificationSetting(ctx context.Context, serverId string, payload map[string]string) error {
	path := dsa.getPath("/servers/" + serverId + "/notificationSettings/ddos/")
	if err := doRequest(ctx, http.MethodPut, path, "", nil, payload); err != nil {
		return err
	}
	return nil
}

func (dsa DedicatedServerApi) PowerCycleServer(ctx context.Context, serverId string) error {
	path := dsa.getPath("/servers/" + serverId + "/powerCycle")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dsa DedicatedServerApi) GetPowerStatus(ctx context.Context, serverId string) (*PowerStatus, error) {
	result := &PowerStatus{}
	path := dsa.getPath("/servers/" + serverId + "/powerInfo")
	if err := doRequest(ctx, http.MethodGet, path, "", result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) PowerOffServer(ctx context.Context, serverId string) error {
	path := dsa.getPath("/servers/" + serverId + "/powerOff")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dsa DedicatedServerApi) PowerOnServer(ctx context.Context, serverId string) error {
	path := dsa.getPath("/servers/" + serverId + "/powerOn")
	return doRequest(ctx, http.MethodPost, path, "")
}

func (dsa DedicatedServerApi) ListOperatingSystems(ctx context.Context, opts DedicatedServerListOperatingSystemsOptions) (*DedicatedServerOperatingSystems, error) {
	result := &DedicatedServerOperatingSystems{}
	path := dsa.getPath("/operatingSystems")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) GetOperatingSystem(ctx context.Context, operatingSystemId, controlPanelId string) (*DedicatedServerOperatingSystem, error) {
	v := url.Values{}
	v.Add("controlPanelId", fmt.Sprint(controlPanelId))
	result := &DedicatedServerOperatingSystem{}
	path := dsa.getPath("/operatingSystems/" + operatingSystemId)
	query := v.Encode()
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListControlPanels(ctx context.Context, opts DedicatedServerListControlPanelsOptions) (*DedicatedServerControlPanels, error) {
	result := &DedicatedServerControlPanels{}
	path := dsa.getPath("/controlPanels")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}

func (dsa DedicatedServerApi) ListRescueImages(ctx context.Context, opts PaginationOptions) (*DedicatedServerRescueImages, error) {
	result := &DedicatedServerRescueImages{}
	path := dsa.getPath("/rescueImages")
	query := options.Encode(opts)
	if err := doRequest(ctx, http.MethodGet, path, query, result); err != nil {
		return result, err
	}
	return result, nil
}
