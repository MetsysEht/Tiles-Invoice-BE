package gormDatabase

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

// Config struct to hold the database configuration
type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	Charset  string
}

// Function to build a DSN string from Config struct
func (c *Config) BuildDSN() string {
	password := url.QueryEscape(c.Password) // URL-encode the password
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.Username, password, c.Host, c.Port, c.DBName, c.Charset)
}

func CreateGormDatabase(c *Config) (*gorm.DB, error) {
	dsn := c.BuildDSN()

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, nil
}
