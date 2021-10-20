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
package acl

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/ACL"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Export ACL group",
	Long:  `Export ACL group of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportACLGroup(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportACLGroup permits to export an ACL group of the centreon server
func ExportACLGroup(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
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
		groups := getAllGroup(debugV)
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
		err, group := getACLGroupInfo(n, debugV)
		if err != nil {
			return err
		}
		if group.Name == "" {
			continue
		}

		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,aclGroup,\"" + group.Name + "\",\"" + group.Alias + "\"\n")
		_, _ = f.WriteString("modify,aclGroup,\"" + group.Name + "\",activate,\"" + group.Activate + "\"\n")

		//Write in the file the contacts
		if len(group.Contact) != 0 {
			for _, m := range group.Contact {
				_, _ = f.WriteString("modify,aclGroup,\"" + group.Name + "\",contact,\"" + m.Name + "\"\n")
			}
		}

		//Write in the file the contacts group
		if len(group.ContactGroup) != 0 {
			for _, m := range group.ContactGroup {
				_, _ = f.WriteString("modify,aclGroup,\"" + group.Name + "\",contactgroup,\"" + m.Name + "\"\n")
			}
		}

		//Write in the file the menu
		if len(group.Menu) != 0 {
			for _, m := range group.Menu {
				_, _ = f.WriteString("modify,aclGroup,\"" + group.Name + "\",menu,\"" + m.Name + "\"\n")
			}
		}

		//Write in the file the action
		if len(group.Action) != 0 {
			for _, m := range group.Action {
				_, _ = f.WriteString("modify,aclGroup,\"" + group.Name + "\",action,\"" + m.Name + "\"\n")
			}
		}

		//Write in the file the resource
		if len(group.Resource) != 0 {
			for _, m := range group.Resource {
				_, _ = f.WriteString("modify,aclGroup,\"" + group.Name + "\",resource,\"" + m.Name + "\"\n")
			}
		}
	}
	return nil
}

//getACLGroupInfo permits to get all informations about an ACL group
func getACLGroupInfo(name string, debugV bool) (error, ACL.ExportGroup) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the ACL group
	err, body := request.GeneriqueCommandV1Post("show", "ACLGROUP", name, "export ACL group", debugV, false, "")
	if err != nil {
		return err, ACL.ExportGroup{}
	}
	var resultGroup ACL.ExportResult
	json.Unmarshal(body, &resultGroup)

	//Get informations
	find := false
	group := ACL.ExportGroup{}
	for _, g := range resultGroup.GroupACL {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			group = g
			find = true
		}
	}
	//Check if the ACL group is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, group
	}

	//Get the contacts of the ACL group
	err, body = request.GeneriqueCommandV1Post("getcontact", "ACLGROUP", name, "export ACL group", debugV, false, "")
	if err != nil {
		return err, ACL.ExportGroup{}
	}
	var resultContact ACL.ExportResultContact
	json.Unmarshal(body, &resultContact)

	//Get the contacts group of the ACL group
	err, body = request.GeneriqueCommandV1Post("getcontactgroup", "ACLGROUP", name, "export ACL group", debugV, false, "")
	if err != nil {
		return err, ACL.ExportGroup{}
	}
	var resultContactGroup ACL.ExportResultContactGroup
	json.Unmarshal(body, &resultContactGroup)

	//Get the menu of the ACL group
	err, body = request.GeneriqueCommandV1Post("getmenu", "ACLGROUP", name, "export ACL group", debugV, false, "")
	if err != nil {
		return err, ACL.ExportGroup{}
	}
	var resultMenu ACL.ExportResultMenu
	json.Unmarshal(body, &resultMenu)

	//Get the action of the ACL group
	err, body = request.GeneriqueCommandV1Post("getaction", "ACLGROUP", name, "export ACL group", debugV, false, "")
	if err != nil {
		return err, ACL.ExportGroup{}
	}
	var resultAction ACL.ExportResultAction
	json.Unmarshal(body, &resultAction)

	//Get the resource of the ACL group
	err, body = request.GeneriqueCommandV1Post("getresource", "ACLGROUP", name, "export ACL group", debugV, false, "")
	if err != nil {
		return err, ACL.ExportGroup{}
	}
	var resultResource ACL.ExportResultResource
	json.Unmarshal(body, &resultResource)

	group.Contact = resultContact.GroupContact
	group.ContactGroup = resultContactGroup.GroupContactGroup
	group.Menu = resultMenu.GroupMenu
	group.Action = resultAction.GroupAction
	group.Resource = resultResource.GroupResource

	return nil, group

}

//getAllGroup permits to find all acl group in the centreon server
func getAllGroup(debugV bool) []ACL.ExportGroup {
	//Get all ACL group
	err, body := request.GeneriqueCommandV1Post("show", "ACLGROUP", "", "export ACL group", debugV, false, "")
	if err != nil {
		return []ACL.ExportGroup{}
	}
	var resultGroup ACL.ExportResult
	json.Unmarshal(body, &resultGroup)

	return resultGroup.GroupACL
}

func init() {
	groupCmd.Flags().StringSliceP("name", "n", []string{}, "ACL group's name (separate by a comma the multiple values)")
	groupCmd.Flags().StringP("file", "f", "ExportACLGroup.csv", "To define the name of the csv file")
	groupCmd.Flags().StringP("regex", "r", "", "The regex to apply on the ACL group's name")

}
