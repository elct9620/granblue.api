package main

import (
	"fmt"
	"os"

	"github.com/elct9620/granblue.api/cmd"
)

// Entrypoint
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
