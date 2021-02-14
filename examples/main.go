package main

import (
	"fmt"
	"log"

	"github.com/fcjr/alert"
)

func main() {
	if err := alert.Message("Example Message", "Example Message Text"); err != nil {
		log.Fatal(err)
	}
	if err := alert.Error("Example Error", "Example Error Text"); err != nil {
		log.Fatal(err)
	}
	resp, err := alert.Question("Example Question", "Yes or No?", "Yes", "No")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %t\n", resp)
}
