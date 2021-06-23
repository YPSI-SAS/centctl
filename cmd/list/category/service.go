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

package category

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "List category service",
	Long:  `List category service of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListGroupService(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListGroupService permits to display the array of service category return by the API
func ListGroupService(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "SC", "", "list category service", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the service category contain into the response body
	category := service.ResultCategory{}
	json.Unmarshal(body, &category)
	finalCategories := category.Categories
	if regex != "" {
		finalCategories = deleteServiceCategory(finalCategories, regex)
	}

	//Sort service category based on their ID
	sort.SliceStable(finalCategories, func(i, j int) bool {
		return strings.ToLower(finalCategories[i].Name) < strings.ToLower(finalCategories[j].Name)
	})

	server := service.CategoryServer{
		Server: service.CategoryInformations{
			Name:       os.Getenv("SERVER"),
			Categories: finalCategories,
		},
	}

	//Display all service category
	displayCategoryService, err := display.CategoryService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayCategoryService)

	return nil
}

func deleteServiceCategory(serviceCategories []service.Category, regex string) []service.Category {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range serviceCategories {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			serviceCategories[index] = s
			index++
		}
	}
	return serviceCategories[:index]
}

func init() {
	serviceCmd.Flags().StringP("regex", "r", "", "The regex to apply on the service category's name")

}
