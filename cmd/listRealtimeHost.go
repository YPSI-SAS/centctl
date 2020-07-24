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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// listRealtimeHostCmd represents the host command
var listRealtimeHostCmd = &cobra.Command{
	Use:   "realtimeHost",
	Short: "List the hosts's realtime informations",
	Long:  `List the hosts's realtime information of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		state, _ := cmd.Flags().GetString("state")
		limit, _ := cmd.Flags().GetInt("limit")
		viewType, _ := cmd.Flags().GetString("viewType")
		poller, _ := cmd.Flags().GetInt("poller")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListRealtimeHost(output, state, limit, viewType, poller, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListRealtimeHost permits to display the array of realtime information of host status return by the API
func ListRealtimeHost(output string, state string, limit int, viewType string, poller int, debugV bool) error {
	state = strings.ToLower(state)
	output = strings.ToLower(output)
	viewType = strings.ToLower(viewType)

	//Verification that the state exists
	if !(state == "up" || state == "down" || state == "unrea" || state == "all") {
		fmt.Println("The states available for host are : up, down, unrea, all")
		os.Exit(1)
	}

	//Verification that the viewType exists
	if viewType != "unhandled" && viewType != "all" {
		fmt.Println("The type view available are : all or unhandled")
		os.Exit(1)
	}

	//Recovery of the response body
	var urlCentreon string
	if poller != 0 {
		urlCentreon = os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&limit=" + strconv.Itoa(limit) + "&instance=" + strconv.Itoa(poller) + "&viewType=" + viewType + "&order=desc&status=" + state + "&fields=id,name,alias,address,state,acknowledged,active_checks,instance"
	} else {
		urlCentreon = os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&limit=" + strconv.Itoa(limit) + "&viewType=" + viewType + "&order=desc&status=" + state + "&fields=id,name,alias,address,state,acknowledged,active_checks,instance"
	}
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.Get()
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list realtimeHost", "", urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the hosts contain into the response body
	var hosts []host.RealtimeHost
	json.Unmarshal(body, &hosts)

	//Sort hosts based on their ID
	sort.SliceStable(hosts, func(i, j int) bool {
		return hosts[i].ID < hosts[j].ID
	})

	server := host.RealtimeServer{
		Server: host.RealtimeInformations{
			Name:  os.Getenv("SERVER"),
			Hosts: hosts,
		},
	}

	//Display all hosts
	displayHost, err := display.RealtimeHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func init() {
	listCmd.AddCommand(listRealtimeHostCmd)
	listRealtimeHostCmd.Flags().StringP("state", "s", "all", "The state of the hosts you want to list (all, up, down, unrea)")
	listRealtimeHostCmd.Flags().IntP("limit", "l", 60, "The number of hosts you want to list")
	listRealtimeHostCmd.Flags().StringP("viewType", "v", "all", "The type of hosts (all or unhandled")
	listRealtimeHostCmd.Flags().IntP("poller", "p", 0, "The ID poller")
}
