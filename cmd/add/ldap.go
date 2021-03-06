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

package add

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// ldapCmd represents the LDAP command
var ldapCmd = &cobra.Command{
	Use:   "LDAP",
	Short: "Add a LDAP configuration",
	Long:  `Add a LDAP configuration into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddLDAP(name, description, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddLDAP permits to add a LDAP in the centreon server
func AddLDAP(name string, description string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + description
	err := request.Add("add", "LDAP", values, "add LDAP", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	ldapCmd.Flags().StringP("name", "n", "", "To define the name of the LDAP configuration")
	ldapCmd.MarkFlagRequired("name")
	ldapCmd.Flags().StringP("description", "d", "", "To define the description of the LDAP configuration")
	ldapCmd.MarkFlagRequired("description")
}
