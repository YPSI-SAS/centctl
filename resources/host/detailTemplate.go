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

package host

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailTemplate represents the caracteristics of a host template
type DetailTemplate struct {
	Name     string `json:"name" yaml:"name"` //Template Name
	ID       string `json:"id" yaml:"id"`     //Template ID
	Alias    string `json:"alias" yaml:"alias"`
	Address  string `json:"address" yaml:"address"`
	Activate string `json:"activate" yaml:"activate"`
}

//DetailResultTemplate represents a host template array
type DetailResultTemplate struct {
	DetailTemplates []DetailTemplate `json:"result" yaml:"result"`
}

//DetailTemplateServer represents a server with informations
type DetailTemplateServer struct {
	Server DetailTemplateInformations `json:"server" yaml:"server"`
}

//DetailTemplateInformations represents the informations of the server
type DetailTemplateInformations struct {
	Name     string          `json:"name" yaml:"name"`
	Template *DetailTemplate `json:"template" yaml:"template"`
}

//StringText permits to display the caracteristics of the host templates to text
func (s DetailTemplateServer) StringText() string {
	var values string
	template := s.Server.Template
	if template != nil {
		elements := [][]string{{"0", "Template host:"}, {"1", "ID: " + (*template).ID}, {"1", "Name: " + (*template).Name}, {"1", "Alias: " + (*template).Alias}, {"1", "Address: " + (*template).Address}, {"1", "Activate: " + (*template).Activate}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "template: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the host ResultTemplate to csv
func (s DetailTemplateServer) StringCSV() string {
	var p []DetailTemplate
	if s.Server.Template != nil {
		p = append(p, *s.Server.Template)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the host ResultTemplate to json
func (s DetailTemplateServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the host ResultTemplate to yaml
func (s DetailTemplateServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
