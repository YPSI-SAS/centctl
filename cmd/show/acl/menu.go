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

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Show one ACL menu's details",
	Long:  `Show one ACL menu's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowACLMenu(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowACLMenu permits to display the details of one ACL menu
func ShowACLMenu(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLMENU", name, "show ACL menu", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL menu contain into the response body
	menusACL := ACL.DetailResultMenu{}
	json.Unmarshal(body, &menusACL)

	//Permits to find the good ACL menu in the array
	var GroupFind ACL.DetailMenu
	for _, v := range menusACL.Menus {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			GroupFind = v
		}
	}

	var server ACL.DetailMenuServer
	if GroupFind.Name != "" {
		//Organization of data
		server = ACL.DetailMenuServer{
			Server: ACL.DetailMenuInformations{
				Name: os.Getenv("SERVER"),
				Menu: &GroupFind,
			},
		}
	} else {
		//Organization of data
		server = ACL.DetailMenuServer{
			Server: ACL.DetailMenuInformations{
				Name: os.Getenv("SERVER"),
				Menu: nil,
			},
		}
	}

	//Display details of the ACL menu
	displayACLMenu, err := display.DetailACLMenu(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLMenu)
	return nil
}

func init() {
	menuCmd.Flags().StringP("name", "n", "", "To define the ACL menu which will show")
	menuCmd.MarkFlagRequired("name")
	menuCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetACLMenuNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
