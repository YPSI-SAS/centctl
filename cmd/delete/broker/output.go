/*
MIT License

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
	"strconv"

	"github.com/spf13/cobra"
)

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Delete a broker output",
	Long:  `Delete a broker output into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		broker, _ := cmd.Flags().GetString("broker")
		id, _ := cmd.Flags().GetInt("ID")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := DeleteBrokerOutput(broker, id, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DeleteBrokerOutput permits to delete a broker output in the centreon server
func DeleteBrokerOutput(broker string, id int, debugV bool) error {
	values := broker + ";" + strconv.Itoa(id)
	err := request.Delete("deloutput", "CENTBROKERCFG", values, "delete broker output", broker+" output "+strconv.Itoa(id), debugV, false, "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	outputCmd.Flags().StringP("broker", "b", "", "To define the name of the broker CFG")
	outputCmd.MarkFlagRequired("broker")
	outputCmd.Flags().IntP("ID", "i", -1, "To define the I/O ID of the object which will delete")
	outputCmd.MarkFlagRequired("ID")
}
