/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"centctl/debug"
	"centctl/display"
	"centctl/poller"
	"centctl/request"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listPollerCmd represents the poller command
var listPollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "List the pollers",
	Long:  `List the pollers wof the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListPoller(output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListPoller permits to display the array of poller return by the API
func ListPoller(output string, debugV bool) error {
	output = strings.ToLower(output)

	//Creation of the request body
	requestBody, err := request.CreateBodyRequest("Show", "instance", "")
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list poller", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the pollers contain into the response body
	pollers := poller.Result{}
	json.Unmarshal(body, &pollers)

	//Sort pollers based on their ID
	sort.SliceStable(pollers.Pollers, func(i, j int) bool {
		return strings.ToLower(pollers.Pollers[i].ID) < strings.ToLower(pollers.Pollers[j].ID)
	})

	server := poller.Server{
		Server: poller.Informations{
			Name:    os.Getenv("SERVER"),
			Pollers: pollers.Pollers,
		},
	}

	//Display all pollers
	displayPoller, err := display.Poller(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayPoller)
	return nil
}

func init() {
	listCmd.AddCommand(listPollerCmd)
}
