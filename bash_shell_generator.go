package main

import (
	"fmt"
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

	safeWriteToFile(out, path)
}
