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
package template

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Show one template service's details",
	Long:  `Show one template service's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowTemplateService(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowTemplateService permits to display the details of one template service
func ShowTemplateService(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "STPL", name, "show template service", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the template service contain into the response body
	templatesHost := service.DetailResultTemplate{}
	json.Unmarshal(body, &templatesHost)

	//Permits to find the good template service in the array
	var TemplateFind service.DetailTemplate
	for _, v := range templatesHost.DetailTemplates {
		if strings.ToLower(v.Description) == strings.ToLower(name) {
			TemplateFind = v
		}
	}

	var server service.DetailTemplateServer
	if TemplateFind.Description != "" {
		//Organization of data
		server = service.DetailTemplateServer{
			Server: service.DetailTemplateInformations{
				Name:     os.Getenv("SERVER"),
				Template: &TemplateFind,
			},
		}
	} else {
		//Organization of data
		server = service.DetailTemplateServer{
			Server: service.DetailTemplateInformations{
				Name:     os.Getenv("SERVER"),
				Template: nil,
			},
		}
	}

	//Display details of the template service
	displayTemplateService, err := display.DetailTemplateService(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTemplateService)
	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "To define the service template which will show")
	serviceCmd.MarkFlagRequired("name")
}
