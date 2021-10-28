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
package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/LDAP"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// ldapCmd represents the LDAP command
var ldapCmd = &cobra.Command{
	Use:   "LDAP",
	Short: "List the LDAP configuration",
	Long:  `List the LDAP configuration of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListLDAP(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListLDAP permits to display the array of LDAP return by the API
func ListLDAP(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "LDAP", "", "list LDAP", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the LDAP contain into the response body
	LDAPs := LDAP.Result{}
	json.Unmarshal(body, &LDAPs)
	finalLDAPs := LDAPs.LDAP
	if regex != "" {
		finalLDAPs = deleteLDAP(finalLDAPs, regex)
	}

	//Sort LDAP based on their ID
	sort.SliceStable(finalLDAPs, func(i, j int) bool {
		return strings.ToLower(finalLDAPs[i].Name) < strings.ToLower(finalLDAPs[j].Name)
	})

	//Organization of data
	server := LDAP.Server{
		Server: LDAP.Informations{
			Name: os.Getenv("SERVER"),
			LDAP: finalLDAPs,
		},
	}

	//Display all LDAP
	displayLDAP, err := display.LDAPs(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayLDAP)

	return nil
}

func deleteLDAP(ldaps []LDAP.LDAP, regex string) []LDAP.LDAP {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range ldaps {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			ldaps[index] = s
			index++
		}
	}
	return ldaps[:index]
}

func init() {
	ldapCmd.Flags().StringP("regex", "r", "", "The regex to apply on the LDAP's name")
}
