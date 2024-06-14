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
package downtime

import (
	"centctl/colorMessage"
	"centctl/request"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add downtime on a service",
	Long:  `Add downtime on a service described right after`,
	Run: func(cmd *cobra.Command, args []string) {
		idS, _ := cmd.Flags().GetInt("idS")
		idH, _ := cmd.Flags().GetInt("idH")
		fixed, _ := cmd.Flags().GetBool("fixed")
		comment, _ := cmd.Flags().GetString("comment")
		duration, _ := cmd.Flags().GetInt("duration")
		startDay, _ := cmd.Flags().GetString("startDay")
		startHour, _ := cmd.Flags().GetString("startHour")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := DowntimeService(idS, idH, startDay, startHour, fixed, duration, comment, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DowntimeService permits to add a downtine on a service in the centreon server
func DowntimeService(idS int, idH int, startDay string, startHour string, fixed bool, duration int, comment string, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	colorGreen := colorMessage.GetColorGreen()
	err := VerifyStartDayAndHour(startDay, startHour)
	if err != nil {
		fmt.Printf(colorRed, "ERROR:")
		fmt.Println(err.Error())
		os.Exit(1)
	}



	timezone := getTimezoneHost(idH, debugV)
	timeStart, timeEnd := GetEndDowntime(startDay, startHour, duration, timezone)
	var requestBody []byte
	requestBody, _ = json.Marshal(map[string]interface{}{
		"comment":    comment,
		"end_time":   timeEnd,
		"start_time": timeStart,
		"is_fixed":   fixed,
		"duration":   duration,
	})

	urlCentreon := "/monitoring/hosts/" + strconv.Itoa(idH) + "/services/" + strconv.Itoa(idS) + "/downtimes"
	err, _ = request.GeneriqueCommandV2Post(urlCentreon, requestBody, "downtime service", debugV)
	if err != nil {
		return err
	}

	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The service `%d` of the host `%d` has a downtime\n", idS, idH)
	return nil
}

func init() {
	serviceCmd.Flags().Int("idS", -1, "ID of the service")
	serviceCmd.MarkFlagRequired("idS")
	serviceCmd.Flags().Int("idH", -1, "ID of the host")
	serviceCmd.MarkFlagRequired("idH")
	serviceCmd.Flags().Bool("fixed", true, "Indicates whether the downtime is fixed or not")
}
