package main

import (
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

Arfuments of sh:
	-a string	Author of the script. default is $USER for your environment (default "eragon")
	-d string	Description of the script (default "Basic shell script")
	-p string	Path for the script to save (default "hello_world.sh")
	-pt		Create script by promting all variables in the termianl
	-t string	Bash script Type (default "hello_world")
	-v string	Version of the script (default "1.0")
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

	// print help string when used without any args
	if len(os.Args) <= 1 {
		fmt.Print(HELP_STRING)
		os.Exit(1)
	}

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
