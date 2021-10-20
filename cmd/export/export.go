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
package export

import (
	"centctl/cmd/export/acl"
	"centctl/cmd/export/category"
	"centctl/cmd/export/group"
	"centctl/cmd/export/template"

	"github.com/spf13/cobra"
)

// Cmd represents the export command
var Cmd = &cobra.Command{
	Use:   "export",
	Short: "Export in a csv file an object",
	Long:  `Export in a csv file an object defined right after.`,
	// Run: func(cmd *cobra.Command, args []string) {	},
}

func init() {
	Cmd.AddCommand(acl.Cmd)
	Cmd.AddCommand(category.Cmd)
	Cmd.AddCommand(contactCmd)
	Cmd.AddCommand(commandCmd)
	Cmd.AddCommand(engineCFGCmd)
	Cmd.AddCommand(group.Cmd)
	Cmd.AddCommand(hostCmd)
	Cmd.AddCommand(ldapCmd)
	Cmd.AddCommand(pollerCmd)
	Cmd.AddCommand(resourceCFGCmd)
	Cmd.AddCommand(serviceCmd)
	Cmd.AddCommand(template.Cmd)
	Cmd.AddCommand(timePeriodCmd)
	Cmd.AddCommand(trapCmd)
	Cmd.AddCommand(vendorCmd)

	Cmd.PersistentFlags().Bool("append", false, "Append the export in the csv file")
	Cmd.PersistentFlags().Bool("all", false, "Export all objects of this type in the csv file")

}
