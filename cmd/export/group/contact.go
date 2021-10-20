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
package group

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/contact"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Export group contact",
	Long:  `Export group contact of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportGroupContact(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportGroupContact permits to export a group contact of the centreon server
func ExportGroupContact(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	//Check if the name of file contains the extension
	if !strings.Contains(file, ".csv") {
		file = file + ".csv"
	}

	//Create the file
	var f *os.File
	var err error
	if appendFile {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		f, err = os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	}
	defer f.Close()
	if err != nil {
		return err
	}

	if all || regex != "" {
		groups := getAllGroupContact(debugV)
		for _, a := range groups {
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
		err, group := getGroupContactInfo(n, debugV)
		if err != nil {
			return err
		}
		if group.Name == "" {
			continue
		}

		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,groupContact,\"" + group.Name + "\",\"" + group.Alias + "\"\n")

		//Write in the file the members
		if len(group.Member) != 0 {
			for _, m := range group.Member {
				_, _ = f.WriteString("modify,groupContact,\"" + group.Name + "\",contact,\"" + m.Name + "\"\n")
			}
		}
	}
	return nil
}

//The arguments impossible to get : linked_acl_group|activate|comment
//getGroupContactInfo permits to get all informations about a contact group
func getGroupContactInfo(name string, debugV bool) (error, contact.ExportGroup) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the contact group
	err, body := request.GeneriqueCommandV1Post("show", "CG", name, "export group contact", debugV, false, "")
	if err != nil {
		return err, contact.ExportGroup{}
	}
	var resultGroup contact.ExportResult
	json.Unmarshal(body, &resultGroup)

	//Get the group
	group := contact.ExportGroup{}
	find := false
	for _, g := range resultGroup.GroupContacts {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			group = g
			find = true
		}
	}

	//Check if the contact group is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, group
	}

	//Get the members of the contact group
	err, body = request.GeneriqueCommandV1Post("getcontact", "CG", name, "export group contact", debugV, false, "")
	if err != nil {
		return err, contact.ExportGroup{}
	}
	var resultMember contact.ExportResultMember
	json.Unmarshal(body, &resultMember)

	group.Member = resultMember.GroupMember

	return nil, group

}

//getAllGroupContact permits to find all contact group in the centreon server
func getAllGroupContact(debugV bool) []contact.ExportGroup {
	//Get all contact group
	err, body := request.GeneriqueCommandV1Post("show", "CG", "", "export group contact", debugV, false, "")
	if err != nil {
		return []contact.ExportGroup{}
	}
	var resultGroup contact.ExportResult
	json.Unmarshal(body, &resultGroup)

	return resultGroup.GroupContacts
}

func init() {
	contactCmd.Flags().StringSliceP("name", "n", []string{}, "Contactgroup's name (separate by a comma the multiple values)")
	contactCmd.Flags().StringP("file", "f", "ExportContactGroup.csv", "To define the name of the csv file")
	contactCmd.Flags().StringP("regex", "r", "", "The regex to apply on the contact group's name")

}
