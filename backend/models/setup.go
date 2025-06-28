package models

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"time"
)

var DBSQLite *gorm.DB
var DBInflux influxdb2.Client
var DBInfluxWriteAPI api.WriteAPIBlocking

// ConnectDatabase initializes connections to the SQLite and InfluxDB databases and configures required migrations.
func ConnectDatabase() {
	var err error

	time.Sleep(3 * time.Second)

	pathSQLite := os.Getenv("DATABASE_SQLITE_URL")
	if pathSQLite == "" {
		log.Fatal("DATABASE_SQLITE_URL environment variable is not set")
	}

	if err := ensureDatabaseDirectoryExists(pathSQLite); err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	if err := ensureDatabaseFileExists(pathSQLite); err != nil {
		log.Fatalf("Failed to create database file: %v", err)
	}

	DBSQLite, err = gorm.Open(sqlite.Open(pathSQLite), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := autoMigrateModels(); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	if err := createTriggers(); err != nil {
		log.Fatalf("Failed to create triggers: %v", err)
	}

	initInfluxDB()

	log.Println("Database connection initialized successfully")
}

func ensureDatabaseDirectoryExists(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if dir == "." || dir == "" {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("Creating database directory: %s", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func ensureDatabaseFileExists(dbPath string) error {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("Creating database file: %s", dbPath)
		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

// autoMigrateModels performs auto-migration for all models
func autoMigrateModels() error {
	models := []interface{}{
		&User{},
		&Device{},
		&Session{},
		&SessionField{},
		&SessionStat{},
	}

	for _, model := range models {
		if err := DBSQLite.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}

// createTriggers initializes database triggers for automatic updates of device last seen and session end times.
// It creates two triggers: one for updating the last seen timestamp of devices, and one for setting session end times.
// Returns an error if the trigger creation fails.
func createTriggers() error {
	err := DBSQLite.Exec(`
    CREATE TRIGGER IF NOT EXISTS update_device_last_seen 
	    AFTER INSERT ON sessions 
    	FOR EACH ROW
	BEGIN
		UPDATE devices 
		SET last_seen = CURRENT_TIMESTAMP,
			version = NEW.version
		WHERE device_id = NEW.device_id;
	END;
    `).Error
	if err != nil {
		return err
	}

	err = DBSQLite.Exec(`
    CREATE TRIGGER IF NOT EXISTS update_session_end_time
	    AFTER UPDATE OF is_active ON sessions
	    FOR EACH ROW
	    WHEN NEW.is_active = 0 AND OLD.is_active = 1
	BEGIN
	    UPDATE sessions 
	    SET end_time = CURRENT_TIMESTAMP 
	    WHERE id = NEW.id;
	END;
    `).Error
	if err != nil {
		return err
	}

	return nil
}

// initInfluxDB initializes InfluxDB connection
func initInfluxDB() {
	influxURL := os.Getenv("INFLUX_URL")
	influxToken := os.Getenv("INFLUX_TOKEN")
	influxOrg := os.Getenv("INFLUX_ORG")
	influxBucket := os.Getenv("INFLUX_BUCKET")

	if influxURL == "" || influxToken == "" || influxOrg == "" || influxBucket == "" {
		log.Println("Warning: InfluxDB configuration incomplete, skipping InfluxDB initialization")
		return
	}

	log.Printf("Connecting to InfluxDB at: %s", influxURL)
	DBInflux = influxdb2.NewClient(influxURL, influxToken)
	DBInfluxWriteAPI = DBInflux.WriteAPIBlocking(influxOrg, influxBucket)
}
