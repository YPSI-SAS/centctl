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

package add

import (
	"centctl/request"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// pollerCmd represents the poller command
var pollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "Add a poller",
	Long:  `Add a poller into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		IPaddress, _ := cmd.Flags().GetString("IPaddress")
		SSHPort, _ := cmd.Flags().GetInt("SSHPort")
		connProtocol, _ := cmd.Flags().GetString("connectionProtocol")
		portConn, _ := cmd.Flags().GetInt("portConnection")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddPoller(name, IPaddress, SSHPort, connProtocol, portConn, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddPoller permits to add a poller in the centreon server
func AddPoller(name string, IPaddress string, SSHPort int, connProtocol string, portConn int, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + IPaddress + ";" + strconv.Itoa(SSHPort) + ";" + connProtocol + ";" + strconv.Itoa(portConn)
	err := request.Add("add", "instance", values, "add poller", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	pollerCmd.Flags().StringP("name", "n", "", "To define the name of the poller")
	pollerCmd.MarkFlagRequired("name")
	pollerCmd.Flags().StringP("IPaddress", "i", "", "To define the IP address of the poller")
	pollerCmd.MarkFlagRequired("IPaddress")
	pollerCmd.Flags().IntP("SSHPort", "s", 22, "To define the SSH port of the poller")
	pollerCmd.Flags().StringP("connectionProtocol", "c", "ZMQ", "To define the gorgone connection protocol (SSH or ZMQ)")
	pollerCmd.Flags().IntP("portConnection", "p", 5556, "To define the gorgone connection port")

}
