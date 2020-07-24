/*
MIT License

Copyright (c) 2020 YPSI SAS
Centctl is developped by : Mélissa Bertin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
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
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listHostCmd represents the host command
var listHostCmd = &cobra.Command{
	Use:   "host",
	Short: "List the hosts",
	Long:  `List the hosts of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListHost(output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListHost permits to display the array of host return by the API
func ListHost(output string, debugV bool) error {
	output = strings.ToLower(output)

	//Creation of the request body
	requestBody, err := request.CreateBodyRequest("Show", "host", "")
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list host", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the hosts contain into the response body
	hosts := host.Result{}
	json.Unmarshal(body, &hosts)

	//Sort hosts based on their ID
	sort.SliceStable(hosts.Hosts, func(i, j int) bool {
		return strings.ToLower(hosts.Hosts[i].ID) < strings.ToLower(hosts.Hosts[i].ID)
	})

	//Organization of data
	server := host.Server{
		Server: host.Informations{
			Name:  os.Getenv("SERVER"),
			Hosts: hosts.Hosts,
		},
	}

	//Display all hosts
	displayHost, err := display.Host(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func init() {
	listCmd.AddCommand(listHostCmd)
}
