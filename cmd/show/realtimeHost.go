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
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// realtimeHostCmd represents the host command
var realtimeHostCmd = &cobra.Command{
	Use:   "realtimeHost",
	Short: "Show one host's realtime details ",
	Long:  `Show one host's realtime details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("hostID")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowRealtimeHost(id, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowRealtimeHost permits to display the details of one host
func ShowRealtimeHost(id int, debugV bool, output string) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := "/monitoring/hosts/" + strconv.Itoa(id)
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "show host", debugV)
	if err != nil {
		return err
	}

	var server host.DetailRealtimeServer
	if len(body) == 0 {
		server = host.DetailRealtimeServer{
			Server: host.DetailRealtimeInformations{
				Name: os.Getenv("SERVER"),
				Host: nil,
			},
		}
	} else {
		//Permits to recover the host contains into the response body
		var hostResult host.DetailRealtimeHost
		json.Unmarshal(body, &hostResult)

		server = host.DetailRealtimeServer{
			Server: host.DetailRealtimeInformations{
				Name: os.Getenv("SERVER"),
				Host: &hostResult,
			},
		}
	}

	//Display details of the host
	displayHost, err := display.DetailRealtimeHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func init() {
	realtimeHostCmd.Flags().IntP("hostID", "i", 0, "ID host")
	realtimeHostCmd.MarkFlagRequired("hostID")
}
