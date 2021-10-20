/*
MIT License

Copyright (c)  2020-2021 YPSI SAS


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
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Trap represents the caracteristics of a Trap
type Trap struct {
	ID           string `json:"id" yaml:"id"`                     //Trap ID
	Name         string `json:"name" yaml:"name"`                 //Trap name
	Oid          string `json:"oid" yaml:"oid"`                   //Trap Oid
	Manufacturer string `json:"manufacturer" yaml:"manufacturer"` //Trap bool state
}

//Result represents a poller array
type Result struct {
	Traps []Trap `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name  string `json:"name" yaml:"name"`
	Traps []Trap `json:"traps" yaml:"traps"`
}

//StringText permits to display the caracteristics of the Traps to text
func (s Server) StringText() string {
	var values string = "Trap list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Traps); i++ {
		values += "ID: " + s.Server.Traps[i].ID + "\t"
		values += "Name: " + s.Server.Traps[i].Name + "\t"
		values += "Oid: " + s.Server.Traps[i].Oid + "\t"
		values += "Manufacturer: " + s.Server.Traps[i].Manufacturer + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the Traps to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Oid,Manufacturer\n"
	for i := 0; i < len(s.Server.Traps); i++ {
		values += s.Server.Name + "," + s.Server.Traps[i].ID + "," + s.Server.Traps[i].Name + "," + s.Server.Traps[i].Oid + "," + s.Server.Traps[i].Manufacturer + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the Traps to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the Traps to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
