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
package broker

import (
	"centctl/request"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Modfiy a broker output",
	Long:  `Modfiy a broker output into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		id, _ := cmd.Flags().GetInt("id")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyBrokerOutput(name, id, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyBrokerOutput permits to modify a borker output in the centreon server
func ModifyBrokerOutput(name string, id int, parameter string, value string, debugV bool, isImport bool, detail bool) error {

	values := name + ";" + strconv.Itoa(id) + ";" + parameter + ";" + value
	err := request.Modify("setoutput", "CENTBROKERCFG", values, "modify broker output", name+" output "+strconv.Itoa(id), parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	outputCmd.Flags().StringP("name", "n", "", "To define the name of the broker CFG")
	outputCmd.MarkFlagRequired("name")
	outputCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetBrokerCFGNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	outputCmd.Flags().IntP("id", "i", -1, "To define the I/O object to be modified")
	outputCmd.MarkFlagRequired("id")
	outputCmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if outputCmd.Flag("name").Value.String() != "" {
			if request.InitAuthentification(cmd) {
				values = request.GetBrokerOutputID(outputCmd.Flag("name").Value.String())
			}
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	outputCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation.")
	outputCmd.MarkFlagRequired("parameter")
	outputCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. ")
	outputCmd.MarkFlagRequired("value")

}
