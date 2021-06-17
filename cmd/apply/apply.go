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

package apply

import (
	"centctl/request"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Cmd represents the apply command
var Cmd = &cobra.Command{
	Use:   "apply",
	Short: "apply configuration on a poller",
	Long:  `apply configuration on a poller`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := Apply(name, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//Apply the poller configuration
func Apply(name string, debugV bool) error {

	client := request.NewClientV1(os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi")
	err := client.ExportConf(name, debugV)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	Cmd.Flags().StringP("name", "n", "", "Name of the poller")
	Cmd.MarkFlagRequired("name")
}
