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
	"centctl/display"
	"centctl/request"
	"centctl/resources/broker"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// outputCmd represents the broker output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Show one broker output's details",
	Long:  `Show one broker output's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		id, _ := cmd.Flags().GetInt("id")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowBrokerOutput(name, id, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowBrokerOutput permits to display the details of one broker output
func ShowBrokerOutput(name string, id int, debugV bool, output string) error {
	output = strings.ToLower(output)

	values := name + ";" + strconv.Itoa(id)
	err, body := request.GeneriqueCommandV1Post("getoutput", "CENTBROKERCFG", values, "show broker output", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the broker output contain into the response body
	brokerOutputs := broker.DetailResultOutput{}
	json.Unmarshal(body, &brokerOutputs)

	server := broker.DetailServerOutput{
		Server: broker.DetailInformationsOutput{
			Name:         os.Getenv("SERVER"),
			BrokerOutput: brokerOutputs.BrokerOutputs,
		},
	}

	//Display details of the broker output
	displayBrokerOutput, err := display.DetailBrokerOutput(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBrokerOutput)
	return nil
}

func init() {
	outputCmd.Flags().StringP("name", "n", "", "To define the name of the broker output")
	outputCmd.MarkFlagRequired("name")
	outputCmd.Flags().IntP("id", "i", -1, "To define the id of the broke output")
	outputCmd.MarkFlagRequired("id")
}
