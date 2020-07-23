/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version",
	Long:  `Show the version of centctl`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("centctl v0.3")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.ResetCommands()
}
