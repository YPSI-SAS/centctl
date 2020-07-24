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
	"centctl/display"
	"centctl/request"
	"centctl/service"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// listServiceCmd represents the service command
var listServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "List the services",
	Long:  `List the services of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ListService(output, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListService permits to display the array of Service return by the API
func ListService(output string, debugV bool) error {
	output = strings.ToLower(output)

	//Creation of the request body
	requestBody, err := request.CreateBodyRequest("Show", "service", "")
	if err != nil {
		return err
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := request.NewClient(urlCentreon)
	statusCode, body, err := client.CentreonCLAPI(requestBody)
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("list service", string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}

	//Permits to recover the services contain into the response body
	services := service.Result{}
	json.Unmarshal(body, &services)

	//Sort services based on their ID
	sort.SliceStable(services.Services, func(i, j int) bool {
		return strings.ToLower(services.Services[i].ServiceID) < strings.ToLower(services.Services[i].ServiceID)
	})

	//Organization of data
	server := service.Server{
		Server: service.Informations{
			Name:     os.Getenv("SERVER"),
			Services: services.Services,
		},
	}

	//Display all pollers
	displayService, err := display.Service(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayService)
	return nil
}

func init() {
	listCmd.AddCommand(listServiceCmd)
}
