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

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "List broker input",
	Long:  `List broker input of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		broker, _ := cmd.Flags().GetString("broker")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListBrokerInput(broker, output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListBrokerInput permits to display the array of broker input return by the API
func ListBrokerInput(brokerName string, output string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("listinput", "CENTBROKERCFG", brokerName, "list broker input", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the broker input contain into the response body
	brokerInput := broker.ResultInput{}
	json.Unmarshal(body, &brokerInput)

	//Sort broker input  based on their ID
	sort.SliceStable(brokerInput.BrokerInputs, func(i, j int) bool {
		return strings.ToLower(brokerInput.BrokerInputs[i].Name) < strings.ToLower(brokerInput.BrokerInputs[j].Name)
	})

	server := broker.ServerInput{
		Server: broker.InformationsInput{
			Name:         os.Getenv("SERVER"),
			BrokerInputs: brokerInput.BrokerInputs,
		},
	}

	//Display all borker input
	displayBrokerInput, err := display.BrokerInput(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBrokerInput)

	return nil
}

func init() {
	inputCmd.Flags().StringP("broker", "b", "", "To define the name of the broker CFG")
	inputCmd.MarkFlagRequired("broker")
}
