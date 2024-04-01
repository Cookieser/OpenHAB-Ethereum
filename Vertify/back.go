package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "net/http"
)

type TemperatureData struct {
    DeviceId    int64   `json:"deviceId"`
    Temperature float64 `json:"temperature"`
    Timestamp   string  `json:"timestamp"`
}

func hashTemperatureData(data TemperatureData) (string, error) {
    dataString := fmt.Sprintf("%d:%f:%s", data.DeviceId, data.Temperature, data.Timestamp)
    hash := sha256.New()
    _, err := hash.Write([]byte(dataString))
    if err != nil {
        return "", err
    }
    hashedData := hex.EncodeToString(hash.Sum(nil))
    return hashedData, nil
}

func hashHandler(w http.ResponseWriter, r *http.Request) {
    // 允许跨域请求
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var tempData TemperatureData
    err := json.NewDecoder(r.Body).Decode(&tempData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedData, err := hashTemperatureData(tempData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(w, hashedData)
}

func main() {
    http.HandleFunc("/hash", hashHandler)
    fmt.Println("Server started at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
