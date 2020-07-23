/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"centctl/debug"
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addHostCmd represents the host command
var addHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Add a host",
	Long:  `Add a host into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		alias, _ := cmd.Flags().GetString("alias")
		addressIP, _ := cmd.Flags().GetString("addressIP")
		template, _ := cmd.Flags().GetString("template")
		poller, _ := cmd.Flags().GetString("poller")
		hostGroup, _ := cmd.Flags().GetString("hostGroupe")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := AddHost(name, alias, addressIP, template, poller, hostGroup, debugV, apply)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddHost permits to add a host in the centreon server
func AddHost(hostName string, hostAlias string, adresseIP string, template string, pollerName string, hostGroup string, debugV bool, apply bool) error {
	//Verification if the hostGroup value exist
	var values string
	if hostGroup == "" {
		values = hostName + ";" + hostAlias + ";" + adresseIP + ";" + template + ";" + pollerName + ";"
	} else {
		values = hostName + ";" + hostAlias + ";" + adresseIP + ";" + template + ";" + pollerName + ";" + hostGroup
	}

	//Creation of the request body
	requestBody, err := request.CreateBodyRequest("add", "host", values)
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("add host", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Verification with the response body that the host was created out
	if string(body) != "{\"result\":[]}" {
		fmt.Println("erreur: ", string(body))
		os.Exit(1)
	}

	fmt.Printf("The host %v is created\n", hostName)

	if apply {
		//Export the poller configuration
		client = request.NewClient(os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi")
		err = client.ExportConf(pollerName, debugV)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	addCmd.AddCommand(addHostCmd)
	addHostCmd.Flags().StringP("name", "n", "", "To define the name of the host")
	addHostCmd.MarkFlagRequired("name")
	addHostCmd.Flags().StringP("alias", "a", "", "To define the alias of the host")
	addHostCmd.MarkFlagRequired("alias")
	addHostCmd.Flags().StringP("addressIP", "i", "", "To define the IP address of the host")
	addHostCmd.MarkFlagRequired("addressIP")
	addHostCmd.Flags().StringP("template", "t", "", "To define the template of the host")
	addHostCmd.MarkFlagRequired("template")
	addHostCmd.Flags().StringP("poller", "p", "", "To define the poller of the host")
	addHostCmd.MarkFlagRequired("poller")
	addHostCmd.Flags().StringP("hostGroup", "g", "", "To define if the contact is in a host group")
	addHostCmd.Flags().Bool("apply", false, "Export configuration of the poller")
}
