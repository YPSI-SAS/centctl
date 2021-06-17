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

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Add downtime on a host",
	Long:  `Add downtime on a host described right after`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("idH")
		withServices, _ := cmd.Flags().GetBool("withServices")
		fixed, _ := cmd.Flags().GetBool("fixed")
		comment, _ := cmd.Flags().GetString("comment")
		duration, _ := cmd.Flags().GetInt("duration")
		startDay, _ := cmd.Flags().GetString("startDay")
		startHour, _ := cmd.Flags().GetString("startHour")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := DowntimeHost(id, startDay, startHour, fixed, duration, comment, withServices, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DowntimeHost permits to add a downtine on a host in the centreon server
func DowntimeHost(id int, startDay string, startHour string, fixed bool, duration int, comment string, withServices bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	colorGreen := colorMessage.GetColorGreen()
	err := VerifyStartDayAndHour(startDay, startHour)
	if err != nil {
		fmt.Printf(colorRed, "ERROR:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	idAuthor, err := GetAuthorId(debugV)
	if err != nil {
		fmt.Printf(colorRed, "ERROR:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	timezone := getTimezoneHost(id, debugV)
	timeStart, timeEnd := GetEndDowntime(startDay, startHour, duration, timezone)
	var requestBody []byte
	requestBody, _ = json.Marshal(map[string]interface{}{
		"comment":       comment,
		"end_time":      timeEnd,
		"start_time":    timeStart,
		"is_fixed":      fixed,
		"duration":      duration,
		"author_id":     idAuthor,
		"with_services": withServices,
	})

	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/hosts/" + strconv.Itoa(id) + "/downtimes"
	err, _ = request.GeneriqueCommandV2Post(urlCentreon, requestBody, "downtime host", debugV)
	if err != nil {
		return err
	}

	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The host `%s` has a downtime\n", strconv.Itoa(id))
	return nil
}

func init() {
	hostCmd.Flags().Int("idH", -1, "ID of the host")
	hostCmd.MarkFlagRequired("idH")
	hostCmd.Flags().Bool("withServices", true, "Indicates whether we should add the downtime on the host-related services or not")
	hostCmd.Flags().Bool("fixed", true, "Indicates whether the downtime is fixed or not")
}
