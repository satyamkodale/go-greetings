package main

import (
	"fmt"

	"github.com/satyamkodale/go-greetings/config"
	"github.com/satyamkodale/go-greetings/greetings"
)

func main() {
	// Call the Greetings function from the greetings package
	fmt.Println("Version : " + config.Version)
	message := greetings.Greeting("Satyam")
	fmt.Println(message)
}
