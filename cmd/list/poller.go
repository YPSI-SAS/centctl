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
	"centctl/resources/poller"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// pollerCmd represents the poller command
var pollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "List the pollers",
	Long:  `List the pollers wof the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListPoller(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListPoller permits to display the array of poller return by the API
func ListPoller(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/beta/platform/topology"
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "list poller", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the result
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return err
	}

	var pollers []poller.Poller
	_, ok := (f.(map[string]interface{}))["graph"]
	if ok {
		//Permits to go down in the JSON response for find list of nodes (list of pollers)
		nodes := ((f.(map[string]interface{}))["graph"].(map[string]interface{}))["nodes"].(map[string]interface{})

		//For each node in nodes, get informations of the poller, create new poller and had this in the list
		for _, v := range nodes {
			var pollerVal poller.Poller
			switch c := v.(type) {
			case map[string]interface{}:
				metadataVal := c["metadata"].(map[string]interface{})
				var metadata poller.Metadata
				pollerVal.Label = c["label"].(string)
				pollerVal.Type = c["type"].(string)
				metadata.Address = metadataVal["address"].(string)
				metadata.CentreonID = metadataVal["centreon-id"].(string)
				if metadataVal["hostname"] != nil {
					metadata.HostName = metadataVal["hostname"].(string)
				}
				pollerVal.Metadata = metadata
			}
			pollers = append(pollers, pollerVal)
		}

		if regex != "" {
			pollers = deletePoller(pollers, regex)
		}

		//Sort pollers based on their ID
		sort.SliceStable(pollers, func(i, j int) bool {
			valI, _ := strconv.Atoi(pollers[i].Metadata.CentreonID)
			valJ, _ := strconv.Atoi(pollers[j].Metadata.CentreonID)
			return valI < valJ
		})
	}

	server := poller.Server{
		Server: poller.Informations{
			Name:    os.Getenv("SERVER"),
			Pollers: pollers,
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

func deletePoller(pollers []poller.Poller, regex string) []poller.Poller {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range pollers {
		matched, err := regexp.MatchString(regex, s.Label)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			pollers[index] = s
			index++
		}
	}
	return pollers[:index]
}

func init() {
	pollerCmd.Flags().StringP("regex", "r", "", "The regex to apply on the poller's name")
}
