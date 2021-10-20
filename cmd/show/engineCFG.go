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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/engineCFG"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// engineCFGCmd represents the engineCFG command
var engineCFGCmd = &cobra.Command{
	Use:   "engineCFG",
	Short: "Show one engineCFG's details",
	Long:  `Show one engineCFG's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowEngineCFG(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowEngineCFG permits to display the details of one cooleanrule
func ShowEngineCFG(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "engineCFG", name, "show engineCFG", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the engineCFGs contain into the response body
	engineCFGs := engineCFG.DetailResultEngineCFG{}
	json.Unmarshal(body, &engineCFGs)

	//Permits to find the good engineCFG in the array
	var EngineCFGFind engineCFG.DetailEngineCFG
	for _, v := range engineCFGs.EngineCFG {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			EngineCFGFind = v
		}
	}

	var server engineCFG.DetailServerEngineCFG
	if EngineCFGFind.Name != "" {
		//Organization of data
		server = engineCFG.DetailServerEngineCFG{
			Server: engineCFG.DetailInformationsEngineCFG{
				Name:      os.Getenv("SERVER"),
				EngineCFG: &EngineCFGFind,
			},
		}
	} else {
		server = engineCFG.DetailServerEngineCFG{
			Server: engineCFG.DetailInformationsEngineCFG{
				Name:      os.Getenv("SERVER"),
				EngineCFG: nil,
			},
		}
	}

	//Display details of the engineCFG
	displayEngineCFG, err := display.DetailEngineCFG(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayEngineCFG)
	return nil
}

func init() {
	engineCFGCmd.Flags().StringP("name", "n", "", "To define the name of the engineCFG")
	engineCFGCmd.MarkFlagRequired("name")
}
