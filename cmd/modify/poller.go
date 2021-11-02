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
package modify

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// pollerCmd represents the poller command
var pollerCmd = &cobra.Command{
	Use:   "poller",
	Short: "Modfiy a poller",
	Long:  `Modfiy a poller into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := ModifyPoller(name, parameter, value, debugV, apply, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyPoller permits to modify a BA in the centreon server
func ModifyPoller(name string, parameter string, value string, debugV bool, apply bool, isImport bool, detail bool) error {

	values := name + ";" + parameter + ";" + value
	err := request.Modify("setparam", "instance", values, "modify poller", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	pollerCmd.Flags().StringP("name", "n", "", "To define the name of the poller to be modified")
	pollerCmd.MarkFlagRequired("name")
	pollerCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetPollerNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	pollerCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation.")
	pollerCmd.MarkFlagRequired("parameter")
	pollerCmd.RegisterFlagCompletionFunc("parameter", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"name", "localhost", "ns_ip_address", "engine_start_command", "engine_stop_command", "engine_restart_command", "engine_reload_command", "nagios_bin", "nagiostats_bin", "ssh_port", "broker_reload_command", "centreonbroker_cfg_path", "centreonbroker_module_path"}, cobra.ShellCompDirectiveDefault
	})
	pollerCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified")
	pollerCmd.MarkFlagRequired("value")
	pollerCmd.Flags().Bool("apply", false, "Export configuration of the poller")

}
