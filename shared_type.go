package leaseweb

type Credentials struct {
	Credentials []Credential `json:"credentials"`
	Metadata    Metadata     `json:"_metadata"`
}

type Credential struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

type Ips struct {
	Ips      []Ip     `json:"ips"`
	Metadata Metadata `json:"_metadata"`
}

type Ip struct {
	FloatingIp       bool               `json:"floatingIp"`
	Gateway          string             `json:"gateway"`
	Ip               string             `json:"ip"`
	MainIp           bool               `json:"mainIp"`
	NetworkType      string             `json:"networkType"`
	NullRouted       bool               `json:"nullRouted"`
	ReverseLookup    string             `json:"reverseLookup"`
	Version          int                `json:"version"`
	Type             string             `json:"type"`
	PrefixLength     int                `json:"prefixLength"`
	Primary          bool               `json:"primary"`
	UnnullingAllowed bool               `json:"unnullingAllowed"`
	EquipmentId      string             `json:"equipmentId"`
	DDOS             IpDdos             `json:"ddos"`
	AssignedContract IpAssignedContract `json:"assignedContract"`
	Subnet           IpSubnet           `json:"subnet"`
}

type IpDdos struct {
	DetectionProfile string `json:"detectionProfile"`
	ProtectionType   string `json:"protectionType"`
}

type IpSubnet struct {
	Id           string `json:"id"`
	NetworkIp    string `json:"networkIp"`
	PrefixLength int    `json:"prefixLength"`
	Gateway      string `json:"gateway"`
}

type IpAssignedContract struct {
	Id string `json:"id"`
}

type NetworkInterfaces struct {
	Internal         NetworkInterface `json:"internal"`
	Public           NetworkInterface `json:"public"`
	RemoteManagement NetworkInterface `json:"remoteManagement"`
}

type NetworkInterface struct {
	Gateway string `json:"gateway"`
	Ip      string `json:"ip"`
	Mac     string `json:"mac"`
	Ports   []Port `json:"ports"`
}

type NullRoutes struct {
	NullRoutes []NullRoute `json:"nullroutes"`
	Metadata   Metadata    `json:"_metadata"`
}

type NullRoute struct {
	Id                   string             `json:"id"`
	Ip                   string             `json:"ip"`
	NulledAt             string             `json:"nulledAt"`
	NulledBy             string             `json:"nulledBy"`
	NullLevel            int                `json:"nullLevel"`
	AutomatedUnnullingAt string             `json:"automatedUnnullingAt"`
	UnnulledAt           string             `json:"unnulledAt"`
	UnnulledBy           string             `json:"unnulledBy"`
	TicketId             string             `json:"ticketId"`
	Comment              string             `json:"comment"`
	EquipmentId          string             `json:"equipmentId"`
	AssignedContract     IpAssignedContract `json:"assignedContract"`
}

type BandWidthMetrics struct {
	Metric   BandWidthMetric `json:"metrics"`
	Metadata MetricMetadata  `json:"_metadata"`
}

type BandWidthMetric struct {
	UpPublic   BasicMetric `json:"UP_PUBLIC"`
	DownPublic BasicMetric `json:"DOWN_PUBLIC"`
}

type DataTrafficMetricsV1 struct {
	Metric struct {
		UpPublic   BasicMetric `json:"UP_PUBLIC"`
		DownPublic BasicMetric `json:"DOWN_PUBLIC"`
	} `json:"metrics"`
	Metadata MetricMetadata `json:"_metadata"`
}

type DataTrafficMetricsV2 struct {
	Metric struct {
		UpPublic   BasicMetric `json:"DATATRAFFIC_UP"`
		DownPublic BasicMetric `json:"DATATRAFFIC_DOWN"`
	} `json:"metrics"`
	Metadata MetricMetadata `json:"_metadata"`
}

type Metadata struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"totalCount"`
}

type TimestampValuePair struct {
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
}

type BasicMetric struct {
	Unit   string               `json:"unit"`
	Values []TimestampValuePair `json:"values"`
}

type MetricMetadata struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Granularity string `json:"granularity"`
	Aggregation string `json:"aggregation"`
}

type NetworkTraffic struct {
	Type             string `json:"type"`
	TrafficType      string `json:"trafficType"`
	DataTrafficUnit  string `json:"datatrafficUnit"`
	DataTrafficLimit int    `json:"datatrafficLimit"`
}

type FeatureAvailability struct {
	Automation       bool `json:"automation"`
	IpmiReboot       bool `json:"ipmiReboot"`
	PowerCycle       bool `json:"powerCycle"`
	PrivateNetwork   bool `json:"privateNetwork"`
	RemoteManagement bool `json:"remoteManagement"`
}

type Location struct {
	Rack  string `json:"rack"`
	Site  string `json:"site"`
	Suite string `json:"suite"`
	Unit  string `json:"unit"`
}

type Port struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

type Contract struct {
	BillingCycle      int               `json:"billingCycle"`
	BillingFrequency  string            `json:"billingFrequency"`
	ContractTerm      int               `json:"contractTerm"`
	Currency          string            `json:"currency"`
	EndsAt            string            `json:"endsAt"`
	StartsAt          string            `json:"startsAt"`
	CustomerId        string            `json:"customerId"`
	DeliveryStatus    string            `json:"deliveryStatus"`
	Id                string            `json:"id"`
	Reference         string            `json:"reference"`
	SalesOrgId        string            `json:"salesOrgId"`
	NetworkTraffic    NetworkTraffic    `json:"networkTraffic"`
	PricePerFrequency float32           `json:"pricePerFrequency"`
	PrivateNetworks   []PrivateNetwork  `json:"privateNetworks"`
	Sla               string            `json:"sla"`
	SoftwareLicenses  []SoftwareLicense `json:"softwareLicenses"`
}

type PrivateNetwork struct {
	Id             string   `json:"id"`
	LinkSpeed      int      `json:"linkSpeed"`
	Status         string   `json:"status"`
	Subnet         string   `json:"subnet"`
	VLanId         string   `json:"vlanId"`
	EquipmentCount int      `json:"equipmentCount"`
	Name           string   `json:"name"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
	Servers        []string `json:"servers"`
}

type SoftwareLicense struct {
	Currency string  `json:"currency"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
}

type DdosNotificationSetting struct {
	Nulling   string `json:"nulling"`
	Scrubbing string `json:"scrubbing"`
}

type BandWidthNotificationSettings struct {
	Settings []NotificationSetting `json:"bandwidthNotificationSettings"`
	Metadata Metadata              `json:"_metadata"`
}

type DataTrafficNotificationSettings struct {
	Settings []NotificationSetting `json:"datatrafficNotificationSettings"`
	Metadata Metadata              `json:"_metadata"`
}

type NotificationSetting struct {
	Actions             []NotificationSettingAction `json:"actions"`
	Frequency           string                      `json:"frequency"`
	Id                  string                      `json:"id"`
	LastCheckedAt       string                      `json:"lastCheckedAt"`
	Threshold           string                      `json:"threshold"`
	ThresholdExceededAt string                      `json:"thresholdExceededAt"`
	Unit                string                      `json:"unit"`
}

type NotificationSettingAction struct {
	LastTriggeredAt string `json:"lastTriggeredAt"`
	Type            string `json:"type"`
}

type Hardware struct {
	Cpu struct {
		Cores int `json:"cores"`
	} `json:"cpu"`

	Memory struct {
		Unit   string `json:"unit"`
		Amount int    `json:"amount"`
	} `json:"memory"`

	Storage struct {
		Unit   string `json:"unit"`
		Amount int    `json:"amount"`
	} `json:"storage"`
}
