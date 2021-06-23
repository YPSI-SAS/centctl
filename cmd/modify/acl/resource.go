/*MIT License

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
LIABILITY, WHETHER IN AN resource OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package acl

import (
	"centctl/request"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Modfiy a ACL resource",
	Long:  `Modfiy a ACL resource into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := ModifyACLResource(name, parameter, value, debugV, apply, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyACLResource permits to modify a ACL resource in the centreon server
func ModifyACLResource(name string, parameter string, value string, debugV bool, apply bool, isImport bool, detail bool) error {
	var action string
	var values string

	switch strings.ToLower(parameter) {
	case "grant_host":
		action = "grant_host"
		values = name + ";" + value
	case "grant_hostgroup":
		action = "grant_hostgroup"
		values = name + ";" + value
	case "grant_servicegroup":
		action = "grant_servicegroup"
		values = name + ";" + value
	case "grant_metaservice":
		action = "grant_metaservice"
		values = name + ";" + value
	case "addhostexclusion":
		action = "addhostexclusion"
		values = name + ";" + value
	case "revoke_host":
		action = "revoke_host"
		values = name + ";" + value
	case "revoke_hostgroup":
		action = "revoke_hostgroup"
		values = name + ";" + value
	case "revoke_servicegroup":
		action = "revoke_servicegroup"
		values = name + ";" + value
	case "revoke_metaservice":
		action = "revoke_metaservice"
		values = name + ";" + value
	case "delhostexclusion":
		action = "delhostexclusion"
		values = name + ";" + value
	case "addfilter_instance":
		action = "addfilter_instance"
		values = name + ";" + value
	case "addfilter_hostcategory":
		action = "addfilter_hostcategory"
		values = name + ";" + value
	case "addfilter_servicecategory":
		action = "addfilter_servicecategory"
		values = name + ";" + value
	case "delfilter_instance":
		action = "delfilter_instance"
		values = name + ";" + value
	case "delfilter_hostcategory":
		action = "delfilter_hostcategory"
		values = name + ";" + value
	case "delfilter_servicecategory":
		action = "delfilter_servicecategory"
		values = name + ";" + value
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value

	}

	err := request.Modify(action, "ACLRESOURCE", values, "modify ACL resource", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	resourceCmd.Flags().StringP("name", "n", "", "To define the name of the ACL resource to be modified")
	resourceCmd.MarkFlagRequired("name")
	resourceCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or command in \"grant and revoke\" section")
	resourceCmd.MarkFlagRequired("parameter")
	resourceCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. Use | for defining multiple resources.")
	resourceCmd.MarkFlagRequired("value")
	resourceCmd.Flags().Bool("apply", false, "Export configuration of the poller")

}
