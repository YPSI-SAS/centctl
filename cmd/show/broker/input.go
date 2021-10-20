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

// inputCmd represents the broker input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Show one broker input's details",
	Long:  `Show one broker input's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		id, _ := cmd.Flags().GetInt("id")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowBrokerInput(name, id, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowBrokerInput permits to display the details of one broker input
func ShowBrokerInput(name string, id int, debugV bool, output string) error {
	output = strings.ToLower(output)

	values := name + ";" + strconv.Itoa(id)
	err, body := request.GeneriqueCommandV1Post("getinput", "CENTBROKERCFG", values, "show broker input", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the broker inputs contain into the response body
	brokerInputs := broker.DetailResultInput{}
	json.Unmarshal(body, &brokerInputs)

	server := broker.DetailServerInput{
		Server: broker.DetailInformationsInput{
			Name:        os.Getenv("SERVER"),
			BrokerInput: brokerInputs.BrokerInputs,
		},
	}

	//Display details of the broker input
	displayBrokerInput, err := display.DetailBrokerInput(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBrokerInput)
	return nil
}

func init() {
	inputCmd.Flags().StringP("name", "n", "", "To define the name of the broker input")
	inputCmd.MarkFlagRequired("name")
	inputCmd.Flags().IntP("id", "i", -1, "To define the id of the broke input")
	inputCmd.MarkFlagRequired("id")
}
