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

package add

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// dependencyCmd represents the dependency command
var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Add a dependency",
	Long:  `Add a dependency into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		typeD, _ := cmd.Flags().GetString("type")
		parentName, _ := cmd.Flags().GetString("parentName")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddDependencie(name, description, typeD, parentName, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddDependencie permits to add a dependency in the centreon server
func AddDependencie(name string, description string, typeD string, parentName string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + description + ";" + typeD + ";" + parentName
	err := request.Add("add", "DEP", values, "add dependency", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	dependencyCmd.Flags().StringP("name", "n", "", "To define the name of the dependency")
	dependencyCmd.MarkFlagRequired("name")
	dependencyCmd.Flags().StringP("description", "d", "", "To define the description of the dependency")
	dependencyCmd.MarkFlagRequired("description")
	dependencyCmd.Flags().StringP("type", "t", "", "To define the type of the dependency (host,hg,sg,service ou meta)")
	dependencyCmd.MarkFlagRequired("type")
	dependencyCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"host", "hg", "sg", "service", "meta"}, cobra.ShellCompDirectiveDefault
	})
	dependencyCmd.Flags().StringP("parentName", "p", "", "To define the name of the parent resource (if type is SERVICE the parentName must be in the form \"HOSTNAME,SERVICENAME\") ")
	dependencyCmd.MarkFlagRequired("parentName")
}
