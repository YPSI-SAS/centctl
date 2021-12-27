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

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Submit a result to a single host",
	Long:  `Submit a result to a single host`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		output, _ := cmd.Flags().GetString("output")
		perfdata, _ := cmd.Flags().GetString("perfdata")
		status, _ := cmd.Flags().GetString("status")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := SubmitStatusHost(id, output, perfdata, status, debugV)
		if err != nil {
			fmt.Println(err)
		}

	},
}

//SubmitStatusHost permits to submit a result to a single host
func SubmitStatusHost(id int, output string, perfdata string, status string, debugV bool) error {
	colorGreen := colorMessage.GetColorGreen()
	colorRed := colorMessage.GetColorRed()

	statusID := convertStatusHost(status)
	if statusID == -1 {
		fmt.Printf(colorRed, "ERROR:")
		fmt.Println("The status is incorrect. The value must be (up or down or unreachable)")
		os.Exit(1)
	}

	var requestBody []byte
	requestBody, _ = json.Marshal(map[string]interface{}{
		"output":           output,
		"status":           statusID,
		"performance_data": perfdata,
	})

	urlCentreon := "/monitoring/hosts/" + strconv.Itoa(id) + "/submit"
	err, _ := request.GeneriqueCommandV2Post(urlCentreon, requestBody, "submit host", debugV)
	if err != nil {
		return err
	}

	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The host `%s` has submit result\n", strconv.Itoa(id))
	return nil
}

func convertStatusHost(status string) int {
	switch status {
	case "up":
		return 0
	case "down":
		return 1
	case "unreachable":
		return 2
	default:
		return -1
	}
}

func init() {
	hostCmd.Flags().IntP("id", "i", -1, "ID of the host")
	hostCmd.MarkFlagRequired("id")
	hostCmd.Flags().String("status", "", "Host status that can be submitted (up, down, unreachable)")
	hostCmd.MarkFlagRequired("status")
	hostCmd.RegisterFlagCompletionFunc("status", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"up", "down", "unreachable"}, cobra.ShellCompDirectiveDefault
	})
}
