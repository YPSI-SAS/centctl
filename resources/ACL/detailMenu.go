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
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailMenu represents the caracteristics of a host Menu
type DetailMenu struct {
	ID       string `json:"id" yaml:"id"`             //Menu ID
	Name     string `json:"name" yaml:"name"`         //Menu Name
	Alias    string `json:"alias" yaml:"alias"`       //Menu Alias
	Comment  string `json:"comment" yaml:"comment"`   //Menu Comment
	Activate string `json:"activate" yaml:"activate"` //Menu Activate

}

//DetailResultMenu represents a host Menu array
type DetailResultMenu struct {
	Menus []DetailMenu `json:"result" yaml:"result"`
}

//DetailMenuServer represents a server with informations
type DetailMenuServer struct {
	Server DetailMenuInformations `json:"server" yaml:"server"`
}

//DetailMenuInformations represents the informations of the server
type DetailMenuInformations struct {
	Name string      `json:"name" yaml:"name"`
	Menu *DetailMenu `json:"menu" yaml:"menu"`
}

//StringText permits to display the caracteristics of the ACL Menus to text
func (s DetailMenuServer) StringText() string {
	var values string = "ACL Menu list for server " + s.Server.Name + ": \n"

	menu := s.Server.Menu
	if menu != nil {
		values += (*menu).ID + "\t"
		values += (*menu).Name + "\t"
		values += (*menu).Alias + "\t"
		values += (*menu).Comment + "\t"
		values += (*menu).Activate + "\n"
	} else {
		values += "menu: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the ACL ResultMenu to csv
func (s DetailMenuServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Activate\n"
	values += s.Server.Name + ","
	menu := s.Server.Menu
	if menu != nil {
		values += "\"" + (*menu).ID + "\"" + "," + "\"" + (*menu).Name + "\"" + "," + "\"" + (*menu).Alias + "\"" + "," + "\"" + (*menu).Comment + "\"" + "," + "\"" + (*menu).Activate + "\"" + "\n"

	} else {
		values += ",,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the ACL ResultMenu to json
func (s DetailMenuServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the ACL ResultMenu to yaml
func (s DetailMenuServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
