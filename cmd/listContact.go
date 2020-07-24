/*
MIT License

Copyright (c) 2020 YPSI SAS
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

package cmd

import (
	"centctl/contact"
	"centctl/debug"
	"centctl/display"
	"centctl/request"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listContactCmd represents the contact command
var listContactCmd = &cobra.Command{
	Use:   "contact",
	Short: "List the contacts",
	Long:  `List the contacts of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListContact(output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListContact permits to display the array of contact return by the API
func ListContact(output string, debugV bool) error {
	output = strings.ToLower(output)

	//Creation of the request body
	requestBody, err := request.CreateBodyRequest("Show", "contact", "")
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list contact", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the contacts contain into the response body
	contacts := contact.Result{}
	json.Unmarshal(body, &contacts)

	//Sort contacts based on their ID
	sort.SliceStable(contacts.Contacts, func(i, j int) bool {
		return strings.ToLower(contacts.Contacts[i].ID) < strings.ToLower(contacts.Contacts[j].ID)
	})

	//Organization of data
	server := contact.Server{
		Server: contact.Informations{
			Name:     os.Getenv("SERVER"),
			Contacts: contacts.Contacts,
		},
	}

	//Display all contacts
	displayContact, err := display.Contact(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayContact)

	return nil
}

func init() {
	listCmd.AddCommand(listContactCmd)
}
