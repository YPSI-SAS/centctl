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
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package group

import (
	"centctl/colorMessage"
	"centctl/request"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Modify group service",
	Long:  `Modify group service of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyGroupService(name, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyGroupService permits to modify a service in the centreon server
func ModifyGroupService(name string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	colorRed := colorMessage.GetColorRed()
	//Creation of the request body
	var values string
	var action string
	switch strings.ToLower(parameter) {
	case "service":
		valueSplit := strings.Split(value, "|")
		if len(valueSplit) < 2 || len(valueSplit) > 2 {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println("The value when parameter is service must be of the form : hostName|serviceDescription")
			os.Exit(1)
		}
		values = name + ";" + valueSplit[0] + "," + valueSplit[1]
		action = "addservice"
	case "hostgroupservice":
		valueSplit := strings.Split(value, "|")
		if len(valueSplit) < 2 || len(valueSplit) > 2 {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println("The value when parameter is hostgroupservice must be of the form : hostgroupName|serviceDescription")
			os.Exit(1)
		}
		values = name + ";" + valueSplit[0] + "," + valueSplit[1]
		action = "addhostgroupservice"
	default:
		values = name + ";" + parameter + ";" + value
		action = "setparam"
	}

	err := request.Modify(action, "SG", values, "modify group service", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "To define the name of the service group to be modified")
	serviceCmd.MarkFlagRequired("name")
	serviceCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: service,hostgroupservice")
	serviceCmd.MarkFlagRequired("parameter")
	serviceCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "alias", "comment", "activate", "service", "hostgroupservice"}, cobra.ShellCompDirectiveDefault
	})
	serviceCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is service or hostgroupservice the value must be of the form : host|service or hostGroup|service")
	serviceCmd.MarkFlagRequired("value")
}
