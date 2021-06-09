/*MIT License

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
	"centctl/display"
	"centctl/request"
	"centctl/resources/ACL"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "List ACL action",
	Long:  `List ACL action of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListACLAction(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListACLAction permits to display the array of ACLs action return by the API
func ListACLAction(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLACTION", "", "list ACL action", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL actions contain into the response body
	actions := ACL.ResultAction{}
	json.Unmarshal(body, &actions)
	finalActions := actions.Actions
	if regex != "" {
		finalActions = deleteAction(finalActions, regex)
	}

	//Sort ACL actions based on their ID
	sort.SliceStable(finalActions, func(i, j int) bool {
		return strings.ToLower(finalActions[i].Name) < strings.ToLower(finalActions[j].Name)
	})

	server := ACL.ActionServer{
		Server: ACL.ActionInformations{
			Name:    os.Getenv("SERVER"),
			Actions: finalActions,
		},
	}

	//Display all ACL actions
	displayACLActions, err := display.ACLAction(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLActions)

	return nil
}

func deleteAction(actions []ACL.Action, regex string) []ACL.Action {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range actions {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			actions[index] = s
			index++
		}
	}
	return actions[:index]
}

func init() {
	actionCmd.Flags().StringP("regex", "r", "", "The regex to apply on the action's name")

}
