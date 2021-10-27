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

	"gopkg.in/yaml.v2"
)

//DetailCategory represents the caracteristics of a host Category
type DetailCategory struct {
	ID      string                 `json:"id" yaml:"id"`     //Category ID
	Name    string                 `json:"name" yaml:"name"` //Category Name
	Alias   string                 `json:"alias" yaml:"alias"`
	Level   string                 `json:"level" yaml:"level"`
	Members []DetailCategoryMember `json:"members" yaml:"members"`
}

//DetailCategoryMember represents the caracteristics of a member
type DetailCategoryMember struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

//DetailResultCategoryMember represents a member array
type DetailResultCategoryMember struct {
	Members []DetailCategoryMember `json:"result" yaml:"result"`
}

//DetailResultCategory represents a host Category array
type DetailResultCategory struct {
	Categories []DetailCategory `json:"result" yaml:"result"`
}

//DetailCategoryServer represents a server with informations
type DetailCategoryServer struct {
	Server DetailCategoryInformations `json:"server" yaml:"server"`
}

//DetailCategoryInformations represents the informations of the server
type DetailCategoryInformations struct {
	Name     string          `json:"name" yaml:"name"`
	Category *DetailCategory `json:"category" yaml:"categoy"`
}

//StringText permits to display the caracteristics of the host categories to text
func (s DetailCategoryServer) StringText() string {
	var values string
	category := s.Server.Category
	if category != nil {
		elements := [][]string{{"0", "Category host:"}, {"1", "ID: " + (*category).ID}, {"1", "Name: " + (*category).Name + "\t" + "Alias: " + (*category).Alias}, {"1", "Level: " + (*category).Level}}
		if len((*category).Members) == 0 {
			elements = append(elements, []string{"1", "Members: []"})
		} else {
			elements = append(elements, []string{"1", "Members:"})
			for _, member := range (*category).Members {
				elements = append(elements, []string{"2", member.Name + " (ID=" + member.ID + ")"})
			}
		}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "category: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the host ResultCategory to csv
func (s DetailCategoryServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Level\n"
	values += s.Server.Name + ","
	category := s.Server.Category
	if category != nil {
		values += "\"" + (*category).ID + "\"" + "," + "\"" + (*category).Name + "\"" + "," + "\"" + (*category).Alias + "\"" + "," + "\"" + (*category).Level + "\"" + "\n"
	} else {
		values += ",,,\n"
	}

	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the host ResultCategory to json
func (s DetailCategoryServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the host ResultCategory to yaml
func (s DetailCategoryServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
