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

package resourceCFG

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailResourceCFG represents the caracteristics of a resourceCFG
type DetailResourceCFG struct {
	ID       string   `json:"id" yaml:"id"`           //resourceCFG ID
	Name     string   `json:"name" yaml:"name"`       //resourceCFG name
	Value    string   `json:"value" yaml:"value"`     //resourceCFG value
	Comment  string   `json:"comment" yaml:"comment"` //resourceCFG comment
	Activate string   `json:"activate"`
	Instance []string `json:"instance"`
}

//DetailResult represents a poller array
type DetailResult struct {
	ResourceCFG []DetailResourceCFG `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name        string             `json:"name" yaml:"name"`
	ResourceCFG *DetailResourceCFG `json:"resourceCFG" yaml:"resourceCFG"`
}

//StringText permits to display the caracteristics of the resourceCFG to text
func (s DetailServer) StringText() string {
	var values string
	resourceCFG := s.Server.ResourceCFG
	if resourceCFG != nil {
		elements := [][]string{{"0", "ResourceCFG:"}}
		elements = append(elements, []string{"1", "ID: " + (*resourceCFG).ID})
		elements = append(elements, []string{"1", "Name: " + (*resourceCFG).Name})
		elements = append(elements, []string{"1", "Value: " + (*resourceCFG).Value})
		elements = append(elements, []string{"1", "Comment: " + (*resourceCFG).Comment})
		elements = append(elements, []string{"1", "Activate: " + (*resourceCFG).Activate})
		if len((*resourceCFG).Instance) != 0 {
			var instances string
			for index, inst := range (*resourceCFG).Instance {
				instances += inst
				if index != len((*resourceCFG).Instance)-1 {
					instances += " | "
				}
			}
			elements = append(elements, []string{"1", "Instances: " + instances})
		} else {
			elements = append(elements, []string{"1", "Instances: []"})
		}

		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "resourceCFG: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the resourceCFG to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,ID,Name,Value,Comment,Activate,Instance\n"
	resourceCFG := s.Server.ResourceCFG
	if resourceCFG != nil {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + (*resourceCFG).ID + "\"" + "," + "\"" + (*resourceCFG).Name + "\"" + "," + "\"" + (*resourceCFG).Value + "\"" + "," + "\"" + (*resourceCFG).Comment + "\"" + "," + "\"" + (*resourceCFG).Activate + "\"" + "," + "\""
		for index, inst := range (*resourceCFG).Instance {
			values += inst
			if index != len((*resourceCFG).Instance)-1 {
				values += "|"
			}
		}
		values += "\"" + "\n"
	} else {
		values += ",,,,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the resourceCFG to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the resourceCFG to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
