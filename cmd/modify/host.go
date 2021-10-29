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
package modify

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
	Short: "Modfiy a host",
	Long:  `Modfiy a host into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := ModifyHost(name, parameter, value, debugV, apply, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyHost permits to modify a host in the centreon server
func ModifyHost(name string, parameter string, value string, debugV bool, apply bool, isImport bool, detail bool) error {
	colorGreen := colorMessage.GetColorGreen()
	colorRed := colorMessage.GetColorRed()
	isTemplate := false
	var action string
	var values string
	object := "host"

	switch strings.ToLower(parameter) {
	case "instance":
		action = "setinstance"
		values = name + ";" + value
	case "template":
		action = "addtemplate"
		values = name + ";" + value
		isTemplate = true
	case "parent":
		action = "addparent"
		values = name + ";" + value
	case "child":
		action = "addchild"
		values = name + ";" + value
	case "contactgroup":
		action = "addcontactgroup"
		values = name + ";" + value
	case "contact":
		action = "addcontact"
		values = name + ";" + value
	case "hostgroup":
		action = "addhostgroup"
		values = name + ";" + value
	case "hostcategorie":
		action = "addmember"
		values = value + ";" + name
		object = "HC"
	case "macro":
		valueSplit := strings.Split(value, "|")
		if len(valueSplit) < 4 || len(valueSplit) > 4 {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println("The new value for macro must be of the form : macroName|macroValue|IsPassword(0 or 1)|macroDescription")
			os.Exit(1)
		}
		action = "setmacro"
		values = name + ";" + valueSplit[0] + ";" + valueSplit[1] + ";" + valueSplit[2] + ";" + valueSplit[3]
		object = "HOST"
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value

	}

	poller := ""
	var err error
	if apply {
		//Find the name of the host poller
		client := request.NewClientV1(os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&search=" + name)
		for poller == "" {
			poller, err = client.NamePollerHost(name, debugV)
			if err != nil {
				return err
			}
		}
	}

	err = request.Modify(action, object, values, "modify host", name, parameter, detail, debugV, apply, poller, isImport)
	if err != nil {
		return err
	}

	if isTemplate {
		err, _ = request.GeneriqueCommandV1Post("APPLYTPL", "HOST", name, "centctl modify host", debugV, apply, poller)
		if err != nil {
			return err
		}

		fmt.Printf(colorGreen, "INFO: ")
		fmt.Printf("The template of the host %v is applied\n", name)
	}

	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the name of the host to be modified")
	hostCmd.MarkFlagRequired("name")
	hostCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: instance,template,parent,child,contactgroup,contact,hostgroup,hostcategorie,macro")
	hostCmd.MarkFlagRequired("parameter")
	hostCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"instance", "template", "parent", "child", "contactgroup", "contact", "hostgroup", "hostcategorie", "macro", "geo_coords", "2d_coords", "3d_coords", "action_url", "activate", "active_checks_enabled", "acknowledgement_timeout", "address", "alias", "check_command", "check_command_arguments", "check_interval", "check_freshness", "check_period", "contact_additive_inheritance", "cg_additive_inheritance", "event_handler", "event_handler_arguments", "event_handler_enabled", "first_notification_delay", "flap_detection_enabled", "flap_detection_options", "host_high_flap_threshold", "host_low_flap_threshold", "icon_image", "icon_image_alt", "max_check_attempts", "name", "notes", "notes_url", "notifications_enabled", "notification_interval", "notification_options", "notification_period", "recovery_notification_delay", "obsess_over_host", "passive_checks_enabled", "retain_nonstatus_information", "retain_status_information", "retry_check_interval", "snmp_community", "snmp_version", "stalking_options", "statusmap_image", "host_notification_options", "timezone", "comment"}, cobra.ShellCompDirectiveDefault
	})
	hostCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is MACRO the value must be of the form : macroName|macroValue|IsPassword(0 or 1)|macroDescription")
	hostCmd.MarkFlagRequired("value")
	hostCmd.Flags().Bool("apply", false, "Export configuration of the poller")
}
