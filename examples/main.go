package main

import (
	"fmt"
	"log"

	"github.com/fcjr/alert"
)

func main() {
	if err := alert.Info("Example Message", "Example Message Text"); err != nil {
		log.Fatal(err)
	}
	if err := alert.Error("Example Error", "Example Error Text"); err != nil {
		log.Fatal(err)
	}
	resp, err := alert.Question("Example Question", "Default or Alternate?", "Default", "Alternate")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %t\n", resp)
}
