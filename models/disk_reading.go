package models

type DiskReading struct {
  ReadingAt string `json:"reading_at"`
  Usage     int32  `json:"usage"`
  Read      int32  `json:"read"`
  Write     int32  `json:"write"`
}
