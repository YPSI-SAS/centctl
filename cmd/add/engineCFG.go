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

// engineCFGCmd represents the engineCFG command
var engineCFGCmd = &cobra.Command{
	Use:   "engineCFG",
	Short: "Add an engineCFG",
	Long:  `Add an engineCFG into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		instance, _ := cmd.Flags().GetString("instance")
		comment, _ := cmd.Flags().GetString("comment")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddEngineCFG(name, instance, comment, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddEngineCFG permits to add a engineCFG in the centreon server
func AddEngineCFG(name string, instance string, comment string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + instance + ";" + comment
	err := request.Add("add", "ENGINECFG", values, "add engineCFG", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	engineCFGCmd.Flags().StringP("name", "n", "", "To define the name of the centreon engine configuration")
	engineCFGCmd.MarkFlagRequired("name")
	engineCFGCmd.Flags().StringP("instance", "i", "", "To define the instance that is linked to engine configuration")
	engineCFGCmd.MarkFlagRequired("instance")
	engineCFGCmd.RegisterFlagCompletionFunc("instance", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetPollerNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	engineCFGCmd.Flags().StringP("comment", "c", "", "To define the comment regarding the engine configuration")
	engineCFGCmd.MarkFlagRequired("comment")
}
