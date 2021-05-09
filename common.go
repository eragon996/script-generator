package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput(key string, def string) string {

	// reader
	reader := bufio.NewReader(os.Stdin)

	// Promt for the value
	fmt.Printf("%s[%s]: ", key, def)
	tmp, err := reader.ReadString('\n')
	if err != nil {
		// retry if this failed
		fmt.Println("Error in parsing input. Try again")
		getInput(key, def)
	}
	if tmp == "\n" {
		return def
	}
	// returing with last new line removed
	return tmp[:len(tmp)-1]
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}
