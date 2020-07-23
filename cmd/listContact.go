/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"centctl/contact"
	"centctl/debug"
	"centctl/display"
	"centctl/request"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listContactCmd represents the contact command
var listContactCmd = &cobra.Command{
	Use:   "contact",
	Short: "List the contacts",
	Long:  `List the contacts of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListContact(output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListContact permits to display the array of contact return by the API
func ListContact(output string, debugV bool) error {
	output = strings.ToLower(output)

	//Creation of the request body
	requestBody, err := request.CreateBodyRequest("Show", "contact", "")
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list contact", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the contacts contain into the response body
	contacts := contact.Result{}
	json.Unmarshal(body, &contacts)

	//Sort contacts based on their ID
	sort.SliceStable(contacts.Contacts, func(i, j int) bool {
		return strings.ToLower(contacts.Contacts[i].ID) < strings.ToLower(contacts.Contacts[j].ID)
	})

	//Organization of data
	server := contact.Server{
		Server: contact.Informations{
			Name:     os.Getenv("SERVER"),
			Contacts: contacts.Contacts,
		},
	}

	//Display all contacts
	displayContact, err := display.Contact(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayContact)

	return nil
}

func init() {
	listCmd.AddCommand(listContactCmd)
}
