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
package group

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
	Short: "Show one group contact's details",
	Long:  `Show one group contact's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowGroupContact(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowGroupContact permits to display the details of one group contact
func ShowGroupContact(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "CG", name, "show group contact", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the group contact contain into the response body
	groupsContact := contact.DetailResultGroup{}
	json.Unmarshal(body, &groupsContact)
	//Permits to find the good group contact in the array
	var GroupFind contact.DetailGroup
	for _, v := range groupsContact.DetailGroups {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			GroupFind = v
		}
	}

	var server contact.DetailGroupServer
	if GroupFind.Name != "" {
		err, body := request.GeneriqueCommandV1Post("getcontact", "CG", GroupFind.Name, "getcontact", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the group contact contain into the response body
		groupContacts := contact.DetailResultGroupContact{}
		json.Unmarshal(body, &groupContacts)

		GroupFind.Contacts = groupContacts.DetailGroupContacts
		//Organization of data
		server = contact.DetailGroupServer{
			Server: contact.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: &GroupFind,
			},
		}
	} else {
		//Organization of data
		server = contact.DetailGroupServer{
			Server: contact.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: nil,
			},
		}
	}

	//Display details of the contact group
	displayGroupContact, err := display.DetailGroupContact(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayGroupContact)
	return nil
}

func init() {
	contactCmd.Flags().StringP("name", "n", "", "To define the contact group which will show")
	contactCmd.MarkFlagRequired("name")
}
