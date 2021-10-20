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
package delete

import (
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// serviceCmd represents the deleteService command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Delete a service",
	Long:  `Delete a service into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		hostName, _ := cmd.Flags().GetString("hostName")
		description, _ := cmd.Flags().GetString("description")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := DeleteService(hostName, description, debugV, apply)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DeleteService permits to delete a host in the centreon server
func DeleteService(hostName string, description string, debugV bool, apply bool) error {
	poller := ""
	var err error
	//Find the name of the host poller
	client := request.NewClientV1(os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&search=" + hostName)
	for poller == "" {
		poller, err = client.NamePollerHost(hostName, debugV)
		if err != nil {
			return err
		}
	}

	//Creation of the request body
	values := hostName + ";" + description

	err = request.Delete("del", "service", values, "delete service", description+" attached to host "+hostName, debugV, apply, poller)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	serviceCmd.Flags().StringP("hostName", "n", "", "To define the host to wich the service is attached")
	serviceCmd.MarkFlagRequired("hostName")
	serviceCmd.Flags().StringP("description", "d", "", "The description of the service which will delete")
	serviceCmd.MarkFlagRequired("description")
	serviceCmd.Flags().Bool("apply", false, "Export configuration of the poller")

}
