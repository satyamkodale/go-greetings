🚀 Go Greetings Module

This is a simple Go module that provides a function to return a greeting message.

📂 Project Structure
go-greetings/
├── go.mod # Module file (**github.com/satyamkodale/go-greetings**)
├── main.go # Main entry point (**must have package main**)
├── greetings/
│ ├── greetings.go # Greetings package (**library**)

🛠️ Installation & Setup

1️⃣ Clone the Repository

git clone https://github.com/satyamkodale/go-greetings.git
cd go-greetings

2️⃣ Initialize Go Module

Ensure the Go module is properly set up:
go mod tidy

3️⃣ Run the Application

go run main.go

🚀 Usage

✅ Import the Package

To use the greetings package in your Go project, import it:
import "github.com/satyamkodale/go-greetings/greetings"

✅ Example Code

package main

import (
"fmt"
"github.com/satyamkodale/go-greetings/greetings"
)

func main() {
fmt.Println("Hello from main..")

    message := greetings.Greetings("Satyam")
    fmt.Println(message)

}

✅ Expected Output
Hello from main..
Hi, Satyam Welcome!

🌟 Contributing

Feel free to fork this repository and submit pull requests. Contributions are welcome!
