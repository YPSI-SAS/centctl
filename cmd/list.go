/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List objects",
	Long:  `List objects of the Centreon Server`,
	/*Run: func(cmd *cobra.Command, args []string) error {},*/
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().String("output", "json", "Type of output (json, yaml, text, csv)")
}
