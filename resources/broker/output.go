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

package broker

import (
	"encoding/json"
	"fmt"

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
	var values string = "BrokerOutput list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.BrokerOutputs); i++ {
		values += "ID: " + s.Server.BrokerOutputs[i].ID + "\t"
		values += "Name: " + s.Server.BrokerOutputs[i].Name + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the BrokerOutputs to csv
func (s ServerOutput) StringCSV() string {
	var values string = "Server,ID,Name\n"
	for i := 0; i < len(s.Server.BrokerOutputs); i++ {
		values += s.Server.Name + "," + s.Server.BrokerOutputs[i].ID + "," + s.Server.BrokerOutputs[i].Name + "\n"
	}
	return fmt.Sprintf(values)
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
