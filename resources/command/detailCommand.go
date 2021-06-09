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

//DetailCommand represents the caracteristics of a Command
type DetailCommand struct {
	ID      string `json:"id" yaml:"id"`     //Command ID
	Name    string `json:"name" yaml:"name"` //Command Name
	Type    string `json:"type" yaml:"type"` //Command type
	CmdLine string `json:"cmd_line" yaml:"cmd_line"`

	Line json.RawMessage `json:"line,omitempty" yaml:"line,omitempty"` //Command line
}

//DetailResult represents a Command array send by the API
type DetailResult struct {
	Commands []DetailCommand `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name    string         `json:"name" yaml:"name"`
	Command *DetailCommand `json:"command" yaml:"command"`
}

//StringText permits to display the caracteristics of the commands to text
func (s DetailServer) StringText() string {
	var values string = "Command list for server " + s.Server.Name + ": \n"

	cmd := s.Server.Command
	if cmd != nil {
		values += "ID: " + (*cmd).ID + "\t"
		values += "Name: " + (*cmd).Name + "\t"
		values += "Type: " + (*cmd).Type + "\t"
		values += "CmdLine: " + (*cmd).CmdLine + "\n"
	} else {
		values += "command: null\n"

	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the commands to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,ID,Name,Type,CmdLine\n"
	values += s.Server.Name + ","
	cmd := s.Server.Command
	if cmd != nil {
		values += (*cmd).ID + ","
		values += (*cmd).Name + ","
		values += (*cmd).Type + ","
		values += (*cmd).CmdLine + "\n"
	} else {
		values += ",,,\n"

	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the commands to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the commands to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
