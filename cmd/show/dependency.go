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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/dependency"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// dependencyCmd represents the dependency command
var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Show one dependency's details",
	Long:  `Show one dependency's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowDependency(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowDependency permits to display the details of one dependency
func ShowDependency(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "DEP", name, "show dependency", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the dependencies contain into the response body
	dependencies := dependency.DetailResult{}
	json.Unmarshal(body, &dependencies)

	//Permits to find the good dependency in the array
	var DependencyFind dependency.DetailDependency
	for _, v := range dependencies.Dependencies {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			DependencyFind = v
		}
	}

	var server dependency.DetailServer
	if DependencyFind.Name != "" {
		//Organization of data
		server = dependency.DetailServer{
			Server: dependency.DetailInformations{
				Name:       os.Getenv("SERVER"),
				Dependency: &DependencyFind,
			},
		}
	} else {
		server = dependency.DetailServer{
			Server: dependency.DetailInformations{
				Name:       os.Getenv("SERVER"),
				Dependency: nil,
			},
		}
	}

	//Display details of the dependency
	displayDependency, err := display.DetailDependency(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayDependency)
	return nil
}

func init() {
	dependencyCmd.Flags().StringP("name", "n", "", "To define the name of the dependency")
	dependencyCmd.MarkFlagRequired("name")
	dependencyCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetDependencyNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
