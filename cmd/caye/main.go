// Package main provides the command to run caye.
// The caye command is just a wrapper that finds and runs the main file in the .caye folder.
// This can be accomplished simply by running, `go run .caye` without using the caye command.
// The intention is that the entire CI pipeline, including status transmission, analytics, etc. is entirely contained in the caye library
// that is used to build test pipelines.
package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Read first argument
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatalln("Invalid number of arguments")
	}

	project := args[0]
	// Attempt to enter .caye folder

	path := filepath.Join(project, ".caye")
	f, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}

	if f.IsDir() != true {
		log.Fatalln("Expected '.caye' to be directory")
	}

	// Attempt to run package
	cmd := exec.Command("go", "run", ".caye/main.go")
	cmd.Dir = project
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(string(out), err)
	}

	log.Println(string(out))
}
