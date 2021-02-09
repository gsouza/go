package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("GG")
	if err != nil {
		log.Fatal(err)
	}
	// Get a greeting message and print it.
	// message := greetings.Hello("Guilherme")
	fmt.Println(message)
}
