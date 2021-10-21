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
	"centctl/request"
	"centctl/resources/contact"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Export template contact",
	Long:  `Export template contact of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		alias, _ := cmd.Flags().GetStringSlice("alias")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportTemplateContact(alias, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportTemplateContact permits to export a contact template of the centreon server
func ExportTemplateContact(alias []string, regex string, file string, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(alias) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag alias or flag all or flag regex ")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		templates := getAllTemplateContact(debugV)
		for _, a := range templates {
			if regex != "" {
				matched, err := regexp.MatchString(regex, a.Alias)
				if err != nil {
					fmt.Printf(colorRed, "ERROR:")
					fmt.Println(err.Error())
					os.Exit(1)
				}
				if matched {
					alias = append(alias, a.Alias)
				}
			} else {
				alias = append(alias, a.Alias)
			}
		}
	}
	for _, n := range alias {
		err, templateContact := getTemplateContactInfo(n, debugV)
		if err != nil {
			return err
		}
		if templateContact.Alias == "" {
			continue
		}

		rand.Seed(time.Now().UnixNano())
		//Write templateContact informations
		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,templateContact,\""+templateContact.Name+"\",\""+templateContact.Alias+"\"\n", file, writeFile)
		request.WriteValues("modify,templateContact,"+templateContact.Alias+",activate,"+templateContact.Activate+"\n", file, writeFile)

	}
	return nil
}

//The arguments impossible to get : all except name and alias
//getTemplateContactInfo permits to get all informations about a contact
func getTemplateContactInfo(alias string, debugV bool) (error, contact.ExportTemplateContact) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "contactTPL", alias, "export contact", debugV, false, "")
	if err != nil {
		return err, contact.ExportTemplateContact{}
	}
	var resultTemplate contact.ExportResultTemplateContact
	json.Unmarshal(body, &resultTemplate)

	template := contact.ExportTemplateContact{}
	find := false
	for _, g := range resultTemplate.TemplateContacts {
		if strings.ToLower(g.Alias) == strings.ToLower(alias) {
			template = g
			find = true
		}
	}
	//Check if the contact template is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + alias)
		return nil, template
	}

	return nil, template

}

//getAllTemplateContact permits to find all contact template in the centreon server
func getAllTemplateContact(debugV bool) []contact.ExportTemplateContact {
	//Get all contact template
	err, body := request.GeneriqueCommandV1Post("show", "contactTPL", "", "export contact", debugV, false, "")
	if err != nil {
		return []contact.ExportTemplateContact{}
	}
	var resultTemplate contact.ExportResultTemplateContact
	json.Unmarshal(body, &resultTemplate)

	return resultTemplate.TemplateContacts
}

func init() {
	contactCmd.Flags().StringSliceP("alias", "a", []string{}, "Template contact's alias (separate by a comma the multiple values)")
	contactCmd.Flags().StringP("regex", "r", "", "The regex to apply on the template contact's alias")

}
