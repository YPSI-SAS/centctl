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
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Show one template host's details",
	Long:  `Show one template host's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowTemplateHost(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowTemplateHost permits to display the details of one template host
func ShowTemplateHost(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "HTPL", name, "show template host", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the template host contain into the response body
	templatesHost := host.DetailResultTemplate{}
	json.Unmarshal(body, &templatesHost)

	//Permits to find the good template host in the array
	var TemplateFind host.DetailTemplate
	for _, v := range templatesHost.DetailTemplates {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			TemplateFind = v
		}
	}

	var server host.DetailTemplateServer
	if TemplateFind.Name != "" {
		//Organization of data
		server = host.DetailTemplateServer{
			Server: host.DetailTemplateInformations{
				Name:     os.Getenv("SERVER"),
				Template: &TemplateFind,
			},
		}
	} else {
		//Organization of data
		server = host.DetailTemplateServer{
			Server: host.DetailTemplateInformations{
				Name:     os.Getenv("SERVER"),
				Template: nil,
			},
		}
	}

	//Display details of the template Host
	displayTemplateHost, err := display.DetailTemplateHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTemplateHost)
	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the host template which will show")
	hostCmd.MarkFlagRequired("name")

}
