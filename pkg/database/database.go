package database

import (
	"fmt"
	"log"
	"os"

	"github.com/charitan-go/project-server/ent"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Client *ent.Client

func connect() error {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	Client, err = ent.Open("postgres", dsn)
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		return err
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		return err
	}

	return err
}

// func migrate() error {
// 	if err := DB.AutoMigrate(&model.Project{}); err != nil {
// 		log.Println("Migrate failed")
// 		return err
// 	}
//
// 	return nil
// }

func seedData() error {
	sqlScript, err := os.ReadFile("resource/data.sql")
	if err != nil {
		log.Printf("Failed to read SQL script: %v", err)
		return err
	}

	// Execute the SQL script
	err = DB.Exec(string(sqlScript)).Error
	if err != nil {
		log.Printf("Failed to seed data: %v", err)
		return err
	}

	log.Println("Data seeded successfully")
	return nil
}

func SetupDatabase() error {
	if err := connect(); err != nil {
		return err
	}

	// if err := migrate(); err != nil {
	// 	return err
	// }

	if err := seedData(); err != nil {
		return err
	}

	return nil
}
