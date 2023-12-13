package main

import (
	"fmt"
	"log"

)

// Get Listener IP
func getIP() {

	promptIP := `Enter the Listener IP: 
--------------------

~: `
	fmt.Println("Enter the Listener IP: ")
	var lhost string
	ip, err := fmt.Scanln(&lhost)
	if err != nil {
		log.Fatal(err)
	}

}


func generateDrop() {

}

func configA1() {

}

func configR1() {

}

func configSecUpdate() {

}

func main() {

}
