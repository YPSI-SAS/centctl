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
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Show one service's details",
	Long:  `Show one service's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowService(name, description, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowService permits to display the details of one service
func ShowService(name string, description string, debugV bool, output string) error {
	output = strings.ToLower(output)

	values := name + ";" + description
	err, body := request.GeneriqueCommandV1Post("show", "service", values, "show service", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the services contain into the response body
	services := service.DetailResult{}
	json.Unmarshal(body, &services)

	//Permits to find the good service in the array
	var ServiceFind service.DetailService
	for _, v := range services.DetailServices {
		if strings.ToLower(v.HostName) == strings.ToLower(name) && strings.ToLower(v.Description) == strings.ToLower(description) {
			ServiceFind = v
		}
	}

	var server service.DetailServer
	if ServiceFind.Description != "" {
		//Organization of data
		server = service.DetailServer{
			Server: service.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Service: &ServiceFind,
			},
		}
	} else {
		server = service.DetailServer{
			Server: service.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Service: nil,
			},
		}
	}

	//Display details of the service
	displayService, err := display.DetailService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayService)
	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "Host's name")
	serviceCmd.MarkFlagRequired("name")
	serviceCmd.Flags().StringP("description", "d", "", "Service's description")
	serviceCmd.MarkFlagRequired("description")
}
