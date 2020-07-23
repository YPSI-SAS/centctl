/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add objects",
	Long:  `Add an object defined right after.`,
	/*Run: func(cmd *cobra.Command, args []string) error {},*/

}

func init() {
	rootCmd.AddCommand(addCmd)
}
