package main

import (
	"bytes"
	"net/http"
	"github.com/a-h/templ"
	"klog.co/plantz/comReader"
	"fmt"
)

func main() {
	dataChannel := make(chan comReader.SensorData)
	prevHumidity := "0"
	prevTemperature := "0"
	prevLight := "0"
  
  // // Start the Read function as a goroutine
	go comReader.Read(dataChannel)
		
	http.Handle("/", templ.Handler(hello("franquito")))
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.WriteHeader(http.StatusOK)


		for {
			// Receive SensorData from the channel
			sensorData := <-dataChannel
			if sensorData.Humidity == "" {
				continue
			}
			var buf bytes.Buffer
			lineChartUpdate(prevHumidity, sensorData.Humidity, prevTemperature, sensorData.Temperature, prevLight, sensorData.Light).Render(r.Context(), &buf)
			content := buf.String()

			fmt.Fprintf(w, "event: sowitroli\ndata: %s\n\n", content)
			fmt.Fprintf(w, "event: humidityUpdate\ndata: %s\n\n", fmt.Sprintf("%s", sensorData.Humidity))
			fmt.Fprintf(w, "event: temperatureUpdate\ndata: %s\n\n", fmt.Sprintf("%s", sensorData.Temperature))
			fmt.Fprintf(w, "event: lightUpdate\ndata: %s\n\n", fmt.Sprintf("%s", sensorData.Light))
			
			prevHumidity = sensorData.Humidity
			prevTemperature = sensorData.Temperature
			prevLight = sensorData.Light
			w.(http.Flusher).Flush()
		}

	})
	http.ListenAndServe(":8080", nil)
}
