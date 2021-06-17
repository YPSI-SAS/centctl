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
	"centctl/resources/poller"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// pollerCmd represents the poller command
var pollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "Show one poller's details",
	Long:  `Show one poller's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowPoller(name, output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowPoller permits to display the poller return by the API
func ShowPoller(name string, output string, debugV bool) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/beta/platform/topology"
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "show poller", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the result
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return err
	}

	var pollerFind poller.DetailPoller
	_, ok := (f.(map[string]interface{}))["graph"]
	if ok {
		//Permits to go down in the JSON response for find list of nodes (list of pollers)
		nodes := ((f.(map[string]interface{}))["graph"].(map[string]interface{}))["nodes"].(map[string]interface{})

		//For each node in nodes, get informations of the poller, create new poller and had this in the list
		for _, v := range nodes {
			var pollerVal poller.DetailPoller
			switch c := v.(type) {
			case map[string]interface{}:
				metadataVal := c["metadata"].(map[string]interface{})
				var metadata poller.DetailMetadata
				pollerVal.Label = c["label"].(string)
				pollerVal.Type = c["type"].(string)
				metadata.Address = metadataVal["address"].(string)
				metadata.CentreonID = metadataVal["centreon-id"].(string)
				if metadataVal["hostname"] != nil {
					metadata.HostName = metadataVal["hostname"].(string)
				}
				pollerVal.Metadata = metadata
			}
			if name == pollerVal.Label {
				pollerFind = pollerVal
			}

		}
	}

	var server poller.DetailServer
	if pollerFind.Label != "" {
		server = poller.DetailServer{
			Server: poller.DetailInformations{
				Name:   os.Getenv("SERVER"),
				Poller: &pollerFind,
			},
		}
	} else {
		server = poller.DetailServer{
			Server: poller.DetailInformations{
				Name:   os.Getenv("SERVER"),
				Poller: nil,
			},
		}
	}

	//Display all pollers
	displayPoller, err := display.DetailPoller(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayPoller)
	return nil
}

func init() {
	pollerCmd.Flags().StringP("name", "n", "", "Poller's name")
	pollerCmd.MarkFlagRequired("name")
}
