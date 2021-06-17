/*MIT License

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
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Show one host's details",
	Long:  `Show one host's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowHost(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowHost permits to display the details of one host
func ShowHost(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "host", name, "show host", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the hosts contain into the response body
	hosts := host.DetailResult{}
	json.Unmarshal(body, &hosts)

	//Permits to find the good host in the array
	var HostFind host.DetailHost
	for _, v := range hosts.DetailHosts {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			HostFind = v
		}
	}

	var server host.DetailServer
	if HostFind.Name != "" {

		err, body := request.GeneriqueCommandV1Post("getparent", "HOST", HostFind.Name, "getparent", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		parents := host.DetailResultHostParent{}
		json.Unmarshal(body, &parents)

		HostFind.Parent = parents.Parents

		err, body = request.GeneriqueCommandV1Post("getchild", "HOST", HostFind.Name, "getchild", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		childs := host.DetailResultHostChild{}
		json.Unmarshal(body, &childs)

		HostFind.Child = childs.Childs

		//Organization of data
		server = host.DetailServer{
			Server: host.DetailInformations{
				Name: os.Getenv("SERVER"),
				Host: &HostFind,
			},
		}
	} else {
		server = host.DetailServer{
			Server: host.DetailInformations{
				Name: os.Getenv("SERVER"),
				Host: nil,
			},
		}
	}

	//Display details of the host
	displayHost, err := display.DetailHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "Host's name")
	hostCmd.MarkFlagRequired("name")
}
