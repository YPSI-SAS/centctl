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

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Acknowledge hosts",
	Long:  `Acknowledge the host described right after`,
	Run: func(cmd *cobra.Command, args []string) {
		comment, _ := cmd.Flags().GetString("comment")
		hostID, _ := cmd.Flags().GetInt("hostID")
		services, _ := cmd.Flags().GetBool("services")
		notify, _ := cmd.Flags().GetBool("notify")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AcknowledgeHost(comment, notify, services, hostID, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AcknowledgeHost permits to acknowledge a host in the centreon server
func AcknowledgeHost(comment string, notify bool, services bool, hostID int, debugV bool) error {
	colorGreen := colorMessage.GetColorGreen()
	var requestBody []byte
	requestBody, _ = json.Marshal(map[string]interface{}{
		"comment":               comment,
		"is_notify_contacts":    notify,
		"is_persistent_comment": true,
		"is_sticky":             true,
		"with_services":         services,
	})

	//Recovery of the response
	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/hosts/" + strconv.Itoa(hostID) + "/acknowledgements"
	err, _ := request.GeneriqueCommandV2Post(urlCentreon, requestBody, "acknowledge host", debugV)
	if err != nil {
		return err
	}

	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The host `%s` is acknowledged\n", strconv.Itoa(hostID))
	return nil
}

func init() {
	hostCmd.Flags().IntP("hostID", "i", -1, "To know the host which must be acknowledge")
	hostCmd.MarkFlagRequired("hostID")
	hostCmd.Flags().Bool("notify", false, "Indicates whether notification is sent to the contacts")
	hostCmd.Flags().Bool("services", false, "Indicates whether we should add the acknowledge on the host-related services")
}
