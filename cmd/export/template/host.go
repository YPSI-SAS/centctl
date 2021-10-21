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
package template

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/host"
	"centctl/resources/service"
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
	Short: "Export template host",
	Long:  `Export template host of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		serviceTpl, _ := cmd.Flags().GetBool("serviceTpl")
		err := ExportTemplateHost(name, regex, file, all, serviceTpl, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportTemplateHost permits to export a host template of the centreon server
func ExportTemplateHost(name []string, regex string, file string, all bool, serviceTpl bool, debugV bool) error {
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
		templates := getAllTemplateHost(debugV)
		for _, a := range templates {
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
		err, templateHost := getHostTemplateInfo(n, debugV)
		if err != nil {
			return err
		}
		if templateHost.Name == "" {
			continue
		}

		//Write templateHost informations
		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,templateHost,\""+templateHost.Name+"\",\""+templateHost.Alias+"\",\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",address,\""+templateHost.Address+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",snmp_community,\""+templateHost.SnmpCommunity+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",snmp_version,\""+templateHost.SnmpVersion+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",timezone,\""+templateHost.Timezone+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",check_command,\""+templateHost.CheckCommand+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",check_command_arguments,\""+templateHost.CheckCommandArguments+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",check_period,\""+templateHost.CheckPeriod+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",max_check_attempts,\""+templateHost.MaxCheckAttempts+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",check_interval,\""+templateHost.CheckInterval+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",retry_check_interval,\""+templateHost.RetryCheckInterval+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",active_checks_enabled,\""+templateHost.ActiveChecksEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",passive_checks_enabled,\""+templateHost.PassiveChecksEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",notifications_enabled,\""+templateHost.NotificationsEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",contact_additive_inheritance,\""+templateHost.ContactAdditiveInheritance+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",cg_additive_inheritance,\""+templateHost.CgAdditiveInheritance+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",notification_options,\""+templateHost.NotificationOptions+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",notification_interval,\""+templateHost.NotificationInterval+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",notification_period,\""+templateHost.NotificationPeriod+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",first_notification_delay,\""+templateHost.FirstNotificationDelay+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",recovery_notification_delay,\""+templateHost.RecoveryNotificationDelay+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",obsess_over_host,\""+templateHost.ObsessOverHost+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",acknowledgement_timeout,\""+templateHost.AcknowledgementTimeout+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",check_freshness,\""+templateHost.CheckFreshness+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",freshness_threshold,\""+templateHost.FreshnessThreshold+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",flap_detection_enabled,\""+templateHost.FlapDetectionEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",low_flap_threshold,\""+templateHost.LowFlapThreshold+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",high_flap_threshold,\""+templateHost.HighFlapThreshold+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",retain_status_information,\""+templateHost.RetainStatusInformation+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",retain_nonstatus_information,\""+templateHost.RetainNonstatusInformation+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",stalking_options,\""+templateHost.StalkingOptions+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",event_handler_enabled,\""+templateHost.EventHandlerEnabled+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",event_handler,\""+templateHost.EventHandler+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",event_handler_arguments,\""+templateHost.EventHandlerArguments+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",action_url,\""+templateHost.ActionURL+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",notes,\""+templateHost.Notes+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",notes_url,\""+templateHost.NotesURL+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",icon_image,\""+templateHost.IconImage+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",icon_image_alt,\""+templateHost.IconImageAlt+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",statusmap_image,\""+templateHost.StatusMapImage+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",2d_coords,\""+templateHost.Coords2d+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",3d_coords,\""+templateHost.Coords3d+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",activate,\""+templateHost.Activate+"\"\n", file, writeFile)
		request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",comment,\""+templateHost.Comment+"\"\n", file, writeFile)

		//Write macros information
		if len(templateHost.Macros) != 0 {
			for _, m := range templateHost.Macros {
				if strings.Contains(m.Value, "\"") {
					m.Value = strings.ReplaceAll(m.Value, "\"", "'")
				}
				request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",macro,\""+m.Name+"|"+m.Value+"|"+m.IsPassword+"|"+m.Description+"\"\n", file, writeFile)
			}
		}

		//Write Templates information
		if len(templateHost.Templates) != 0 {
			for _, t := range templateHost.Templates {
				request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",template,\""+t.Name+"\"\n", file, writeFile)
			}
		}

		//Write ContactGroups information
		if len(templateHost.ContactGroups) != 0 {
			for _, c := range templateHost.ContactGroups {
				request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",contactgroup,\""+c.Name+"\"\n", file, writeFile)
			}
		}

		//Write Contacts information
		if len(templateHost.Contacts) != 0 {
			for _, c := range templateHost.Contacts {
				request.WriteValues("modify,templateHost,\""+templateHost.Name+"\",contact,\""+c.Name+"\"\n", file, writeFile)
			}
		}
	}
	if serviceTpl {
		for _, n := range name {
			serviceTpl := getAllTemplateServiceRelatedToTemplateHost(debugV, n)
			if len(serviceTpl) != 0 {
				err := ExportTemplateService(serviceTpl, "", file, false, debugV)
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

//The arguments impossible to get : linked_service_template|host_category
//getHostTemplateInfo permits to get all informations about a template host
func getHostTemplateInfo(name string, debugV bool) (error, host.ExportTemplateHost) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the host template
	values := name + ";name|alias|address|snmp_community|snmp_version|timezone|check_command|check_command_arguments|check_period|max_check_attempts" +
		"|check_interval|retry_check_interval|active_checks_enabled|passive_checks_enabled|notifications_enabled" +
		"|contact_additive_inheritance|cg_additive_inheritance|notification_options|notification_interval|notification_period" +
		"|first_notification_delay|recovery_notification_delay|obsess_over_host|acknowledgement_timeout|check_freshness|freshness_threshold" +
		"|flap_detection_enabled|low_flap_threshold|high_flap_threshold|retain_status_information|retain_nonstatus_information|stalking_options" +
		"|event_handler_enabled|event_handler|event_handler_arguments|notes_url|action_url|notes|icon_image|icon_image_alt|statusmap_image|2d_coords" +
		"|3d_coords|activate|comment"
	err, body := request.GeneriqueCommandV1Post("getparam", "HOST", values, "export template host", debugV, false, "")
	if err != nil {
		return err, host.ExportTemplateHost{}
	}
	var resultHost host.ExportHostTemplateResult
	json.Unmarshal(body, &resultHost)

	//Check if the host template is found
	if len(resultHost.HostTemplates) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, host.ExportTemplateHost{}
	}

	//Get the template of the host template
	err, body = request.GeneriqueCommandV1Post("gettemplate", "HTPL", name, "export template host", debugV, false, "")
	if err != nil {
		return err, host.ExportTemplateHost{}
	}
	var resultTemplate host.ExportResultHostTemplateTemplate
	json.Unmarshal(body, &resultTemplate)

	//Get the macro of the host template
	err, body = request.GeneriqueCommandV1Post("getmacro", "HTPL", name, "export template host", debugV, false, "")
	if err != nil {
		return err, host.ExportTemplateHost{}
	}
	var resultMacro host.ExportResultHostTemplateMacro
	json.Unmarshal(body, &resultMacro)

	//Get the contact group of the host template
	err, body = request.GeneriqueCommandV1Post("getcontactgroup", "HTPL", name, "export template host", debugV, false, "")
	if err != nil {
		return err, host.ExportTemplateHost{}
	}
	var resultContactGroup host.ExportResultHostTemplateContactGroup
	json.Unmarshal(body, &resultContactGroup)

	//Get the contact of the host
	err, body = request.GeneriqueCommandV1Post("getcontact", "HTPL", name, "export template host", debugV, false, "")
	if err != nil {
		return err, host.ExportTemplateHost{}
	}
	var resultContact host.ExportResultHostTemplateContact
	json.Unmarshal(body, &resultContact)

	//Get the host template
	host := resultHost.HostTemplates[0]
	host.Macros = resultMacro.Macros
	host.Templates = resultTemplate.Templates
	host.ContactGroups = resultContactGroup.ContactGroups
	host.Contacts = resultContact.Contacts

	return nil, host

}

//getAllTemplateHost permits to find all host template in the centreon server
func getAllTemplateHost(debugV bool) []host.ExportHostTemplate {
	//Get all host template
	err, body := request.GeneriqueCommandV1Post("show", "HTPL", "", "export template host", debugV, false, "")
	if err != nil {
		return []host.ExportHostTemplate{}
	}
	var resultTemplate host.ExportResultHostTemplate
	json.Unmarshal(body, &resultTemplate)

	return resultTemplate.Templates
}

func getAllTemplateServiceRelatedToTemplateHost(debugV bool, name string) []string {
	err, body := request.GeneriqueCommandV1Post("show", "STPL", "", "export template host", debugV, false, "")
	if err != nil {
		return []string{}
	}
	var resultTemplate service.ResultTemplate
	json.Unmarshal(body, &resultTemplate)
	result := []string{}
	for _, tpl := range resultTemplate.Templates {
		err, body = request.GeneriqueCommandV1Post("gethosttemplate", "STPL", tpl.Description, "export template host", debugV, false, "")
		if err != nil {
			return []string{}
		}
		var resultTplHost host.ResultTemplate
		json.Unmarshal(body, &resultTplHost)
		for _, tplHost := range resultTplHost.Templates {
			if strings.ToLower(name) == strings.ToLower(tplHost.Name) {
				result = append(result, tpl.Description)
			}

		}

	}

	return result
}

func init() {
	hostCmd.Flags().StringSliceP("name", "n", []string{}, "Host template's name (separate by a comma the multiple values)")
	hostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host template's name")
	hostCmd.Flags().Bool("serviceTpl", false, "Export all services templates related to this host template")

}
