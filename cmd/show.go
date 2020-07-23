/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show with details one host or service",
	Long:  `Show one object's details of the Centreon Server`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.PersistentFlags().String("output", "json", "Type of output (json, yaml, text, csv)")
}
