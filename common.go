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

func safeWriteToFile(data string, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// file does not exist hence create one
		// Only allow owner to execute the script 0700
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
		checkError(err, fmt.Sprintf("Uanble to open file %s", path))
		defer f.Close()
		f.WriteString(data)
	} else {
		// file exist
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Files already exist in path %s. Do you want to continue [y/n]", path)
		tmp, err := reader.ReadString('\n')
		checkError(err, "Error reading the input")
		tmp = tmp[:len(tmp)-1]
		if tmp == "Y" || tmp == "y" {
			// Only allow owner to execute the script 0700
			f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
			checkError(err, fmt.Sprintf("Uanble to open file %s", path))
			defer f.Close()
			f.WriteString(data)
		} else {
			os.Exit(1)
		}
	}

}
