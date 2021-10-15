package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB for connect the DB
func ConnectDB() {
	var err error
	// p := config.Config("DB_PORT")
	// port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error to connect mysql")
	}

	// Connection URL to connect to Mysql Database
	dsn := "root:@tcp(127.0.0.1:3306)/sebat-go?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc-Local",config.Config("DB_USER"),config.Config("DB_PASSWORD"), port, config.Config("DB_NAME"))

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection Opened to Database")
}