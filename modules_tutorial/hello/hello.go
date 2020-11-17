package main

import (
    "fmt"

	"example.com/greetings" 
	//our own module created before... 
	//but this link won't work if we don't modify the module
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}