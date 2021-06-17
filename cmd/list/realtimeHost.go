/*
MIT License

Copyright (c)  2020-2021 YPSI SAS
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

package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// realtimeHostCmd represents the host command
var realtimeHostCmd = &cobra.Command{
	Use:   "realtimeHost",
	Short: "List the hosts's realtime informations",
	Long:  `List the hosts's realtime information of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		state, _ := cmd.Flags().GetString("state")
		limit, _ := cmd.Flags().GetInt("limit")
		viewType, _ := cmd.Flags().GetString("viewType")
		poller, _ := cmd.Flags().GetInt("poller")
		regex, _ := cmd.Flags().GetString("regex")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListRealtimeHost(output, state, limit, viewType, poller, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListRealtimeHost permits to display the array of realtime information of host return by the API
func ListRealtimeHost(output string, state string, limit int, viewType string, poller int, regex string, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	state = strings.ToLower(state)
	output = strings.ToLower(output)
	viewType = strings.ToLower(viewType)

	//Verify that the state exists
	if !(state == "up" || state == "down" || state == "unrea" || state == "all" || state == "pending") {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The states available for host are : up, down, unrea, pending, all")
		os.Exit(1)
	}

	//Verify that the viewType exists
	if viewType != "unhandled" && viewType != "all" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The type view available are : all or unhandled")
		os.Exit(1)
	}

	//Conversion of the viewType
	viewTypeSearch := ""
	switch viewType {
	case "unhandled":
		viewTypeSearch = "[\"unhandled_problems\"]"
	case "all":
		viewTypeSearch = "[\"all\"]"
	}

	//Conversion of the state
	stateSearch := ""
	switch state {
	case "up":
		stateSearch = "[\"UP\"]"
	case "down":
		stateSearch = "[\"DOWN\"]"
	case "unrea":
		stateSearch = "[\"UNREACHABLE\"]"
	case "pending":
		stateSearch = "[\"PENDING\"]"
	default:
		stateSearch = "[\"UP\",\"DOWN\",\"UNREACHABLE\",\"PENDING\"]"

	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/resources?limit=" + strconv.Itoa(limit) + "&types=[\"host\"]&statuses=" + stateSearch + "&states=" + viewTypeSearch
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "list realtimeHost", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the array result into the body
	var hostResult host.RealtimeResultBodyV2
	json.Unmarshal(body, &hostResult)

	//Permits to get the list of hosts
	var hosts = hostResult.ListHosts
	if regex != "" {
		hosts = deleteRealtimeHost(hosts, regex)
	}

	//Recovery the host's pollerID
	for i, h := range hosts {
		pollerID, _ := request.IDPollerHost(h.ID, debugV)
		hosts[i].PollerID = pollerID
	}

	//Get final host based on pollerID
	var finalHosts = hosts
	if poller != -1 {
		finalHosts = findAndDeleteHost(hosts, poller)
	}

	//Sort hosts based on their ID
	sort.SliceStable(finalHosts, func(i, j int) bool {
		return finalHosts[i].ID < finalHosts[j].ID
	})

	server := host.RealtimeServerV2{
		Server: host.RealtimeInformationsV2{
			Name:  os.Getenv("SERVER"),
			Hosts: finalHosts,
		},
	}

	//Display all hosts
	displayHost, err := display.RealtimeHostV2(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func findAndDeleteHost(hosts []host.RealtimeHostV2, poller int) []host.RealtimeHostV2 {
	index := 0
	for _, h := range hosts {
		if h.PollerID == poller {
			hosts[index] = h
			index++
		}
	}
	return hosts[:index]
}

func deleteRealtimeHost(hosts []host.RealtimeHostV2, regex string) []host.RealtimeHostV2 {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, h := range hosts {
		matched, err := regexp.MatchString(regex, h.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			hosts[index] = h
			index++
		}
	}
	return hosts[:index]
}

func init() {
	realtimeHostCmd.Flags().StringP("state", "s", "all", "The state of the hosts you want to list (all, up, down, pending, unrea)")
	realtimeHostCmd.Flags().IntP("limit", "l", 60, "The number of hosts you want to list")
	realtimeHostCmd.Flags().StringP("viewType", "v", "all", "The type of hosts (all or unhandled")
	realtimeHostCmd.Flags().IntP("poller", "p", -1, "The ID poller")
	realtimeHostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host's name")

}
