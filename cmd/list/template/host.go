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
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "List template host",
	Long:  `List template host of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListTemplateHost(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListTemplateHost permits to display the array of host template return by the API
func ListTemplateHost(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "HTPL", "", "list template host", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the host templates contain into the response body
	templates := host.ResultTemplate{}
	json.Unmarshal(body, &templates)
	finalTemplates := templates.Templates
	if regex != "" {
		finalTemplates = deleteHostTemplate(finalTemplates, regex)
	}

	//Sort host templates based on their ID
	sort.SliceStable(finalTemplates, func(i, j int) bool {
		return strings.ToLower(finalTemplates[i].Name) < strings.ToLower(finalTemplates[j].Name)
	})

	server := host.TemplateServer{
		Server: host.TemplateInformations{
			Name:      os.Getenv("SERVER"),
			Templates: finalTemplates,
		},
	}

	//Display all host templates
	displayTemplateHost, err := display.TemplateHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTemplateHost)

	return nil
}

func deleteHostTemplate(hostTemplate []host.Template, regex string) []host.Template {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range hostTemplate {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			hostTemplate[index] = s
			index++
		}
	}
	return hostTemplate[:index]
}

func init() {
	hostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host template's name")

}
