/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export configuration on a poller",
	Long:  `Export configuration on a poller`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := Export(name, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//Export the poller configuration
func Export(name string, debugV bool) error {

	client := request.NewClient(os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi")
	err := client.ExportConf(name, debugV)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP("name", "n", "", "Name of the poller")
	exportCmd.MarkFlagRequired("name")
}
