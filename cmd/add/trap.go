/*MIT License

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

// trapCmd represents the trap command
var trapCmd = &cobra.Command{
	Use:   "trap",
	Short: "Add a trap",
	Long:  `Add a trap into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		oid, _ := cmd.Flags().GetString("oid")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddTrap(name, oid, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddTrap permits to add a trap in the centreon server
func AddTrap(name string, oid string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + oid
	err := request.Add("add", "TRAP", values, "add trap", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	trapCmd.Flags().StringP("name", "n", "", "To define the name of the trap")
	trapCmd.MarkFlagRequired("name")
	trapCmd.Flags().StringP("oid", "o", "", "To define the oid of the SNMP trap")
	trapCmd.MarkFlagRequired("oid")
}
