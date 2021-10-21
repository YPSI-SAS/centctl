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
package submit

import (
	"github.com/spf13/cobra"
)

// Cmd represents the submit command
var Cmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a result to an object",
	Long:  `Submit a result to an object`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	Cmd.AddCommand(hostCmd)

	Cmd.PersistentFlags().StringP("output", "o", "", "Output result of the check send")
	Cmd.MarkPersistentFlagRequired("output")
	Cmd.PersistentFlags().StringP("perfdata", "p", "", "Performance data result of the check send")
	Cmd.MarkPersistentFlagRequired("perfdata")

	Cmd.PersistentFlags().String("status", "", "Host status that can be submitted (up, down, unreachable)")
	Cmd.MarkPersistentFlagRequired("status")
}
