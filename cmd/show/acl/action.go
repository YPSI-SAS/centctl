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
	"centctl/display"
	"centctl/request"
	"centctl/resources/ACL"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// actionCmd represents the action command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Show one ACL action's details",
	Long:  `Show one ACL action's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowACLAction(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowACLAction permits to display the details of one ACL action
func ShowACLAction(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLACTION", name, "show ACL action", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL action contain into the response body
	actionsACL := ACL.DetailResultAction{}
	json.Unmarshal(body, &actionsACL)

	//Permits to find the good ACL action in the array
	var ActionFind ACL.DetailAction
	for _, v := range actionsACL.Actions {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			ActionFind = v
		}
	}

	var server ACL.DetailActionServer
	if ActionFind.Name != "" {
		//Organization of data
		server = ACL.DetailActionServer{
			Server: ACL.DetailActionInformations{
				Name:   os.Getenv("SERVER"),
				Action: &ActionFind,
			},
		}
	} else {
		//Organization of data
		server = ACL.DetailActionServer{
			Server: ACL.DetailActionInformations{
				Name:   os.Getenv("SERVER"),
				Action: nil,
			},
		}
	}

	//Display details of the ACL Action
	displayACLAction, err := display.DetailACLAction(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLAction)
	return nil
}

func init() {
	actionCmd.Flags().StringP("name", "n", "", "To define the ACL action which will show")
	actionCmd.MarkFlagRequired("name")
}
