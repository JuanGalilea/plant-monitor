package main

import (
	"net/http"
	"github.com/a-h/templ"
	"klog.co/plantz/comReader"
	"fmt"
)

func main() {
	dataChannel := make(chan comReader.SensorData)

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
			
			fmt.Fprintf(w, "event: humidityUpdate\ndata: %s\n\n", fmt.Sprintf("%s", sensorData.Humidity))
			fmt.Fprintf(w, "event: temperatureUpdate\ndata: %s\n\n", fmt.Sprintf("%s", sensorData.Temperature))
			fmt.Fprintf(w, "event: lightUpdate\ndata: %s\n\n", fmt.Sprintf("%s", sensorData.Light))
			w.(http.Flusher).Flush()
		}

	})
	http.ListenAndServe(":8080", nil)
}
