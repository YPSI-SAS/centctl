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

package broker

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Add a broker input",
	Long:  `Add a broker input into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		broker, _ := cmd.Flags().GetString("broker")
		objectName, _ := cmd.Flags().GetString("objectName")
		objectNature, _ := cmd.Flags().GetString("objectNature")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddBrokerInput(broker, objectName, objectNature, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddBrokerInput permits to add a broker input in the centreon server
func AddBrokerInput(broker string, objectName string, objectNature string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := broker + ";" + objectName + ";" + objectNature
	err := request.Add("addinput", "CENTBROKERCFG", values, "add broker input", broker+" input "+objectName, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	inputCmd.Flags().StringP("broker", "b", "", "To define the name of the broker CFG")
	inputCmd.MarkFlagRequired("broker")
	inputCmd.Flags().StringP("objectName", "o", "", "To define the broker of the I/O object")
	inputCmd.MarkFlagRequired("objectName")
	inputCmd.Flags().StringP("objectNature", "n", "", "To define the nature of the I/O object")
	inputCmd.MarkFlagRequired("objectNature")
}
