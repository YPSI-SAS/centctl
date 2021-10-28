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
	"sort"
	"strings"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//Group represents the caracteristics of a host Group
type Group struct {
	Name string `json:"name" yaml:"name"` //Group Name
	ID   string `json:"id" yaml:"id"`     //Group ID
}

//ResultGroup represents a host Group array
type ResultGroup struct {
	Groups []Group `json:"result" yaml:"result"`
}

//GroupServer represents a server with informations
type GroupServer struct {
	Server GroupInformations `json:"server" yaml:"server"`
}

//GroupInformations represents the informations of the server
type GroupInformations struct {
	Name   string  `json:"name" yaml:"name"`
	Groups []Group `json:"groups" yaml:"groups"`
}

//StringText permits to display the caracteristics of the host groups to text
func (s GroupServer) StringText() string {
	sort.SliceStable(s.Server.Groups, func(i, j int) bool {
		return strings.ToLower(s.Server.Groups[i].Name) < strings.ToLower(s.Server.Groups[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name"})
	for i := 0; i < len(s.Server.Groups); i++ {
		table = append(table, []string{s.Server.Groups[i].ID, s.Server.Groups[i].Name})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the host ResultGroup to csv
func (s GroupServer) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.Groups)
	return string(b)
}

//StringJSON permits to display the caracteristics of the host ResultGroup to json
func (s GroupServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the host ResultGroup to yaml
func (s GroupServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
