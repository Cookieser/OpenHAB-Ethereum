
package DBoperation

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

// TemperatureRecord holds the temperature data
type TemperatureRecord struct {
	DeviceID  int
	Timestamp int64
	Value     int
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
    fmt.Println("Successfully connected to the database!")
    return db, nil
}











// insertTemperatureRecord inserts a single temperature record into the database
func InsertTemperatureRecord(db *sql.DB, record TemperatureRecord) error {
	_, err := db.Exec("INSERT INTO Temperature (Device_ID, Timestamp, Value) VALUES (?, ?, ?)",
		record.DeviceID, record.Timestamp, record.Value)
	if err != nil {
		return err
	}
	return nil
}


