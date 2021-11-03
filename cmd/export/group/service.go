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
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Export group service",
	Long:  `Export group service of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportGroupService(name, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportGroupService permits to export a group service of the centreon server
func ExportGroupService(name []string, regex string, file string, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		groups := getAllGroupService(debugV)
		for _, a := range groups {
			if regex != "" {
				matched, err := regexp.MatchString(regex, a.Name)
				if err != nil {
					fmt.Printf(colorRed, "ERROR:")
					fmt.Println(err.Error())
					os.Exit(1)
				}
				if matched {
					name = append(name, a.Name)
				}
			} else {
				name = append(name, a.Name)
			}
		}
	}
	for _, n := range name {
		err, group := getGroupServiceInfo(n, debugV)
		if err != nil {
			return err
		}
		if group.Name == "" {
			continue
		}

		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,groupService,\""+group.Name+"\",\""+group.Alias+"\"\n", file, writeFile)
		request.WriteValues("modify,groupService,\""+group.Name+"\",activate,\""+group.Activate+"\"\n", file, writeFile)
		request.WriteValues("modify,groupService,\""+group.Name+"\",geo_coords,\""+group.GeoCoords+"\"\n", file, writeFile)
		request.WriteValues("modify,groupService,\""+group.Name+"\",comment,\""+group.Comment+"\"\n", file, writeFile)

		//Write in the file the members service and hostgroup service
		if len(group.Services) != 0 {
			for _, m := range group.Services {
				request.WriteValues("modify,groupService,\""+group.Name+"\",service,\""+m.HostName+"|"+m.ServiceDescription+"\"\n", file, writeFile)
			}
		}
		if len(group.HostGroupServices) != 0 {
			for _, m := range group.HostGroupServices {
				request.WriteValues("modify,groupService,\""+group.Name+"\",hostgroupservice,\""+m.HostgroupName+"|"+m.ServiceDescription+"\"\n", file, writeFile)
			}
		}
	}
	return nil
}

//getGroupServiceInfo permits to get all informations about a service group
func getGroupServiceInfo(name string, debugV bool) (error, service.ExportGroup) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the service group
	values := name + ";name|alias|comment|activate|geo_coords"
	err, body := request.GeneriqueCommandV1Post("getparam", "SG", values, "export group service", debugV, false, "")
	if err != nil {
		return err, service.ExportGroup{}
	}
	var resultGroup service.ExportResult
	json.Unmarshal(body, &resultGroup)

	//Check if the service group is found
	if len(resultGroup.GroupServices) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, service.ExportGroup{}
	}

	//Get the members service of the service group
	err, body = request.GeneriqueCommandV1Post("getservice", "SG", name, "export group service", debugV, false, "")
	if err != nil {
		return err, service.ExportGroup{}
	}
	var resultService service.ExportResultService
	json.Unmarshal(body, &resultService)

	//Get the members hostgroup service of the service group
	err, body = request.GeneriqueCommandV1Post("gethostgroupservice", "SG", name, "export group service", debugV, false, "")
	if err != nil {
		return err, service.ExportGroup{}
	}
	var resultHostGroupService service.ExportResultHostGroupServices
	json.Unmarshal(body, &resultHostGroupService)

	group := resultGroup.GroupServices[0]
	group.Services = resultService.GroupServices
	group.HostGroupServices = resultHostGroupService.HostGroupServices

	return nil, group
}

//getAllGroupService permits to find all service group in the centreon server
func getAllGroupService(debugV bool) []service.ExportGroup {
	//Get all service group
	err, body := request.GeneriqueCommandV1Post("show", "SG", "", "export group service", debugV, false, "")
	if err != nil {
		return []service.ExportGroup{}
	}
	var resultGroup service.ExportResult
	json.Unmarshal(body, &resultGroup)

	return resultGroup.GroupServices
}

func init() {
	serviceCmd.Flags().StringSliceP("name", "n", []string{}, "Servicegroup's name (separate by a comma the multiple values)")
	serviceCmd.Flags().StringP("regex", "r", "", "The regex to apply on the service group's name")

}
