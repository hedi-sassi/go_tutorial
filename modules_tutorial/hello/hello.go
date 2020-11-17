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

	names := []string{"Bob", "Terry", "Barry", "Larry"}

    // Get a greeting message and print it.
	messages, err := greetings.Hellos(names)
	
	if(err != nil){
		log.Fatal(err)
	}

    fmt.Println(messages)
}