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

package list

import (
	"centctl/cmd/list/acl"
	"centctl/cmd/list/broker"
	"centctl/cmd/list/category"
	"centctl/cmd/list/group"
	"centctl/cmd/list/template"

	"github.com/spf13/cobra"
)

// Cmd represents the list command
var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List objects",
	Long:  `List objects of the Centreon Server`,
	/*Run: func(cmd *cobra.Command, args []string) error {},*/
}

func init() {
	Cmd.AddCommand(acl.Cmd)
	Cmd.AddCommand(broker.Cmd)
	Cmd.AddCommand(category.Cmd)
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
	Cmd.AddCommand(realtimePollerCmd)
	Cmd.AddCommand(resourceCFGCmd)
	Cmd.AddCommand(serviceCmd)
	Cmd.AddCommand(template.Cmd)
	Cmd.AddCommand(timePeriodCmd)
	Cmd.AddCommand(trapCmd)
	Cmd.AddCommand(vendorCmd)

	Cmd.PersistentFlags().String("output", "json", "Type of output (json, yaml, text, csv)")
	Cmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "text", "yaml", "csv"}, cobra.ShellCompDirectiveDefault
	})
}
