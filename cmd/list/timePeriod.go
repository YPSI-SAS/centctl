/*
MIT License

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
package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/timePeriod"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// timePeriodCmd represents the timePeriod command
var timePeriodCmd = &cobra.Command{
	Use:   "timePeriod",
	Short: "List the time periods",
	Long:  `List the time periods of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListTimePeriod(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListTimePeriod permits to display the array of time periods return by the API
func ListTimePeriod(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "TP", "", "list timePeriod", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the time periods contain into the response body
	timePeriods := timePeriod.Result{}
	json.Unmarshal(body, &timePeriods)
	finalTimePeriods := timePeriods.TimePeriods
	if regex != "" {
		finalTimePeriods = deleteTimePeriod(finalTimePeriods, regex)
	}

	//Sort time periods based on their ID
	sort.SliceStable(finalTimePeriods, func(i, j int) bool {
		valI, _ := strconv.Atoi(finalTimePeriods[i].ID)
		valJ, _ := strconv.Atoi(finalTimePeriods[j].ID)
		return valI < valJ
	})

	//Organization of data
	server := timePeriod.Server{
		Server: timePeriod.Informations{
			Name:        os.Getenv("SERVER"),
			TimePeriods: finalTimePeriods,
		},
	}

	//Display all time periods
	displayTimePeriod, err := display.TimePeriod(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTimePeriod)

	return nil
}

func deleteTimePeriod(timePeriods []timePeriod.TimePeriod, regex string) []timePeriod.TimePeriod {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range timePeriods {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			timePeriods[index] = s
			index++
		}
	}
	return timePeriods[:index]
}

func init() {
	timePeriodCmd.Flags().StringP("regex", "r", "", "The regex to apply on the timePeriod's name")
}
