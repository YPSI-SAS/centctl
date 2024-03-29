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

// dependencyCmd represents the dependency command
var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Modfiy a dependency",
	Long:  `Modfiy a dependency into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		operation, _ := cmd.Flags().GetString("operation")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyDependency(name, parameter, value, operation, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyDependency permits to modify a dependency in the centreon server
func ModifyDependency(name string, parameter string, value string, operation string, debugV bool, isImport bool, detail bool) error {
	var action string
	var values string
	isDefault := false
	operation = strings.ToLower(operation)

	switch strings.ToLower(parameter) {
	case "parent":
		isDefault = true
	case "child":
		isDefault = true
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value
	}

	if isDefault {
		action = operation + strings.ToLower(parameter)
		values = name + ";" + value
	}

	err := request.Modify(action, "DEP", values, "modify dependency", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	dependencyCmd.Flags().StringP("name", "n", "", "To define the name of the dependency to be modified")
	dependencyCmd.MarkFlagRequired("name")
	dependencyCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetDependencyNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	dependencyCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: parent,child")
	dependencyCmd.MarkFlagRequired("parameter")
	dependencyCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "description", "inherits_parent", "execution_failure_criteria", "notification_failure_criteria", "comment", "parent", "child"}, cobra.ShellCompDirectiveDefault
	})
	dependencyCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	dependencyCmd.MarkFlagRequired("value")
	dependencyCmd.Flags().StringP("operation", "o", "", "To define the operation: add, del")
	dependencyCmd.MarkFlagRequired("operation")
	dependencyCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"add", "del"}, cobra.ShellCompDirectiveDefault
	})
}
