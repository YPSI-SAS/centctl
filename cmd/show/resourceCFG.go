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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/resourceCFG"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// resourceCFGCmd represents the resourceCFG command
var resourceCFGCmd = &cobra.Command{
	Use:   "resourceCFG",
	Short: "Show one resourceCFG's details",
	Long:  `Show one resourceCFG's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowResourceCFG(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowResourceCFG permits to display the details of one resourceCFG
func ShowResourceCFG(name string, debugV bool, output string) error {
	output = strings.ToLower(output)
	nameVal := strings.Replace(name, "$", "", -1)
	err, body := request.GeneriqueCommandV1Post("show", "RESOURCECFG", nameVal, "show resourceCFG", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the booleanrules contain into the response body
	resourceCFGs := resourceCFG.DetailResult{}
	json.Unmarshal(body, &resourceCFGs)

	//Permits to find the good resourceCFG in the array
	var ResourceCFGFind resourceCFG.DetailResourceCFG

	for _, v := range resourceCFGs.ResourceCFG {
		nameValFind := strings.Replace(v.Name, "$", "", -1)
		if strings.ToLower(nameValFind) == strings.ToLower(nameVal) {
			ResourceCFGFind = v
		}
	}

	var server resourceCFG.DetailServer
	if ResourceCFGFind.Name != "" {
		//Organization of data
		server = resourceCFG.DetailServer{
			Server: resourceCFG.DetailInformations{
				Name:        os.Getenv("SERVER"),
				ResourceCFG: &ResourceCFGFind,
			},
		}
	} else {
		server = resourceCFG.DetailServer{
			Server: resourceCFG.DetailInformations{
				Name:        os.Getenv("SERVER"),
				ResourceCFG: nil,
			},
		}
	}

	//Display details of the resourceCFG
	displayResourceCFG, err := display.DetailResourceCFG(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayResourceCFG)
	return nil
}

func init() {
	resourceCFGCmd.Flags().StringP("name", "n", "", "To define the name of the resourceCFG")
	resourceCFGCmd.MarkFlagRequired("name")
}
