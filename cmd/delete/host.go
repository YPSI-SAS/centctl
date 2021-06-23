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

package delete

import (
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// hostCmd represents the deleteHost command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Delete a host ",
	Long:  `Delete a host into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := DeleteHost(name, debugV, apply)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DeleteHost permits to delete a host in the centreon server
func DeleteHost(name string, debugV bool, apply bool) error {
	poller := ""
	var err error
	//Find the name of the host poller
	client := request.NewClientV1(os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&search=" + name)
	for poller == "" {
		poller, err = client.NamePollerHost(name, debugV)
		if err != nil {
			return err
		}
	}

	err = request.Delete("del", "host", name, "delete host", name, debugV, apply, poller)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the host which will delete")
	hostCmd.MarkFlagRequired("name")
	hostCmd.Flags().Bool("apply", false, "Export configuration of the poller")

}
