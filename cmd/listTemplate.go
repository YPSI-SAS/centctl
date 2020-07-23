/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"centctl/debug"
	"centctl/display"
	"centctl/host"
	"centctl/request"
	"centctl/service"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listTemplateCmd represents the template command
var listTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "List hosts's and services's templates",
	Long:  `ListList hosts's and services's templates of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		object, _ := cmd.Flags().GetString("object")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListTemplate(output, object, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListTemplate permits to display the array of object template return by the API
func ListTemplate(output string, object string, debugV bool) error {
	output = strings.ToLower(output)
	object = strings.ToLower(object)

	//Verification that the object exists and create object centreon based on the object entered by the user
	objectCentreon := ""
	if object == "service" {
		objectCentreon = "STPL"
	} else if object == "host" {
		objectCentreon = "HTPL"
	} else {
		fmt.Println("The objects availables are: service and host ")
		os.Exit(1)
	}

	//Creation of the request body
	requestBody, err := request.CreateBodyRequest("Show", objectCentreon, "")
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list template", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Treatment of the response body based on the object entered by the user
	if object == "service" {
		//Permits to recover the service templates contain into the response body
		templates := service.ResultTemplate{}
		json.Unmarshal(body, &templates)

		//Sort service templates based on their ID
		sort.SliceStable(templates.Templates, func(i, j int) bool {
			return strings.ToLower(templates.Templates[i].Description) < strings.ToLower(templates.Templates[j].Description)
		})

		server := service.TemplateServer{
			Server: service.TemplateInformations{
				Name:      os.Getenv("SERVER"),
				Templates: templates.Templates,
			},
		}

		//Display all service templates
		displayTemplateService, err := display.TemplateService(output, server)
		if err != nil {
			return err
		}
		fmt.Println(displayTemplateService)
	} else {
		//Permits to recover the host templates contain into the response body
		templates := host.ResultTemplate{}
		json.Unmarshal(body, &templates)

		//Sort host templates based on their ID
		sort.SliceStable(templates.Templates, func(i, j int) bool {
			return strings.ToLower(templates.Templates[i].Name) < strings.ToLower(templates.Templates[j].Name)
		})

		server := host.TemplateServer{
			Server: host.TemplateInformations{
				Name:      os.Getenv("SERVER"),
				Templates: templates.Templates,
			},
		}

		//Display all host templates
		displayTemplateHost, err := display.TemplateHost(output, server)
		if err != nil {
			return err
		}
		fmt.Println(displayTemplateHost)
	}
	return nil
}

func init() {
	listCmd.AddCommand(listTemplateCmd)
	listTemplateCmd.Flags().StringP("object", "o", "", "To list the object templates (host or service)")
	listTemplateCmd.MarkFlagRequired("object")
}
