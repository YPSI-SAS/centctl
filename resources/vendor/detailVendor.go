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

package vendor

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailVendor represents the caracteristics of a Vendor
type DetailVendor struct {
	ID    string `json:"id" yaml:"id"`       //Vendor ID
	Name  string `json:"name" yaml:"name"`   //Vendor name
	Alias string `json:"alias" yaml:"alias"` //Vendor Alias
}

//DetailResult represents a poller array
type DetailResult struct {
	Vendors []DetailVendor `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name   string        `json:"name" yaml:"name"`
	Vendor *DetailVendor `json:"vendor" yaml:"vendor"`
}

//StringText permits to display the caracteristics of the Vendors to text
func (s DetailServer) StringText() string {
	var values string
	vendor := s.Server.Vendor
	if vendor != nil {
		elements := [][]string{{"0", "Vendor:"}}
		elements = append(elements, []string{"1", "ID: " + (*vendor).ID})
		elements = append(elements, []string{"1", "Name: " + (*vendor).Name})
		elements = append(elements, []string{"1", "Alias: " + (*vendor).Alias})
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "vendor: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the Vendors to csv
func (s DetailServer) StringCSV() string {
	var p []DetailVendor
	if s.Server.Vendor != nil {
		p = append(p, *s.Server.Vendor)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the Vendors to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the Vendors to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
