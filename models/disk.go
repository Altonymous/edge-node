package models

type Disk struct {
  UUID         string        `json:"uuid"`
  Name         string        `json:"name"`
  MaximumSize  int32         `json:"maximum_size"`
  Kind         string        `json:"kind"`
  Thin         bool          `json:"uuid"`
  DiskReadings []DiskReading `json:"disk_readings"`
}
