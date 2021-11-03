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
package template

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Delete template contact",
	Long:  `Delete template contact of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		alias, _ := cmd.Flags().GetString("alias")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := DeleteTemplateContact(alias, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DeleteTemplateContact permits to delete a service template in the centreon server
func DeleteTemplateContact(alias string, debugV bool) error {
	err := request.Delete("del", "CONTACTTPL", alias, "delete template contact", alias, debugV, false, "")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	contactCmd.Flags().StringP("alias", "a", "", "To define the contact template which will delete")
	contactCmd.MarkFlagRequired("alias")
	contactCmd.RegisterFlagCompletionFunc("alias", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetTemplateContactAlias()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
