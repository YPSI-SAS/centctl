/*
MIT License

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
	Short: "Delete a broker logger",
	Long:  `Delete a broker logger into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		broker, _ := cmd.Flags().GetString("broker")
		id, _ := cmd.Flags().GetInt("ID")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := DeleteBrokerLogger(broker, id, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DeleteBrokerLogger permits to delete a broker logger in the centreon server
func DeleteBrokerLogger(broker string, id int, debugV bool) error {
	values := broker + ";" + strconv.Itoa(id)
	err := request.Delete("dellogger", "CENTBROKERCFG", values, "delete broker logger", broker+" logger "+strconv.Itoa(id), debugV, false, "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	loggerCmd.Flags().StringP("broker", "b", "", "To define the name of the broker CFG")
	loggerCmd.MarkFlagRequired("broker")
	loggerCmd.Flags().IntP("ID", "i", -1, "To define the I/O ID of the object which will delete")
	loggerCmd.MarkFlagRequired("ID")
}
