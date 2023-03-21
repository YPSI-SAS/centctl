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
	"strings"

	"github.com/spf13/cobra"
)

// engineCFGCmd represents the engineCFG command
var engineCFGCmd = &cobra.Command{
	Use:   "engineCFG",
	Short: "Modfiy a engineCFG",
	Long:  `Modfiy a engineCFG into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		operation, _ := cmd.Flags().GetString("operation")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyEngineCFG(name, parameter, value, operation, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyEngineCFG permits to modify a engineCFG in the centreon server
func ModifyEngineCFG(name string, parameter string, value string, operation string, debugV bool, isImport bool, detail bool) error {
	var action string
	var values string
	operation = strings.ToLower(operation)

	switch strings.ToLower(parameter) {
	case "brokermodule":
		action = operation + strings.ToLower(parameter)
		values = name + ";" + value
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value

	}

	err := request.Modify(action, "ENGINECFG", values, "modify engineCFG", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	engineCFGCmd.Flags().StringP("name", "n", "", "To define the name of the engine configuration to be modified")
	engineCFGCmd.MarkFlagRequired("name")
	engineCFGCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetEngineCFGNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	engineCFGCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: brokermodule")
	engineCFGCmd.MarkFlagRequired("parameter")
	engineCFGCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"brokermodule", "nagios_name", "instance", "nagios_activate"}, cobra.ShellCompDirectiveDefault
	})
	engineCFGCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	engineCFGCmd.MarkFlagRequired("value")
	engineCFGCmd.Flags().StringP("operation", "o", "", "To define the operation: add, del")
	engineCFGCmd.MarkFlagRequired("operation")
	engineCFGCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"add", "del"}, cobra.ShellCompDirectiveDefault
	})
}
