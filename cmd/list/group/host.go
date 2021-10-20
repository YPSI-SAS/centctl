/*MIT License

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

package group

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
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "List group host",
	Long:  `List group host of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListGroupHost(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListGroupHost permits to display the array of host group return by the API
func ListGroupHost(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "HG", "", "list group host", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the host groups contain into the response body
	groups := host.ResultGroup{}
	json.Unmarshal(body, &groups)
	finalGroups := groups.Groups
	if regex != "" {
		finalGroups = deleteHostGroup(finalGroups, regex)
	}

	//Sort host groups based on their ID
	sort.SliceStable(finalGroups, func(i, j int) bool {
		return strings.ToLower(finalGroups[i].Name) < strings.ToLower(finalGroups[j].Name)
	})

	server := host.GroupServer{
		Server: host.GroupInformations{
			Name:   os.Getenv("SERVER"),
			Groups: finalGroups,
		},
	}

	//Display all host groups
	displayGroupHost, err := display.GroupHost(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayGroupHost)

	return nil
}

func deleteHostGroup(hostGroup []host.Group, regex string) []host.Group {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range hostGroup {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			hostGroup[index] = s
			index++
		}
	}
	return hostGroup[:index]
}

func init() {
	hostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host group's name")

}
