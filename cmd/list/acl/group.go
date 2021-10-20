/*MIT License

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

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "List ACL group",
	Long:  `List ACL group of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListACLGroup(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListACLGroup permits to display the array of ACLs group return by the API
func ListACLGroup(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLGROUP", "", "list ACL group", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL groups contain into the response body
	groups := ACL.ResultGroup{}
	json.Unmarshal(body, &groups)
	finalGroups := groups.Groups
	if regex != "" {
		finalGroups = deleteGroup(finalGroups, regex)
	}

	//Sort ACL groups based on their ID
	sort.SliceStable(finalGroups, func(i, j int) bool {
		return strings.ToLower(finalGroups[i].Name) < strings.ToLower(finalGroups[j].Name)
	})

	server := ACL.GroupServer{
		Server: ACL.GroupInformations{
			Name:   os.Getenv("SERVER"),
			Groups: finalGroups,
		},
	}

	//Display all ACL groups
	displayACLGroup, err := display.ACLGroup(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLGroup)

	return nil
}

func deleteGroup(groups []ACL.Group, regex string) []ACL.Group {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range groups {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			groups[index] = s
			index++
		}
	}
	return groups[:index]
}

func init() {
	groupCmd.Flags().StringP("regex", "r", "", "The regex to apply on the group's name")

}
