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
	"centctl/cmd/delete/acl"
	"centctl/cmd/delete/broker"
	"centctl/cmd/delete/category"
	"centctl/cmd/delete/group"
	"centctl/cmd/delete/template"

	"github.com/spf13/cobra"
)

// Cmd represents the delete command
var Cmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete objects",
	Long:  `Delete objects defined right after.`,
	/*Run: func(cmd *cobra.Command, args []string) error {},*/

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
	Cmd.AddCommand(resourceCFGCmd)
	Cmd.AddCommand(serviceCmd)
	Cmd.AddCommand(template.Cmd)
	Cmd.AddCommand(timePeriodCmd)
	Cmd.AddCommand(trapCmd)
	Cmd.AddCommand(vendorCmd)

}
