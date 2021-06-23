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

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "List ACL menu",
	Long:  `List ACL menu of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListACLMenu(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListACLMenu permits to display the array of ACLs menu return by the API
func ListACLMenu(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLMENU", "", "list ACL menu", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL menus contain into the response body
	menus := ACL.ResultMenu{}
	json.Unmarshal(body, &menus)
	finalMenus := menus.Menus
	if regex != "" {
		finalMenus = deleteMenu(finalMenus, regex)
	}

	//Sort ACL menus based on their ID
	sort.SliceStable(finalMenus, func(i, j int) bool {
		return strings.ToLower(finalMenus[i].Name) < strings.ToLower(finalMenus[j].Name)
	})

	server := ACL.MenuServer{
		Server: ACL.MenuInformations{
			Name:  os.Getenv("SERVER"),
			Menus: finalMenus,
		},
	}

	//Display all ACL menus
	displayACLMenu, err := display.ACLMenu(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLMenu)

	return nil
}

func deleteMenu(menus []ACL.Menu, regex string) []ACL.Menu {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range menus {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			menus[index] = s
			index++
		}
	}
	return menus[:index]
}

func init() {
	menuCmd.Flags().StringP("regex", "r", "", "The regex to apply on the menu's name")

}
