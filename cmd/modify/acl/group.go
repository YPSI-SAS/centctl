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
package acl

import (
	"centctl/colorMessage"
	"centctl/request"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Modfiy a ACL group",
	Long:  `Modfiy a ACL group into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		operation, _ := cmd.Flags().GetString("operation")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := ModifyACLGroup(name, parameter, value, operation, debugV, apply, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyACLGroup permits to modify a ACL group in the centreon server
func ModifyACLGroup(name string, parameter string, value string, operation string, debugV bool, apply bool, isImport bool, detail bool) error {
	colorRed := colorMessage.GetColorRed()
	var action string
	var values string
	isDefault := false

	operation = strings.ToLower(operation)
	if operation != "add" && operation != "del" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The operation's value must be : add or del")
		os.Exit(1)
	}

	switch strings.ToLower(parameter) {
	case "action":
		isDefault = true
	case "menu":
		isDefault = true
	case "resource":
		isDefault = true
	case "contactgroup":
		isDefault = true
	case "contact":
		isDefault = true
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value

	}

	if isDefault {
		action = operation + strings.ToLower(parameter)
		values = name + ";" + value
	}

	err := request.Modify(action, "ACLGROUP", values, "modify ACL group", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	groupCmd.Flags().StringP("name", "n", "", "To define the name of the ACL group to be modified")
	groupCmd.MarkFlagRequired("name")
	groupCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetACLGroupNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	groupCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: action,menu,resource,contact,contactgroup")
	groupCmd.MarkFlagRequired("parameter")
	groupCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "alias", "activate", "action", "menu", "resource", "contact", "contactgroup"}, cobra.ShellCompDirectiveDefault
	})
	groupCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	groupCmd.MarkFlagRequired("value")
	groupCmd.Flags().Bool("apply", false, "Export configuration of the poller")
	groupCmd.Flags().StringP("operation", "o", "", "To define the operation: add, del")
	groupCmd.MarkFlagRequired("operation")
	groupCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"add", "del"}, cobra.ShellCompDirectiveDefault
	})
}
