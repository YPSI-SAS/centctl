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
	"sort"
	"strings"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//BrokerInput represents the caracteristics of a BrokerInput
type BrokerInput struct {
	ID   string `json:"id" yaml:"id"`     //BrokerInput ID
	Name string `json:"name" yaml:"name"` //BrokerInput name
}

//ResultInput represents a poller array
type ResultInput struct {
	BrokerInputs []BrokerInput `json:"result" yaml:"result"`
}

//ServerInput represents a server with informations
type ServerInput struct {
	Server InformationsInput `json:"server" yaml:"server"`
}

//InformationsInput represents the informations of the server
type InformationsInput struct {
	Name         string        `json:"name" yaml:"name"`
	BrokerInputs []BrokerInput `json:"broker_inputs" yaml:"broker_inputs"`
}

//StringText permits to display the caracteristics of the BrokerInputs to text
func (s ServerInput) StringText() string {
	sort.SliceStable(s.Server.BrokerInputs, func(i, j int) bool {
		return strings.ToLower(s.Server.BrokerInputs[i].Name) < strings.ToLower(s.Server.BrokerInputs[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name"})
	for i := 0; i < len(s.Server.BrokerInputs); i++ {
		table = append(table, []string{s.Server.BrokerInputs[i].ID, s.Server.BrokerInputs[i].Name})
	}
	values := resources.TableListWithHeader(table)
	return values

}

//StringCSV permits to display the caracteristics of the BrokerInputs to csv
func (s ServerInput) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.BrokerInputs)
	return string(b)
}

//StringJSON permits to display the caracteristics of the BrokerInputs to json
func (s ServerInput) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerInputs to yaml
func (s ServerInput) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
