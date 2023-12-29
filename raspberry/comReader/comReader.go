package comReader

import (
  "log"
  "strings"
  "github.com/tarm/serial"
  "fmt"
  "bufio"
)

type SensorData struct {
  Humidity    string `json:"humidity"`
  Temperature string `json:"temperature"`
  Light       string `json:"light"`
}

func Read(dataChannel chan SensorData) {
  c := &serial.Config{Name: "COM3", Baud: 9600,ReadTimeout: 1, Size:8}
  stream, err := serial.OpenPort(c)
  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(stream)
  for scanner.Scan() {
  data := strings.TrimRight(scanner.Text(), "\r\n")
  info := strings.Split(data, ",")
  
  if len(info) == 3 {
    sensorData := SensorData{
      Humidity:    info[0],
      Temperature: info[1],
      Light:       info[2],
    }
    dataChannel <- sensorData
  }
      fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
 
}