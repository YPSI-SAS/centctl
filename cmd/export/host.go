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
	"centctl/cmd/export/group"
	"centctl/cmd/export/template"
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Export host",
	Long:  `Export host of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		services, _ := cmd.Flags().GetBool("services")
		groups, _ := cmd.Flags().GetBool("groups")
		hostTpl, _ := cmd.Flags().GetBool("hostTpl")
		serviceTpl, _ := cmd.Flags().GetBool("serviceTpl")
		commands, _ := cmd.Flags().GetBool("commands")
		err := ExportHost(name, regex, file, all, debugV, services, groups, hostTpl, serviceTpl, commands)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportHost permits to export a host of the centreon server
func ExportHost(name []string, regex string, file string, all bool, debugV bool, services bool, groups bool, hostTpl bool, serviceTpl bool, commands bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		hosts := getAllHost(debugV)
		for _, a := range hosts {
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
		err, host := getHostInfo(n, debugV)
		if err != nil {
			return err
		}
		if host.Name == "" {
			continue
		}

		if groups {
			if len(host.HostGroups) != 0 {
				for _, h := range host.HostGroups {
					err := group.ExportGroupHost([]string{h.Name}, "", file, false, debugV)
					if err != nil {
						return err
					}
				}
			}

		}

		if hostTpl {
			if len(host.Templates) != 0 {
				for _, h := range host.Templates {
					err := template.ExportTemplateHost([]string{h.Name}, "", file, false, serviceTpl, debugV)
					if err != nil {
						return err
					}
				}
			}

		}

		if commands {
			if host.CheckCommand != "" {
				err := ExportCommand([]string{host.CheckCommand}, "", "", file, false, debugV)
				if err != nil {
					return err
				}
			}

		}

		//Write host informations
		request.WriteValues("\n", file, writeFile)
		if len(host.Templates) != 0 {
			request.WriteValues("add,host,\""+host.Name+"\",\""+host.Alias+"\",\""+host.Address+"\",\""+host.Templates[0].Name+"\",\""+host.Instance.Name+"\",\n", file, writeFile)
		} else {
			request.WriteValues("add,host,\""+host.Name+"\",\""+host.Alias+"\",\""+host.Address+"\",,\""+host.Instance.Name+"\",\n", file, writeFile)
		}
		request.WriteValues("modify,host,\""+host.Name+"\",snmp_community,\""+host.SnmpCommunity+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",snmp_version,\""+host.SnmpVersion+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",timezone,\""+host.Timezone+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",check_command,\""+host.CheckCommand+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",check_command_arguments,\""+host.CheckCommandArguments+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",check_period,\""+host.CheckPeriod+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",max_check_attempts,\""+host.MaxCheckAttempts+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",check_interval,\""+host.CheckInterval+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",retry_check_interval,\""+host.RetryCheckInterval+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",active_checks_enabled,\""+host.ActiveChecksEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",passive_checks_enabled,\""+host.PassiveChecksEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",notifications_enabled,\""+host.NotificationsEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",contact_additive_inheritance,\""+host.ContactAdditiveInheritance+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",cg_additive_inheritance,\""+host.CgAdditiveInheritance+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",notification_options,\""+host.NotificationOptions+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",notification_interval,\""+host.NotificationInterval+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",notification_period,\""+host.NotificationPeriod+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",first_notification_delay,\""+host.FirstNotificationDelay+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",recovery_notification_delay,\""+host.RecoveryNotificationDelay+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",obsess_over_host,\""+host.ObsessOverHost+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",acknowledgement_timeout,\""+host.AcknowledgementTimeout+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",check_freshness,\""+host.CheckFreshness+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",freshness_threshold,\""+host.FreshnessThreshold+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",flap_detection_enabled,\""+host.FlapDetectionEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",low_flap_threshold,\""+host.LowFlapThreshold+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",high_flap_threshold,\""+host.HighFlapThreshold+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",retain_status_information,\""+host.RetainStatusInformation+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",retain_nonstatus_information,\""+host.RetainNonstatusInformation+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",stalking_options,\""+host.StalkingOptions+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",event_handler_enabled,\""+host.EventHandlerEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",event_handler,\""+host.EventHandler+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",event_handler_arguments,\""+host.EventHandlerArguments+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",action_url,\""+host.ActionURL+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",notes,\""+host.Notes+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",notes_url,\""+host.NotesURL+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",icon_image,\""+host.IconImage+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",icon_image_alt,\""+host.IconImageAlt+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",statusmap_image,\""+host.StatusMapImage+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",geo_coords,\""+host.GeoCoords+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",2d_coords,\""+host.Coords2d+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",3d_coords,\""+host.Coords3d+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",activate,\""+host.Activate+"\"\n", file, writeFile)
		request.WriteValues("modify,host,\""+host.Name+"\",comment,\""+host.Comment+"\"\n", file, writeFile)

		//Write macros information
		if len(host.Macros) != 0 {
			for _, m := range host.Macros {
				if strings.Contains(m.Value, "\"") {
					m.Value = strings.ReplaceAll(m.Value, "\"", "'")
				}
				request.WriteValues("modify,host,\""+host.Name+"\",macro,\""+m.Name+"|"+m.Value+"|"+m.IsPassword+"|"+m.Description+"\"\n", file, writeFile)
			}
		}

		//Write Templates information
		if len(host.Templates) != 0 {
			for _, t := range host.Templates {
				request.WriteValues("modify,host,\""+host.Name+"\",template,\""+t.Name+"\"\n", file, writeFile)
			}
		}

		//Write Parents information
		if len(host.Parents) != 0 {
			for _, p := range host.Parents {
				request.WriteValues("modify,host,\""+host.Name+"\",parent,\""+p.Name+"\"\n", file, writeFile)
			}
		}

		//Write Childs information
		if len(host.Childs) != 0 {
			for _, c := range host.Childs {
				request.WriteValues("modify,host,\""+host.Name+"\",child,\""+c.Name+"\"\n", file, writeFile)
			}
		}

		//Write ContactGroups information
		if len(host.ContactGroups) != 0 {
			for _, c := range host.ContactGroups {
				request.WriteValues("modify,host,\""+host.Name+"\",contactgroup,\""+c.Name+"\"\n", file, writeFile)
			}
		}

		//Write Contacts information
		if len(host.Contacts) != 0 {
			for _, c := range host.Contacts {
				request.WriteValues("modify,host,\""+host.Name+"\",contact,\""+c.Name+"\"\n", file, writeFile)
			}
		}

		//Write HostGroups information
		if len(host.HostGroups) != 0 {
			for _, h := range host.HostGroups {
				request.WriteValues("modify,host,\""+host.Name+"\",hostgroup,\""+h.Name+"\"\n", file, writeFile)
			}
		}
	}
	if services {
		err := ExportService([]string{}, file, name, true, debugV)
		if err != nil {
			return err
		}
	}

	return nil
}

//The arguments impossible to get : hostcategorie
//getHostInfo permits to get all informations about a host
func getHostInfo(name string, debugV bool) (error, host.ExportHost) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the host
	values := name + ";name|alias|address|snmp_community|snmp_version|timezone|check_command|check_command_arguments|check_period|max_check_attempts" +
		"|check_interval|retry_check_interval|active_checks_enabled|passive_checks_enabled|notifications_enabled" +
		"|contact_additive_inheritance|cg_additive_inheritance|notification_options|notification_interval|notification_period" +
		"|first_notification_delay|recovery_notification_delay|obsess_over_host|acknowledgement_timeout|check_freshness|freshness_threshold" +
		"|flap_detection_enabled|low_flap_threshold|high_flap_threshold|retain_status_information|retain_nonstatus_information|stalking_options" +
		"|event_handler_enabled|event_handler|event_handler_arguments|notes_url|action_url|notes|icon_image|icon_image_alt|statusmap_image|geo_coords|2d_coords" +
		"|3d_coords|activate|comment"
	err, body := request.GeneriqueCommandV1Post("getparam", "HOST", values, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultHost host.ExportHostResult
	json.Unmarshal(body, &resultHost)

	//Check if the host is found
	if len(resultHost.Hosts) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, host.ExportHost{}
	}

	//Get the instance of the host
	err, body = request.GeneriqueCommandV1Post("showinstance", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultInstance host.ExportResultHostInstance
	json.Unmarshal(body, &resultInstance)

	//Get the template of the host
	err, body = request.GeneriqueCommandV1Post("gettemplate", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultTemplate host.ExportResultHostTemplate
	json.Unmarshal(body, &resultTemplate)

	//Get the macro of the host
	err, body = request.GeneriqueCommandV1Post("getmacro", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultMacro host.ExportResultHostMacro
	json.Unmarshal(body, &resultMacro)

	//Get the parent of the host
	err, body = request.GeneriqueCommandV1Post("getparent", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultParent host.ExportResultHostParent
	json.Unmarshal(body, &resultParent)

	//Get the child of the host
	err, body = request.GeneriqueCommandV1Post("getchild", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultChild host.ExportResultHostChild
	json.Unmarshal(body, &resultChild)

	//Get the contact group of the host
	err, body = request.GeneriqueCommandV1Post("getcontactgroup", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultContactGroup host.ExportResultHostContactGroup
	json.Unmarshal(body, &resultContactGroup)

	//Get the contact of the host
	err, body = request.GeneriqueCommandV1Post("getcontact", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultContact host.ExportResultHostContact
	json.Unmarshal(body, &resultContact)

	//Get the hostgroup of the host
	err, body = request.GeneriqueCommandV1Post("gethostgroup", "HOST", name, "export host", debugV, false, "")
	if err != nil {
		return err, host.ExportHost{}
	}
	var resultHostGroup host.ExportResultHostHostGroup
	json.Unmarshal(body, &resultHostGroup)

	//Get the host informations
	host := resultHost.Hosts[0]
	host.Instance = resultInstance.Instances[0]
	host.Macros = resultMacro.Macros
	host.Templates = resultTemplate.Templates
	host.ContactGroups = resultContactGroup.ContactGroups
	host.Contacts = resultContact.Contacts
	host.HostGroups = resultHostGroup.HostGroups
	host.Parents = resultParent.Parents
	host.Childs = resultChild.Childs

	return nil, host

}

//getAllHost permits to find all host in the centreon server
func getAllHost(debugV bool) []host.ExportHost {
	//Get all host
	err, body := request.GeneriqueCommandV1Post("show", "host", "", "export host", debugV, false, "")
	if err != nil {
		return []host.ExportHost{}
	}
	var resultHost host.ExportHostResult
	json.Unmarshal(body, &resultHost)

	return resultHost.Hosts
}

func init() {
	hostCmd.Flags().StringSliceP("name", "n", []string{}, "Host's name (separate by a comma the multiple values)")
	hostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host's name")
	hostCmd.Flags().Bool("services", false, "Export all services related to this host")
	hostCmd.Flags().Bool("groups", false, "Export all hostgroups related to this host")
	hostCmd.Flags().Bool("hostTpl", false, "Export all host templates related to this host")
	hostCmd.Flags().Bool("serviceTpl", false, "Export all services templates related to the host template")
	hostCmd.Flags().Bool("commands", false, "Export all commands related to the host")

}
