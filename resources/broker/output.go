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

//BrokerOutput represents the caracteristics of a BrokerOutput
type BrokerOutput struct {
	ID   string `json:"id" yaml:"id"`     //BrokerOutput ID
	Name string `json:"name" yaml:"name"` //BrokerOutput name
}

//ResultOutput represents a poller array
type ResultOutput struct {
	BrokerOutputs []BrokerOutput `json:"result" yaml:"result"`
}

//ServerOutput represents a server with informations
type ServerOutput struct {
	Server InformationsOutput `json:"server" yaml:"server"`
}

//InformationsOutput represents the informations of the server
type InformationsOutput struct {
	Name          string         `json:"name" yaml:"name"`
	BrokerOutputs []BrokerOutput `json:"broker_outputs" yaml:"broker_outputs"`
}

//StringText permits to display the caracteristics of the BrokerOutputs to text
func (s ServerOutput) StringText() string {
	sort.SliceStable(s.Server.BrokerOutputs, func(i, j int) bool {
		return strings.ToLower(s.Server.BrokerOutputs[i].Name) < strings.ToLower(s.Server.BrokerOutputs[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name"})
	for i := 0; i < len(s.Server.BrokerOutputs); i++ {
		table = append(table, []string{s.Server.BrokerOutputs[i].ID, s.Server.BrokerOutputs[i].Name})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the BrokerOutputs to csv
func (s ServerOutput) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.BrokerOutputs)
	return string(b)
}

//StringJSON permits to display the caracteristics of the BrokerOutputs to json
func (s ServerOutput) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerOutputs to yaml
func (s ServerOutput) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
