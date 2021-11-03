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
package export

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/trap"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// trapCmd represents the trap command
var trapCmd = &cobra.Command{
	Use:   "trap",
	Short: "Export trap",
	Long:  `Export trap of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportTrap(name, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportTrap permits to export a trap of the centreon server
func ExportTrap(name []string, regex string, file string, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex ")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		templates := getAllTrap(debugV)
		for _, a := range templates {
			if regex != "" {
				matched, err := regexp.MatchString(regex, a.Name)
				if err != nil {
					fmt.Printf(colorRed, "ERROR:")
					fmt.Println(err.Error())
					os.Exit(1)
				}
				if matched {
					name = append(name, a.Name)
				}
			} else {
				name = append(name, a.Name)
			}
		}
	}
	for _, n := range name {
		err, trap := getTrapInfo(n, debugV)
		if err != nil {
			return err
		}
		if trap.Name == "" {
			continue
		}

		//Write trap informations
		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,trap,\""+trap.Name+"\",\""+trap.Oid+"\"\n", file, writeFile)
		request.WriteValues("modify,trap,\""+trap.Name+"\",vendor,\""+trap.Manufacturer+"\"\n", file, writeFile)

		//Write Matchings information
		if len(trap.Matchings) != 0 {
			for _, b := range trap.Matchings {
				request.WriteValues("modify,trap,\""+trap.Name+"\",matching,\""+b.String+";"+b.Regexp+";"+b.Status+"\"\n", file, writeFile)
			}
		}
	}
	return nil
}

//The arguments impossible to get : all elements in setparam table
//getTrapInfo permits to get all informations about a trap
func getTrapInfo(name string, debugV bool) (error, trap.ExportTrap) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "trap", name, "export trap", debugV, false, "")
	if err != nil {
		return err, trap.ExportTrap{}
	}
	var resultTrap trap.ExportResultTrap
	json.Unmarshal(body, &resultTrap)

	trapFind := trap.ExportTrap{}
	find := false
	for _, g := range resultTrap.Traps {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			trapFind = g
			find = true
		}
	}
	//Check if the trap  is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, trapFind
	}

	//Get the BA of the trap
	err, body = request.GeneriqueCommandV1Post("getmatching", "trap", name, "export trap", debugV, false, "")
	if err != nil {
		return err, trap.ExportTrap{}
	}
	var resultMatchings trap.ExportResultTrapMatching
	json.Unmarshal(body, &resultMatchings)

	trapFind.Matchings = resultMatchings.Matchings

	return nil, trapFind

}

//getAllTrap permits to find all trap in the centreon server
func getAllTrap(debugV bool) []trap.ExportTrap {
	//Get all trap
	err, body := request.GeneriqueCommandV1Post("show", "trap", "", "export trap", debugV, false, "")
	if err != nil {
		return []trap.ExportTrap{}
	}
	var resultTrap trap.ExportResultTrap
	json.Unmarshal(body, &resultTrap)

	return resultTrap.Traps
}

func init() {
	trapCmd.Flags().StringSliceP("name", "n", []string{}, "trap's name (separate by a comma the multiple values)")
	trapCmd.Flags().StringP("regex", "r", "", "The regex to apply on the trap's name")

}
