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
	"sort"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//Resource represents the caracteristics of a host Resource
type Resource struct {
	ID       string `json:"id" yaml:"id"`             //Resource ID
	Name     string `json:"name" yaml:"name"`         //Resource Name
	Alias    string `json:"alias" yaml:"alias"`       //Resource Alias
	Comment  string `json:"comment" yaml:"comment"`   //Resource Comment
	Activate string `json:"activate" yaml:"activate"` //Resource Activate

}

//ResultResource represents a host Resource array
type ResultResource struct {
	Resources []Resource `json:"result" yaml:"result"`
}

//ResourceServer represents a server with informations
type ResourceServer struct {
	Server ResourceInformations `json:"server" yaml:"server"`
}

//ResourceInformations represents the informations of the server
type ResourceInformations struct {
	Name      string     `json:"name" yaml:"name"`
	Resources []Resource `json:"resources" yaml:"resources"`
}

//StringText permits to display the caracteristics of the ACL Resources to text
func (s ResourceServer) StringText() string {
	sort.SliceStable(s.Server.Resources, func(i, j int) bool {
		return strings.ToLower(s.Server.Resources[i].Name) < strings.ToLower(s.Server.Resources[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Alias", "Comment", "Activate"})
	for i := 0; i < len(s.Server.Resources); i++ {
		table = append(table, []string{s.Server.Resources[i].ID, s.Server.Resources[i].Name, s.Server.Resources[i].Alias, s.Server.Resources[i].Comment, s.Server.Resources[i].Activate})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the ACL resource to csv
func (s ResourceServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Activate\n"
	for i := 0; i < len(s.Server.Resources); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + s.Server.Resources[i].ID + "\"" + "," + "\"" + s.Server.Resources[i].Name + "\"" + "," + "\"" + s.Server.Resources[i].Alias + "\"" + "," + "\"" + s.Server.Resources[i].Comment + "\"" + "," + "\"" + s.Server.Resources[i].Activate + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the ACL resource to json
func (s ResourceServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the ACL resource to yaml
func (s ResourceServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
