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
	"centctl/resources/vendor"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// vendorCmd represents the vendor command
var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "List the vendors",
	Long:  `List the vendors of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListVendor(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListVendor permits to display the array of vendors return by the API
func ListVendor(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "vendor", "", "list vendor", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the vendors contain into the response body
	vendors := vendor.Result{}
	json.Unmarshal(body, &vendors)
	finalVendors := vendors.Vendors
	if regex != "" {
		finalVendors = deleteVendor(finalVendors, regex)
	}

	//Sort vendors based on their ID
	sort.SliceStable(finalVendors, func(i, j int) bool {
		valI, _ := strconv.Atoi(finalVendors[i].ID)
		valJ, _ := strconv.Atoi(finalVendors[j].ID)
		return valI < valJ
	})

	//Organization of data
	server := vendor.Server{
		Server: vendor.Informations{
			Name:    os.Getenv("SERVER"),
			Vendors: finalVendors,
		},
	}

	//Display all vendors
	displayVendor, err := display.Vendor(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayVendor)

	return nil
}

func deleteVendor(vendors []vendor.Vendor, regex string) []vendor.Vendor {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range vendors {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			vendors[index] = s
			index++
		}
	}
	return vendors[:index]
}

func init() {
	vendorCmd.Flags().StringP("regex", "r", "", "The regex to apply on the vendor's name")
}
