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
LIABILITY, WHETHER IN AN menu OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package acl

import (
	"centctl/request"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Modfiy a ACL menu",
	Long:  `Modfiy a ACL menu into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := ModifyACLMenu(name, parameter, value, debugV, apply, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyACLMenu permits to modify a ACL menu in the centreon server
func ModifyACLMenu(name string, parameter string, value string, debugV bool, apply bool, isImport bool, detail bool) error {
	var action string
	var values string

	switch strings.ToLower(parameter) {
	case "grantrw":
		action = "grantrw"
		values = name + ";1;" + value
	case "grantro":
		action = "grantro"
		values = name + ";1;" + value
	case "revoke":
		action = "revoke"
		values = name + ";1;" + value
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value

	}

	err := request.Modify(action, "ACLMENU", values, "modify ACL menu", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	menuCmd.Flags().StringP("name", "n", "", "To define the name of the ACL menu to be modified")
	menuCmd.MarkFlagRequired("name")
	menuCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: grantrw, grantro,revoke")
	menuCmd.MarkFlagRequired("parameter")
	menuCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is grantrw or grantro or revoke the value must be of the form : menuName or menuName;subMenuName")
	menuCmd.MarkFlagRequired("value")
	menuCmd.Flags().Bool("apply", false, "Export configuration of the poller")

}
