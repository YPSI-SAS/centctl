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

//Action represents the caracteristics of a host Action
type Action struct {
	ID          string `json:"id" yaml:"id"`                   //Action ID
	Name        string `json:"name" yaml:"name"`               //Action Name
	Description string `json:"description" yaml:"description"` //Action Description
	Activate    string `json:"activate" yaml:"activate"`       //Action Activate

}

//ResultAction represents a host Action array
type ResultAction struct {
	Actions []Action `json:"result" yaml:"result"`
}

//ActionServer represents a server with informations
type ActionServer struct {
	Server ActionInformations `json:"server" yaml:"server"`
}

//ActionInformations represents the informations of the server
type ActionInformations struct {
	Name    string   `json:"name" yaml:"name"`
	Actions []Action `json:"actions" yaml:"actions"`
}

//StringText permits to display the caracteristics of the ACL actions to text
func (s ActionServer) StringText() string {
	sort.SliceStable(s.Server.Actions, func(i, j int) bool {
		return strings.ToLower(s.Server.Actions[i].Name) < strings.ToLower(s.Server.Actions[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Description", "Activate"})
	for i := 0; i < len(s.Server.Actions); i++ {
		table = append(table, []string{s.Server.Actions[i].ID, s.Server.Actions[i].Name, s.Server.Actions[i].Description, s.Server.Actions[i].Activate})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the ACL actions to csv
func (s ActionServer) StringCSV() string {
	var values string = "Server,ID,Name,Description,Activate\n"
	for i := 0; i < len(s.Server.Actions); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + s.Server.Actions[i].ID + "\"" + "," + "\"" + s.Server.Actions[i].Name + "\"" + "," + "\"" + s.Server.Actions[i].Description + "\"" + "," + "\"" + s.Server.Actions[i].Activate + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the ACL actions to json
func (s ActionServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the ACL actions to yaml
func (s ActionServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
