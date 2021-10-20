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
	"centctl/display"
	"centctl/request"
	"centctl/resources/ACL"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Show one ACL group's details",
	Long:  `Show one ACL group's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowACLGroup(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowACLGroup permits to display the details of one ACL group
func ShowACLGroup(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLGROUP", name, "show ACL group", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL group contain into the response body
	groupsACL := ACL.DetailResultGroup{}
	json.Unmarshal(body, &groupsACL)

	//Permits to find the good ACL group in the array
	var GroupFind ACL.DetailGroup
	for _, v := range groupsACL.Groups {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			GroupFind = v
		}
	}

	var server ACL.DetailGroupServer
	if GroupFind.Name != "" {
		//Organization of data
		server = ACL.DetailGroupServer{
			Server: ACL.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: &GroupFind,
			},
		}
	} else {
		//Organization of data
		server = ACL.DetailGroupServer{
			Server: ACL.DetailGroupInformations{
				Name:  os.Getenv("SERVER"),
				Group: nil,
			},
		}
	}

	//Display details of the ACL Group
	displayACLGroup, err := display.DetailACLGroup(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLGroup)
	return nil
}

func init() {
	groupCmd.Flags().StringP("name", "n", "", "To define the ACL group which will show")
	groupCmd.MarkFlagRequired("name")
}
