package main

import (
	"github.com/pit-cli/pit/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "pit"}
	rootCmd.AddCommand(cmd.CmdSearch)
	rootCmd.Execute()
}
