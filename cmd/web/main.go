package main

import (
	"os"

	"Sharykhin/rent-car/api/web"
)

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "3000"
	}

	web.Start(serverPort)
}
