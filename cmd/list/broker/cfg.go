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
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/broker"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// cfgCmd represents the CFG command
var cfgCmd = &cobra.Command{
	Use:   "CFG",
	Short: "List broker CFG",
	Long:  `List broker CFG of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListBrokerCFG(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListBrokerCFG permits to display the array of broker CFG return by the API
func ListBrokerCFG(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "CENTBROKERCFG", "", "list broker CFG", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the broker CFG contain into the response body
	brokerCFG := broker.ResultCFG{}
	json.Unmarshal(body, &brokerCFG)
	finalBrokerCFG := brokerCFG.BrokerCFGs
	if regex != "" {
		finalBrokerCFG = deleteBrokerCFG(finalBrokerCFG, regex)
	}

	//Sort broker CFG  based on their ID
	sort.SliceStable(finalBrokerCFG, func(i, j int) bool {
		return strings.ToLower(finalBrokerCFG[i].Name) < strings.ToLower(finalBrokerCFG[j].Name)
	})

	server := broker.ServerCFG{
		Server: broker.InformationsCFG{
			Name:       os.Getenv("SERVER"),
			BrokerCFGs: finalBrokerCFG,
		},
	}

	//Display all broker CFG
	displayBrokerCFG, err := display.BrokerCFG(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBrokerCFG)

	return nil
}

func deleteBrokerCFG(borkerCFGs []broker.BrokerCFG, regex string) []broker.BrokerCFG {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range borkerCFGs {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			borkerCFGs[index] = s
			index++
		}
	}
	return borkerCFGs[:index]
}

func init() {
	cfgCmd.Flags().StringP("regex", "r", "", "The regex to apply on the broker cfg's name")

}
