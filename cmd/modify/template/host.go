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
package template

import (
	"centctl/colorMessage"
	"centctl/request"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Modify template host",
	Long:  `Modify template host of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		operation, _ := cmd.Flags().GetString("operation")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyTemplateHost(name, parameter, value, operation, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyTemplateHost permits to modify a host in the centreon server
func ModifyTemplateHost(name string, parameter string, value string, operation string, debugV bool, isImport bool, detail bool) error {
	colorRed := colorMessage.GetColorRed()
	var action string
	var values string
	isDefault := false
	object := "HTPL"

	operation = strings.ToLower(operation)
	if operation != "add" && operation != "del" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The operation's value must be : add or del")
		os.Exit(1)
	}

	switch strings.ToLower(parameter) {
	case "template":
		isDefault = true
	case "linkedservice":
		action = operation + "hosttemplate"
		values = value + ";" + name
		object = "STPL"
	case "category":
		action = operation + "member"
		values = value + ";" + name
		object = "HC"
	case "contactgroup":
		isDefault = true
	case "contact":
		isDefault = true
	case "macro":
		if operation == "add" {
			valueSplit := strings.Split(value, "|")
			if len(valueSplit) != 4 {
				fmt.Printf(colorRed, "ERROR: ")
				fmt.Println("The new value for macro must be of the form : macroName|macroValue|IsPassword(0 or 1)|macroDescription")
				os.Exit(1)
			}
			action = "setmacro"
			values = name + ";" + valueSplit[0] + ";" + valueSplit[1] + ";" + valueSplit[2] + ";" + valueSplit[3]
		} else {
			action = "delmacro"
			values = name + ";" + value
		}

	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value

	}

	if isDefault {
		action = operation + strings.ToLower(parameter)
		values = name + ";" + value
	}

	err := request.Modify(action, object, values, "modify template host", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the name of the host template to be modified")
	hostCmd.MarkFlagRequired("name")
	hostCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetTemplateHostNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	hostCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: template,contactgroup,contact,linkedservice,category,macro")
	hostCmd.MarkFlagRequired("parameter")
	hostCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "localhost", "ns_ip_address", "ns_activate", "engine_start_command", "engine_stop_command", "engine_restart_command", "engine_reload_command", "nagios_bin", "template", "contactgroup", "contact", "linkedservice", "category", "macro", "nagiostats_bin", "ssh_port", "broker_reload_command", "centreonbroker_module_path", "centreonbroker_cfg_path"}, cobra.ShellCompDirectiveDefault
	})
	hostCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is MACRO the value must be of the form : macroName|macroValue|IsPassword(0 or 1)|macroDescription")
	hostCmd.MarkFlagRequired("value")
	hostCmd.Flags().StringP("operation", "o", "", "To define the operation: add, del")
	hostCmd.MarkFlagRequired("operation")
	hostCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"add", "del"}, cobra.ShellCompDirectiveDefault
	})
}
