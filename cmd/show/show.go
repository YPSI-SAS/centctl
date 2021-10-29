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

package show

import (
	"centctl/cmd/show/acl"
	"centctl/cmd/show/broker"
	"centctl/cmd/show/category"
	"centctl/cmd/show/group"
	"centctl/cmd/show/template"

	"github.com/spf13/cobra"
)

// Cmd represents the show command
var Cmd = &cobra.Command{
	Use:   "show",
	Short: "Show with details one host or service",
	Long:  `Show one object's details of the Centreon Server`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	Cmd.AddCommand(acl.Cmd)
	Cmd.AddCommand(broker.Cmd)
	Cmd.AddCommand(category.Cmd)
	Cmd.AddCommand(centreonProxyCmd)
	Cmd.AddCommand(commandCmd)
	Cmd.AddCommand(contactCmd)
	Cmd.AddCommand(dependencyCmd)
	Cmd.AddCommand(engineCFGCmd)
	Cmd.AddCommand(group.Cmd)
	Cmd.AddCommand(hostCmd)
	Cmd.AddCommand(ldapCmd)
	Cmd.AddCommand(pollerCmd)
	Cmd.AddCommand(realtimeHostCmd)
	Cmd.AddCommand(realtimeServiceCmd)
	Cmd.AddCommand(resourceCFGCmd)
	Cmd.AddCommand(serviceCmd)
	Cmd.AddCommand(template.Cmd)
	Cmd.AddCommand(timelineHostCmd)
	Cmd.AddCommand(timelineServiceCmd)
	Cmd.AddCommand(timePeriodCmd)
	Cmd.AddCommand(trapCmd)
	Cmd.AddCommand(vendorCmd)

	Cmd.PersistentFlags().String("output", "json", "Type of output (json, yaml, text, csv)")
	Cmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "text", "yaml", "csv"}, cobra.ShellCompDirectiveDefault
	})
}
