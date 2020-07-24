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
	"centctl/request"
	"centctl/service"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// listRealtimeServiceCmd represents the realtimeService command
var listRealtimeServiceCmd = &cobra.Command{
	Use:   "realtimeService",
	Short: "List the services's realtime informations",
	Long:  `List the services's realtime informations of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		state, _ := cmd.Flags().GetString("state")
		limit, _ := cmd.Flags().GetInt("limit")
		viewType, _ := cmd.Flags().GetString("viewType")
		poller, _ := cmd.Flags().GetInt("poller")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListRealtimeService(output, state, limit, viewType, poller, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListRealtimeService permits to display the array of Service return by the API
func ListRealtimeService(output string, state string, limit int, viewType string, poller int, debugV bool) error {
	state = strings.ToLower(state)
	output = strings.ToLower(output)
	viewType = strings.ToLower(viewType)

	//Verification that the viewType exists
	if viewType != "unhandled" && viewType != "all" {
		fmt.Println("The type view available are : all or unhandled")
		os.Exit(1)
	}

	//Verification that the state exists
	if !(state == "critical" || state == "ok" || state == "warning" || state == "all" || state == "unknown" || state == "pending") {
		fmt.Println("The state available for service are : ok, warning, critical, unknown, pending, all")
		os.Exit(1)
	}

	//Recovery of the response body
	var urlCentreon string
	if poller != 0 {
		urlCentreon = os.Getenv("URL") + "/api/index.php?action=list&object=centreon_realtime_services&limit=" + strconv.Itoa(limit) + "&instance=" + strconv.Itoa(poller) + "&viewType=" + viewType + "&sortType=name&order=desc&status=" + state + "&fields=id,description,host_id,host_name,state,output,acknowledged,active_checks"
	} else {
		urlCentreon = os.Getenv("URL") + "/api/index.php?action=list&object=centreon_realtime_services&limit=" + strconv.Itoa(limit) + "&viewType=" + viewType + "&sortType=name&order=desc&status=" + state + "&fields=id,description,host_id,host_name,state,output,acknowledged,active_checks"
	}
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.Get()

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list realtimeService", "", urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the services contain into the response body
	var services []service.RealtimeService
	json.Unmarshal(body, &services)

	//Sort services based on their ID
	sort.SliceStable(services, func(i, j int) bool {
		return services[i].ServiceID < services[j].ServiceID
	})

	server := service.RealtimeServer{
		Server: service.RealtimeInformations{
			Name:     os.Getenv("SERVER"),
			Services: services,
		},
	}

	//Display all services
	displayService, err := display.RealtimeService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayService)
	return nil
}

func init() {
	listCmd.AddCommand(listRealtimeServiceCmd)
	listRealtimeServiceCmd.Flags().StringP("state", "s", "all", "The state of the hosts you want to list (all, warning, critical, ok, unknown)")
	listRealtimeServiceCmd.Flags().IntP("limit", "l", 60, "The number of hosts you want to list")
	listRealtimeServiceCmd.Flags().StringP("viewType", "v", "all", "The type of services (all or unhandled")
	listRealtimeServiceCmd.Flags().IntP("poller", "p", 0, "The ID poller")
}
