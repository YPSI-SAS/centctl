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
package template

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/contact"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Show one template contact's details",
	Long:  `Show one template contact's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		alias, _ := cmd.Flags().GetString("alias")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowTemplateContact(alias, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowTemplateContact permits to display the details of one template contact
func ShowTemplateContact(alias string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "CONTACTTPL", alias, "show template contact", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the template contact contain into the response body
	templatesContact := contact.DetailResultTemplate{}
	json.Unmarshal(body, &templatesContact)

	//Permits to find the good template contact in the array
	var TemplateFind contact.DetailTemplate
	for _, v := range templatesContact.DetailTemplates {
		if strings.ToLower(v.Alias) == strings.ToLower(alias) {
			TemplateFind = v
		}
	}

	var server contact.DetailTemplateServer
	if TemplateFind.Alias != "" {
		//Organization of data
		server = contact.DetailTemplateServer{
			Server: contact.DetailTemplateInformations{
				Name:     os.Getenv("SERVER"),
				Template: &TemplateFind,
			},
		}
	} else {
		//Organization of data
		server = contact.DetailTemplateServer{
			Server: contact.DetailTemplateInformations{
				Name:     os.Getenv("SERVER"),
				Template: nil,
			},
		}
	}

	//Display details of the template contact
	displayTemplateContact, err := display.DetailTemplateContact(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTemplateContact)
	return nil
}

func init() {
	contactCmd.Flags().StringP("alias", "a", "", "To define the contact template which will show")
	contactCmd.MarkFlagRequired("alias")
}
