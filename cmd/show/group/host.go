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
package group

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Show one group host's details",
	Long:  `Show one group host's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowGroupHost(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowGroupHost permits to display the details of one group host
func ShowGroupHost(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "HG", name, "show group host", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the group host contain into the response body
	groupsHost := host.DetailResultGroup{}
	json.Unmarshal(body, &groupsHost)

	//Permits to find the good group host in the array
	var GroupFind host.DetailGroup
	for _, v := range groupsHost.DetailGroups {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			GroupFind = v
		}
	}

	var server host.DetailGroupServer
	if GroupFind.Name != "" {

		err, body := request.GeneriqueCommandV1Post("getmember", "HG", GroupFind.Name, "getmember", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		members := host.DetailResultGroupMember{}
		json.Unmarshal(body, &members)

		GroupFind.Members = members.Members

		//Organization of data
		server = host.DetailGroupServer{
			Server: host.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: &GroupFind,
			},
		}
	} else {
		//Organization of data
		server = host.DetailGroupServer{
			Server: host.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: nil,
			},
		}
	}

	//Display details of the group host
	displayGroupHost, err := display.DetailGroupHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayGroupHost)
	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the host group which will show")
	hostCmd.MarkFlagRequired("name")
}
