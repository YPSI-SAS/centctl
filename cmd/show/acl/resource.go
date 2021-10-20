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
	"centctl/display"
	"centctl/request"
	"centctl/resources/ACL"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Show one ACL resource's details",
	Long:  `Show one ACL resource's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowACLResource(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowACLResource permits to display the details of one ACL resource
func ShowACLResource(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLRESOURCE", name, "show ACL resource", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL resource contain into the response body
	menusACL := ACL.DetailResultResource{}
	json.Unmarshal(body, &menusACL)

	//Permits to find the good ACL resource in the array
	var GroupFind ACL.DetailResource
	for _, v := range menusACL.Resources {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			GroupFind = v
		}
	}

	var server ACL.DetailResourceServer
	if GroupFind.Name != "" {
		//Organization of data
		server = ACL.DetailResourceServer{
			Server: ACL.DetailResourceInformations{
				Name:     os.Getenv("SERVER"),
				Resource: &GroupFind,
			},
		}
	} else {
		//Organization of data
		server = ACL.DetailResourceServer{
			Server: ACL.DetailResourceInformations{
				Name:     os.Getenv("SERVER"),
				Resource: nil,
			},
		}
	}

	//Display details of the ACL resource
	displayACLResource, err := display.DetailACLResource(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLResource)
	return nil
}

func init() {
	resourceCmd.Flags().StringP("name", "n", "", "To define the ACL resource which will show")
	resourceCmd.MarkFlagRequired("name")
}
