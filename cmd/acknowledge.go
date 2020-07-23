/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"github.com/spf13/cobra"
)

// acknowledgeCmd represents the acknowledge command
var acknowledgeCmd = &cobra.Command{
	Use:   "acknowledge",
	Short: "Acknowledge hosts/services",
	Long:  `Acknowledge hosts or services`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(acknowledgeCmd)
	acknowledgeCmd.PersistentFlags().StringP("comment", "c", "", "To define a comment for the acknowledgment")
	acknowledgeCmd.MarkPersistentFlagRequired("comment")
}
