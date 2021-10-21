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
package export

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/contact"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// contactCmd represents the contact command
var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Export contact",
	Long:  `Export contact of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		alias, _ := cmd.Flags().GetStringSlice("alias")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportContact(alias, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportContact permits to export a contact of the centreon server
func ExportContact(alias []string, regex string, file string, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(alias) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag alias or flag all or flag regex ")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		templates := getAllContact(debugV)
		for _, a := range templates {
			if regex != "" {
				matched, err := regexp.MatchString(regex, a.Alias)
				if err != nil {
					fmt.Printf(colorRed, "ERROR:")
					fmt.Println(err.Error())
					os.Exit(1)
				}
				if matched {
					alias = append(alias, a.Alias)
				}
			} else {
				alias = append(alias, a.Alias)
			}
		}
	}
	for _, n := range alias {
		err, contact := getContactInfo(n, debugV)
		if err != nil {
			return err
		}
		if contact.Alias == "" {
			continue
		}

		rand.Seed(time.Now().UnixNano())
		//Write contact informations
		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,contact,\""+contact.Name+"\","+contact.Alias+","+contact.Email+","+randSeq(10)+",\n", file, writeFile)
		request.WriteValues("modify,contact,"+contact.Alias+",pager,"+contact.Pager+"\n", file, writeFile)
		request.WriteValues("modify,contact,"+contact.Alias+",access,"+contact.GuiAccess+"\n", file, writeFile)
		request.WriteValues("modify,contact,"+contact.Alias+",admin,"+contact.Admin+"\n", file, writeFile)
		request.WriteValues("modify,contact,"+contact.Alias+",activate,"+contact.Activate+"\n", file, writeFile)

	}
	return nil
}

//The arguments impossible to get : all in setparam table
//getContactInfo permits to get all informations about a contact
func getContactInfo(alias string, debugV bool) (error, contact.ExportContact) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "contact", alias, "export contact", debugV, false, "")
	if err != nil {
		return err, contact.ExportContact{}
	}
	var ExportResultContact contact.ExportResultContact
	json.Unmarshal(body, &ExportResultContact)

	contact := contact.ExportContact{}
	find := false
	for _, g := range ExportResultContact.Contacts {
		if strings.ToLower(g.Alias) == strings.ToLower(alias) {
			contact = g
			find = true
		}
	}
	//Check if the contact  is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + alias)
		return nil, contact
	}

	return nil, contact

}

//getAllContact permits to find all contact in the centreon server
func getAllContact(debugV bool) []contact.ExportContact {
	//Get all contact
	err, body := request.GeneriqueCommandV1Post("show", "contact", "", "export contact", debugV, false, "")
	if err != nil {
		return []contact.ExportContact{}
	}
	var resultContact contact.ExportResultContact
	json.Unmarshal(body, &resultContact)

	return resultContact.Contacts
}

var vals = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?$*)(][}{#")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = vals[rand.Intn(len(vals))]
	}
	return string(b)
}

func init() {
	contactCmd.Flags().StringSliceP("alias", "a", []string{}, "contact's alias (separate by a comma the multiple values)")
	contactCmd.Flags().StringP("regex", "r", "", "The regex to apply on the contact's alias")

}
