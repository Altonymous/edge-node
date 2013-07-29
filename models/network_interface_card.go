package models

type NetworkInterfaceCard struct {
  UUID                         string                        `json:"uuid"`
  Name                         string                        `json:"name"`
  MacAddress                   string                        `json:"mac_address"`
  IPAddress                    string                        `json:"ip_address"`
  NetworkInterfaceCardReadings []NetworkInterfaceCardReading `json:"network_interface_card_readings"`
}
