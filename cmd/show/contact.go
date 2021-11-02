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

package show

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
	Short: "Show one contact's details",
	Long:  `Show one contact's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		alias, _ := cmd.Flags().GetString("alias")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowContact(alias, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowContact permits to display the details of one contact
func ShowContact(alias string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "contact", alias, "show contact", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the contacts contain into the response body
	contacts := contact.DetailResult{}
	json.Unmarshal(body, &contacts)

	//Permits to find the good contact in the array
	var ContactFind contact.DetailContact
	for _, v := range contacts.DetailContacts {
		if strings.ToLower(v.Alias) == strings.ToLower(alias) {
			ContactFind = v
		}
	}

	var server contact.DetailServer
	if ContactFind.Alias != "" {
		server = contact.DetailServer{
			Server: contact.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Contact: &ContactFind,
			},
		}
	} else {
		server = contact.DetailServer{
			Server: contact.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Contact: nil,
			},
		}
	}

	//Display details of the contact
	displayContact, err := display.DetailContact(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayContact)
	return nil

}

func init() {
	contactCmd.Flags().StringP("alias", "a", "", "Contact's alias")
	contactCmd.MarkFlagRequired("alias")
	contactCmd.RegisterFlagCompletionFunc("alias", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetContactAlias()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
