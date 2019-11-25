package main

import (
	"fmt"
	"os"
)

// GetGreeting returns a greeting addressed to the given name­.
func GetGreeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	name := "Undefined"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(GetGreeting(name))
}
