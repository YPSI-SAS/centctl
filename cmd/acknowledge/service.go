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

package acknowledge

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
	Short: "Acknowledge services",
	Long:  `Acknowledge the service described right after`,
	Run: func(cmd *cobra.Command, args []string) {
		comment, _ := cmd.Flags().GetString("comment")
		hostID, _ := cmd.Flags().GetInt("hostID")
		serviceID, _ := cmd.Flags().GetInt("serviceID")
		notify, _ := cmd.Flags().GetBool("notify")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AcknowledgeService(comment, notify, hostID, serviceID, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AcknowledgeService permits to acknowledge a service in the centreon server
func AcknowledgeService(comment string, notify bool, hostID int, serviceID int, debugV bool) error {
	colorGreen := colorMessage.GetColorGreen()
	var requestBody []byte
	requestBody, _ = json.Marshal(map[string]interface{}{
		"comment":               comment,
		"is_notify_contacts":    notify,
		"is_persistent_comment": true,
		"is_sticky":             true,
	})

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/hosts/" + strconv.Itoa(hostID) + "/services/" + strconv.Itoa(serviceID) + "/acknowledgements"
	err, _ := request.GeneriqueCommandV2Post(urlCentreon, requestBody, "acknowledge service", debugV)
	if err != nil {
		return err
	}

	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The service `%s` of the host `%s` is acknowledged\n", strconv.Itoa(serviceID), strconv.Itoa(hostID))
	return nil
}

func init() {
	serviceCmd.Flags().IntP("hostID", "i", -1, "To know the host which must be acknowledge")
	serviceCmd.MarkFlagRequired("hostID")
	serviceCmd.Flags().IntP("serviceID", "s", -1, "To know the host which must be acknowledge")
	serviceCmd.MarkFlagRequired("serviceID")
	serviceCmd.Flags().Bool("notify", false, "Indicates whether notification is sent to the contacts")

}
