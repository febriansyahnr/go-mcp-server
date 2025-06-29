package config

type CBSetting struct {
	MaxRequests  uint32  `json:"maxRequests"`
	Interval     int64   `json:"interval"`
	Timeout      int64   `json:"timeout"`
	RequestCount uint32  `json:"requestCount"`
	FailureRatio float64 `json:"failureRatio"`
}
