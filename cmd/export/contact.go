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
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		alias, _ := cmd.Flags().GetStringSlice("alias")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportContact(alias, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportContact permits to export a contact of the centreon server
func ExportContact(alias []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(alias) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag alias or flag all or flag regex ")
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
		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,contact,\"" + contact.Name + "\"," + contact.Alias + "," + contact.Email + "," + randSeq(10) + ",\n")
		_, _ = f.WriteString("modify,contact," + contact.Alias + ",pager," + contact.Pager + "\n")
		_, _ = f.WriteString("modify,contact," + contact.Alias + ",access," + contact.GuiAccess + "\n")
		_, _ = f.WriteString("modify,contact," + contact.Alias + ",admin," + contact.Admin + "\n")
		_, _ = f.WriteString("modify,contact," + contact.Alias + ",activate," + contact.Activate + "\n")

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
	contactCmd.Flags().StringP("file", "f", "ExportContact.csv", "To define the name of the csv file")
	contactCmd.Flags().StringP("regex", "r", "", "The regex to apply on the contact's alias")

}
