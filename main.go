// main.go

package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

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

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", showIndexPage)
	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

	port, err := findAvailablePort()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Found opened port: %v\n", port)
	router.Run(port)

}
