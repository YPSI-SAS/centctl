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
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "List the hosts",
	Long:  `List the hosts of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListHost(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListHost permits to display the array of host return by the API
func ListHost(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "host", "", "list host", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the hosts contain into the response body
	hosts := host.Result{}
	json.Unmarshal(body, &hosts)
	finalHosts := hosts.Hosts
	if regex != "" {
		finalHosts = deleteHost(finalHosts, regex)
	}

	//Sort hosts based on their ID
	sort.SliceStable(finalHosts, func(i, j int) bool {
		valI, _ := strconv.Atoi(finalHosts[i].ID)
		valJ, _ := strconv.Atoi(finalHosts[j].ID)
		return valI < valJ
	})

	//Organization of data
	server := host.Server{
		Server: host.Informations{
			Name:  os.Getenv("SERVER"),
			Hosts: finalHosts,
		},
	}

	//Display all hosts
	displayHost, err := display.Host(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayHost)
	return nil
}

func deleteHost(hosts []host.Host, regex string) []host.Host {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, h := range hosts {
		matched, err := regexp.MatchString(regex, h.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			hosts[index] = h
			index++
		}
	}
	return hosts[:index]
}

func init() {
	hostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host's name")

}
