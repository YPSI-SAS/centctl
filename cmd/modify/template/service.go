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

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Modify template service",
	Long:  `Modify template service of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		operation, _ := cmd.Flags().GetString("operation")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyTemplateService(name, parameter, value, operation, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyTemplateService permits to modify a service in the centreon server
func ModifyTemplateService(name string, parameter string, value string, operation string, debugV bool, isImport bool, detail bool) error {
	colorRed := colorMessage.GetColorRed()
	var action string
	var values string
	isDefault := false

	operation = strings.ToLower(operation)
	if operation != "add" && operation != "del" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The operation's value must be : add or del")
		os.Exit(1)
	}

	switch strings.ToLower(parameter) {
	case "trap":
		isDefault = true
	case "category":
		isDefault = true
	case "contactgroup":
		isDefault = true
	case "contact":
		isDefault = true
	case "linkedhost":
		action = operation + "hosttemplate"
		values = name + ";" + value
	case "macro":
		if operation == "add" {
			valueSplit := strings.Split(value, "|")
			if len(valueSplit) < 4 || len(valueSplit) > 4 {
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

	//Creation of the request body
	err := request.Modify(action, "STPL", values, "modify template service", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "To define the name of the service template to be modified")
	serviceCmd.MarkFlagRequired("name")
	serviceCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetTemplateServiceNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	serviceCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: trap,category,contactgroup,contact,linkedhost,macro")
	serviceCmd.MarkFlagRequired("parameter")
	serviceCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"activate", "description", "alias", "template", "is_volatile", "check_period", "check_command", "check_command_arguments", "trap", "category", "contactgroup", "contact", "linkedhost", "macro", "max_check_attempts", "normal_check_interval", "retry_check_interval", "active_checks_enabled", "passive_checks_enabled", "contact_additive_inheritance", "cg_additive_inheritance", "notification_interval", "notification_period", "notification_options", "first_notification_delay", "recovery_notification_delay", "parallelize_check", "obsess_over_service", "check_freshness", "freshness_threshold", "event_handler_enabled", "flap_detection_enabled", "process_perf_data", "retain_status_information", "retain_nonstatus_information", "stalking_options", "event_handler", "event_handler_arguments", "notes", "notes_url", "action_url", "icon_image", "icon_image_alt", "graphtemplate", "comment", "service_notification_options"}, cobra.ShellCompDirectiveDefault
	})
	serviceCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is MACRO the value must be of the form : macroName|macroValue|IsPassword(0 or 1)|macroDescription")
	serviceCmd.MarkFlagRequired("value")
	serviceCmd.Flags().StringP("operation", "o", "", "To define the operation: add, del")
	serviceCmd.MarkFlagRequired("operation")
	serviceCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"add", "del"}, cobra.ShellCompDirectiveDefault
	})
}
