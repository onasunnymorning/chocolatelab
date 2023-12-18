package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/onasunnymorning/chocolatelab/app/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// @title Chocolate Lab API
// @version 0.1.0
// @license.name Geoffrey De Prins all rights reserved
func main() {
	godotenv.Load()
	db := SetupDB()
	h := api.NewBaseHandler(db)

	r := gin.Default()
	api.SetupRoutes(r, h)

	r.Run(":" + os.Getenv("API_PORT"))

}

func SetupDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	gormDB, err := gorm.Open(postgres.Open("postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName))
	if err != nil {
		errMsg := err.Error()
		// If the database does not exist, create it and retry
		if errMsg == fmt.Sprintf("failed to connect to `host=%s user=postgres database=%s`: server error (FATAL: database \"%s\" does not exist (SQLSTATE 3D000))", dbHost, dbName, dbName) {
			log.Println("Database does not exist. Creating...")
			CreateDB(dbUser, dbPass, dbHost, dbPort, dbName)
			gormDB, err := gorm.Open(postgres.Open("postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName))
			if err != nil {
				log.Fatal(err)
			}
			return gormDB
		} else {
			// Otherwise, log the error and exit
			log.Fatal(err)
		}
	}

	return gormDB
}

func CreateDB(dbUser, dbPass, dbHost, dbPort, dbName string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err = db.Exec(createDatabaseCommand)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database created: %s", dbName)
}
