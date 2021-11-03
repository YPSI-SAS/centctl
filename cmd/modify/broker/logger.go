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

// loggerCmd represents the logger command
var loggerCmd = &cobra.Command{
	Use:   "logger",
	Short: "Modfiy a broker logger",
	Long:  `Modfiy a broker logger into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		id, _ := cmd.Flags().GetInt("id")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyBrokerLogger(name, id, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyBrokerLogger permits to modify a borker logger in the centreon server
func ModifyBrokerLogger(name string, id int, parameter string, value string, debugV bool, isImport bool, detail bool) error {

	values := name + ";" + strconv.Itoa(id) + ";" + parameter + ";" + value
	err := request.Modify("setlogger", "CENTBROKERCFG", values, "modify broker logger", name+" logger "+strconv.Itoa(id), parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	loggerCmd.Flags().StringP("name", "n", "", "To define the name of the broker CFG")
	loggerCmd.MarkFlagRequired("name")
	loggerCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetBrokerCFGNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	loggerCmd.Flags().IntP("id", "i", -1, "To define the I/O object to be modified")
	loggerCmd.MarkFlagRequired("id")
	loggerCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation.")
	loggerCmd.MarkFlagRequired("parameter")
	loggerCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. ")
	loggerCmd.MarkFlagRequired("value")
	loggerCmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if loggerCmd.Flag("name").Value.String() != "" {
			if request.InitAuthentification(cmd) {
				values = request.GetBrokerLoggerID(loggerCmd.Flag("name").Value.String())
			}
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
