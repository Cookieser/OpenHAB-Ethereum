package main

import (

    "myproject/BlockChain"
    "myproject/DBoperation"
    "database/sql"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"


)





func main() {

    db, err := DBoperation.InitDB("root:root@tcp(localhost:3306)/DataCollection")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    
    BlockChain.Init();

    http.HandleFunc("/temperature", makeHandleTemperature(db))
    fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}

func makeHandleTemperature(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {






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

    var data BlockChain.TemperatureData
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
    
    
    record := DBoperation.TemperatureRecord{
    DeviceId:    data.DeviceId,
    Timestamp:   data.Timestamp,
    Temperature: data.Temperature,
}

// Insert the record into the database
err = DBoperation.InsertTemperatureRecord(db, record)
if err != nil {
    log.Printf("Failed to insert temperature record %+v: %v\n", record, err)
    http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
    return
    }
    
    

/*
err = BlockChain.SendHashToContract(data.DeviceId, hashedData)
    if err != nil {
        log.Fatalf("Error sending hash to contract: %v\n", err)
    }

*/


    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Data received and processed"))
}
}


func hashTemperatureData(data BlockChain.TemperatureData) (string, error) {
    formattedTemperature := fmt.Sprintf("%.6f", data.Temperature)
    fmt.Println("formattedTemperature:", data.DeviceId)
        fmt.Println("formattedTemperature:", formattedTemperature)
            fmt.Println("formattedTemperature:", data.Timestamp)
    dataString := fmt.Sprintf("%d:%s:%s", data.DeviceId, formattedTemperature, data.Timestamp)
    hash := sha256.New()
    _, err := hash.Write([]byte(dataString))
    if err != nil {
        return "", err
    }
    hashedData := hex.EncodeToString(hash.Sum(nil))
    return hashedData, nil
}




