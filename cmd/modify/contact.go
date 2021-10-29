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

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Modfiy a contact",
	Long:  `Modfiy a contact into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyContact(name, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyContact permits to modify a contact in the centreon server
func ModifyContact(name string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	//Creation of the request body
	values := name + ";" + parameter + ";" + value

	err := request.Modify("setparam", "contact", values, "modify contact", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	contactCmd.Flags().StringP("name", "n", "", "To define the name of the contact to be modified")
	contactCmd.MarkFlagRequired("name")
	contactCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation.")
	contactCmd.MarkFlagRequired("parameter")
	contactCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "alias", "comment", "email", "password", "access", "language", "admin", "authtype", "hostnotifcmd", "svcnotifcmd", "hostnotifperiod", "svcnotifperiod", "hostnotifopt", "servicenotifopt", "address1", "address2", "address3", "address4", "address5", "address6", "ldap_dn", "enable_notifications", "autologin_key", "template", "timezone"}, cobra.ShellCompDirectiveDefault
	})
	contactCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	contactCmd.MarkFlagRequired("value")
}
