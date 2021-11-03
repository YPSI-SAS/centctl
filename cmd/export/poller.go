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
package export

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/poller"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// pollerCmd represents the poller command
var pollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "Export poller",
	Long:  `Export poller of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportPoller(name, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportPoller permits to export a poller of the centreon server
func ExportPoller(name []string, regex string, file string, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex ")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all || regex != "" {
		templates := getAllPoller(debugV)
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
		err, poller := getPollerInfo(n, debugV)
		if err != nil {
			return err
		}
		if poller.Name == "" {
			continue
		}

		//Write poller informations
		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,poller,\""+poller.Name+"\",\""+poller.IPAddress+"\","+poller.SSHPort+","+poller.GorgonePorotocol+","+poller.GorgonePort+"\n", file, writeFile)
		request.WriteValues("modify,poller,\""+poller.Name+"\",localhost,\""+poller.Localhost+"\"\n", file, writeFile)
		request.WriteValues("modify,poller,\""+poller.Name+"\",ns_activate,\""+poller.Activate+"\"\n", file, writeFile)
		request.WriteValues("modify,poller,\""+poller.Name+"\",engine_restart_command,\""+poller.EngineRestartCmd+"\"\n", file, writeFile)
		request.WriteValues("modify,poller,\""+poller.Name+"\",engine_reload_command,\""+poller.EngineReloadCmd+"\"\n", file, writeFile)
		request.WriteValues("modify,poller,\""+poller.Name+"\",broker_reload_command,\""+poller.BorkerReloadCmd+"\"\n", file, writeFile)
		request.WriteValues("modify,poller,\""+poller.Name+"\",nagios_bin,\""+poller.Bin+"\"\n", file, writeFile)
		request.WriteValues("modify,poller,\""+poller.Name+"\",nagiostats_bin,\""+poller.StatsBin+"\"\n", file, writeFile)

	}

	return nil
}

//The arguments impossible to get : element in setparam table
//getPollerInfo permits to get all informations about a poller
func getPollerInfo(name string, debugV bool) (error, poller.ExportPoller) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "instance", name, "export poller", debugV, false, "")
	if err != nil {
		return err, poller.ExportPoller{}
	}
	var resultPoller poller.ExportResultPoller
	json.Unmarshal(body, &resultPoller)

	poller := poller.ExportPoller{}
	find := false
	for _, g := range resultPoller.Pollers {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			poller = g
			find = true
		}
	}
	//Check if the poller  is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, poller
	}

	return nil, poller

}

//getAllPoller permits to find all Poller in the centreon server
func getAllPoller(debugV bool) []poller.ExportPoller {
	//Get all poller
	err, body := request.GeneriqueCommandV1Post("show", "instance", "", "export poller", debugV, false, "")
	if err != nil {
		return []poller.ExportPoller{}
	}
	var resultPoller poller.ExportResultPoller
	json.Unmarshal(body, &resultPoller)

	return resultPoller.Pollers
}

func init() {
	pollerCmd.Flags().StringSliceP("name", "n", []string{}, "poller's name (separate by a comma the multiple values)")
	pollerCmd.Flags().StringP("regex", "r", "", "The regex to apply on the poller's name")

}
