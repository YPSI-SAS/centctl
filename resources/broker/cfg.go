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

package broker

import (
	"centctl/resources"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//BrokerCFG represents the caracteristics of a BrokerCFG
type BrokerCFG struct {
	ID       string `json:"config id" yaml:"config id"`     //BrokerCFG ID
	Name     string `json:"config name" yaml:"config name"` //BrokerCFG name
	Instance string `json:"instance" yaml:"instance"`       //BrokerCFG instance
}

//ResultCFG represents a poller array
type ResultCFG struct {
	BrokerCFGs []BrokerCFG `json:"result" yaml:"result"`
}

//ServerCFG represents a server with informations
type ServerCFG struct {
	Server InformationsCFG `json:"server" yaml:"server"`
}

//InformationsCFG represents the informations of the server
type InformationsCFG struct {
	Name       string      `json:"name" yaml:"name"`
	BrokerCFGs []BrokerCFG `json:"broker_cfgs" yaml:"broker_cfgs"`
}

//StringText permits to display the caracteristics of the BrokerCFGs to text
func (s ServerCFG) StringText() string {
	sort.SliceStable(s.Server.BrokerCFGs, func(i, j int) bool {
		return strings.ToLower(s.Server.BrokerCFGs[i].Name) < strings.ToLower(s.Server.BrokerCFGs[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Instance"})
	for i := 0; i < len(s.Server.BrokerCFGs); i++ {
		table = append(table, []string{s.Server.BrokerCFGs[i].ID, s.Server.BrokerCFGs[i].Name, s.Server.BrokerCFGs[i].Instance})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the BrokerCFGs to csv
func (s ServerCFG) StringCSV() string {
	var values string = "Server,ID,Name,Instance\n"
	for i := 0; i < len(s.Server.BrokerCFGs); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + s.Server.BrokerCFGs[i].ID + "\"" + "," + "\"" + s.Server.BrokerCFGs[i].Name + "\"" + "," + "\"" + s.Server.BrokerCFGs[i].Instance + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the BrokerCFGs to json
func (s ServerCFG) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerCFGs to yaml
func (s ServerCFG) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
