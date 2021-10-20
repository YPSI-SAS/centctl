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
	"centctl/display"
	"centctl/request"
	"centctl/resources/broker"

	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// loggerCmd represents the logger command
var loggerCmd = &cobra.Command{
	Use:   "logger",
	Short: "List broker logger",
	Long:  `List broker logger of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		broker, _ := cmd.Flags().GetString("broker")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListBrokerLogger(broker, output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListBrokerLogger permits to display the array of broker logger return by the API
func ListBrokerLogger(brokerName string, output string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("listlogger", "CENTBROKERCFG", brokerName, "list broker logger", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the broker logger contain into the response body
	brokerLogger := broker.ResultLogger{}
	json.Unmarshal(body, &brokerLogger)

	//Sort broker logger  based on their ID
	sort.SliceStable(brokerLogger.BrokerLoggers, func(i, j int) bool {
		return strings.ToLower(brokerLogger.BrokerLoggers[i].Name) < strings.ToLower(brokerLogger.BrokerLoggers[j].Name)
	})

	server := broker.ServerLogger{
		Server: broker.InformationsLogger{
			Name:          os.Getenv("SERVER"),
			BrokerLoggers: brokerLogger.BrokerLoggers,
		},
	}

	//Display all broker logger
	displayBrokerLogger, err := display.BrokerLogger(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBrokerLogger)

	return nil
}

func init() {
	loggerCmd.Flags().StringP("broker", "b", "", "To define the name of the broker CFG")
	loggerCmd.MarkFlagRequired("broker")
}
