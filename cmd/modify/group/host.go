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
package group

import (
	"centctl/request"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Modify group host",
	Long:  `Modify group host of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyGroupHost(name, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyGroupHost permits to modify a host in the centreon server
func ModifyGroupHost(name string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	//Creation of the request body
	var values string
	var action string
	if strings.ToLower(parameter) == "member" {
		values = name + ";" + value
		action = "addmember"
	} else {
		values = name + ";" + parameter + ";" + value
		action = "setparam"
	}

	err := request.Modify(action, "HG", values, "modify group host", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the name of the host group to be modified")
	hostCmd.MarkFlagRequired("name")
	hostCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: member")
	hostCmd.MarkFlagRequired("parameter")
	hostCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	hostCmd.MarkFlagRequired("value")
}
