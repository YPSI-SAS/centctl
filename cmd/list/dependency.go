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

package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/dependency"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// dependencyCmd represents the dependency command
var dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "List the dependencies",
	Long:  `List the dependencies of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		err := ListDependencies(output, regex, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListDependencies permits to display the array of dependency return by the API
func ListDependencies(output string, regex string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "DEP", "", "list dependency", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the dependencies contain into the response body
	dependencies := dependency.Result{}
	json.Unmarshal(body, &dependencies)
	finalDependencies := dependencies.Dependencies
	if regex != "" {
		finalDependencies = deleteDependency(finalDependencies, regex)
	}

	//Sort dependencies based on their ID
	sort.SliceStable(finalDependencies, func(i, j int) bool {
		valI, _ := strconv.Atoi(finalDependencies[i].ID)
		valJ, _ := strconv.Atoi(finalDependencies[j].ID)
		return valI < valJ
	})

	//Organization of data
	server := dependency.Server{
		Server: dependency.Informations{
			Name:         os.Getenv("SERVER"),
			Dependencies: finalDependencies,
		},
	}

	//Display all dependencies
	displayDependencies, err := display.Dependency(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayDependencies)

	return nil
}

func deleteDependency(dependencies []dependency.Dependency, regex string) []dependency.Dependency {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range dependencies {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			dependencies[index] = s
			index++
		}
	}
	return dependencies[:index]
}

func init() {
	dependencyCmd.Flags().StringP("regex", "r", "", "The regex to apply on the dependency's name")
}
