package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/charitan-go/project-server/ent"
)

// var DB *gorm.DB
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
		log.Fatalf("failed opening connection to postgres: %v", err)
		return err
	}
	defer Client.Close()
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return err
	}

	return nil

	//
	// var err error
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err == nil {
	// 	log.Println("Connect to database success")
	// }
	//
	// return err
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
	// sqlScript, err := os.ReadFile("resource/data.sql")
	// if err != nil {
	// 	log.Fatalf("Failed to read SQL script: %v", err)
	// }
	//
	// // Execute the SQL script
	// err = DB.Exec(string(sqlScript)).Error
	// if err != nil {
	// 	log.Fatalf("Failed to seed data: %v", err)
	// 	return err
	// }
	//
	// log.Println("Data seeded successfully")
	// return nil

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
