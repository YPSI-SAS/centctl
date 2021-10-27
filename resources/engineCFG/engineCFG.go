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

package engineCFG

import (
	"centctl/resources"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//EngineCFG represents the caracteristics of a EngineCFG
type EngineCFG struct {
	ID       string `json:"nagios id" yaml:"nagios id"`           //EngineCFG ID
	Name     string `json:"nagios name" yaml:"nagios name"`       //EngineCFG name
	Comment  string `json:"nagios comment" yaml:"nagios comment"` //EngineCFG comment
	Instance string `json:"instance" yaml:"instance"`             //EngineCFG instance
}

//ResultEngineCFG represents a poller array
type ResultEngineCFG struct {
	EngineCFG []EngineCFG `json:"result" yaml:"result"`
}

//ServerEngineCFG represents a server with informations
type ServerEngineCFG struct {
	Server InformationsEngineCFG `json:"server" yaml:"server"`
}

//InformationsEngineCFG represents the informations of the server
type InformationsEngineCFG struct {
	Name      string      `json:"name" yaml:"name"`
	EngineCFG []EngineCFG `json:"engine_cfgs" yaml:"engine_cfgs"`
}

//StringText permits to display the caracteristics of the EngineCFG to text
func (s ServerEngineCFG) StringText() string {
	sort.SliceStable(s.Server.EngineCFG, func(i, j int) bool {
		return strings.ToLower(s.Server.EngineCFG[i].Name) < strings.ToLower(s.Server.EngineCFG[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Instance", "Comment"})
	for i := 0; i < len(s.Server.EngineCFG); i++ {
		table = append(table, []string{s.Server.EngineCFG[i].ID, s.Server.EngineCFG[i].Name, s.Server.EngineCFG[i].Instance, s.Server.EngineCFG[i].Comment})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the EngineCFG to csv
func (s ServerEngineCFG) StringCSV() string {
	var values string = "Server,ID,Name,Instance,Comment\n"
	for i := 0; i < len(s.Server.EngineCFG); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + s.Server.EngineCFG[i].ID + "\"" + "," + "\"" + s.Server.EngineCFG[i].Name + "\"" + "," + "\"" + s.Server.EngineCFG[i].Instance + "\"" + "," + "\"" + s.Server.EngineCFG[i].Comment + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the EngineCFG to json
func (s ServerEngineCFG) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the EngineCFG to yaml
func (s ServerEngineCFG) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
