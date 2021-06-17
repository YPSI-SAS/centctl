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
	"centctl/resources/trap"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// trapCmd represents the trap command
var trapCmd = &cobra.Command{
	Use:   "trap",
	Short: "List the traps",
	Long:  `List the traps of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListTrap(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListTrap permits to display the array of traps return by the API
func ListTrap(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "trap", "", "list trap", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the traps contain into the response body
	traps := trap.Result{}
	json.Unmarshal(body, &traps)
	finalTraps := traps.Traps
	if regex != "" {
		finalTraps = deleteTrap(finalTraps, regex)
	}

	//Sort traps based on their ID
	sort.SliceStable(finalTraps, func(i, j int) bool {
		valI, _ := strconv.Atoi(finalTraps[i].ID)
		valJ, _ := strconv.Atoi(finalTraps[j].ID)
		return valI < valJ
	})

	//Organization of data
	server := trap.Server{
		Server: trap.Informations{
			Name:  os.Getenv("SERVER"),
			Traps: finalTraps,
		},
	}

	//Display all traps
	displayTrap, err := display.Trap(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTrap)

	return nil
}

func deleteTrap(traps []trap.Trap, regex string) []trap.Trap {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range traps {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			traps[index] = s
			index++
		}
	}
	return traps[:index]
}

func init() {
	trapCmd.Flags().StringP("regex", "r", "", "The regex to apply on the trap's name")
}
