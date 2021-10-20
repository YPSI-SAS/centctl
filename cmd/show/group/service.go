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
package group

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
	Short: "Show one group service's details",
	Long:  `Show one group service's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowGroupService(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowGroupService permits to display the details of one group service
func ShowGroupService(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "SG", name, "show group service", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the group service contain into the response body
	groupsService := service.DetailResultGroup{}
	json.Unmarshal(body, &groupsService)

	//Permits to find the good group service in the array
	var GroupFind service.DetailGroup
	for _, v := range groupsService.DetailGroups {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			GroupFind = v
		}
	}

	var server service.DetailGroupServer
	if GroupFind.Name != "" {
		err, body := request.GeneriqueCommandV1Post("getservice", "SG", GroupFind.Name, "getservice", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		services := service.DetailResultGroupService{}
		json.Unmarshal(body, &services)

		GroupFind.Services = services.Services

		err, body = request.GeneriqueCommandV1Post("gethostgroupservice", "SG", GroupFind.Name, "gethostgroupservice", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		serviceHostGroup := service.DetailResultHostGroupService{}
		json.Unmarshal(body, &serviceHostGroup)

		GroupFind.HostGroupServices = serviceHostGroup.Services
		//Organization of data
		server = service.DetailGroupServer{
			Server: service.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: &GroupFind,
			},
		}
	} else {
		//Organization of data
		server = service.DetailGroupServer{
			Server: service.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: nil,
			},
		}
	}

	//Display details of the group service
	displayGroupService, err := display.DetailGroupService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayGroupService)
	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "To define the service group which will show")
	serviceCmd.MarkFlagRequired("name")
}
