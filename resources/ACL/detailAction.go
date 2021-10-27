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

package ACL

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailAction represents the caracteristics of a host Action
type DetailAction struct {
	ID          string `json:"id" yaml:"id"`                   //Action ID
	Name        string `json:"name" yaml:"name"`               //Action Name
	Description string `json:"description" yaml:"description"` //Action Description
	Activate    string `json:"activate" yaml:"activate"`       //Action Activate

}

//DetailResultAction represents a host Action array
type DetailResultAction struct {
	Actions []DetailAction `json:"result" yaml:"result"`
}

//DetailActionServer represents a server with informations
type DetailActionServer struct {
	Server DetailActionInformations `json:"server" yaml:"server"`
}

//DetailActionInformations represents the informations of the server
type DetailActionInformations struct {
	Name   string        `json:"name" yaml:"name"`
	Action *DetailAction `json:"action" yaml:"action"`
}

//StringText permits to display the caracteristics of the ACL actions to text
func (s DetailActionServer) StringText() string {
	var values string
	action := s.Server.Action
	if action != nil {
		elements := [][]string{{"0", "ACL action:"}, {"1", "ID: " + (*action).ID}, {"1", "Name: " + (*action).Name}, {"1", "Description: " + (*action).Description}, {"1", "Activate: " + (*action).Activate}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "action: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the ACL actions to csv
func (s DetailActionServer) StringCSV() string {
	var values string = "Server,ID,Name,Description,Activate\n"
	values += s.Server.Name + ","
	action := s.Server.Action
	if action != nil {
		values += "\"" + (*action).ID + "\"" + "," + "\"" + (*action).Name + "\"" + "," + "\"" + (*action).Description + "\"" + "," + "\"" + (*action).Activate + "\"" + "\n"

	} else {
		values += ",,,\n"
	}

	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the ACL actions to json
func (s DetailActionServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the ACL actions to yaml
func (s DetailActionServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
