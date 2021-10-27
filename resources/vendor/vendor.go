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

package vendor

import (
	"centctl/resources"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//Vendor represents the caracteristics of a Vendor
type Vendor struct {
	ID    string `json:"id" yaml:"id"`       //Vendor ID
	Name  string `json:"name" yaml:"name"`   //Vendor name
	Alias string `json:"alias" yaml:"alias"` //Vendor Alias
}

//Result represents a poller array
type Result struct {
	Vendors []Vendor `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name    string   `json:"name" yaml:"name"`
	Vendors []Vendor `json:"vendors" yaml:"vendors"`
}

//StringText permits to display the caracteristics of the Vendors to text
func (s Server) StringText() string {
	sort.SliceStable(s.Server.Vendors, func(i, j int) bool {
		return strings.ToLower(s.Server.Vendors[i].Name) < strings.ToLower(s.Server.Vendors[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Alias"})
	for i := 0; i < len(s.Server.Vendors); i++ {
		table = append(table, []string{s.Server.Vendors[i].ID, s.Server.Vendors[i].Name, s.Server.Vendors[i].Alias})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the Vendors to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Alias\n"
	for i := 0; i < len(s.Server.Vendors); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + s.Server.Vendors[i].ID + "\"" + "," + "\"" + s.Server.Vendors[i].Name + "\"" + "," + "\"" + s.Server.Vendors[i].Alias + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the Vendors to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the Vendors to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
