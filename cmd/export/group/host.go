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
package group

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/host"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Export group host",
	Long:  `Export group host of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		appendFile, _ := cmd.Flags().GetBool("append")
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportGroupHost(name, regex, file, appendFile, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportGroupHost permits to export a group host of the centreon server
func ExportGroupHost(name []string, regex string, file string, appendFile bool, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	//Check if the name of file contains the extension
	if !strings.Contains(file, ".csv") {
		file = file + ".csv"
	}

	//Create the file
	var f *os.File
	var err error
	if appendFile {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		f, err = os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	}
	defer f.Close()
	if err != nil {
		return err
	}

	if all || regex != "" {
		groups := getAllGroupHost(debugV)
		for _, a := range groups {
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
		err, group := getGroupHostInfo(n, debugV)
		if err != nil {
			return err
		}
		if group.Name == "" {
			continue
		}

		_, _ = f.WriteString("\n")
		_, _ = f.WriteString("add,groupHost,\"" + group.Name + "\",\"" + group.Alias + "\"\n")
		_, _ = f.WriteString("modify,groupHost,\"" + group.Name + "\",notes,\"" + group.Notes + "\"\n")
		_, _ = f.WriteString("modify,groupHost,\"" + group.Name + "\",notes_url,\"" + group.NotesURL + "\"\n")
		_, _ = f.WriteString("modify,groupHost,\"" + group.Name + "\",action_url,\"" + group.ActionURL + "\"\n")
		_, _ = f.WriteString("modify,groupHost,\"" + group.Name + "\",activate,\"" + group.Activate + "\"\n")

		//Problem SQL Syntax when the images are imported after
		// _, _ = f.WriteString("modify,groupHost," + group.Name + ",icon_image," + group.IconImage + "\n")
		// _, _ = f.WriteString("modify,groupHost," + group.Name + ",map_icon_image," + group.MapIconImage + "\n")

		_, _ = f.WriteString("modify,groupHost,\"" + group.Name + "\",geo_coords,\"" + group.GeoCoords + "\"\n")
		_, _ = f.WriteString("modify,groupHost,\"" + group.Name + "\",comment,\"" + group.Comment + "\"\n")

		//Write in the file the members
		if len(group.Member) != 0 {
			for _, m := range group.Member {
				_, _ = f.WriteString("modify,groupHost,\"" + group.Name + "\",member,\"" + m.Name + "\"\n")
			}
		}
	}
	return nil
}

//The arguments impossible to get : rrd_retention
//getGroupHostInfo permits to get all informations about a host group
func getGroupHostInfo(name string, debugV bool) (error, host.ExportGroup) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the host group
	values := name + ";name|alias|comment|activate|notes|notes_url|action_url|icon_image|map_icon_image|geo_coords"
	err, body := request.GeneriqueCommandV1Post("getparam", "HG", values, "export group host", debugV, false, "")
	if err != nil {
		return err, host.ExportGroup{}
	}
	var resultGroup host.ExportResult
	json.Unmarshal(body, &resultGroup)

	//Check if the host group is found
	if len(resultGroup.GroupHosts) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, host.ExportGroup{}
	}

	//Get the members of the host group
	err, body = request.GeneriqueCommandV1Post("getmember", "HG", name, "export group host", debugV, false, "")
	if err != nil {
		return err, host.ExportGroup{}
	}
	var resultMember host.ExportResultMember
	json.Unmarshal(body, &resultMember)

	//Get the group and the member
	group := resultGroup.GroupHosts[0]
	group.Member = resultMember.GroupMember

	return nil, group

}

//getAllGroupHost permits to find all host group in the centreon server
func getAllGroupHost(debugV bool) []host.ExportGroup {
	//Get all host group
	err, body := request.GeneriqueCommandV1Post("show", "HG", "", "export group host", debugV, false, "")
	if err != nil {
		return []host.ExportGroup{}
	}
	var resultGroup host.ExportResult
	json.Unmarshal(body, &resultGroup)

	return resultGroup.GroupHosts
}

func init() {
	hostCmd.Flags().StringSliceP("name", "n", []string{}, "Hostgroup's name (separate by a comma the multiple values)")
	hostCmd.Flags().StringP("file", "f", "ExportHostGroup.csv", "To define the name of the csv file")
	hostCmd.Flags().StringP("regex", "r", "", "The regex to apply on the host group's name")

}
