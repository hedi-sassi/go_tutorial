package main

import (
    "fmt"
	"log"
	"example.com/greetings" 
	//our own module created before... 
	//but this link won't work if we don't modify the module
)

func main() {

	log.SetPrefix("Error, greetings: ")
	log.SetFlags(0)

	//seed the rand function
	greetings.Realinit()

    // Get a greeting message and print it.
	message, err := greetings.Hello("Bob")
	
	if(err != nil){
		log.Fatal(err)
	}

    fmt.Println(message)
}