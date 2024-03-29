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

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailBrokerCFG represents the caracteristics of a BrokerCFG
type DetailBrokerCFG struct {
	ID       string `json:"config id" yaml:"config id"`     //BrokerCFG ID
	Name     string `json:"config name" yaml:"config name"` //BrokerCFG name
	Instance string `json:"instance" yaml:"instance"`       //BrokerCFG instance
}

//DetailResultCFG represents a poller array
type DetailResultCFG struct {
	BrokerCFGs []DetailBrokerCFG `json:"result" yaml:"result"`
}

//DetailServerCFG represents a server with informations
type DetailServerCFG struct {
	Server DetailInformationsCFG `json:"server" yaml:"server"`
}

//DetailInformationsCFG represents the informations of the server
type DetailInformationsCFG struct {
	Name      string           `json:"name" yaml:"name"`
	BrokerCFG *DetailBrokerCFG `json:"broker_cfg" yaml:"broker_cfg"`
}

//StringText permits to display the caracteristics of the BrokerCFGs to text
func (s DetailServerCFG) StringText() string {
	var values string

	brokerCFG := s.Server.BrokerCFG
	if brokerCFG != nil {
		elements := [][]string{{"0", "Broker CFG:"}, {"1", "ID: " + (*brokerCFG).ID}, {"1", "Name: " + (*brokerCFG).Name}, {"1", "Instance: " + (*brokerCFG).Instance}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "brokerCFG: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the BrokerCFGs to csv
func (s DetailServerCFG) StringCSV() string {
	var p []DetailBrokerCFG
	if s.Server.BrokerCFG != nil {
		p = append(p, *s.Server.BrokerCFG)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the BrokerCFGs to json
func (s DetailServerCFG) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerCFGs to yaml
func (s DetailServerCFG) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
