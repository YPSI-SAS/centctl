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

	"gopkg.in/yaml.v2"
)

//DetailParametersOutput represents the caracteristics of a BrokerOutput
type DetailParametersOutput struct {
	ParamKey   string `json:"parameter key" yaml:"parameter key"`     //BrokerOutput param key
	ParamValue string `json:"parameter value" yaml:"parameter value"` //BrokerOutput param value
}

type DetailBrokerOutput struct {
	ID         string                   `json:"ID" yaml:"ID"`
	BrokerName string                   `json:"broker_name" yaml:"broker_name"`
	Parameters []DetailParametersOutput `json:"parameters" yaml:"parameters"`
}

//DetailResultOutput represents a poller array
type DetailResultOutput struct {
	BrokerOutputs []DetailParametersOutput `json:"result" yaml:"result"`
}

//DetailServerOutput represents a server with informations
type DetailServerOutput struct {
	Server DetailInformationsOutput `json:"server" yaml:"server"`
}

//DetailInformationsOutput represents the informations of the server
type DetailInformationsOutput struct {
	Name         string             `json:"name" yaml:"name"`
	BrokerOutput DetailBrokerOutput `json:"broker_output" yaml:"broker_output"`
}

//StringText permits to display the caracteristics of the BrokerOutputs to text
func (s DetailServerOutput) StringText() string {
	var values string
	for i := 0; i < len(s.Server.BrokerOutput.Parameters); i++ {
		values += "ID: " + s.Server.BrokerOutput.Parameters[i].ParamKey + "\t"
		values += "Name: " + s.Server.BrokerOutput.Parameters[i].ParamValue + "\n"
	}
	elements := [][]string{{"0", "Broker output:"}, {"1", "ID: " + s.Server.BrokerOutput.ID}, {"1", "BrokerName: " + s.Server.BrokerOutput.BrokerName}}
	elements = append(elements, []string{"1", "Parameters:"})
	sort.SliceStable(s.Server.BrokerOutput.Parameters, func(i, j int) bool {
		return strings.ToLower(s.Server.BrokerOutput.Parameters[i].ParamKey) < strings.ToLower(s.Server.BrokerOutput.Parameters[j].ParamKey)
	})
	for _, params := range s.Server.BrokerOutput.Parameters {
		elements = append(elements, []string{"2", params.ParamKey + " \t(value=" + params.ParamValue + ")"})
	}
	items := resources.GenerateListItems(elements, "")
	values = resources.BulletList(items)

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the BrokerOutputs to csv
func (s DetailServerOutput) StringCSV() string {
	var values string = "Server,outputID,brokerName,Key,Value\n"
	for i := 0; i < len(s.Server.BrokerOutput.Parameters); i++ {
		values += s.Server.Name + "," + s.Server.BrokerOutput.ID + "," + s.Server.BrokerOutput.BrokerName + ","
		values += "\"" + s.Server.BrokerOutput.Parameters[i].ParamKey + "\"" + "," + "\"" + s.Server.BrokerOutput.Parameters[i].ParamValue + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the BrokerOutputs to json
func (s DetailServerOutput) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerOutputs to yaml
func (s DetailServerOutput) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
