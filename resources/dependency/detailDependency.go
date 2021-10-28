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

package dependency

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailDependency represents the caracteristics of a Dependency
type DetailDependency struct {
	ID                          string `json:"id" yaml:"id"`                           //Dependency ID
	Name                        string `json:"name" yaml:"name"`                       //Dependency Name
	Description                 string `json:"description" yaml:"description"`         //Dependency description
	InheritsParent              string `json:"inherits_parent" yaml:"inherits_parent"` //Dependency inherits_parent
	ExecutionFailureCriteria    string `json:"execution_failure_criteria"`
	NotificationFailureCriteria string `json:"notification_failure_criteria"`
}

//DetailResult represents a Dependency array send by the API
type DetailResult struct {
	Dependencies []DetailDependency `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name       string            `json:"name" yaml:"name"`
	Dependency *DetailDependency `json:"dependency" yaml:"dependency"`
}

//StringText permits to display the caracteristics of the Dependencies to text
func (s DetailServer) StringText() string {
	var values string
	dependency := s.Server.Dependency
	if dependency != nil {
		elements := [][]string{{"0", "Dependency:"}, {"1", "ID: " + (*dependency).ID}, {"1", "Name: " + (*dependency).Name}, {"1", "Description: " + (*dependency).Description}, {"1", "Inherits parent: " + (*dependency).InheritsParent}, {"1", "Execution Failure Criteria: " + (*dependency).ExecutionFailureCriteria}, {"1", "Notification Failure Criteria: " + (*dependency).NotificationFailureCriteria}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "dependency: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the Dependencies to csv
func (s DetailServer) StringCSV() string {
	var p []DetailDependency
	if s.Server.Dependency != nil {
		p = append(p, *s.Server.Dependency)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the Dependencies to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the Dependencies to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
