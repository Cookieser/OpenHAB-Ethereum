package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)


type TemperatureData struct {
	DeviceId    int64  `json:"deviceId"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Timestamp   string  `json:"timestamp"`
}


var num int64= 0;

func simulateTemperatureSensor() TemperatureData {

	temperature := -10.0 + rand.Float64()*(50.0) 
	currentTime := time.Now().UTC()
	num += 1

	return TemperatureData{
		DeviceId:    num,
		Temperature: temperature,
		Unit:        "C",
		Timestamp:   currentTime.Format(time.RFC3339),
	}
}


func sendTemperatureData(data TemperatureData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/temperature", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Data sent to server, response status:", resp.Status)
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) 


	for {
	
		temperatureData := simulateTemperatureSensor()

	
		err := sendTemperatureData(temperatureData)
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}

		
		time.Sleep(5 * time.Second)
	}
}
