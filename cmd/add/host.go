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

package add

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Add a host",
	Long:  `Add a host into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		alias, _ := cmd.Flags().GetString("alias")
		IPaddress, _ := cmd.Flags().GetString("IPaddress")
		template, _ := cmd.Flags().GetString("template")
		poller, _ := cmd.Flags().GetString("poller")
		hostGroup, _ := cmd.Flags().GetString("hostGroupe")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := AddHost(name, alias, IPaddress, template, poller, hostGroup, debugV, apply, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddHost permits to add a host in the centreon server
func AddHost(hostName string, hostAlias string, adresseIP string, template string, pollerName string, hostGroup string, debugV bool, apply bool, isImport bool) error {
	//Get values for POST
	var values string
	if hostGroup == "" {
		values = hostName + ";" + hostAlias + ";" + adresseIP + ";" + template + ";" + pollerName + ";"
	} else {
		values = hostName + ";" + hostAlias + ";" + adresseIP + ";" + template + ";" + pollerName + ";" + hostGroup
	}

	err := request.Add("add", "host", values, "add host", hostName, debugV, isImport, apply, pollerName, "")
	if err != nil {
		return err
	}

	values = hostName
	err = request.Add("APPLYTPL", "host", values, "add host", hostName, debugV, isImport, apply, pollerName, "The template of the host "+hostName+" is applied")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the name of the host")
	hostCmd.MarkFlagRequired("name")
	hostCmd.Flags().StringP("alias", "a", "", "To define the alias of the host")
	hostCmd.MarkFlagRequired("alias")
	hostCmd.Flags().StringP("IPaddress", "i", "", "To define the IP address of the host")
	hostCmd.MarkFlagRequired("IPaddress")
	hostCmd.Flags().StringP("template", "t", "", "To define the template of the host")
	hostCmd.MarkFlagRequired("template")
	hostCmd.Flags().StringP("poller", "p", "", "To define the poller of the host")
	hostCmd.MarkFlagRequired("poller")
	hostCmd.Flags().StringP("hostGroup", "g", "", "To define if the contact is in a host group")
	hostCmd.Flags().Bool("apply", false, "Export configuration of the poller")
}
