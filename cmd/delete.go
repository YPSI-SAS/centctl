/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete objects",
	Long:  `Delete objects defined right after.`,
	/*Run: func(cmd *cobra.Command, args []string) error {},*/

}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
