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
package category

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Export category host",
	Long:  `Export category host of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportCategoryHost(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportCategoryHost permits to export a category host of the centreon server
func ExportCategoryHost(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
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
		categories := getAllCategoryHost(debugV)
		for _, a := range categories {
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
		err, category := getCategoryHostInfo(n, debugV)
		if err != nil {
			return err
		}
		if category.Name == "" {
			continue
		}

		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,categoryhost,\"" + category.Name + "\",\"" + category.Alias + "\"\n")

		//Write in the file the members
		if len(category.Member) != 0 {
			for _, m := range category.Member {
				_, _ = f.WriteString("modify,categoryhost,\"" + category.Name + "\",member,\"" + m.Name + "\"\n")
			}
		}
	}
	return nil
}

//The arguments impossible to get : activate,comment
//getCategoryHostInfo permits to get all informations about a host category
func getCategoryHostInfo(name string, debugV bool) (error, host.ExportCategory) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the host category
	err, body := request.GeneriqueCommandV1Post("show", "HC", name, "export category host", debugV, false, "")
	if err != nil {
		return err, host.ExportCategory{}
	}
	var resultCategory host.ExportResultCategory
	json.Unmarshal(body, &resultCategory)

	//Get the category
	find := false
	category := host.ExportCategory{}
	for _, g := range resultCategory.CategoryHosts {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			category = g
			find = true
		}
	}

	//Check if the host category is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, category
	}

	//Get the members of the host category
	err, body = request.GeneriqueCommandV1Post("getmember", "HC", name, "export category host", debugV, false, "")
	if err != nil {
		return err, host.ExportCategory{}
	}
	var resultMember host.ExportResultCategoryMember
	json.Unmarshal(body, &resultMember)

	category.Member = resultMember.GroupMember

	return nil, category

}

//getAllCategoryHost permits to find all host category in the centreon server
func getAllCategoryHost(debugV bool) []host.ExportCategory {
	//Get all host category
	err, body := request.GeneriqueCommandV1Post("show", "HC", "", "export host", debugV, false, "")
	if err != nil {
		return []host.ExportCategory{}
	}
	var resultCategory host.ExportResultCategory
	json.Unmarshal(body, &resultCategory)

	return resultCategory.CategoryHosts
}

func init() {
	hostCmd.Flags().StringSliceP("name", "n", []string{}, "Host category's name (separate by a comma the multiple values)")
	hostCmd.Flags().StringP("file", "f", "ExportHostCategory.csv", "To define the name of the csv file")
	hostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host category's name")

}
