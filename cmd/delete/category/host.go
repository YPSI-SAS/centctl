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
package category

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Delete category host",
	Long:  `Delete category host of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := DeleteCategoryHost(name, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DeleteCategoryHost permits to delete a host category in the centreon server
func DeleteCategoryHost(name string, debugV bool) error {

	err := request.Delete("del", "HC", name, "delete category host", name, debugV, false, "")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the host category which will delete")
	hostCmd.MarkFlagRequired("name")
}
