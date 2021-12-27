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
	urlCentreon := "/configuration/monitoring-servers?search={\"name\":\"" + name + "\"}"
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "show poller", debugV)
	if err != nil {
		return err
	}

	var server poller.DetailServer
	if len(body) == 0 {
		server = poller.DetailServer{
			Server: poller.DetailInformations{
				Name:   os.Getenv("SERVER"),
				Poller: nil,
			},
		}
	} else {
		//Permits to recover the poller contains into the response body
		var pollerResult poller.ResultDetailPoller
		json.Unmarshal(body, &pollerResult)

		server = poller.DetailServer{
			Server: poller.DetailInformations{
				Name:   os.Getenv("SERVER"),
				Poller: &pollerResult.Pollers[0],
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
	pollerCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetPollerNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
