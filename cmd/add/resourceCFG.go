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

package add

import (
	"centctl/cmd/modify"
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/resourceCFG"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// resourceCFGCmd represents the resourceCFG command
var resourceCFGCmd = &cobra.Command{
	Use:   "resourceCFG",
	Short: "Add a resourceCFG",
	Long:  `Add a resourceCFG into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		value, _ := cmd.Flags().GetString("value")
		instance, _ := cmd.Flags().GetString("instance")
		comment, _ := cmd.Flags().GetString("comment")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddResourceCFG(name, value, instance, comment, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddResourceCFG permits to add a resourceCFG in the centreon server
func AddResourceCFG(name string, value string, instance string, comment string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + value + ";" + instance + ";" + comment
	err := request.Add("add", "RESOURCECFG", values, "add resourceCFG", name, debugV, isImport, false, "", "")
	if err != nil {
		if strings.Contains(err.Error(), "already tied to instance") {
			err, resource := getResourceID(name, debugV)
			if err != nil {
				return err
			}
			for _, inst := range resource.Instance {
				instance += "|" + inst
			}
			err = modify.ModifyResourceCFG(resource.ID, "instance", instance, debugV, isImport, true)
			if err != nil {
				return err
			}
		}
		return err
	}
	return nil
}

func getResourceID(name string, debugV bool) (error, resourceCFG.DetailResourceCFG) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "RESOURCECFG", name, "list resourceCFG", debugV, false, "")
	if err != nil {
		return err, resourceCFG.DetailResourceCFG{}
	}

	//Permits to recover the resourceCFG contain into the response body
	resourceCFGs := resourceCFG.DetailResult{}
	json.Unmarshal(body, &resourceCFGs)
	//Check if the resourceCFG is found
	if len(resourceCFGs.ResourceCFG) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, resourceCFG.DetailResourceCFG{}
	}
	return nil, resourceCFGs.ResourceCFG[0]
}

func init() {
	resourceCFGCmd.Flags().StringP("name", "n", "", "To define the macro name of the resourceCFG")
	resourceCFGCmd.MarkFlagRequired("name")
	resourceCFGCmd.Flags().StringP("value", "v", "", "To define the macro value of the resourceCFG")
	resourceCFGCmd.MarkFlagRequired("value")
	resourceCFGCmd.Flags().StringP("instance", "i", "", "To define the instance of the resourceCFG")
	resourceCFGCmd.MarkFlagRequired("instance")
	resourceCFGCmd.RegisterFlagCompletionFunc("instance", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetPollerNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	resourceCFGCmd.Flags().StringP("comment", "c", "", "To define the comment of the resourceCFG")
	resourceCFGCmd.MarkFlagRequired("comment")
}
