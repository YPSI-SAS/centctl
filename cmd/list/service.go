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

package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/service"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "List the services",
	Long:  `List the services of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListService(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListService permits to display the array of services return by the API
func ListService(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "service", "", "list service", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the services contain into the response body
	services := service.Result{}
	json.Unmarshal(body, &services)
	finalServices := services.Services
	if regex != "" {
		finalServices = deleteService(finalServices, regex)
	}
	//Sort services based on their ID
	sort.SliceStable(finalServices, func(i, j int) bool {
		valI, _ := strconv.Atoi(finalServices[i].ServiceID)
		valJ, _ := strconv.Atoi(finalServices[j].ServiceID)
		return valI < valJ
	})

	//Organization of data
	server := service.Server{
		Server: service.Informations{
			Name:     os.Getenv("SERVER"),
			Services: finalServices,
		},
	}

	//Display all services
	displayService, err := display.Service(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayService)
	return nil
}

func deleteService(services []service.Service, regex string) []service.Service {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range services {
		matched, err := regexp.MatchString(regex, s.Description)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			services[index] = s
			index++
		}
	}
	return services[:index]
}

func init() {
	serviceCmd.Flags().StringP("regex", "r", "", "The regex to apply on the service's name")

}
