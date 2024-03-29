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

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailEngineCFG represents the caracteristics of a EngineCFG
type DetailEngineCFG struct {
	ID       string `json:"nagios id" yaml:"nagios id"`           //EngineCFG ID
	Name     string `json:"nagios name" yaml:"nagios name"`       //EngineCFG name
	Comment  string `json:"nagios comment" yaml:"nagios comment"` //EngineCFG comment
	Instance string `json:"instance" yaml:"instance"`             //EngineCFG instance
}

//DetailResultEngineCFG represents a poller array
type DetailResultEngineCFG struct {
	EngineCFG []DetailEngineCFG `json:"result" yaml:"result"`
}

//DetailServerEngineCFG represents a server with informations
type DetailServerEngineCFG struct {
	Server DetailInformationsEngineCFG `json:"server" yaml:"server"`
}

//DetailInformationsEngineCFG represents the informations of the server
type DetailInformationsEngineCFG struct {
	Name      string           `json:"name" yaml:"name"`
	EngineCFG *DetailEngineCFG `json:"engine_cfg" yaml:"engine_cfg"`
}

//StringText permits to display the caracteristics of the EngineCFG to text
func (s DetailServerEngineCFG) StringText() string {
	var values string

	engineCFG := s.Server.EngineCFG
	if engineCFG != nil {
		elements := [][]string{{"0", "EngineCFG:"}, {"1", "ID: " + (*engineCFG).ID}, {"1", "Name: " + (*engineCFG).Name}, {"1", "Instance: " + (*engineCFG).Instance}, {"1", "Comment: " + (*engineCFG).Comment}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "engineCFG: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the EngineCFG to csv
func (s DetailServerEngineCFG) StringCSV() string {
	var p []DetailEngineCFG
	if s.Server.EngineCFG != nil {
		p = append(p, *s.Server.EngineCFG)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the EngineCFG to json
func (s DetailServerEngineCFG) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the EngineCFG to yaml
func (s DetailServerEngineCFG) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
