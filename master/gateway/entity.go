package gateway

type ComputeResponse struct {
	DownstreamService string  `json:"downstreamService"`
	Sum               float64 `json:"sum"`
	ExecutionTimeMs   int64   `json:"execution_time_ms"`
}
