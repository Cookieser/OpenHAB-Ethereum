
package DBoperation

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

// TemperatureRecord holds the temperature data


type TemperatureRecord struct {
    DeviceId    int64   `json:"deviceId"`
    Temperature float64 `json:"temperature"`
    Timestamp   string  `json:"timestamp"` // This matches the client's format.
}


// initDB initializes and verifies a connection to the database.
// It returns a database connection pool and an error if any.
func InitDB(dataSourceName string) (*sql.DB, error) {
    // Open a database connection.
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        // Handle the error in a way that's appropriate for your application.
        // Here, we're returning the error to the caller.
        return nil, fmt.Errorf("error opening connection: %w", err)
    }

    // Verify the connection with a ping.
    if err := db.Ping(); err != nil {
        // It's important to close the database connection if we're not going to use it,
        // because failing to ping the database successfully means the connection is not usable.
        db.Close()
        return nil, fmt.Errorf("error connecting to the database: %w", err)
    }

    // If everything is okay, return the database connection.
    fmt.Println("Connection to Database -------------------------------- Success!!!")
    return db, nil
}




// insertTemperatureRecord inserts a single temperature record into the database
func InsertTemperatureRecord(db *sql.DB, record TemperatureRecord) error {
	formattedTemperature := fmt.Sprintf("%.6f", record.Temperature)
	   // fmt.Println("formattedTemperature:", formattedTemperature)
	_, err := db.Exec("INSERT INTO Temperature (Device_ID, Timestamp, Value) VALUES (?, ?, ?)",
		record.DeviceId, record.Timestamp, formattedTemperature)
	if err != nil {
		return err
	}
	return nil
}


