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
package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/engineCFG"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// engineCFGCmd represents the engineCFG command
var engineCFGCmd = &cobra.Command{
	Use:   "engineCFG",
	Short: "List the enginesCFG",
	Long:  `List the enginesCFG of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListEngineCFG(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListEngineCFG permits to display the array of enginesCFG return by the API
func ListEngineCFG(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ENGINECFG", "", "list engineCFG", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the enginesCFG contain into the response body
	enginesCFGs := engineCFG.ResultEngineCFG{}
	json.Unmarshal(body, &enginesCFGs)
	finalEngineCFGs := enginesCFGs.EngineCFG
	if regex != "" {
		finalEngineCFGs = deleteEngineCFG(finalEngineCFGs, regex)
	}

	//Sort enginesCFG based on their ID
	sort.SliceStable(finalEngineCFGs, func(i, j int) bool {
		return strings.ToLower(finalEngineCFGs[i].Name) < strings.ToLower(finalEngineCFGs[j].Name)
	})

	//Organization of data
	server := engineCFG.ServerEngineCFG{
		Server: engineCFG.InformationsEngineCFG{
			Name:      os.Getenv("SERVER"),
			EngineCFG: finalEngineCFGs,
		},
	}

	//Display all enginesCFG
	displayEngineCFG, err := display.EngineCFG(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayEngineCFG)

	return nil
}

func deleteEngineCFG(engineCFGs []engineCFG.EngineCFG, regex string) []engineCFG.EngineCFG {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range engineCFGs {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			engineCFGs[index] = s
			index++
		}
	}
	return engineCFGs[:index]
}

func init() {
	engineCFGCmd.Flags().StringP("regex", "r", "", "The regex to apply on the engineCFG's name")
}
