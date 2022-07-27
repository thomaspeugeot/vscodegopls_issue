package main

import (
	"log"

	"github.com/thomaspeugeot/vscodegopls_issue/go/models"
)

// imports

// Main docs
func main() {

	hello := (&models.Hello{Name: "Bonjour"})

	log.Println("Hello world : ")
	models.PrintNameOfValue(hello)
}
