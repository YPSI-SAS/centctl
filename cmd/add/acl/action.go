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
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Add ACL action",
	Long:  `Add ACL action of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddACLAction(name, description, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddACLAction permits to add a ACL action in the centreon server
func AddACLAction(name string, description string, debugV bool, isImport bool) error {
	//Get values for POST
	values := name + ";" + description

	err := request.Add("add", "ACLACTION", values, "add acl action", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	actionCmd.Flags().StringP("name", "n", "", "To define the name of the ACL action")
	actionCmd.MarkFlagRequired("name")
	actionCmd.Flags().StringP("description", "d", "", "To define the description of the ACL action")
	actionCmd.MarkFlagRequired("description")
}
