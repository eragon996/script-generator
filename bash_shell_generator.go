package main

import (
	"fmt"
	"os"
)

func sh(shell_type string, path string, desc string, version string, author string, promt bool) {

	switch shell_type {

	case "hello_world":
		if !promt {
			sh_hello_world(path, desc, author, version)
		}
		fmt.Printf("Creating script type: %s, path: %s ...\n", shell_type, path)
		desc = getInput("description", desc)
		version = getInput("version", version)
		author = getInput("user", author)
		sh_hello_world(path, desc, author, version)
	}
}

func sh_hello_world(path string, desc string, author string, version string) {

	out := fmt.Sprintf(
		`#!/bin/bash
#
# Description: %s
# Author: %s
# Version: %s
# -------------------------------------------

echo "Hello World!"
`,
		desc, author, version)
	// Only allow owner to execute the script 0700
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0700)
	checkError(err, fmt.Sprintf("Uanble to open file %s", path))
	defer f.Close()
	f.WriteString(out)
	fmt.Println("existing")
}
