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
	"sort"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//Category represents the caracteristics of a host Category
type Category struct {
	ID    string `json:"id" yaml:"id"`     //Category ID
	Name  string `json:"name" yaml:"name"` //Category Name
	Alias string `json:"alias" yaml:"alias"`
	Level string `json:"level" yaml:"level"`
}

//ResultCategory represents a host Category array
type ResultCategory struct {
	Categories []Category `json:"result" yaml:"result"`
}

//CategoryServer represents a server with informations
type CategoryServer struct {
	Server CategoryInformations `json:"server" yaml:"server"`
}

//CategoryInformations represents the informations of the server
type CategoryInformations struct {
	Name       string     `json:"name" yaml:"name"`
	Categories []Category `json:"categories" yaml:"categories"`
}

//StringText permits to display the caracteristics of the host categories to text
func (s CategoryServer) StringText() string {
	sort.SliceStable(s.Server.Categories, func(i, j int) bool {
		return strings.ToLower(s.Server.Categories[i].Name) < strings.ToLower(s.Server.Categories[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Alias", "Level"})
	for i := 0; i < len(s.Server.Categories); i++ {
		table = append(table, []string{s.Server.Categories[i].ID, s.Server.Categories[i].Name, s.Server.Categories[i].Alias, s.Server.Categories[i].Level})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the host ResultCategory to csv
func (s CategoryServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Level\n"
	for i := 0; i < len(s.Server.Categories); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + s.Server.Categories[i].ID + "\"" + "," + "\"" + s.Server.Categories[i].Name + "\"" + "," + "\"" + s.Server.Categories[i].Alias + "\"" + "," + "\"" + s.Server.Categories[i].Level + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the host ResultCategory to json
func (s CategoryServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the host ResultCategory to yaml
func (s CategoryServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
