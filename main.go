package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not set in this environment")
	}

	fmt.Println("Port:", portString)
}
