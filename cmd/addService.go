/*
MIT License

Copyright (c) 2020 YPSI SAS
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

package cmd

import (
	"centctl/debug"
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addServiceCmd represents the service command
var addServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add a service",
	Long:  `Add a service into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		hostName, _ := cmd.Flags().GetString("hostName")
		description, _ := cmd.Flags().GetString("description")
		template, _ := cmd.Flags().GetString("template")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := AddService(hostName, description, template, debugV, apply)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddService permits to add a service in the centreon server
func AddService(hostName string, description string, template string, debugV bool, apply bool) error {
	//Creation of the request body
	values := hostName + ";" + description + ";" + template
	requestBody, err := request.CreateBodyRequest("add", "SERVICE", values)
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("add service", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Verification with the response body that the service was created out
	if string(body) != "{\"result\":[]}" {
		fmt.Println("erreur: ", string(body))
		os.Exit(1)
	}

	fmt.Printf("The service %v is created\n", description)

	if apply {
		//Find the name of the host poller
		client = request.NewClient(os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&search=" + hostName)
		poller := ""
		for poller == "" {
			poller, err = client.NamePollerHost(hostName, debugV)
			if err != nil {
				return err
			}
		}

		//Export the poller configuration
		client = request.NewClient(os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi")
		err = client.ExportConf(poller, debugV)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	addCmd.AddCommand(addServiceCmd)
	addServiceCmd.Flags().StringP("hostName", "n", "", "To define the host to wich the service is attached")
	addServiceCmd.MarkFlagRequired("hostName")
	addServiceCmd.Flags().StringP("description", "d", "", "The description of the service")
	addServiceCmd.MarkFlagRequired("description")
	addServiceCmd.Flags().StringP("template", "t", "", "To define the template to wich the service is attached")
	addServiceCmd.MarkFlagRequired("template")
	addServiceCmd.Flags().Bool("apply", false, "Export configuration of the poller")
}
