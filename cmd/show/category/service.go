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
package category

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
	Short: "Show one category service's details",
	Long:  `Show one category service's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowCategoryService(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowCategoryService permits to display the details of one category service
func ShowCategoryService(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "SC", name, "show category service", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the category service contain into the response body
	categoriesService := service.DetailResultCategory{}
	json.Unmarshal(body, &categoriesService)

	//Permits to find the good category service in the array
	var CategoryFind service.DetailCategory
	for _, v := range categoriesService.Categories {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			CategoryFind = v
		}
	}

	var server service.DetailCategoryServer
	if CategoryFind.Name != "" {
		err, body := request.GeneriqueCommandV1Post("getservice", "SC", CategoryFind.Name, "getservice", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		services := service.DetailResultCategoryService{}
		json.Unmarshal(body, &services)

		CategoryFind.Services = services.Services

		err, body = request.GeneriqueCommandV1Post("getservicetemplate", "SC", CategoryFind.Name, "getservicetemplate", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		serviceTemplates := service.DetailResultCategoryServiceTemplate{}
		json.Unmarshal(body, &serviceTemplates)

		CategoryFind.ServiceTemplates = serviceTemplates.ServiceTemplates

		//Organization of data
		server = service.DetailCategoryServer{
			Server: service.DetailCategoryInformations{
				Name:     os.Getenv("SERVER"),
				Category: &CategoryFind,
			},
		}
	} else {
		//Organization of data
		server = service.DetailCategoryServer{
			Server: service.DetailCategoryInformations{
				Name:     os.Getenv("SERVER"),
				Category: nil,
			},
		}
	}

	//Display details of the category service
	displayCategoryService, err := display.DetailCategoryService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayCategoryService)
	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "To define the service category which will show")
	serviceCmd.MarkFlagRequired("name")
	serviceCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetCategoryServiceNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
