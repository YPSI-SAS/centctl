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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/timePeriod"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// timePeriodCmd represents the timePeriod command
var timePeriodCmd = &cobra.Command{
	Use:   "timePeriod",
	Short: "Show one timePeriod's details",
	Long:  `Show one timePeriod's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowTimePeriod(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowTimePeriod permits to display the details of one time period
func ShowTimePeriod(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "TP", name, "show timePeriod", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the booleanrules contain into the response body
	timePeriods := timePeriod.DetailResult{}
	json.Unmarshal(body, &timePeriods)

	//Permits to find the good timePeriod in the array
	var TimePeriodFind timePeriod.DetailTimePeriod
	for _, v := range timePeriods.TimePeriods {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			TimePeriodFind = v
		}
	}

	var server timePeriod.DetailServer
	if TimePeriodFind.Name != "" {

		err, body := request.GeneriqueCommandV1Post("getexception", "TP", TimePeriodFind.Name, "getexception", debugV, false, "")
		if err != nil {
			return err
		}

		//Permits to recover the exception contain into the response body
		exceptions := timePeriod.DetailResultException{}
		json.Unmarshal(body, &exceptions)

		TimePeriodFind.Exceptions = exceptions.TimePeriodExceptions
		//Organization of data
		server = timePeriod.DetailServer{
			Server: timePeriod.DetailInformations{
				Name:        os.Getenv("SERVER"),
				TimePeriods: &TimePeriodFind,
			},
		}
	} else {
		server = timePeriod.DetailServer{
			Server: timePeriod.DetailInformations{
				Name:        os.Getenv("SERVER"),
				TimePeriods: nil,
			},
		}
	}

	//Display details of the timePeriod
	displayTimePeriod, err := display.DetailTimePeriod(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTimePeriod)
	return nil
}

func init() {
	timePeriodCmd.Flags().StringP("name", "n", "", "To define the name of the timePeriod")
	timePeriodCmd.MarkFlagRequired("name")
}
