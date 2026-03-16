package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from Go!")
	fmt.Println("Name: moh@shukra")
	fmt.Println("Focus: backend services and clean architecture")
	fmt.Printf("Generated: %s\n", time.Now().UTC().Format(time.RFC3339))
}
