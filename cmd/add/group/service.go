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

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add group service",
	Long:  `Add group service of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		alias, _ := cmd.Flags().GetString("alias")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddGroupService(name, alias, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddGroupService permits to add a service group in the centreon server
func AddGroupService(name string, alias string, debugV bool, isImport bool) error {
	//Get values for POST
	values := name + ";" + alias

	err := request.Add("add", "SG", values, "add group service", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "To define the name of the service group")
	serviceCmd.MarkFlagRequired("name")
	serviceCmd.Flags().StringP("alias", "a", "", "To define the alias of the service group")
	serviceCmd.MarkFlagRequired("alias")
}
