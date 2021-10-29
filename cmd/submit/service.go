/*MIT License

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
package submit

import (
	"centctl/colorMessage"
	"centctl/request"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// serviceCmd represents the submit/service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Submit a result to a single service",
	Long:  `Submit a result to a single service`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		serviceID, _ := cmd.Flags().GetInt("serviceID")
		output, _ := cmd.Flags().GetString("output")
		perfdata, _ := cmd.Flags().GetString("perfdata")
		status, _ := cmd.Flags().GetString("status")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := SubmitStatusService(id, serviceID, output, perfdata, status, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//SubmitStatusService permits to submit a result to a single service
func SubmitStatusService(id int, serviceID int, output string, perfdata string, status string, debugV bool) error {
	colorGreen := colorMessage.GetColorGreen()
	colorRed := colorMessage.GetColorRed()

	statusID := convertStatusService(status)
	if statusID == -1 {
		fmt.Printf(colorRed, "ERROR:")
		fmt.Println("The status is incorrect. The value must be (ok or warning or critical or unknown)")
		os.Exit(1)
	}

	var requestBody []byte
	requestBody, _ = json.Marshal(map[string]interface{}{
		"output":           output,
		"status":           statusID,
		"performance_data": perfdata,
	})

	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/hosts/" + strconv.Itoa(id) + "/services/" + strconv.Itoa(serviceID) + "/submit"
	err, _ := request.GeneriqueCommandV2Post(urlCentreon, requestBody, "submit service", debugV)
	if err != nil {
		return err
	}

	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The service `%s` of the host `%s` has submit result\n", strconv.Itoa(serviceID), strconv.Itoa(id))
	return nil
}

func convertStatusService(status string) int {
	switch status {
	case "ok":
		return 0
	case "warning":
		return 1
	case "critical":
		return 2
	case "unknown":
		return 3
	default:
		return -1
	}
}

func init() {
	serviceCmd.Flags().IntP("id", "i", -1, "ID of the host")
	serviceCmd.MarkFlagRequired("id")
	serviceCmd.Flags().IntP("serviceID", "s", -1, "ID of the service")
	serviceCmd.MarkFlagRequired("serviceID")
	serviceCmd.Flags().String("status", "", "Service status that can be submitted (ok or warning or critical or unknown)")
	serviceCmd.MarkFlagRequired("status")
	serviceCmd.RegisterFlagCompletionFunc("status", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"ok", "warning", "critical", "unknown"}, cobra.ShellCompDirectiveDefault
	})
}
