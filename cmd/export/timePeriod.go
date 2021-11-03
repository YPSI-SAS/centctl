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
	"centctl/resources/timePeriod"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// timePeriodCmd represents the timePeriod command
var timePeriodCmd = &cobra.Command{
	Use:   "timePeriod",
	Short: "Export timePeriod",
	Long:  `Export timePeriod of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportTimePeriod(name, regex, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportTimePeriod permits to export a timePeriod of the centreon server
func ExportTimePeriod(name []string, regex string, file string, all bool, debugV bool) error {
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
		templates := getAllTimePeriod(debugV)
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
		err, timePeriod := getTimePeriodInfo(n, debugV)
		if err != nil {
			return err
		}
		if timePeriod.Name == "" {
			continue
		}

		//Write timePeriod informations
		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,timePeriod,\""+timePeriod.Name+"\",\""+strings.ReplaceAll(timePeriod.Alias, "\"", "\"\"")+"\"\n", file, writeFile)
		request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",sunday,\""+timePeriod.Sunday+"\"\n", file, writeFile)
		request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",monday,\""+timePeriod.Monday+"\"\n", file, writeFile)
		request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",tuesday,\""+timePeriod.Tuesday+"\"\n", file, writeFile)
		request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",thursday,\""+timePeriod.Thursday+"\"\n", file, writeFile)
		request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",wednesday,\""+timePeriod.Wednesday+"\"\n", file, writeFile)
		request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",friday,\""+timePeriod.Friday+"\"\n", file, writeFile)
		request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",saturday,\""+timePeriod.Saturday+"\"\n", file, writeFile)

		//Write Exceptions information
		if len(timePeriod.Exceptions) != 0 {
			for _, b := range timePeriod.Exceptions {
				request.WriteValues("modify,timePeriod,\""+timePeriod.Name+"\",exception,\""+b.Days+";"+b.Timerange+"\"\n", file, writeFile)
			}
		}
	}
	return nil
}

//getTimePeriodInfo permits to get all informations about a timePeriod
func getTimePeriodInfo(name string, debugV bool) (error, timePeriod.ExportTimePeriod) {
	colorRed := colorMessage.GetColorRed()

	err, body := request.GeneriqueCommandV1Post("show", "TP", name, "export timePeriod", debugV, false, "")
	if err != nil {
		return err, timePeriod.ExportTimePeriod{}
	}
	var resultTimePeriod timePeriod.ExportResultTimePeriod
	json.Unmarshal(body, &resultTimePeriod)

	timeperiod := timePeriod.ExportTimePeriod{}
	find := false
	for _, g := range resultTimePeriod.TimePeriods {
		if strings.ToLower(g.Name) == strings.ToLower(name) {
			timeperiod = g
			find = true
		}
	}
	//Check if the timePeriod is found
	if !find {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, timeperiod
	}

	//Get the BA of the timePeriod
	err, body = request.GeneriqueCommandV1Post("getexception", "TP", name, "export timePeriod", debugV, false, "")
	if err != nil {
		return err, timePeriod.ExportTimePeriod{}
	}
	var resultExceptions timePeriod.ExportResultTimePeriodExecption
	json.Unmarshal(body, &resultExceptions)

	timeperiod.Exceptions = resultExceptions.Exceptions

	return nil, timeperiod

}

//getAllTimePeriod permits to find all timePeriod in the centreon server
func getAllTimePeriod(debugV bool) []timePeriod.ExportTimePeriod {
	//Get all timePeriod
	err, body := request.GeneriqueCommandV1Post("show", "TP", "", "export timePeriod", debugV, false, "")
	if err != nil {
		return []timePeriod.ExportTimePeriod{}
	}
	var resultTimePeriod timePeriod.ExportResultTimePeriod
	json.Unmarshal(body, &resultTimePeriod)

	return resultTimePeriod.TimePeriods
}

func init() {
	timePeriodCmd.Flags().StringSliceP("name", "n", []string{}, "timePeriod's name (separate by a comma the multiple values)")
	timePeriodCmd.Flags().StringP("regex", "r", "", "The regex to apply on the timePeriod's name")

}
