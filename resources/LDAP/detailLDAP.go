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

package LDAP

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailLDAP represents the caracteristics of a LDAP
type DetailLDAP struct {
	ID          string             `json:"id" yaml:"id"`                   //LDAP ID
	Name        string             `json:"name" yaml:"name"`               //LDAP name
	Description string             `json:"description" yaml:"description"` //LDAP Description
	Status      string             `json:"status" yaml:"status"`           //LDAP Status
	Servers     []DetailLDAPServer `json:"servers" yaml:"servers"`
}

//DetailLDAPServer represents the caracteristics of a member
type DetailLDAPServer struct {
	ID      string `json:"id" yaml:"id"`
	Address string `json:"address" yaml:"address"`
	Port    string `json:"port" yaml:"port"`
	SSL     string `json:"ssl" yaml:"ssl"`
	TLS     string `json:"tls" yaml:"tls"`
	Order   string `json:"order" yaml:"order"`
}

//DetailResultServer represents a member array
type DetailResultServer struct {
	Servers []DetailLDAPServer `json:"result" yaml:"result"`
}

//DetailResult represents a poller array
type DetailResult struct {
	LDAP []DetailLDAP `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name string      `json:"name" yaml:"name"`
	LDAP *DetailLDAP `json:"ldap" yaml:"ldap"`
}

//StringText permits to display the caracteristics of the LDAP to text
func (s DetailServer) StringText() string {
	var values string = "LDAP list for server " + s.Server.Name + ": \n"
	ldap := s.Server.LDAP
	if ldap != nil {
		values += "ID: " + (*ldap).ID + "\t"
		values += "Name: " + (*ldap).Name + "\t"
		values += "Status: " + (*ldap).Status + "\t"
		values += "Description: " + (*ldap).Description + "\n"
	} else {
		values += "LDAP: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the LDAP to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,ID,Name,Status,Description\n"
	values += s.Server.Name + ","
	ldap := s.Server.LDAP
	if ldap != nil {
		values += "\"" + (*ldap).ID + "\"" + "," + "\"" + (*ldap).Name + "\"" + "," + "\"" + (*ldap).Status + "\"" + "," + "\"" + (*ldap).Description + "\"" + "\n"
	} else {
		values += ",,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the LDAP to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the LDAP to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
