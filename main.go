package main

import (
  "encoding/json"
  "fmt"
  zmq "github.com/alecthomas/gozmq"
  "github.com/altonymous/api.6fusion.com/models"
  "github.com/nu7hatch/gouuid"
)

func main() {
  var jsonBlob = []byte(`{
    "uuid": "123",
    "cpu_count": 123,
    "cpu_speed": 123.2,
    "maximum_memory": 123,
    "readings": [
        {
            "reading_at": "12/12/2012T12:12:12Z",
            "cpu_usage": 123,
            "memory_bytes": 123
        },
        {
            "reading_at": "12/12/2012T12:12:12Z",
            "cpu_usage": 123,
            "memory_bytes": 123
        }
    ],
    "disks": [
        {
            "id": 123,
            "uuid": "123",
            "name": "123",
            "maximum_size": 123,
            "kind": "123",
            "thin": 123,
            "disk_readings": [
                {
                    "reading_at": "12/12/2012T12:12:12Z",
                    "usage": 123,
                    "read": 123,
                    "write": 123
                },
                {
                    "reading_at": "12/12/2012T12:12:12Z",
                    "usage": 123,
                    "read": 123,
                    "write": 123
                }
            ]
        }
    ],
    "network_interface_cards": [
        {
            "id": 123,
            "uuid": "123",
            "name": "123",
            "mac_address": "123",
            "ip_address": "123",
            "network_interface_card_readings": [
                {
                    "reading_at": "12/12/2012T12:12:12Z",
                    "receive": 123,
                    "transmit": 123
                },
                {
                    "reading_at": "12/12/2012T12:12:12Z",
                    "receive": 123,
                    "transmit": 123
                }
            ]
        }
    ]
}`)

  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REQ)
  socket.Connect("tcp://127.0.0.1:5000")
  socket.Connect("tcp://127.0.0.1:6000")

  var machine models.Machine
  err := json.Unmarshal(jsonBlob, &machine)
  if err != nil {
    fmt.Println("error: ", err)
  }

  for i := 0; i < 10000; i++ {
    u4, _ := uuid.NewV4()
    machine.UUID = u4.String()

    machineJson, err := json.Marshal(machine)
    if err != nil {
      fmt.Println("error:", err)
    }

    // msg := fmt.Sprintf(, i)
    socket.Send([]byte(machineJson), 0)
    // println("Sending", msg)
    socket.Recv(0)
  }
}
