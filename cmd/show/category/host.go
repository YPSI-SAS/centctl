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
	Short: "Show one category host's details",
	Long:  `Show one category host's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowCategoryHost(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowCategoryHost permits to display the details of one category host
func ShowCategoryHost(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "HC", name, "show category host", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the category host contain into the response body
	categoriesHost := host.DetailResultCategory{}
	json.Unmarshal(body, &categoriesHost)

	//Permits to find the good category host in the array
	var CategoryFind host.DetailCategory
	for _, v := range categoriesHost.Categories {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			CategoryFind = v
		}
	}

	var server host.DetailCategoryServer
	if CategoryFind.Name != "" {

		err, body := request.GeneriqueCommandV1Post("getmember", "HC", CategoryFind.Name, "getmember", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		members := host.DetailResultCategoryMember{}
		json.Unmarshal(body, &members)

		CategoryFind.Members = members.Members

		//Organization of data
		server = host.DetailCategoryServer{
			Server: host.DetailCategoryInformations{
				Name:     os.Getenv("SERVER"),
				Category: &CategoryFind,
			},
		}
	} else {
		//Organization of data
		server = host.DetailCategoryServer{
			Server: host.DetailCategoryInformations{
				Name:     os.Getenv("SERVER"),
				Category: nil,
			},
		}
	}

	//Display details of the category host
	displayCategoryHost, err := display.DetailCategoryHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayCategoryHost)
	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the host category which will show")
	hostCmd.MarkFlagRequired("name")
	hostCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetCategoryHostNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
