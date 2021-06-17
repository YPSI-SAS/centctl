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

package command

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Command represents the caracteristics of a Command
type Command struct {
	ID      string `json:"id" yaml:"id"`     //Command ID
	Name    string `json:"name" yaml:"name"` //Command Name
	Type    string `json:"type" yaml:"type"` //Command type
	CmdLine string `json:"cmd_line" yaml:"cmd_line"`

	Line json.RawMessage `json:"line,omitempty" yaml:"line,omitempty"` //Command line
}

//Result represents a Command array send by the API
type Result struct {
	Commands []Command `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name     string    `json:"name" yaml:"name"`
	Commands []Command `json:"commands" yaml:"commands"`
}

//StringText permits to display the caracteristics of the commands to text
func (s Server) StringText() string {
	var values string = "Command list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Commands); i++ {
		values += "ID: " + s.Server.Commands[i].ID + "\t"
		values += "Name: " + s.Server.Commands[i].Name + "\t"
		values += "Type: " + s.Server.Commands[i].Type + "\t"
		values += "CmdLine: " + s.Server.Commands[i].CmdLine + "\n"

	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the commands to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Type,CmdLine\n"
	for i := 0; i < len(s.Server.Commands); i++ {
		values += s.Server.Name + ","
		values += s.Server.Commands[i].ID + ","
		values += s.Server.Commands[i].Name + ","
		values += s.Server.Commands[i].Type + ","
		values += s.Server.Commands[i].CmdLine + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the commands to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the commands to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
