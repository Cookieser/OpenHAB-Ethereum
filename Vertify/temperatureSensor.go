package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

// TemperatureData represents the structure for storing temperature sensor data.
type TemperatureData struct {
    DeviceID    int64   `json:"deviceId"`
    Temperature float64 `json:"temperature"`
    Unit        string  `json:"unit"`
    Timestamp   string  `json:"timestamp"`
}

// num is a global counter used to assign a unique ID to each temperature record.
var num int64 = 0

// simulateTemperatureSensor generates a random temperature between -10 and 40 degrees Celsius,
// along with a timestamp, and increments the device ID.
func simulateTemperatureSensor() TemperatureData {
    temperature := -10.0 + rand.Float64()*50.0
    currentTime := time.Now().UTC()
    num++

    return TemperatureData{
        DeviceID:    num,
        Temperature: temperature,
        Unit:        "C",
        Timestamp:   currentTime.Format(time.RFC3339),
    }
}

// sendTemperatureData serializes the temperature data to JSON and sends it to a server.
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

// main initializes the random seed and continuously simulates and sends temperature data every 5 seconds.
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

