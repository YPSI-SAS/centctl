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

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Export ACL menu",
	Long:  `Export ACL menu of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportACLMenu(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportACLMenu permits to export an ACL menu of the centreon server
func ExportACLMenu(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
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
		menus := getAllMenu(debugV)
		for _, a := range menus {
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
		err, menu := getACLMenuInfo(n, debugV)
		if err != nil {
			return err
		}
		if menu.Name == "" {
			continue
		}

		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,aclMenu,\"" + menu.Name + "\",\"" + menu.Alias + "\"\n")
		_, _ = f.WriteString("modify,aclMenu,\"" + menu.Name + "\",activate,\"" + menu.Activate + "\"\n")
		_, _ = f.WriteString("modify,aclMenu,\"" + menu.Name + "\",comment,\"" + menu.Comment + "\"\n")
	}
	return nil
}

//The arguments impossible to get : all accessible pages
//getACLMenuInfo permits to get all informations about an ACL menu
func getACLMenuInfo(name string, debugV bool) (error, ACL.ExportMenu) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the ACL menu
	err, body := request.GeneriqueCommandV1Post("show", "ACLMENU", name, "export ACL menu", debugV, false, "")
	if err != nil {
		return err, ACL.ExportMenu{}
	}
	var resultMenu ACL.ExportMenuResult
	json.Unmarshal(body, &resultMenu)

	//Get informations
	find := false
	menu := ACL.ExportMenu{}
	for _, m := range resultMenu.MenuACL {
		if strings.ToLower(m.Name) == strings.ToLower(name) {
			menu = m
			find = true
		}
	}

	//Check if the ACL menu is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
	}

	return nil, menu

}

//getAllMenu permits to find all acl menu in the centreon server
func getAllMenu(debugV bool) []ACL.ExportMenu {
	//Get all ACL menu
	err, body := request.GeneriqueCommandV1Post("show", "ACLMENU", "", "export ACL menu", debugV, false, "")
	if err != nil {
		return []ACL.ExportMenu{}
	}
	var resultMenu ACL.ExportMenuResult
	json.Unmarshal(body, &resultMenu)

	return resultMenu.MenuACL
}

func init() {
	menuCmd.Flags().StringSliceP("name", "n", []string{}, "ACL menu's name (separate by a comma the multiple values)")
	menuCmd.Flags().StringP("file", "f", "ExportACLMenu.csv", "To define the name of the csv file")
	menuCmd.Flags().StringP("regex", "r", "", "The regex to apply on the ACL menu's name")
}
