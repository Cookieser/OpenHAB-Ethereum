package DBoperation

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

// TemperatureRecord holds the temperature data with annotations for JSON serialization.
type TemperatureRecord struct {
	DeviceId    int64   `json:"deviceId"`    // Unique identifier for the device
	Temperature float64 `json:"temperature"` // Recorded temperature value
	Timestamp   string  `json:"timestamp"`   // Timestamp of the record in client's format
}

// InitDB initializes and verifies a connection to the database.
// It returns a database connection pool (*sql.DB) and an error if any occurs during initialization.
func InitDB(dataSourceName string) (*sql.DB, error) {
	// Open a database connection using the MySQL driver.
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %w", err)
	}

	// Verify the database connection with a ping to ensure it is active.
	if err := db.Ping(); err != nil {
		db.Close() // Close the database connection if the ping fails
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Print a success message to console when connection is established.
	fmt.Println("Connection to Database successful")
	return db, nil
}

// InsertTemperatureRecord inserts a single temperature record into the database.
// It takes a database connection pool and a TemperatureRecord as arguments.
// Returns an error if the insertion fails.
func InsertTemperatureRecord(db *sql.DB, record TemperatureRecord) error {
	// Format temperature to six decimal places for consistency in database.
	formattedTemperature := fmt.Sprintf("%.6f", record.Temperature)

	// Execute the SQL command to insert the record into the Temperature table.
	_, err := db.Exec("INSERT INTO Temperature (Device_ID, Timestamp, Value) VALUES (?, ?, ?)",
		record.DeviceId, record.Timestamp, formattedTemperature)
	if err != nil {
		return fmt.Errorf("error inserting temperature record: %w", err)
	}
	return nil
}
