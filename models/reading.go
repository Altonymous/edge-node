package models

type Reading struct {
  ReadingAt   string `json:"reading_at"`
  Interval    int32  `json:"interval"`
  CPUUsage    int32  `json:"cpu_usage"`
  MemoryBytes int32  `json:"memory_bytes"`
}
