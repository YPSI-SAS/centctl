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

//Menu represents the caracteristics of a host Menu
type Menu struct {
	ID       string `json:"id" yaml:"id"`             //Menu ID
	Name     string `json:"name" yaml:"name"`         //Menu Name
	Alias    string `json:"alias" yaml:"alias"`       //Menu Alias
	Comment  string `json:"comment" yaml:"comment"`   //Menu Comment
	Activate string `json:"activate" yaml:"activate"` //Menu Activate

}

//ResultMenu represents a host Menu array
type ResultMenu struct {
	Menus []Menu `json:"result" yaml:"result"`
}

//MenuServer represents a server with informations
type MenuServer struct {
	Server MenuInformations `json:"server" yaml:"server"`
}

//MenuInformations represents the informations of the server
type MenuInformations struct {
	Name  string `json:"name" yaml:"name"`
	Menus []Menu `json:"menus" yaml:"menus"`
}

//StringText permits to display the caracteristics of the ACL Menus to text
func (s MenuServer) StringText() string {
	var values string = "ACL Menu list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Menus); i++ {
		values += s.Server.Menus[i].ID + "\t"
		values += s.Server.Menus[i].Name + "\t"
		values += s.Server.Menus[i].Alias + "\t"
		values += s.Server.Menus[i].Comment + "\t"
		values += s.Server.Menus[i].Activate + "\n"

	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the ACL ResultMenu to csv
func (s MenuServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Activate\n"
	for i := 0; i < len(s.Server.Menus); i++ {
		values += s.Server.Name + "," + s.Server.Menus[i].ID + "," + s.Server.Menus[i].Name + "," + s.Server.Menus[i].Alias + "," + s.Server.Menus[i].Comment + "," + s.Server.Menus[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the ACL ResultMenu to json
func (s MenuServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the ACL ResultMenu to yaml
func (s MenuServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
