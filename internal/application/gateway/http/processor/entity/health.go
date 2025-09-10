package entity

type HealthCheck struct {
	Failling        bool  `json:"failling"`
	MinResponseTime int32 `json:minResponseTime`
}
