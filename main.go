package main

import (
	"fmt"

	"github.com/satyamkodale/go-greetings/config"
	"github.com/satyamkodale/go-greetings/greetings"
)

func main() {
	fmt.Println("Version : " + config.Version)
	// Call the Greetings function from the greetings package
	message := greetings.Greeting("Satyam")
	fmt.Println(message)
}
