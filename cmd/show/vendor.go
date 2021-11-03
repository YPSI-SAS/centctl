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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/vendor"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// vendorCmd represents the vendor command
var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Show one vendor's details",
	Long:  `Show one vendor's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowVendor(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowVendor permits to display the details of one cooleanrule
func ShowVendor(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "vendor", name, "show vendor", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the vendors contain into the response body
	vendors := vendor.DetailResult{}
	json.Unmarshal(body, &vendors)

	//Permits to find the good vendor in the array
	var VendorFind vendor.DetailVendor
	for _, v := range vendors.Vendors {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			VendorFind = v
		}
	}

	var server vendor.DetailServer
	if VendorFind.Name != "" {
		//Organization of data
		server = vendor.DetailServer{
			Server: vendor.DetailInformations{
				Name:   os.Getenv("SERVER"),
				Vendor: &VendorFind,
			},
		}
	} else {
		server = vendor.DetailServer{
			Server: vendor.DetailInformations{
				Name:   os.Getenv("SERVER"),
				Vendor: nil,
			},
		}
	}

	//Display details of the vendor
	displayVendor, err := display.DetailVendor(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayVendor)
	return nil
}

func init() {
	vendorCmd.Flags().StringP("name", "n", "", "To define the name of the vendor")
	vendorCmd.MarkFlagRequired("name")
	vendorCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetVendorNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
