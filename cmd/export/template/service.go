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
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Export template service",
	Long:  `Export template service of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportTemplateService(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportTemplateService permits to export a service template of the centreon server
func ExportTemplateService(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	//Check if the name of file contains the extension
	if !strings.Contains(file, ".csv") {
		file = file + ".csv"
	}

	//Create the file
	var f *os.File
	var err error
	if appendFile {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		f, err = os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	}
	defer f.Close()
	if err != nil {
		return err
	}

	if all || regex != "" {
		templates := getAllTemplateService(debugV)
		for _, a := range templates {
			if regex != "" {
				matched, err := regexp.MatchString(regex, a.Description)
				if err != nil {
					fmt.Printf(colorRed, "ERROR:")
					fmt.Println(err.Error())
					os.Exit(1)
				}
				if matched {
					name = append(name, a.Description)
				}
			} else {
				name = append(name, a.Description)
			}
		}
	}
	for _, n := range name {
		err, templateService := getServiceTemplateInfo(n, debugV)
		if err != nil {
			return err
		}
		if templateService.Description == "" {
			continue
		}

		//Write templateService informations
		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,templateService,\"" + templateService.Description + "\",\"" + templateService.Alias + "\",\"" + templateService.Template + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",check_command,\"" + templateService.CheckCommand + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",check_command_arguments,\"" + templateService.CheckCommandArguments + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",check_period,\"" + templateService.CheckPeriod + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",max_check_attempts,\"" + templateService.MaxCheckAttempts + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",normal_check_interval,\"" + templateService.NormalCheckInterval + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",retry_check_interval,\"" + templateService.RetryCheckInterval + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",active_checks_enabled,\"" + templateService.ActiveChecksEnabled + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",passive_checks_enabled,\"" + templateService.PassiveChecksEnabled + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",is_volatile,\"" + templateService.IsVolatile + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",notifications_enabled,\"" + templateService.NotificationsEnabled + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",contact_additive_inheritance,\"" + templateService.ContactAdditiveInheritance + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",cg_additive_inheritance,\"" + templateService.CgAdditiveInheritance + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",notification_options,\"" + templateService.NotificationOptions + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",notification_interval,\"" + templateService.NotificationInterval + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",notification_period,\"" + templateService.NotificationPeriod + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",first_notification_delay,\"" + templateService.FirstNotificationDelay + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",obsess_over_service,\"" + templateService.ObsessOverService + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",check_freshness,\"" + templateService.CheckFreshness + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",freshness_threshold,\"" + templateService.FreshnessThreshold + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",flap_detection_enabled,\"" + templateService.FlapDetectionEnabled + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",retain_status_information,\"" + templateService.RetainStatusInformation + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",retain_nonstatus_information,\"" + templateService.RetainNonstatusInformation + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",event_handler_enabled,\"" + templateService.EventHandlerEnabled + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",event_handler,\"" + templateService.EventHandler + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",event_handler_arguments,\"" + templateService.EventHandlerArguments + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",action_url,\"" + templateService.ActionURL + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",notes,\"" + templateService.Notes + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",notes_url,\"" + templateService.NotesURL + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",icon_image,\"" + templateService.IconImage + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",icon_image_alt,\"" + templateService.IconImageAlt + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",activate,\"" + templateService.Activate + "\"\n")
		_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",comment,\"" + templateService.Comment + "\"\n")

		//Write macros information
		if len(templateService.Macros) != 0 {
			for _, m := range templateService.Macros {
				if strings.Contains(m.Value, "\"") {
					m.Value = strings.ReplaceAll(m.Value, "\"", "'")
				}
				if strings.Contains(m.Name, "$_SERVICE") {
					m.Name = m.Name[9 : len(m.Name)-1]
				}
				_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",macro,\"" + m.Name + "|" + m.Value + "|" + m.IsPassword + "|" + m.Description + "\"\n")
			}
		}

		//Write ContactGroups information
		if len(templateService.ContactGroups) != 0 {
			for _, c := range templateService.ContactGroups {
				_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",contactgroup,\"" + c.Name + "\"\n")
			}
		}

		//Write Contacts information
		if len(templateService.Contacts) != 0 {
			for _, c := range templateService.Contacts {
				_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",contact,\"" + c.Name + "\"\n")
			}
		}

		//Write Traps information
		if len(templateService.Traps) != 0 {
			for _, t := range templateService.Traps {
				_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",trap,\"" + t.Name + "\"\n")
			}
		}

		//Write Categories information
		if len(templateService.Categories) != 0 {
			for _, c := range templateService.Categories {
				_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",category,\"" + c.Name + "\"\n")
			}
		}

		//Write HostTemplates information
		if len(templateService.HostTemplates) != 0 {
			for _, h := range templateService.HostTemplates {
				_, _ = f.WriteString("modify,templateService,\"" + templateService.Description + "\",linkedhost,\"" + h.Name + "\"\n")
			}
		}
	}
	return nil
}

//The arguments impossible to get : recovery_notification_delay|acknowledgement_timeout|low_flap_threshold|high_flap_threshold|stalking_options|graphtemplate
//getServiceTemplateInfo permits to get all informations about a template Service
func getServiceTemplateInfo(name string, debugV bool) (error, service.ExportTemplateService) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the service template
	values := name + ";description|alias|template|check_command|check_command_arguments|check_period|max_check_attempts|normal_check_interval|" +
		"retry_check_interval|active_checks_enabled|passive_checks_enabled|is_volatile|notifications_enabled|contact_additive_inheritance|" +
		"cg_additive_inheritance|notification_options|notification_interval|notification_period|first_notification_delay|obsess_over_service|" +
		"check_freshness|freshness_threshold|flap_detection_enabled|retain_status_information|retain_nonstatus_information|event_handler_enabled|" +
		"event_handler|event_handler_arguments|notes|icon_image|icon_image_alt|notes_url|action_url|activate|comment"
	err, body := request.GeneriqueCommandV1Post("getparam", "STPL", values, "export template service", debugV, false, "")
	if err != nil {
		return err, service.ExportTemplateService{}
	}
	var resultServiceTemplate service.ExportServiceTemplateResult
	json.Unmarshal(body, &resultServiceTemplate)

	//Check if the service template is found
	if len(resultServiceTemplate.TemplateServices) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, service.ExportTemplateService{}
	}

	//Get the macro of the service template
	err, body = request.GeneriqueCommandV1Post("getmacro", "STPL", name, "export template service", debugV, false, "")
	if err != nil {
		return err, service.ExportTemplateService{}
	}
	var resultMacro service.ExportResultServiceTemplateMacro
	json.Unmarshal(body, &resultMacro)

	//Get the contact group of the service template
	err, body = request.GeneriqueCommandV1Post("getcontactgroup", "STPL", name, "export template service", debugV, false, "")
	if err != nil {
		return err, service.ExportTemplateService{}
	}
	var resultContactGroup service.ExportResultServiceTemplateContactGroup
	json.Unmarshal(body, &resultContactGroup)

	//Get the contact of the service template
	err, body = request.GeneriqueCommandV1Post("getcontact", "STPL", name, "export template service", debugV, false, "")
	if err != nil {
		return err, service.ExportTemplateService{}
	}
	var resultContact service.ExportResultServiceTemplateContact
	json.Unmarshal(body, &resultContact)

	//Get the trap of the service template
	err, body = request.GeneriqueCommandV1Post("gettrap", "STPL", name, "export template service", debugV, false, "")
	if err != nil {
		return err, service.ExportTemplateService{}
	}
	var resultTrap service.ExportResultServiceTemplateTrap
	json.Unmarshal(body, &resultTrap)

	//Get the category of the service template
	err, body = request.GeneriqueCommandV1Post("getcategory", "STPL", name, "export template service", debugV, false, "")
	if err != nil {
		return err, service.ExportTemplateService{}
	}
	var resultCategory service.ExportResultServiceTemplateCategory
	json.Unmarshal(body, &resultCategory)

	//Get the host template of the service template
	err, body = request.GeneriqueCommandV1Post("gethosttemplate", "STPL", name, "export template service", debugV, false, "")
	if err != nil {
		return err, service.ExportTemplateService{}
	}
	var resultHostTemplate service.ExportResultServiceTemplateHostTemplate
	json.Unmarshal(body, &resultHostTemplate)

	//Get the  and the member
	serviceTemplate := resultServiceTemplate.TemplateServices[0]
	serviceTemplate.Macros = resultMacro.Macros
	serviceTemplate.ContactGroups = resultContactGroup.ContactGroups
	serviceTemplate.Contacts = resultContact.Contacts
	serviceTemplate.Traps = resultTrap.Traps
	serviceTemplate.Categories = resultCategory.Categories
	serviceTemplate.HostTemplates = resultHostTemplate.HostTemplates

	return nil, serviceTemplate

}

//getAllTemplateService permits to find all service template in the centreon server
func getAllTemplateService(debugV bool) []service.ExportTemplateService {
	//Get all service template
	err, body := request.GeneriqueCommandV1Post("show", "STPL", "", "export template service", debugV, false, "")
	if err != nil {
		return []service.ExportTemplateService{}
	}
	var resultTemplate service.ExportServiceTemplateResult
	json.Unmarshal(body, &resultTemplate)

	return resultTemplate.TemplateServices
}

func init() {
	serviceCmd.Flags().StringSliceP("name", "n", []string{}, "Service template's name (separate by a comma the multiple values)")
	serviceCmd.Flags().StringP("file", "f", "ExportServiceTemplate.csv", "To define the name of the csv file")
	serviceCmd.Flags().StringP("regex", "r", "", "The regex to apply on the service template's name")

}
