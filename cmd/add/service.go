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

package add

import (
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add a service",
	Long:  `Add a service into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		hostName, _ := cmd.Flags().GetString("hostName")
		description, _ := cmd.Flags().GetString("description")
		template, _ := cmd.Flags().GetString("template")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		err := AddService(hostName, description, template, debugV, apply, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddService permits to add a service in the centreon server
func AddService(hostName string, description string, template string, debugV bool, apply bool, isImport bool) error {
	poller := ""
	if apply {
		//Find the name of the host poller
		var err error
		client := request.NewClientV1(os.Getenv("URL") + "/api/index.php?object=centreon_realtime_hosts&action=list&search=" + hostName)
		for poller == "" {
			poller, err = client.NamePollerHost(hostName, debugV)
			if err != nil {
				return err
			}
		}
	}

	//Creation of the request body
	values := hostName + ";" + description + ";" + template

	err := request.Add("add", "SERVICE", values, "add service", description+" attached to host "+hostName, debugV, isImport, apply, poller, "")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	serviceCmd.Flags().StringP("hostName", "n", "", "To define the host to wich the service is attached")
	serviceCmd.MarkFlagRequired("hostName")
	serviceCmd.RegisterFlagCompletionFunc("hostName", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetHostNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	serviceCmd.Flags().StringP("description", "d", "", "The description of the service")
	serviceCmd.MarkFlagRequired("description")
	serviceCmd.Flags().StringP("template", "t", "", "To define the template to wich the service is attached")
	serviceCmd.MarkFlagRequired("template")
	serviceCmd.RegisterFlagCompletionFunc("template", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetTemplateServiceNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
	serviceCmd.Flags().Bool("apply", false, "Export configuration of the poller")
}
