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

package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/contact"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "List the contacts",
	Long:  `List the contacts of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListContact(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListContact permits to display the array of contact return by the API
func ListContact(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "contact", "", "list contact", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the contacts contain into the response body
	contacts := contact.Result{}
	json.Unmarshal(body, &contacts)
	finalContacts := contacts.Contacts
	if regex != "" {
		finalContacts = deleteContact(finalContacts, regex)
	}

	//Sort contacts based on their ID
	sort.SliceStable(finalContacts, func(i, j int) bool {
		valI, _ := strconv.Atoi(finalContacts[i].ID)
		valJ, _ := strconv.Atoi(finalContacts[j].ID)
		return valI < valJ
	})

	//Organization of data
	server := contact.Server{
		Server: contact.Informations{
			Name:     os.Getenv("SERVER"),
			Contacts: finalContacts,
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

func deleteContact(contacts []contact.Contact, regex string) []contact.Contact {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range contacts {
		matched, err := regexp.MatchString(regex, s.Alias)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			contacts[index] = s
			index++
		}
	}
	return contacts[:index]
}

func init() {
	contactCmd.Flags().StringP("regex", "r", "", "The regex to apply on the contact's alias")
}
