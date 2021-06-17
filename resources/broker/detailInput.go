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
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailBrokerInput represents the caracteristics of a BrokerInput
type DetailBrokerInput struct {
	ParamKey   string `json:"parameter key" yaml:"parameter key"`     //BrokerInput param key
	ParamValue string `json:"parameter value" yaml:"parameter value"` //BrokerInput param value
}

//DetailResultInput represents a poller array
type DetailResultInput struct {
	BrokerInputs []DetailBrokerInput `json:"result" yaml:"result"`
}

//DetailServerInput represents a server with informations
type DetailServerInput struct {
	Server DetailInformationsInput `json:"server" yaml:"server"`
}

//DetailInformationsInput represents the informations of the server
type DetailInformationsInput struct {
	Name        string              `json:"name" yaml:"name"`
	BrokerInput []DetailBrokerInput `json:"broker_input" yaml:"broker_input"`
}

//StringText permits to display the caracteristics of the BrokerInputs to text
func (s DetailServerInput) StringText() string {
	var values string = "BrokerInput list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.BrokerInput); i++ {
		values += "ID: " + s.Server.BrokerInput[i].ParamKey + "\t"
		values += "Name: " + s.Server.BrokerInput[i].ParamValue + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the BrokerInputs to csv
func (s DetailServerInput) StringCSV() string {
	var values string = "Server,ID,Name\n"
	values += s.Server.Name + ","
	for i := 0; i < len(s.Server.BrokerInput); i++ {
		values += s.Server.BrokerInput[i].ParamKey + "," + s.Server.BrokerInput[i].ParamValue + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the BrokerInputs to json
func (s DetailServerInput) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerInputs to yaml
func (s DetailServerInput) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
