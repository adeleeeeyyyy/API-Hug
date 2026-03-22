package tester

import "time"

type Result struct {
	Duration   time.Duration
	StatusCode int
	Error      error
}

type Stats struct {
	Environtment  string        `json:"environtment"`
	URL           string        `json:"url"`
	TotalRequests int           `json:"total_requests"`
	SuccessCount  int           `json:"success_count"`
	ErrorCount    int           `json:"error_count"`
	TotalTime     time.Duration `json:"total_time_ns"`
	MinTime       time.Duration `json:"min_time_ns"`
	MaxTime       time.Duration `json:"max_time_ns"`
	AvgTime       time.Duration `json:"avg_time_ns"`
}

type Comparison struct {
	LocalStats Stats `json:"local_stats"`
	ServerStats Stats `json:"server_stats"`
	DifferenceMs float64 `json:"difference_ms"`
	FasterEnv string `json:"faster_env"`
}