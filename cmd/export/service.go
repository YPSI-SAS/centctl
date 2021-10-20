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
package export

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Export service",
	Long:  `Export service of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		name, _ := cmd.Flags().GetStringSlice("name")
		hostFilter, _ := cmd.Flags().GetStringSlice("hostFilter")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportService(name, file, hostFilter, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportService permits to export a service of the centreon server
func ExportService(name []string, file string, hostFilter []string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all")
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

	if all {
		services := getAllService(debugV)
		for _, s := range services {
			if len(hostFilter) != 0 {
				for _, h := range hostFilter {
					if strings.ToLower(h) == strings.ToLower(s.HostName) {
						name = append(name, s.HostName+"|"+s.Description)
					}
				}
			} else {
				name = append(name, s.HostName+"|"+s.Description)
			}
		}
	}
	for _, n := range name {
		var hostName string
		var descriptionService string
		if strings.Contains(n, "|") {
			nameSplit := strings.Split(n, "|")
			hostName = nameSplit[0]
			descriptionService = nameSplit[1]
		} else {
			colorRed := colorMessage.GetColorRed()
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println("The flag name must be of the form: hostName|serviceDescription ")
			os.Exit(1)
		}

		err, service := getServiceInfo(hostName, descriptionService, debugV)
		if err != nil {
			return err
		}
		if service.Description == "" {
			continue
		}

		//Write service informations
		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,service,\"" + service.HostName + "\",\"" + service.Description + "\",\"" + service.Template + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",check_command,\"" + service.CheckCommand + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",check_command_arguments,\"" + service.CheckCommandArguments + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",check_period,\"" + service.CheckPeriod + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",max_check_attempts,\"" + service.MaxCheckAttempts + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",normal_check_interval,\"" + service.NormalCheckInterval + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",retry_check_interval,\"" + service.RetryCheckInterval + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",active_checks_enabled,\"" + service.ActiveChecksEnabled + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",passive_checks_enabled,\"" + service.PassiveChecksEnabled + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",is_volatile,\"" + service.IsVolatile + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",notifications_enabled,\"" + service.NotificationsEnabled + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",contact_additive_inheritance,\"" + service.ContactAdditiveInheritance + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",cg_additive_inheritance,\"" + service.CgAdditiveInheritance + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",notification_options,\"" + service.NotificationOptions + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",notification_interval,\"" + service.NotificationInterval + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",notification_period,\"" + service.NotificationPeriod + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",first_notification_delay,\"" + service.FirstNotificationDelay + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",obsess_over_service,\"" + service.ObsessOverService + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",check_freshness,\"" + service.CheckFreshness + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",freshness_threshold,\"" + service.FreshnessThreshold + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",flap_detection_enabled,\"" + service.FlapDetectionEnabled + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",retain_status_information,\"" + service.RetainStatusInformation + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",retain_nonstatus_information,\"" + service.RetainNonstatusInformation + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",event_handler_enabled,\"" + service.EventHandlerEnabled + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",event_handler,\"" + service.EventHandler + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",event_handler_arguments,\"" + service.EventHandlerArguments + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",action_url,\"" + service.ActionURL + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",notes,\"" + service.Notes + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",notes_url,\"" + service.NotesURL + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",icon_image,\"" + service.IconImage + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",icon_image_alt,\"" + service.IconImageAlt + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",activate,\"" + service.Activate + "\"\n")
		_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",comment,\"" + service.Comment + "\"\n")

		//Write macros information
		if len(service.Macros) != 0 {
			for _, m := range service.Macros {
				if strings.Contains(m.Value, "\"") {
					m.Value = strings.ReplaceAll(m.Value, "\"", "'")
				}
				_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",macro,\"" + m.Name + "|" + m.Value + "|" + m.IsPassword + "|" + m.Description + "\"\n")
			}
		}

		//Write Hosts information
		if len(service.Hosts) != 0 {
			for _, h := range service.Hosts {
				_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",host,\"" + h.Name + "\"\n")
			}
		}

		//Write ContactGroups information
		if len(service.ContactGroups) != 0 {
			for _, c := range service.ContactGroups {
				_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",contactgroup,\"" + c.Name + "\"\n")
			}
		}

		//Write Contacts information
		if len(service.Contacts) != 0 {
			for _, c := range service.Contacts {
				_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",contact,\"" + c.Name + "\"\n")
			}
		}

		//Write ServiceGroups information
		if len(service.ServiceGroups) != 0 {
			for _, s := range service.ServiceGroups {
				_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",servicegroup,\"" + s.Name + "\"\n")
			}
		}

		//Write Traps information
		if len(service.Traps) != 0 {
			for _, t := range service.Traps {
				_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",trap,\"" + t.Name + "\"\n")
			}
		}

		//Write Categories information
		if len(service.Categories) != 0 {
			for _, c := range service.Categories {
				_, _ = f.WriteString("modify,service,\"" + service.HostName + "\",\"" + service.Description + "\",category,\"" + c.Name + "\"\n")
			}
		}
	}
	return nil
}

//The arguments impossible to get : recovery_notification_delay|acknowledgement_timeout|low_flap_threshold|high_flap_threshold|stalking_options|graphtemplate|geo_coords
//getServiceInfo permits to get all informations about a service
func getServiceInfo(hostName string, serviceDescription string, debugV bool) (error, service.ExportService) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the service
	values := hostName + ";" + serviceDescription + ";description|template|check_command|check_command_arguments|check_period|max_check_attempts|normal_check_interval|retry_check_interval|" +
		"active_checks_enabled|passive_checks_enabled|is_volatile|notifications_enabled|contact_additive_inheritance|cg_additive_inheritance|" +
		"notification_options|notification_interval|notification_period|first_notification_delay|obsess_over_service|check_freshness|freshness_threshold|" +
		"flap_detection_enabled|retain_status_information|retain_nonstatus_information|event_handler_enabled|event_handler|event_handler_arguments|notes|" +
		"icon_image|icon_image_alt|notes_url|action_url|activate|comment"
	err, body := request.GeneriqueCommandV1Post("getparam", "service", values, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultService service.ExportServiceResult
	json.Unmarshal(body, &resultService)

	//Check if the service  is found
	if len(resultService.Services) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + hostName + "|" + serviceDescription)
		return nil, service.ExportService{}
	}

	//Get the host of the service
	err, body = request.GeneriqueCommandV1Post("gethost", "service", hostName+";"+serviceDescription, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultHost service.ExportResultServiceHost
	json.Unmarshal(body, &resultHost)

	//Get the macro of the service
	err, body = request.GeneriqueCommandV1Post("getmacro", "service", hostName+";"+serviceDescription, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultMacro service.ExportResultServiceMacro
	json.Unmarshal(body, &resultMacro)

	//Get the contact group of the service
	err, body = request.GeneriqueCommandV1Post("getcontactgroup", "service", hostName+";"+serviceDescription, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultContactGroup service.ExportResultServiceContactGroup
	json.Unmarshal(body, &resultContactGroup)

	//Get the contact of the service
	err, body = request.GeneriqueCommandV1Post("getcontact", "service", hostName+";"+serviceDescription, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultContact service.ExportResultServiceContact
	json.Unmarshal(body, &resultContact)

	//Get the servicegroup of the service
	err, body = request.GeneriqueCommandV1Post("getservicegroup", "service", hostName+";"+serviceDescription, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultServiceGroup service.ExportResultServiceServiceGroup
	json.Unmarshal(body, &resultServiceGroup)

	//Get the trap of the service
	err, body = request.GeneriqueCommandV1Post("gettrap", "service", hostName+";"+serviceDescription, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultTrap service.ExportResultServiceTrap
	json.Unmarshal(body, &resultTrap)

	//Get the category of the service
	err, body = request.GeneriqueCommandV1Post("getcategory", "service", hostName+";"+serviceDescription, "export service", debugV, false, "")
	if err != nil {
		return err, service.ExportService{}
	}
	var resultCategory service.ExportResultServiceCategory
	json.Unmarshal(body, &resultCategory)

	//Get the service informations
	service := resultService.Services[0]
	service.HostName = hostName
	service.Hosts = resultHost.Hosts
	service.Macros = resultMacro.Macros
	service.ContactGroups = resultContactGroup.ContactGroups
	service.Contacts = resultContact.Contacts
	service.ServiceGroups = resultServiceGroup.ServiceGroups
	service.Traps = resultTrap.Traps
	service.Categories = resultCategory.Categories

	return nil, service

}

//getAllService permits to find all service in the centreon server
func getAllService(debugV bool) []service.ExportService {
	//Get all service
	err, body := request.GeneriqueCommandV1Post("show", "service", "", "export service", debugV, false, "")
	if err != nil {
		return []service.ExportService{}
	}
	var resultService service.ExportServiceResult
	json.Unmarshal(body, &resultService)

	return resultService.Services
}

func init() {
	serviceCmd.Flags().StringSliceP("name", "n", []string{}, "Host's name|Service's name (example: rtr-Paris|CPU)(separate by a comma the multiple values)")
	serviceCmd.Flags().StringSliceP("hostFilter", "o", []string{}, "To define the name of the hosts to which the services belong (separate by a comma the multiple values)")
	serviceCmd.Flags().StringP("file", "f", "ExportService.csv", "To define the name of the csv file")
}
