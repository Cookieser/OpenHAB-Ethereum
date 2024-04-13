package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "net/http"
)

// TemperatureData defines the structure for incoming temperature data
type TemperatureData struct {
    DeviceId    int64   `json:"deviceId"`
    Temperature float64 `json:"temperature"`
    Timestamp   string  `json:"timestamp"`
}

// hashTemperatureData creates a SHA-256 hash of the temperature data
func hashTemperatureData(data TemperatureData) (string, error) {
    // Create a data string from the temperature data
    dataString := fmt.Sprintf("%d:%f:%s", data.DeviceId, data.Temperature, data.Timestamp)
    
    // Initialize a new SHA-256 hash generator
    hash := sha256.New()
    
    // Write the data string as bytes to the hash generator
    if _, err := hash.Write([]byte(dataString)); err != nil {
        return "", err
    }
    
    // Generate the hash and encode it as a hexadecimal string
    hashedData := hex.EncodeToString(hash.Sum(nil))
    return hashedData, nil
}

// hashHandler handles the HTTP requests to hash temperature data
func hashHandler(w http.ResponseWriter, r *http.Request) {
    // Setup CORS headers to allow all origins, with POST method and Content-Type header
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Only allow POST method, reject all other methods
    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // Decode JSON from the request body into a TemperatureData struct
    var tempData TemperatureData
    if err := json.NewDecoder(r.Body).Decode(&tempData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Hash the temperature data
    hashedData, err := hashTemperatureData(tempData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the hashed data as a response
    fmt.Fprint(w, hashedData)
}

func main() {
    // Set up the HTTP server route and handler
    http.HandleFunc("/hash", hashHandler)
    
    // Start the HTTP server
    fmt.Println("Server started at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}

