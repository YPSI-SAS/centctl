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

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailParameter represents the caracteristics of a BrokerInput
type DetailParameter struct {
	ParamKey   string `json:"parameter key" yaml:"parameter key"`     //BrokerInput param key
	ParamValue string `json:"parameter value" yaml:"parameter value"` //BrokerInput param value
}
type DetailBrokerInput struct {
	ID         string           `json:"ID" yaml:"ID"`
	BrokerName string           `json:"broker_name" yaml:"broker_name"`
	Parameters DetailParameters `json:"parameters" yaml:"parameters"`
}

type DetailParameters []DetailParameter

func (t DetailParameters) MarshalCSV() ([]byte, error) {
	var value string
	for i, param := range t {
		value += param.ParamKey + ":" + param.ParamValue
		if i < len(t)-1 {
			value += ","
		}
	}
	return []byte(value), nil
}

//DetailResultInput represents a poller array
type DetailResultInput struct {
	BrokerInputs []DetailParameter `json:"result" yaml:"result"`
}

//DetailServerInput represents a server with informations
type DetailServerInput struct {
	Server DetailInformationsInput `json:"server" yaml:"server"`
}

//DetailInformationsInput represents the informations of the server
type DetailInformationsInput struct {
	Name        string            `json:"name" yaml:"name"`
	BrokerInput DetailBrokerInput `json:"broker_input" yaml:"broker_input"`
}

//StringText permits to display the caracteristics of the BrokerInputs to text
func (s DetailServerInput) StringText() string {
	var values string
	for i := 0; i < len(s.Server.BrokerInput.Parameters); i++ {
		values += "ID: " + s.Server.BrokerInput.Parameters[i].ParamKey + "\t"
		values += "Name: " + s.Server.BrokerInput.Parameters[i].ParamValue + "\n"
	}
	elements := [][]string{{"0", "Broker input:"}, {"1", "ID: " + s.Server.BrokerInput.ID}, {"1", "BrokerName: " + s.Server.BrokerInput.BrokerName}}
	elements = append(elements, []string{"1", "Parameters:"})
	sort.SliceStable(s.Server.BrokerInput.Parameters, func(i, j int) bool {
		return strings.ToLower(s.Server.BrokerInput.Parameters[i].ParamKey) < strings.ToLower(s.Server.BrokerInput.Parameters[j].ParamKey)
	})
	for _, params := range s.Server.BrokerInput.Parameters {
		elements = append(elements, []string{"2", params.ParamKey + " \t(value=" + params.ParamValue + ")"})
	}
	items := resources.GenerateListItems(elements, "")
	values = resources.BulletList(items)

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the BrokerInputs to csv
func (s DetailServerInput) StringCSV() string {
	p := []DetailBrokerInput{s.Server.BrokerInput}
	b, _ := csvutil.Marshal(p)
	return string(b)
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
