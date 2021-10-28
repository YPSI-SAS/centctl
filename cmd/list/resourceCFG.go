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
	"centctl/resources/resourceCFG"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// resourceCFGCmd represents the resourceCFG command
var resourceCFGCmd = &cobra.Command{
	Use:   "resourceCFG",
	Short: "List the resourceCFG",
	Long:  `List the resourceCFG of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListResourceCFG(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListResourceCFG permits to display the array of resourceCFG return by the API
func ListResourceCFG(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "RESOURCECFG", "", "list resourceCFG", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the resourceCFG contain into the response body
	resourceCFGs := resourceCFG.Result{}
	json.Unmarshal(body, &resourceCFGs)
	finalResourceCFG := resourceCFGs.ResourceCFG
	if regex != "" {
		finalResourceCFG = deleteResourceCFG(finalResourceCFG, regex)
	}

	//Sort resourceCFG based on their ID
	sort.SliceStable(finalResourceCFG, func(i, j int) bool {
		return strings.ToLower(finalResourceCFG[i].Name) < strings.ToLower(finalResourceCFG[j].Name)
	})

	//Organization of data
	server := resourceCFG.Server{
		Server: resourceCFG.Informations{
			Name:        os.Getenv("SERVER"),
			ResourceCFG: finalResourceCFG,
		},
	}

	//Display all resourceCFG
	displayResourceCFG, err := display.ResourceCFG(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayResourceCFG)

	return nil
}

func deleteResourceCFG(resourceCFGs []resourceCFG.ResourceCFG, regex string) []resourceCFG.ResourceCFG {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range resourceCFGs {
		name := strings.ReplaceAll(s.Name, "$", "")
		matched, err := regexp.MatchString(regex, name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			resourceCFGs[index] = s
			index++
		}
	}
	return resourceCFGs[:index]
}

func init() {
	resourceCFGCmd.Flags().StringP("regex", "r", "", "The regex to apply on the resourceCFG's name")
}
