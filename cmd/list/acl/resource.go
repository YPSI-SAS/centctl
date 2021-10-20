/*MIT License

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
LIABILITY, WHETHER IN AN ANCTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package acl

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/ACL"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "List ACL resource",
	Long:  `List ACL resource of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListACLResource(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListACLResource permits to display the array of ACL resource return by the API
func ListACLResource(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "ACLRESOURCE", "", "list ACL resource", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the ACL resources contain into the response body
	resources := ACL.ResultResource{}
	json.Unmarshal(body, &resources)
	finalResources := resources.Resources
	if regex != "" {
		finalResources = deleteResource(finalResources, regex)
	}

	//Sort ACL resources based on their ID
	sort.SliceStable(finalResources, func(i, j int) bool {
		return strings.ToLower(finalResources[i].Name) < strings.ToLower(finalResources[j].Name)
	})

	server := ACL.ResourceServer{
		Server: ACL.ResourceInformations{
			Name:      os.Getenv("SERVER"),
			Resources: finalResources,
		},
	}

	//Display all ACL resources
	displayACLResource, err := display.ACLResource(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayACLResource)

	return nil
}

func deleteResource(resources []ACL.Resource, regex string) []ACL.Resource {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range resources {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			resources[index] = s
			index++
		}
	}
	return resources[:index]
}

func init() {
	resourceCmd.Flags().StringP("regex", "r", "", "The regex to apply on the resource's name")

}
