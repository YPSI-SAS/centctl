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

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "List broker output",
	Long:  `List broker output of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		broker, _ := cmd.Flags().GetString("broker")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListBrokerOutput(broker, output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListBrokerOutput permits to display the array of broker output return by the API
func ListBrokerOutput(brokerName string, output string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("listoutput", "CENTBROKERCFG", brokerName, "list broker output", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the broker output contain into the response body
	brokerOutput := broker.ResultOutput{}
	json.Unmarshal(body, &brokerOutput)

	//Sort broker output based on their ID
	sort.SliceStable(brokerOutput.BrokerOutputs, func(i, j int) bool {
		return strings.ToLower(brokerOutput.BrokerOutputs[i].Name) < strings.ToLower(brokerOutput.BrokerOutputs[j].Name)
	})

	server := broker.ServerOutput{
		Server: broker.InformationsOutput{
			Name:          os.Getenv("SERVER"),
			BrokerOutputs: brokerOutput.BrokerOutputs,
		},
	}

	//Display all broker output
	displayBrokerOutput, err := display.BrokerOutput(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBrokerOutput)

	return nil
}

func init() {
	outputCmd.Flags().StringP("broker", "b", "", "To define the name of the broker CFG")
	outputCmd.MarkFlagRequired("broker")
}
