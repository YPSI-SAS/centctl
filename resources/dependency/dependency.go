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
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Dependency represents the caracteristics of a Dependency
type Dependency struct {
	ID                          string `json:"id" yaml:"id"`                           //Dependency ID
	Name                        string `json:"name" yaml:"name"`                       //Dependency Name
	Description                 string `json:"description" yaml:"description"`         //Dependency description
	InheritsParent              string `json:"inherits_parent" yaml:"inherits_parent"` //Dependency inherits_parent
	ExecutionFailureCriteria    string `json:"execution_failure_criteria"`
	NotificationFailureCriteria string `json:"notification_failure_criteria"`
}

//Result represents a Dependency array send by the API
type Result struct {
	Dependencies []Dependency `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name         string       `json:"name" yaml:"name"`
	Dependencies []Dependency `json:"dependencies" yaml:"dependencies"`
}

//StringText permits to display the caracteristics of the Dependencies to text
func (s Server) StringText() string {
	var values string = "Dependency list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Dependencies); i++ {
		values += "ID: " + s.Server.Dependencies[i].ID + "\t"
		values += "Name: " + s.Server.Dependencies[i].Name + "\t"
		values += "Description: " + s.Server.Dependencies[i].Description + "\t"
		values += "Inherits parent: " + s.Server.Dependencies[i].InheritsParent + "\t"
		values += "Execution Failure Criteria: " + s.Server.Dependencies[i].ExecutionFailureCriteria + "\t"
		values += "Notification Failure Criteria: " + s.Server.Dependencies[i].NotificationFailureCriteria + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the Dependencies to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Description,InheritsParent,ExecutionFailureCriteria,NotificationFailureCriteria\n"
	for i := 0; i < len(s.Server.Dependencies); i++ {
		values += s.Server.Name + "," + s.Server.Dependencies[i].ID + "," + s.Server.Dependencies[i].Name + "," + s.Server.Dependencies[i].Description + "," + s.Server.Dependencies[i].InheritsParent + "," + s.Server.Dependencies[i].ExecutionFailureCriteria + "," + s.Server.Dependencies[i].NotificationFailureCriteria + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the Dependencies to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the Dependencies to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
