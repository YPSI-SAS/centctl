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

// trapCmd represents the trap command
var trapCmd = &cobra.Command{
	Use:   "trap",
	Short: "Modfiy a trap",
	Long:  `Modfiy a trap into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyTrap(name, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyTrap permits to modify a trap in the centreon server
func ModifyTrap(name string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	var values string
	var action string

	switch parameter {
	case "matching":
		action = "addmatching"
		values = name + ";" + value
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value
	}

	err := request.Modify(action, "trap", values, "modify trap", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	trapCmd.Flags().StringP("name", "n", "", "To define the name of the trap to be modified")
	trapCmd.MarkFlagRequired("name")
	trapCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in the list: matching")
	trapCmd.MarkFlagRequired("parameter")
	trapCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"matching", "name", "comment", "output", "oid", "status", "vendor", "matching_mode", "reschedule_svc_enable", "execution_command", "execution_command", "submit_result_enable"}, cobra.ShellCompDirectiveDefault
	})
	trapCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is MATCHING the value must be of the form: stringMarch;regularExpression;status")
	trapCmd.MarkFlagRequired("value")
}
