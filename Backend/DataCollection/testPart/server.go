package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/sha256"
	"encoding/hex"
)


type TemperatureData struct {
	DeviceId    int64  `json:"deviceId"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Timestamp   string  `json:"timestamp"`
}


func handleTemperature(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()


	var data TemperatureData
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}


	hashedData, err := hashTemperatureData(data)
	if err != nil {
		fmt.Println("Error hashing data:", err)
		http.Error(w, "Error hashing data", http.StatusInternalServerError)
		return
	}


	fmt.Printf("Received data: %+v\n", data)
	fmt.Println("Hashed Data:", hashedData)


	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received and hashed"))
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

func main() {
	http.HandleFunc("/temperature", handleTemperature)


	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}