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
package modify

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// resourceCFGCmd represents the resourceCFG command
var resourceCFGCmd = &cobra.Command{
	Use:   "resourceCFG",
	Short: "Modfiy a resourceCFG",
	Long:  `Modfiy a resourceCFG into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyResourceCFG(id, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyResourceCFG permits to modify a resourceCFG in the centreon server
func ModifyResourceCFG(id string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	//Creation of the request body
	values := id + ";" + parameter + ";" + value

	err := request.Modify("setparam", "resourceCFG", values, "modify resourceCFG", id, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	resourceCFGCmd.Flags().StringP("id", "i", "", "To define the id of the resourceCFG to be modified")
	resourceCFGCmd.MarkFlagRequired("id")
	resourceCFGCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation")
	resourceCFGCmd.MarkFlagRequired("parameter")
	resourceCFGCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	resourceCFGCmd.MarkFlagRequired("value")
}
