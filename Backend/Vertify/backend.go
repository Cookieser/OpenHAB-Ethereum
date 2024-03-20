package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "net/http"
)


func hashTemperatureData(deviceId int, temperature float64, timestamp string) string {
    dataString := fmt.Sprintf("%d:%.6f:%s", deviceId, temperature, timestamp)
    hash := sha256.New()
    hash.Write([]byte(dataString))
    hashedData := hex.EncodeToString(hash.Sum(nil))
    return hashedData
}


func handleRequest(w http.ResponseWriter, r *http.Request) {
    
    r.ParseForm() 
    deviceId := 2 
    temperature := 29.806522 
    timestamp := "2024-03-20T02:34:47Z" 
    
    
    formattedTemperature := fmt.Sprintf("%.6f", temperature)

    dataString := fmt.Sprintf("%d:%s:%s", deviceId, formattedTemperature, timestamp)
    hash := sha256.New()
    _, err := hash.Write([]byte(dataString))
    if err != nil {
        return 
    }
    hashedData := hex.EncodeToString(hash.Sum(nil))
    

    

 
    fmt.Fprintf(w, "<html><body><p>Hash: %s</p></body></html>", hashedData)
}

func main() {
    http.HandleFunc("/", handleRequest) 
    http.ListenAndServe(":8080", nil) 
}
