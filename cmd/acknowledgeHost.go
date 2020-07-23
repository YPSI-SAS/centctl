/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"centctl/debug"
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// acknowledgeHostCmd represents the host command
var acknowledgeHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Acknowledge hosts",
	Long:  `Acknowledge the host described right after`,
	Run: func(cmd *cobra.Command, args []string) {
		comment, _ := cmd.Flags().GetString("comment")
		hostName, _ := cmd.Flags().GetString("hostName")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AcknowledgeHost(comment, hostName, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AcknowledgeHost permits to acnowledge a host in the centreon server
func AcknowledgeHost(comment string, hostName string, debugV bool) error {
	//Creation of the request body
	values := "HOST;" + hostName + ";" + comment + ";2;0;1"
	requestBody, err := request.CreateBodyRequest("add", "RTACKNOWLEDGEMENT", values)
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi")
	statusCode, body, err := client.CentreonCLAPI(requestBody)

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("acknowledge host", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Verification with the response body that the acknowledge was carried out
	if string(body) != "{\"result\":[]}" {
		fmt.Println("erreur: ", string(body))
		os.Exit(1)
	}

	fmt.Printf("The host `%v` is acknowledged\n", hostName)
	return nil
}

func init() {
	acknowledgeCmd.AddCommand(acknowledgeHostCmd)
	acknowledgeHostCmd.Flags().StringP("hostName", "n", "", "To know the host which must be acknowledge")
	acknowledgeHostCmd.MarkFlagRequired("hostName")
}
