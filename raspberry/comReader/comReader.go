package comReader

import (
  "log"
  "strings"
  "time"
  "github.com/tarm/serial"
)

type SensorData struct {
  Humidity    string `json:"humidity"`
  Temperature string `json:"temperature"`
  Light       string `json:"light"`
}

func Read(dataChannel chan SensorData) {
  c := &serial.Config{Name: "COM3", Baud: 9600}
  s, err := serial.OpenPort(c)
  if err != nil {
    log.Fatal(err)
  }
  buf := make([]byte, 128)
  for {
    n, err := s.Read(buf)
    if err != nil {
      log.Fatal(err)
    }
    data := string(buf[:n])
    // remove \r\n from the end of the string
    data = strings.TrimRight(data, "\r\n")
    info := strings.Split(data, ",")
    
    if len(info) == 3 {
      sensorData := SensorData{
        Humidity:    info[0],
        Temperature: info[1],
        Light:       info[2],
      }
      dataChannel <- sensorData
    }

    time.Sleep(1000 * time.Millisecond)
  }
}
