package entity

type HealthCheck struct {
	Failing bool
	MinResponseTime int32
}