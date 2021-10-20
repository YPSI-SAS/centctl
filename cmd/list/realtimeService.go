/*
MIT License

Copyright (c)  2020-2021 YPSI SAS


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
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// realtimeServiceCmd represents the realtimeService command
var realtimeServiceCmd = &cobra.Command{
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
		regex, _ := cmd.Flags().GetString("regex")
		err := ListRealtimeService(output, state, limit, viewType, poller, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListRealtimeService permits to display the array of Service return by the API
func ListRealtimeService(output string, state string, limit int, viewType string, poller int, regex string, debugV bool) error {
	colorRed := colorMessage.GetColorRed()

	state = strings.ToLower(state)
	output = strings.ToLower(output)
	viewType = strings.ToLower(viewType)

	//Verify that the viewType exists
	if viewType != "unhandled" && viewType != "all" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The type view available are : all or unhandled")
		os.Exit(1)
	}

	//Verify that the state exists
	if !(state == "critical" || state == "ok" || state == "warning" || state == "all" || state == "unknown" || state == "pending") {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The state available for service are : ok, warning, critical, unknown, pending, all")
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
	case "ok":
		stateSearch = "[\"OK\"]"
	case "critical":
		stateSearch = "[\"CRITICAL\"]"
	case "warning":
		stateSearch = "[\"WARNING\"]"
	case "pending":
		stateSearch = "[\"PENDING\"]"
	case "unknown":
		stateSearch = "[\"UNKNOWN\"]"
	default:
		stateSearch = "[\"OK\",\"CRITICAL\",\"WARNING\",\"UNKNOWN\",\"PENDING\"]"

	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/resources?limit=" + strconv.Itoa(limit) + "&types=[\"service\"]&statuses=" + stateSearch + "&states=" + viewTypeSearch
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "list realtimeService", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the array result into the body
	var servicesResult service.RealtimeResultBody
	json.Unmarshal(body, &servicesResult)

	//Permits to get the list of services
	var services = servicesResult.ListServices
	if regex != "" {
		services = deleteRealtimeService(services, regex)
	}

	//Recovery the parent's pollerID
	for i, s := range services {
		pollerID, _ := request.IDPollerHost(s.Parent.ID, debugV)
		services[i].Parent.PollerID = pollerID
	}

	//Get final service based on pollerID
	var finalServices = services
	if poller != -1 {
		finalServices = findAndDeleteService(services, poller)
	}

	//Sort services based on their ID
	sort.SliceStable(finalServices, func(i, j int) bool {
		return finalServices[i].ServiceID < finalServices[j].ServiceID
	})

	server := service.RealtimeServer{
		Server: service.RealtimeInformations{
			Name:     os.Getenv("SERVER"),
			Services: finalServices,
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

func findAndDeleteService(services []service.RealtimeService, poller int) []service.RealtimeService {
	index := 0
	for _, s := range services {
		if s.Parent.PollerID == poller {
			services[index] = s
			index++
		}
	}
	return services[:index]
}

func deleteRealtimeService(services []service.RealtimeService, regex string) []service.RealtimeService {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, h := range services {
		matched, err := regexp.MatchString(regex, h.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			services[index] = h
			index++
		}
	}
	return services[:index]
}

func init() {
	realtimeServiceCmd.Flags().StringP("state", "s", "all", "The state of the hosts you want to list (all, warning, critical, ok, unknown, pending)")
	realtimeServiceCmd.Flags().IntP("limit", "l", 60, "The number of hosts you want to list")
	realtimeServiceCmd.Flags().StringP("viewType", "v", "all", "The type of services (all or unhandled")
	realtimeServiceCmd.Flags().IntP("poller", "p", -1, "The ID poller")
	realtimeServiceCmd.Flags().StringP("regex", "r", "", "The regex to apply on the service's name")

}
