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

package ACL

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailResource represents the caracteristics of a host Resource
type DetailResource struct {
	ID       string `json:"id" yaml:"id"`             //Resource ID
	Name     string `json:"name" yaml:"name"`         //Resource Name
	Alias    string `json:"alias" yaml:"alias"`       //Resource Alias
	Comment  string `json:"comment" yaml:"comment"`   //Resource Comment
	Activate string `json:"activate" yaml:"activate"` //Resource Activate

}

//DetailResultResource represents a host Resource array
type DetailResultResource struct {
	Resources []DetailResource `json:"result" yaml:"result"`
}

//DetailResourceServer represents a server with informations
type DetailResourceServer struct {
	Server DetailResourceInformations `json:"server" yaml:"server"`
}

//DetailResourceInformations represents the informations of the server
type DetailResourceInformations struct {
	Name     string          `json:"name" yaml:"name"`
	Resource *DetailResource `json:"resource" yaml:"resource"`
}

//StringText permits to display the caracteristics of the ACL Resources to text
func (s DetailResourceServer) StringText() string {
	var values string = "ACL Resource list for server " + s.Server.Name + ": \n"

	resource := s.Server.Resource
	if resource != nil {
		values += (*resource).ID + "\t"
		values += (*resource).Name + "\t"
		values += (*resource).Alias + "\t"
		values += (*resource).Comment + "\t"
		values += (*resource).Activate + "\n"
	} else {
		values += "resource: null \n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the ACL resource to csv
func (s DetailResourceServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Activate\n"
	values += s.Server.Name + ","
	resource := s.Server.Resource
	if resource != nil {
		values += (*resource).ID + "," + (*resource).Name + "," + (*resource).Alias + "," + (*resource).Comment + "," + (*resource).Activate + "\n"

	} else {
		values += ",,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the ACL resource to json
func (s DetailResourceServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the ACL resource to yaml
func (s DetailResourceServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
