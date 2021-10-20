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
package export

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/resourceCFG"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// resourceCFGCmd represents the resourceCFG command
var resourceCFGCmd = &cobra.Command{
	Use:   "resourceCFG",
	Short: "Export resourceCFG",
	Long:  `Export resourceCFG of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportResourceCFG(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportResourceCFG permits to export a resourceCFG of the centreon server
func ExportResourceCFG(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
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
		templates := getAllResourceCFG(debugV)
		for _, a := range templates {
			if regex != "" {

				matched, err := regexp.MatchString(regex, strings.ReplaceAll(a.Name, "$", ""))
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
		err, resourceCFG := getResourceCFGInfo(n, debugV)
		if err != nil {
			return err
		}
		if resourceCFG.Name == "" {
			continue
		}

		//Write resourceCFG informations
		if len(resourceCFG.Instance) > 0 {
			switch resourceCFG.Instance[0] {
			case '"':
				if err := json.Unmarshal(resourceCFG.Instance, &resourceCFG.PollerFinal); err != nil {
					return err
				}
			case '[':
				var s []string
				if err := json.Unmarshal(resourceCFG.Instance, &s); err != nil {
					return err
				}
				resourceCFG.PollerFinal = strings.Join(s, "|")
			}
		}
		resourceCFG.Instance = json.RawMessage{}
		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,resourceCFG,\"" + resourceCFG.Name + "\",\"" + resourceCFG.Value + "\",\"" + resourceCFG.PollerFinal + "\",\"" + resourceCFG.Comment + "\"\n")

	}

	return nil
}

//The arguments impossible to get : element in setparam table
//getResourceCFGInfo permits to get all informations about a resourceCFG
func getResourceCFGInfo(name string, debugV bool) (error, resourceCFG.ExportResourceCFG) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "resourcecfg", name, "export resourceCFG", debugV, false, "")
	if err != nil {
		return err, resourceCFG.ExportResourceCFG{}
	}
	var resultResourceCFG resourceCFG.ExportResultResourceCFG
	json.Unmarshal(body, &resultResourceCFG)

	resourceCFG := resourceCFG.ExportResourceCFG{}
	find := false
	for _, g := range resultResourceCFG.ResourceCFGs {
		if !strings.Contains(name, "$") {
			name = "$" + name + "$"
		}
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			resourceCFG = g
			find = true
		}
	}
	//Check if the resourceCFG is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, resourceCFG
	}

	return nil, resourceCFG

}

//getAllResourceCFG permits to find all resourceCFG in the centreon server
func getAllResourceCFG(debugV bool) []resourceCFG.ExportResourceCFG {
	//Get all resourceCFG
	err, body := request.GeneriqueCommandV1Post("show", "resourcecfg", "", "export resourceCFG", debugV, false, "")
	if err != nil {
		return []resourceCFG.ExportResourceCFG{}
	}
	var resultResourceCFG resourceCFG.ExportResultResourceCFG
	json.Unmarshal(body, &resultResourceCFG)

	return resultResourceCFG.ResourceCFGs
}

func init() {
	resourceCFGCmd.Flags().StringSliceP("name", "n", []string{}, "resourceCFG's name (separate by a comma the multiple values)")
	resourceCFGCmd.Flags().StringP("file", "f", "ExportResourceCFG.csv", "To define the name of the csv file")
	resourceCFGCmd.Flags().StringP("regex", "r", "", "The regex to apply on the resourceCFG's name")

}
