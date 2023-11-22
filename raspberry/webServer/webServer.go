package webServer

import (
  "fmt"
	"github.com/labstack/echo/v4"
	"klog.co/plantz/comReader"
)

func Start() {
	// Create a channel for receiving data
	dataChannel := make(chan comReader.SensorData)

	// Start the Read function as a goroutine
	go comReader.Read(dataChannel)

	e := echo.New()
  e.Static("/", "webServer/static")

	// Configure route for SSE endpoint
	// https://echo.labstack.com/docs/cookbook/streaming-response
	e.GET("/events", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")

		for {
			// Receive SensorData from the channel
			sensorData := <-dataChannel

			sendEvent(c, "humidityUpdate", fmt.Sprintf("%s", sensorData.Humidity))
			sendEvent(c, "temperatureUpdate", fmt.Sprintf("%s", sensorData.Temperature))
			sendEvent(c, "lightUpdate", fmt.Sprintf("%s", sensorData.Light))
		}
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func sendEvent(c echo.Context, eventName string, data string) {
	fmt.Fprintf(c.Response(), "event: %s\ndata: %s\n\n", eventName, data)
	c.Response().Flush()
}