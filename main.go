// main.go

package main

import (
	"errors"
	"fmt"
	"net"
	"semaphore_gin_labs/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func findAvailablePort() (string, error) {
	// Define the port range you want to check
	startPort := 8080
	endPort := 8300

	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			_ = listener.Close()
			return address, nil
		}
	}

	return "", errors.New("No available port in the range")
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Article{})
	// Create
	db.Create(&models.Article{ID: 1, Title: "Article 1", Content: "Article 1 body"})
	db.Create(&models.Article{ID: 2, Title: "Article 2", Content: "Article 2 body"})
	db.Create(&models.Article{ID: 3, Title: "Article 3", Content: "Article 3 body"})
	db.Create(&models.Article{ID: 4, Title: "Article 4", Content: "Article 4 body"})

	CreateUrlMappings()

	port, err := findAvailablePort()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found opened port: %v\n", port)
	Router.Run(port)

}
