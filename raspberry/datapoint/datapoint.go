package datapoint

import (
	"klog.co/plantz/comReader"
	"strconv"
	// "fmt"
)

func normalizeHumidity(humidity string) string {
	humidityFloat, err := strconv.ParseFloat(humidity, 64)
	if err != nil {
		// fmt.Println(err)
		return "1"
	}
	return strconv.FormatFloat(humidityFloat/70, 'f', 4, 64)
}

func normalizeTemperature(temperature string) string {
	temperatureFloat, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		// fmt.Println(err)
		return "1"
	}
	return strconv.FormatFloat(temperatureFloat/51, 'f', 4, 64)
}

func normalizeLight(light string) string {
	lightFloat, err := strconv.ParseFloat(light, 64)
	if err != nil {
		// fmt.Println(err)
		return "1"
	}
	return strconv.FormatFloat(lightFloat/1024, 'f', 3, 64)
}


func RenderDataPoint(actual comReader.SensorData, previous comReader.SensorData) string {
	previousHumidity := normalizeHumidity(previous.Humidity)
	previousTemperature := normalizeTemperature(previous.Temperature)
	previousLight := normalizeLight(previous.Light)
	
	currentHumidity := normalizeHumidity(actual.Humidity)
	currentTemperature := normalizeTemperature(actual.Temperature)	
	currentLight := normalizeLight(actual.Light)

	// print(previousLight,"-",currentLight,"\n",previousHumidity,"-",currentHumidity,"\n",previousTemperature,"-",currentTemperature,"\n")
	return "<tr><td style='--start: "+ previousHumidity +"; --end: "+ currentHumidity +"'><span class='data'> " + currentHumidity + "</span></td><td style='--start: "+previousTemperature+"; --end: "+ currentTemperature+"'><span class='data'> "+currentTemperature+" </span></td><td style='--start: "+previousLight+"; --end: "+currentLight+"'><span class='data'> "+currentLight+" </span></td></tr>"
}