package main

import (
	"log"

	"github.com/spf13/cobra"
)

// Command
var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "dummy example server",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
