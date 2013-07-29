package models

type NetworkInterfaceCardReading struct {
  ReadingAt string `json:"reading_at"`
  Receive   int32  `json:"receive"`
  Transmit  int32  `json:"transmit"`
}
