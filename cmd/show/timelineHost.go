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

package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// timelineHostCmd represents the host command
var timelineHostCmd = &cobra.Command{
	Use:   "timelineHost",
	Short: "List the host's timeline informations",
	Long:  `List the host's timelien information of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		id, _ := cmd.Flags().GetInt("id")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListTimelineHost(output, id, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListTimelineHost permits to display the array of timeline information of host return by the API
func ListTimelineHost(output string, id int, debugV bool) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/hosts/" + strconv.Itoa(id) + "/timeline"
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "show timelineHost", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the array result into the body
	var hostResult host.TimelineHostResult
	json.Unmarshal(body, &hostResult)

	//Permits to get the list of hosts
	var hosts = hostResult.DetailTimelineHosts

	//Sort hosts based on their ID
	sort.SliceStable(hosts, func(i, j int) bool {
		return hosts[i].ID < hosts[j].ID
	})

	server := host.DetailTimelineServer{
		Server: host.DetailTimelineInformations{
			Name:         os.Getenv("SERVER"),
			TimelineHost: hosts,
		},
	}

	//Display all hosts
	displayHost, err := display.DetailTimelineHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func init() {
	timelineHostCmd.Flags().IntP("id", "i", -1, "Host's ID")
	timelineHostCmd.MarkFlagRequired("id")
}
