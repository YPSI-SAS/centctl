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

package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// timelineServiceCmd represents the service command
var timelineServiceCmd = &cobra.Command{
	Use:   "timelineService",
	Short: "List the service's timeline informations",
	Long:  `List the service's timelien information of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		idH, _ := cmd.Flags().GetInt("idH")
		idS, _ := cmd.Flags().GetInt("idS")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListTimelineService(output, idH, idS, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListTimelineService permits to display the array of timeline information of service return by the API
func ListTimelineService(output string, idH int, idS int, debugV bool) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/hosts/" + strconv.Itoa(idH) + "/services/" + strconv.Itoa(idS) + "/timeline"
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "show timelineService", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the array result into the body
	var hostResult service.TimelineServiceResult
	json.Unmarshal(body, &hostResult)

	//Permits to get the list of services
	var services = hostResult.DetailTimelineServices

	//Sort services based on their ID
	sort.SliceStable(services, func(i, j int) bool {
		return services[i].ID < services[j].ID
	})

	server := service.DetailTimelineServer{
		Server: service.DetailTimelineInformations{
			Name:            os.Getenv("SERVER"),
			TimelineService: services,
		},
	}

	//Display all services
	displayService, err := display.DetailTimelineService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayService)
	return nil
}

func init() {
	timelineServiceCmd.Flags().Int("idH", -1, "Host's ID")
	timelineServiceCmd.MarkFlagRequired("idH")
	timelineServiceCmd.Flags().Int("idS", -1, "Service's ID")
	timelineServiceCmd.MarkFlagRequired("idS")
}
