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

package trap

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailTrap represents the caracteristics of a Trap
type DetailTrap struct {
	ID           string `json:"id" yaml:"id"`                     //Trap ID
	Name         string `json:"name" yaml:"name"`                 //Trap name
	Oid          string `json:"oid" yaml:"oid"`                   //Trap Oid
	Manufacturer string `json:"manufacturer" yaml:"manufacturer"` //Trap bool state
}

//DetailResult represents a poller array
type DetailResult struct {
	Traps []DetailTrap `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name string      `json:"name" yaml:"name"`
	Trap *DetailTrap `json:"trap" yaml:"trap"`
}

//StringText permits to display the caracteristics of the Traps to text
func (s DetailServer) StringText() string {
	var values string

	trap := s.Server.Trap
	if trap != nil {
		elements := [][]string{{"0", "Trap:"}}
		elements = append(elements, []string{"1", "ID: " + (*trap).ID})
		elements = append(elements, []string{"1", "Name: " + (*trap).Name})
		elements = append(elements, []string{"1", "Oid: " + (*trap).Oid})
		elements = append(elements, []string{"1", "Manufacturer: " + (*trap).Manufacturer})
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "trap: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the Traps to csv
func (s DetailServer) StringCSV() string {
	var p []DetailTrap
	if s.Server.Trap != nil {
		p = append(p, *s.Server.Trap)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the Traps to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the Traps to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
