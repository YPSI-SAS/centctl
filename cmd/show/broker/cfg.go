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
	"strings"

	"github.com/spf13/cobra"
)

// cfgCmd represents the broker CFG command
var cfgCmd = &cobra.Command{
	Use:   "CFG",
	Short: "Show one broker CFG's details",
	Long:  `Show one broker CFG's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowBrokerCFG(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowBrokerCFG permits to display the details of one broker CFG
func ShowBrokerCFG(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "CENTBROKERCFG", name, "show broker CFG", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the brokerCGFs contain into the response body
	brokerCGFs := broker.DetailResultCFG{}
	json.Unmarshal(body, &brokerCGFs)

	//Permits to find the good broker CFG in the array
	var BrokerCFGFind broker.DetailBrokerCFG
	for _, v := range brokerCGFs.BrokerCFGs {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			BrokerCFGFind = v
		}
	}

	var server broker.DetailServerCFG
	if BrokerCFGFind.Name != "" {
		//Organization of data
		server = broker.DetailServerCFG{
			Server: broker.DetailInformationsCFG{
				Name:      os.Getenv("SERVER"),
				BrokerCFG: &BrokerCFGFind,
			},
		}
	} else {
		server = broker.DetailServerCFG{
			Server: broker.DetailInformationsCFG{
				Name:      os.Getenv("SERVER"),
				BrokerCFG: nil,
			},
		}
	}

	//Display details of the broker CFG
	displayBrokerCFG, err := display.DetailBrokerCFG(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBrokerCFG)
	return nil
}

func init() {
	cfgCmd.Flags().StringP("name", "n", "", "To define the name of the broker CFG")
	cfgCmd.MarkFlagRequired("name")
	cfgCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetBrokerCFGNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
