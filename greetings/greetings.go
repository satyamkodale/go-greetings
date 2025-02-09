package greetings

import (
	"fmt"

	"rsc.io/quote"
)

func Greeting(name string) string {
	// the := operator is a shortcut for declaring and initializing a variable in one line
	// in Go, every declared variable must be used, or the compiler will throw an error.
	// var message1 string
	// message1 = fmt.Sprintf("Hi, %v. Welcome!", name)
	// println(message1)
	message := fmt.Sprintf("Hi, %v! Welcome! ", name)
	randomMessage := fmt.Sprintf(quote.Go())
	finalMessage := message + "\n" + randomMessage
	return finalMessage
}

// func main() {
// 	fmt.Println(Greetings("Satyam"))
// }

// PS C:\Users\Satyam\OneDrive\Desktop\Goooooooo\go-greetings> go run greetings.go
// package command-line-arguments is not a main package
// PS C:\Users\Satyam\OneDrive\Desktop\Goooooooo\go-greetings>

// -> SOLUTION
// Since greetings.go is part of package greetings, you cannot run it directly using go run greetings.go. Instead, you need a main.go file that imports and calls greetings.Greetings().
