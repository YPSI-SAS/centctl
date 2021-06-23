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
	"centctl/resources/vendor"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// vendorCmd represents the vendor command
var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Export vendor",
	Long:  `Export vendor of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportVendor(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportVendor permits to export a vendor of the centreon server
func ExportVendor(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
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
		templates := getAllVendor(debugV)
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
		err, vendor := getVendorInfo(n, debugV)
		if err != nil {
			return err
		}
		if vendor.Name == "" {
			continue
		}

		//Write vendor informations
		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,vendor,\"" + vendor.Name + "\",\"" + vendor.Alias + "\"\n")

	}
	return nil
}

//The arguments impossible to get : description
//getVendorInfo permits to get all informations about a vendor
func getVendorInfo(name string, debugV bool) (error, vendor.ExportVendor) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "vendor", name, "export vendor", debugV, false, "")
	if err != nil {
		return err, vendor.ExportVendor{}
	}
	var resultVendor vendor.ExportResultVendor
	json.Unmarshal(body, &resultVendor)

	vendor := vendor.ExportVendor{}
	find := false
	for _, g := range resultVendor.Vendors {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			vendor = g
			find = true
		}
	}
	//Check if the vendor  is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, vendor
	}

	return nil, vendor

}

//getAllVendor permits to find all vendor in the centreon server
func getAllVendor(debugV bool) []vendor.ExportVendor {
	//Get all vendor
	err, body := request.GeneriqueCommandV1Post("show", "vendor", "", "export vendor", debugV, false, "")
	if err != nil {
		return []vendor.ExportVendor{}
	}
	var resultVendor vendor.ExportResultVendor
	json.Unmarshal(body, &resultVendor)

	return resultVendor.Vendors
}

func init() {
	vendorCmd.Flags().StringSliceP("name", "n", []string{}, "vendor's name (separate by a comma the multiple values)")
	vendorCmd.Flags().StringP("file", "f", "ExportVendor.csv", "To define the name of the csv file")
	vendorCmd.Flags().StringP("regex", "r", "", "The regex to apply on the vendor's name")

}
