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

	"github.com/spf13/cobra"
)

// cfgCmd represents the brokerCFG command
var cfgCmd = &cobra.Command{
	Use:   "CFG",
	Short: "Add a broker CFG",
	Long:  `Add a broker CFG into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		instance, _ := cmd.Flags().GetString("instance")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddBrokerCFG(name, instance, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddBrokerCFG permits to add a brokerCFG in the centreon server
func AddBrokerCFG(name string, instance string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + instance
	err := request.Add("add", "CENTBROKERCFG", values, "add broker CFG", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	cfgCmd.Flags().StringP("name", "n", "", "To define the name of the broker CFG")
	cfgCmd.MarkFlagRequired("name")
	cfgCmd.Flags().StringP("instance", "i", "", "To define the instance that is linked to broker CFG")
	cfgCmd.MarkFlagRequired("instance")
}
