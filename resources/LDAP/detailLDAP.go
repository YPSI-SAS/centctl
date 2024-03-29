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
	"centctl/resources"
	"encoding/json"
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailLDAP represents the caracteristics of a LDAP
type DetailLDAP struct {
	ID          string            `json:"id" yaml:"id"`                   //LDAP ID
	Name        string            `json:"name" yaml:"name"`               //LDAP name
	Description string            `json:"description" yaml:"description"` //LDAP Description
	Status      string            `json:"status" yaml:"status"`           //LDAP Status
	Servers     DetailLDAPServers `json:"servers" yaml:"servers"`
}

type DetailLDAPServers []DetailLDAPServer

func (t DetailLDAPServers) MarshalCSV() ([]byte, error) {
	var value string
	for i, server := range t {
		value += server.ID + "|" + server.Address + "|" + server.Port + "|" + server.SSL + "|" + server.TLS + "|" + server.Order
		if i < len(t)-1 {
			value += ","
		}
	}
	return []byte(value), nil
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
	var values string
	ldap := s.Server.LDAP
	if ldap != nil {
		elements := [][]string{{"0", "LDAP:"}}
		elements = append(elements, []string{"1", "ID: " + (*ldap).ID})
		elements = append(elements, []string{"1", "Name: " + (*ldap).Name})
		elements = append(elements, []string{"1", "Status: " + (*ldap).Status})
		elements = append(elements, []string{"1", "Description: " + (*ldap).Description})
		if len((*ldap).Servers) == 0 {
			elements = append(elements, []string{"1", "Servers: []"})
		} else {
			elements = append(elements, []string{"1", "Servers:"})
			for _, server := range (*ldap).Servers {
				elements = append(elements, []string{"2", "ID: " + server.ID})
				elements = append(elements, []string{"3", "Address: " + server.Address})
				elements = append(elements, []string{"3", "Port: " + server.Port})
				elements = append(elements, []string{"3", "SSL: " + server.SSL})
				elements = append(elements, []string{"3", "TLS: " + server.TLS})
				elements = append(elements, []string{"3", "Order: " + server.Order})
			}
		}

		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "LDAP: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the LDAP to csv
func (s DetailServer) StringCSV() string {
	var p []DetailLDAP
	if s.Server.LDAP != nil {
		p = append(p, *s.Server.LDAP)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
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
