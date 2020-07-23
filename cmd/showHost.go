/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"centctl/debug"
	"centctl/display"
	"centctl/host"
	"centctl/request"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// showHostCmd represents the host command
var showHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Show one host's details ",
	Long:  `Show one host's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowHost(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowHost permits to display the details of one host
func ShowHost(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&search=" + name + "&fields=id,name,alias,address,state,state_type,output,max_check_attempts,check_attempt,last_check,last_state_change,last_hard_state_change,acknowledged,active_checks,instance,criticality,passive_checks,notify"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.Get()
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("show host", "", urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the hosts contain into the response body
	var hosts []host.DetailHost
	json.Unmarshal(body, &hosts)

	if len(hosts) == 0 {
		fmt.Println("no host with this name")
		os.Exit(1)
	}

	//Permits to find the good host in the array
	var hostFind host.DetailHost
	for _, v := range hosts {
		if v.Name == name {
			hostFind = v
		}
	}

	server := host.DetailServer{
		Server: host.DetailInformations{
			Name: os.Getenv("SERVER"),
			Host: hostFind,
		},
	}

	//Display detail of the host
	displayHost, err := display.DetailHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func init() {
	showCmd.AddCommand(showHostCmd)
	showHostCmd.Flags().StringP("name", "n", "", "Name host")
	showHostCmd.MarkFlagRequired("name")
}
