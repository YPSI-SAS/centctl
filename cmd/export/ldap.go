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
package export

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/LDAP"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// ldapCmd represents the LDAP command
var ldapCmd = &cobra.Command{
	Use:   "LDAP",
	Short: "Export LDAP",
	Long:  `Export LDAP of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportLDAP(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportLDAP permits to export a LDAP of the centreon server
func ExportLDAP(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex ")
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
		templates := getAllLDAP(debugV)
		for _, a := range templates {
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
		err, LDAP := getLDAPInfo(n, debugV)
		if err != nil {
			return err
		}
		if LDAP.Name == "" {
			continue
		}

		//Write LDAP informations
		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,LDAP,\"" + LDAP.Name + "\",\"" + LDAP.Description + "\"\n")
		_, _ = f.WriteString("modify,LDAP,\"" + LDAP.Name + "\",enable,\"" + LDAP.Status + "\"\n")

		//Write Server information
		if len(LDAP.Servers) != 0 {
			for _, c := range LDAP.Servers {
				_, _ = f.WriteString("modify,LDAP,\"" + LDAP.Name + "\",server,\"" + c.Address + ";" + c.Port + ";" + c.SSL + ";" + c.TLS + "\"\n")
			}
		}

	}

	return nil
}

//The arguments impossible to get : all in setparam table
//getLDAPInfo permits to get all informations about a LDAP
func getLDAPInfo(name string, debugV bool) (error, LDAP.ExportLDAP) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "LDAP", name, "export LDAP", debugV, false, "")
	if err != nil {
		return err, LDAP.ExportLDAP{}
	}
	var resultLDAP LDAP.ExportResultLDAP
	json.Unmarshal(body, &resultLDAP)

	ldap := LDAP.ExportLDAP{}
	find := false
	for _, g := range resultLDAP.LDAPs {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			ldap = g
			find = true
		}
	}
	//Check if the LDAP  is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, ldap
	}

	//Get the server of the LDAP
	err, body = request.GeneriqueCommandV1Post("showserver", "LDAP", name, "export LDAP", debugV, false, "")
	if err != nil {
		return err, LDAP.ExportLDAP{}
	}
	var resultServer LDAP.ExportResultLDAPServer
	json.Unmarshal(body, &resultServer)

	ldap.Servers = resultServer.Server

	return nil, ldap

}

//getAllLDAP permits to find all LDAP in the centreon server
func getAllLDAP(debugV bool) []LDAP.ExportLDAP {
	//Get all LDAP
	err, body := request.GeneriqueCommandV1Post("show", "LDAP", "", "export LDAP", debugV, false, "")
	if err != nil {
		return []LDAP.ExportLDAP{}
	}
	var resultLDAP LDAP.ExportResultLDAP
	json.Unmarshal(body, &resultLDAP)

	return resultLDAP.LDAPs
}

func init() {
	ldapCmd.Flags().StringSliceP("name", "n", []string{}, "LDAP's name (separate by a comma the multiple values)")
	ldapCmd.Flags().StringP("file", "f", "ExportLDAP.csv", "To define the name of the csv file")
	ldapCmd.Flags().StringP("regex", "r", "", "The regex to apply on the LDAP's name")

}
