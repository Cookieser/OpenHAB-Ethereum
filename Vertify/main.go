package main

import (
    "crypto/sha256"
    "database/sql"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "myproject/BlockChain"
    "myproject/DBoperation"
    "github.com/spf13/viper"
)

func main() {
    // Initialize configuration using Viper
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("..")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    // Database connection string setup
    username := viper.GetString("database.username")
    password := viper.GetString("database.password")
    host := viper.GetString("database.host")
    port := viper.GetInt("database.port")
    dbname := viper.GetString("database.dbname")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        username, password, host, port, dbname)

    // Initialize database connection
    db, err := DBoperation.InitDB(dsn)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()

    // Initialize Blockchain connection
    BlockChain.Init()

    // Set up HTTP server
    http.HandleFunc("/temperature", makeHandleTemperature(db))
    fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

// makeHandleTemperature creates an HTTP handler function for temperature data processing
func makeHandleTemperature(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
            return
        }

        // Read and parse request body
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error reading request body", http.StatusInternalServerError)
            return
        }
        defer r.Body.Close()

        var data BlockChain.TemperatureData
        if err := json.Unmarshal(body, &data); err != nil {
            http.Error(w, "Error parsing request body", http.StatusBadRequest)
            return
        }

        // Process temperature data
        hashedData, err := hashTemperatureData(data)
        if err != nil {
            log.Printf("Error hashing data: %v", err)
            http.Error(w, "Error hashing data", http.StatusInternalServerError)
            return
        }

        fmt.Printf("Received data: %+v\n", data)
        fmt.Println("Hashed Data:", hashedData)

        // Database record insertion
        record := DBoperation.TemperatureRecord{
            DeviceId:    data.DeviceId,
            Timestamp:   data.Timestamp,
            Temperature: data.Temperature,
        }
        if err := DBoperation.InsertTemperatureRecord(db, record); err != nil {
            log.Printf("Failed to insert temperature record %+v: %v", record, err)
            http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
            return
        }

        // Send hash to blockchain contract
        if err := BlockChain.SendHashToContract(data.DeviceId, hashedData); err != nil {
            log.Fatalf("Error sending hash to contract: %v", err)
            // Note: Using log.Fatalf here will exit the program; consider how to handle errors gracefully
        }

        // Send successful response
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Data received and processed"))
    }
}

// hashTemperatureData hashes the temperature data using SHA-256
func hashTemperatureData(data BlockChain.TemperatureData) (string, error) {
    formattedTemperature := fmt.Sprintf("%.6f", data.Temperature)
    dataString := fmt.Sprintf("%d:%s:%s", data.DeviceId, formattedTemperature, data.Timestamp)
    hash := sha256.New()
    if _, err := hash.Write([]byte(dataString)); err != nil {
        return "", err
    }
    hashedData := hex.EncodeToString(hash.Sum(nil))
    return hashedData, nil
}

