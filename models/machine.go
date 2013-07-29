package models

import (
  "encoding/json"
  "log"
)

type Machine struct {
  UUID                  string                 `json:"uuid"`
  CPUCount              int32                  `json:"cpu_count"`
  CPUSpeed              float64                `json:"cpu_speed"`
  MaximumMemory         int32                  `json:"maximum_memory"`
  GuestAgent            bool                   `json:"guest_agent"`
  PowerState            string                 `json:"power_state"`
  VirtualName           string                 `json:"virtual_name"`
  Readings              []Reading              `json:"readings"`
  Disks                 []Disk                 `json:"disks"`
  NetworkInterfaceCards []NetworkInterfaceCard `json:"network_interface_cards"`
}

func (this *Machine) Save() {
  log.Println("Machine.Save")
  conn := pool.Get()
  defer conn.Close()

  jsonMachine, _ := json.Marshal(this)

  err := conn.Send("SET", "machine:"+this.UUID, jsonMachine)
  if err != nil {
    log.Println("[ERROR] main (dbPool): ", err)
    panic(err)
  }
  err = conn.Flush()
  if err != nil {
    log.Println("[ERROR] main (dbPool): ", err)
    panic(err)
  }
}
