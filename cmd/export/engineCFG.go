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
	"centctl/resources/engineCFG"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// engineCFGCmd represents the engineCFG command
var engineCFGCmd = &cobra.Command{
	Use:   "engineCFG",
	Short: "Export engineCFG",
	Long:  `Export engineCFG of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportEngineCFG(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportEngineCFG permits to export a engineCFG of the centreon server
func ExportEngineCFG(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex ")
		os.Exit(1)
	}

	//Check if the name of file contains the extension
	if !strings.Contains(file, ".csv") {
		file = file + ".csv"
	}

	//Create the file
	var f *os.File
	var err error
	if appendFile {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		f, err = os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	}
	defer f.Close()
	if err != nil {
		return err
	}

	if all || regex != "" {
		templates := getAllEngineCFG(debugV)
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
		err, engineCFG := getEngineCFGInfo(n, debugV)
		if err != nil {
			return err
		}
		if engineCFG.Name == "" {
			continue
		}

		//Write engineCFG informations
		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,engineCFG,\"" + engineCFG.Name + "\",\"" + engineCFG.Instance + "\",\"" + engineCFG.Comment + "\"\n")

	}

	return nil
}

//The arguments impossible to get : all elements except name, instance and comment
//getEngineCFGInfo permits to get all informations about a engineCFG
func getEngineCFGInfo(name string, debugV bool) (error, engineCFG.ExportEngineCFG) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "engineCFG", name, "export engineCFG", debugV, false, "")
	if err != nil {
		return err, engineCFG.ExportEngineCFG{}
	}
	var resultEngineCFG engineCFG.ExportResultEngineCFG
	json.Unmarshal(body, &resultEngineCFG)

	engineCFG := engineCFG.ExportEngineCFG{}
	find := false
	for _, g := range resultEngineCFG.EngineCFGs {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			engineCFG = g
			find = true
		}
	}
	//Check if the engineCFG  is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, engineCFG
	}

	return nil, engineCFG

}

//getAllEngineCFG permits to find all engineCFG in the centreon server
func getAllEngineCFG(debugV bool) []engineCFG.ExportEngineCFG {
	//Get all engineCFG
	err, body := request.GeneriqueCommandV1Post("show", "engineCFG", "", "export engineCFG", debugV, false, "")
	if err != nil {
		return []engineCFG.ExportEngineCFG{}
	}
	var resultEngineCFG engineCFG.ExportResultEngineCFG
	json.Unmarshal(body, &resultEngineCFG)

	return resultEngineCFG.EngineCFGs
}

func init() {
	engineCFGCmd.Flags().StringSliceP("name", "n", []string{}, "engineCFG's name (separate by a comma the multiple values)")
	engineCFGCmd.Flags().StringP("file", "f", "ExportEngineCFG.csv", "To define the name of the csv file")
	engineCFGCmd.Flags().StringP("regex", "r", "", "The regex to apply on the engineCFG's name")

}
