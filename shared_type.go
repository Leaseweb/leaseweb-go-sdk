package leaseweb

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

type BandWidthMetrics struct {
	Metric   BandWidthMetric `json:"metrics"`
	Metadata MetricMetadata  `json:"_metadata"`
}

type BandWidthMetric struct {
	UpPublic   BasicMetric `json:"UP_PUBLIC"`
	DownPublic BasicMetric `json:"DOWN_PUBLIC"`
}

type NetworkTraffic struct {
	Type             string `json:"type"`
	TrafficType      string `json:"trafficType"`
	DataTrafficUnit  string `json:"datatrafficUnit"`
	DataTrafficLimit int    `json:"datatrafficLimit"`
}
