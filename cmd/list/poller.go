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
	"centctl/resources/poller"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// pollerCmd represents the poller command
var pollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "List the pollers",
	Long:  `List the pollers wof the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		limit, _ := cmd.Flags().GetInt("limit")
		err := ListPoller(output, limit, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListPoller permits to display the array of poller return by the API
func ListPoller(output string, limit int, regex string, debugV bool) error {
	output = strings.ToLower(output)

	//Recovery of the response body
	urlCentreon := "/configuration/monitoring-servers?limit=" + strconv.Itoa(limit)
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "list poller", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the array result into the body
	var pollerResult poller.ResultPoller
	json.Unmarshal(body, &pollerResult)
	finalPollers := pollerResult.Pollers
	if regex != "" {
		finalPollers = deletePoller(finalPollers, regex)
	}

	//Sort hosts based on their ID
	sort.SliceStable(finalPollers, func(i, j int) bool {
		return strings.ToLower(finalPollers[i].Name) < strings.ToLower(finalPollers[j].Name)
	})
	server := poller.Server{
		Server: poller.Informations{
			Name:    os.Getenv("SERVER"),
			Pollers: finalPollers,
		},
	}

	//Display all pollers
	displayPoller, err := display.Poller(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayPoller)
	return nil
}

func deletePoller(pollers []poller.Poller, regex string) []poller.Poller {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range pollers {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			pollers[index] = s
			index++
		}
	}
	return pollers[:index]
}

func init() {
	pollerCmd.Flags().IntP("limit", "l", 10, "The number of pollers you want to list")
	pollerCmd.Flags().StringP("regex", "r", "", "The regex to apply on the poller's name")
}
