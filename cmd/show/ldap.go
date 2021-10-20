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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/LDAP"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ldapCmd represents the LDAP command
var ldapCmd = &cobra.Command{
	Use:   "LDAP",
	Short: "Show one LDAP's configuration details",
	Long:  `Show one LDAP's configuration details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowLDAP(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowLDAP permits to display the details of one LDAP
func ShowLDAP(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "LDAP", name, "show LDAP", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the LDAPs contain into the response body
	LDAPs := LDAP.DetailResult{}
	json.Unmarshal(body, &LDAPs)

	//Permits to find the good LDAP in the array
	var LDAPFind LDAP.DetailLDAP
	for _, v := range LDAPs.LDAP {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			LDAPFind = v
		}
	}

	var server LDAP.DetailServer
	if LDAPFind.Name != "" {
		err, body := request.GeneriqueCommandV1Post("showserver", "LDAP", LDAPFind.Name, "showserver", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the member contain into the response body
		servers := LDAP.DetailResultServer{}
		json.Unmarshal(body, &servers)

		LDAPFind.Servers = servers.Servers
		//Organization of data
		server = LDAP.DetailServer{
			Server: LDAP.DetailInformations{
				Name: os.Getenv("SERVER"),
				LDAP: &LDAPFind,
			},
		}
	} else {
		server = LDAP.DetailServer{
			Server: LDAP.DetailInformations{
				Name: os.Getenv("SERVER"),
				LDAP: nil,
			},
		}
	}

	//Display details of the LDAP
	displayLDAP, err := display.DetailLDAP(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayLDAP)
	return nil
}

func init() {
	ldapCmd.Flags().StringP("name", "n", "", "To define the name of the LDAP")
	ldapCmd.MarkFlagRequired("name")
}
