package main

import (
	"fmt"

	"example.com/wakeup"
)

func main() {
	// Get a greeting message and print it.
	message := wakeup.WakeUp("ddd")
	fmt.Println(message)
}
