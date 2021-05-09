package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const (
	HELP_STRING = `
Usage of sgen:
	sgen <command>  [arguments] [path]

The commands are:
	sh		generate bash shell scripts
	py		generate python scripts
	help		print this message

`
	CMD_HELP_STRING = `sgen <command>  [arguments] [path]
Valid commands are 'sh', 'py' and 'help'
`
)

func main() {

	// env variables used
	USER := os.Getenv("USER")

	// bash subcommand
	bashCmd := flag.NewFlagSet("sh", flag.ExitOnError)
	bashType := bashCmd.String("t", "hello_world", "Bash script Type")
	bashPath := bashCmd.String("p", "hello_world.sh", "Path for the script to save")
	bashDesc := bashCmd.String("d", "Basic shell script", "Description of the script")
	bashVersion := bashCmd.String("v", "1.0", "Version of the script")
	bashAuthor := bashCmd.String("a", USER, "Author of the script. default is $USER for your environment")
	bashPromt := bashCmd.Bool("pt", false, "Create script by promting all variables in the termianl")

	// python subcommand
	pythonCmd := flag.NewFlagSet("py", flag.ExitOnError)
	pythonType := pythonCmd.String("type", "hello_world", "Bash script type")

	// help subcommand
	hellpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	// print help strin
	// subcommand check
	if len(os.Args) < 2 {
		fmt.Print(CMD_HELP_STRING)
		os.Exit(1)
	}

	switch os.Args[1] {

	case "sh":
		bashCmd.Parse(os.Args[2:])
		sh(*bashType, *bashPath, *bashDesc, *bashVersion, *bashAuthor, *bashPromt)

	case "py":
		pythonCmd.Parse(os.Args[2:])
		fmt.Println("sh command", *pythonType)

	case "help":
		hellpCmd.Parse(os.Args[2:])
		fmt.Printf(HELP_STRING)

	default:
		os.Exit(1)
	}

}

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

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}

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
