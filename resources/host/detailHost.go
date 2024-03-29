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

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailHost represents the caracteristics of a host
type DetailHost struct {
	ID       string `json:"id" yaml:"id"`             //Host ID
	Name     string `json:"name" yaml:"name"`         //Host name
	Alias    string `json:"alias" yaml:"alias" `      //Host alias
	Address  string `json:"address" yaml:"address"`   //Host address
	Activate string `json:"activate" yaml:"activate"` //If the host is activate or not

	Parent DetailHostParents `json:"parents" yaml:"parents"`
	Child  DetailHostChilds  `json:"children" yaml:"children"`
}

type DetailHostParents []DetailHostParent

func (t DetailHostParents) MarshalCSV() ([]byte, error) {
	var value string
	for i, parent := range t {
		value += parent.ID + "|" + parent.Name
		if i < len(t)-1 {
			value += ","
		}
	}
	return []byte(value), nil
}

type DetailHostChilds []DetailHostParent

func (t DetailHostChilds) MarshalCSV() ([]byte, error) {
	var value string
	for i, child := range t {
		value += child.ID + "|" + child.Name
		if i < len(t)-1 {
			value += ","
		}
	}
	return []byte(value), nil
}

//DetailResult represents a host Group array
type DetailResult struct {
	DetailHosts []DetailHost `json:"result" yaml:"result"`
}

//DetailHostParent represents the caracteristics of a parent
type DetailHostParent struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

//DetailResultHostParent represents a member array
type DetailResultHostParent struct {
	Parents DetailHostParents `json:"result" yaml:"result"`
}

//DetailHostChild represents the caracteristics of a child
type DetailHostChild struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

//DetailResultHostChild represents a member array
type DetailResultHostChild struct {
	Childs DetailHostChilds `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name string      `json:"name" yaml:"name"`
	Host *DetailHost `json:"host" yaml:"host"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s DetailServer) StringText() string {
	var values string
	host := s.Server.Host
	if host != nil {
		elements := [][]string{{"0", "Host:"}, {"1", "ID: " + (*host).ID}, {"1", "Name: " + (*host).Name + "\t" + "Alias: " + (*host).Alias}, {"1", "IP address: " + (*host).Address}, {"1", "Activate: " + (*host).Activate}}
		if len((*host).Child) == 0 {
			elements = append(elements, []string{"1", "Child: []"})
		} else {
			elements = append(elements, []string{"1", "Child:"})
			for _, child := range (*host).Child {
				elements = append(elements, []string{"2", child.Name + " (ID=" + child.ID + ")"})
			}
		}
		if len((*host).Parent) == 0 {
			elements = append(elements, []string{"1", "Parent: []"})
		} else {
			elements = append(elements, []string{"1", "Parent:"})
			for _, parent := range (*host).Parent {
				elements = append(elements, []string{"2", parent.Name + " (ID=" + parent.ID + ")"})
			}
		}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "Host: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s DetailServer) StringCSV() string {
	var p []DetailHost
	if s.Server.Host != nil {
		p = append(p, *s.Server.Host)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
