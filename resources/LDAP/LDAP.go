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

package LDAP

import (
	"centctl/resources"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//LDAP represents the caracteristics of a LDAP
type LDAP struct {
	ID          string `json:"id" yaml:"id"`                   //LDAP ID
	Name        string `json:"name" yaml:"name"`               //LDAP name
	Description string `json:"description" yaml:"description"` //LDAP Description
	Status      string `json:"status" yaml:"status"`           //LDAP Status
}

//Result represents a poller array
type Result struct {
	LDAP []LDAP `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name string `json:"name" yaml:"name"`
	LDAP []LDAP `json:"ldaps" yaml:"ldaps"`
}

//StringText permits to display the caracteristics of the LDAP to text
func (s Server) StringText() string {
	sort.SliceStable(s.Server.LDAP, func(i, j int) bool {
		return strings.ToLower(s.Server.LDAP[i].Name) < strings.ToLower(s.Server.LDAP[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Status", "Description"})
	for i := 0; i < len(s.Server.LDAP); i++ {
		table = append(table, []string{s.Server.LDAP[i].ID, s.Server.LDAP[i].Name, s.Server.LDAP[i].Status, s.Server.LDAP[i].Description})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the LDAP to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Status,Description\n"
	for i := 0; i < len(s.Server.LDAP); i++ {
		values += s.Server.Name + "," + s.Server.LDAP[i].ID + "," + s.Server.LDAP[i].Name + "," + s.Server.LDAP[i].Status + "," + s.Server.LDAP[i].Description + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the LDAP to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the LDAP to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
