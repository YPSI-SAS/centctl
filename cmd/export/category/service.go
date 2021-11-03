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
package category

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Export category service",
	Long:  `Export category service of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportCategoryService(name, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportCategoryService permits to export a category service of the centreon server
func ExportCategoryService(name []string, regex string, file string, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		categories := getAllCategoryService(debugV)
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
		err, category := getCategoryServiceInfo(n, debugV)
		if err != nil {
			return err
		}
		if category.Name == "" {
			continue
		}

		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,categoryService,\""+category.Name+"\",\""+category.Description+"\"\n", file, writeFile)

		//Write in the file the service members
		if len(category.Services) != 0 {
			for _, s := range category.Services {
				request.WriteValues("modify,categoryService,\""+category.Name+"\",service,\""+s.HostName+"|"+s.ServiceDescription+"\"\n", file, writeFile)
			}
		}

		//Write in the file the service template members
		if len(category.ServiceTemplates) != 0 {
			for _, s := range category.ServiceTemplates {
				request.WriteValues("modify,categoryService,\""+category.Name+"\",servicetemplate,\""+s.ServiceTemplateDescription+"\"\n", file, writeFile)
			}
		}
	}

	return nil
}

//The arguments impossible to get : activate
//getCategoryServiceInfo permits to get all informations about a service category
func getCategoryServiceInfo(name string, debugV bool) (error, service.ExportCategory) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the service category
	err, body := request.GeneriqueCommandV1Post("show", "SC", name, "export category service", debugV, false, "")
	if err != nil {
		return err, service.ExportCategory{}
	}
	var resultCategory service.ExportResultCategory
	json.Unmarshal(body, &resultCategory)

	//Get the category
	find := false
	category := service.ExportCategory{}
	for _, g := range resultCategory.CategoryServices {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			category = g
			find = true
		}
	}
	//Check if the service category is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, category
	}

	//Get the members service of the service category
	err, body = request.GeneriqueCommandV1Post("getservice", "SC", name, "export category service", debugV, false, "")
	if err != nil {
		return err, service.ExportCategory{}
	}
	var resultServiceMember service.ExportResultCategoryService
	json.Unmarshal(body, &resultServiceMember)

	//Get the members service template of the service category
	err, body = request.GeneriqueCommandV1Post("getservicetemplate", "SC", name, "export category service", debugV, false, "")
	if err != nil {
		return err, service.ExportCategory{}
	}
	var resultServiceTemplateMember service.ExportResultCategoryServiceTemplate
	json.Unmarshal(body, &resultServiceTemplateMember)

	category.ServiceTemplates = resultServiceTemplateMember.CategoryServiceTemplate
	category.Services = resultServiceMember.CategoryService

	return nil, category

}

//getAllCategoryService permits to find all service category in the centreon server
func getAllCategoryService(debugV bool) []service.ExportCategory {
	//Get all service category
	err, body := request.GeneriqueCommandV1Post("show", "SC", "", "export category service", debugV, false, "")
	if err != nil {
		return []service.ExportCategory{}
	}
	var resultCategory service.ExportResultCategory
	json.Unmarshal(body, &resultCategory)

	return resultCategory.CategoryServices
}

func init() {
	serviceCmd.Flags().StringSliceP("name", "n", []string{}, "Service category's name (separate by a comma the multiple values)")
	serviceCmd.Flags().StringP("regex", "r", "", "The regex to apply on the service category's name")

}
