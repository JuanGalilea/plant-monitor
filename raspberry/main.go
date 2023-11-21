package main

import (
  "log"
  "time"
  "github.com/tarm/serial"
  "strings"
)

func main() {
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

    info := strings.Split(data, ",")

    if len(info) == 3 {
      log.Println("Humidity: ", info[0])
      log.Println("Temperature: ", info[1])
      log.Println("LDR: ", info[2])
    }      

    time.Sleep(1000 * time.Millisecond)
	}
}