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

package contact

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailContact represents the caracteristics of a contact
type DetailContact struct {
	ID        string `json:"id" yaml:"id"`       //Contact ID
	Name      string `json:"name" yaml:"name"`   //Contact Name
	Alias     string `json:"alias" yaml:"alias"` //Contact Alias
	Email     string `json:"email" yaml:"email"` //Contact Email
	Pager     string `json:"pager" yaml:"pager"`
	GuiAccess string `json:"gui access" yaml:"gui access"`
	Admin     string `json:"admin" yaml:"admin"`
	Activate  string `json:"activate" yaml:"activate"`
}

//DetailResult represents a contact array send by the API
type DetailResult struct {
	DetailContacts []DetailContact `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name    string         `json:"name" yaml:"name"`
	Contact *DetailContact `json:"contact" yaml:"contact"`
}

//StringText permits to display the caracteristics of the contacts to text
func (s DetailServer) StringText() string {
	var values string
	contact := s.Server.Contact
	if contact != nil {
		elements := [][]string{{"0", "Contact:"}, {"1", "ID: " + (*contact).ID}, {"1", "Name: " + (*contact).Name}, {"1", "Alias: " + (*contact).Alias}, {"1", "Email: " + (*contact).Email}, {"1", "Pager: " + (*contact).Pager}, {"1", "GuiAcces: " + (*contact).GuiAccess}, {"1", "Admin: " + (*contact).Admin}, {"1", "Activate: " + (*contact).Activate}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "contact: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the contacts to csv
func (s DetailServer) StringCSV() string {
	var p []DetailContact
	if s.Server.Contact != nil {
		p = append(p, *s.Server.Contact)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the contacts to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the contacts to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
