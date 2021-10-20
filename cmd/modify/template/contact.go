/*MIT License

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
package template

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Modify template contact",
	Long:  `Modify template contact of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyTemplateContact(name, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyTemplateContact permits to modify a contact template in the centreon server
func ModifyTemplateContact(name string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	//Creation of the request body
	values := name + ";" + parameter + ";" + value

	err := request.Modify("setparam", "CONTACTTPL", values, "modify template contact", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	contactCmd.Flags().StringP("name", "n", "", "To define the name of the contact template to be modified")
	contactCmd.MarkFlagRequired("name")
	contactCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation.")
	contactCmd.MarkFlagRequired("parameter")
	contactCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	contactCmd.MarkFlagRequired("value")
}
