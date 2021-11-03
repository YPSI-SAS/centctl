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

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Export ACL action",
	Long:  `Export ACL action of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		name, _ := cmd.Flags().GetStringSlice("name")
		regex, _ := cmd.Flags().GetString("regex")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportACLAction(name, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportACLAction permits to export a ACL action of the centreon server
func ExportACLAction(name []string, regex string, file string, all bool, debugV bool) error {
	writeFile := false
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		actions := getAllAction(debugV)
		for _, a := range actions {
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
		err, action := getACLActionInfo(n, debugV)
		if err != nil {
			return err
		}
		if action.Name == "" {
			continue
		}
		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,aclAction,\""+action.Name+"\",\""+action.Description+"\"\n", file, writeFile)
		request.WriteValues("modify,aclAction,\""+action.Name+"\",activate,\""+action.Activate+"\"\n", file, writeFile)

	}

	return nil
}

//The arguments impossible to get : all action grant
//getACLActionInfo permits to get all informations about an ACL action
func getACLActionInfo(name string, debugV bool) (error, ACL.ExportAction) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the ACL action
	err, body := request.GeneriqueCommandV1Post("show", "ACLACTION", name, "export ACL action", debugV, false, "")
	if err != nil {
		return err, ACL.ExportAction{}
	}
	var resultAction ACL.ExportActionResult
	json.Unmarshal(body, &resultAction)

	find := false
	//Get informations
	action := ACL.ExportAction{}
	for _, m := range resultAction.ActionACL {
		if strings.ToLower(m.Name) == strings.ToLower(name) {
			action = m
			find = true
		}
	}

	//Check if the ACL action is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
	}
	return nil, action

}

//getAllAction permits to find all acl Action in the centreon server
func getAllAction(debugV bool) []ACL.ExportAction {
	//Get all ACL action
	err, body := request.GeneriqueCommandV1Post("show", "ACLACTION", "", "export ACL action", debugV, false, "")
	if err != nil {
		return []ACL.ExportAction{}
	}
	var resultAction ACL.ExportActionResult
	json.Unmarshal(body, &resultAction)

	return resultAction.ActionACL
}

func init() {
	actionCmd.Flags().StringSliceP("name", "n", []string{}, "ACL action's name (separate by a comma the multiple values)")
	actionCmd.Flags().StringP("regex", "r", "", "The regex to apply on the ACL action's name")

}
