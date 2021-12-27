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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// realtimeServiceCmd represents the realtimeService command
var realtimeServiceCmd = &cobra.Command{
	Use:   "realtimeService",
	Short: "Show one service's realtime details ",
	Long:  `Show one service realtime of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		hostID, _ := cmd.Flags().GetInt("hostID")
		serviceID, _ := cmd.Flags().GetInt("serviceID")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowRealtimeService(hostID, serviceID, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowRealtimeService permits to display the details of one service
func ShowRealtimeService(hostID int, serviceID int, debugV bool, output string) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := "/monitoring/hosts/" + strconv.Itoa(hostID) + "/services/" + strconv.Itoa(serviceID)
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "show service", debugV)
	if err != nil {
		return err
	}

	var server service.DetailRealtimeServer
	if len(body) != 0 {
		//Permits to recover the service contain into the response body
		var serviceResult service.DetailRealtimeService
		json.Unmarshal(body, &serviceResult)

		server = service.DetailRealtimeServer{
			Server: service.DetailRealtimeInformations{
				Name:    os.Getenv("SERVER"),
				Service: &serviceResult,
			},
		}
	} else {
		server = service.DetailRealtimeServer{
			Server: service.DetailRealtimeInformations{
				Name:    os.Getenv("SERVER"),
				Service: nil,
			},
		}
	}

	//Display details of the service
	displayService, err := display.DetailRealtimeService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayService)
	return nil
}

func init() {
	realtimeServiceCmd.Flags().IntP("hostID", "i", 0, "ID of the host wich the service is attached")
	realtimeServiceCmd.Flags().IntP("serviceID", "s", 0, "ID of the service")
	realtimeServiceCmd.MarkFlagRequired("hostID")
	realtimeServiceCmd.MarkFlagRequired("serviceID")
}
