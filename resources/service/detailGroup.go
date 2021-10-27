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

package service

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailGroup represents the caracteristics of a service Group
type DetailGroup struct {
	ID                string                        `json:"id" yaml:"id"`
	Name              string                        `json:"name" yaml:"name"` //Group name
	Alias             string                        `json:"alias" yaml:"alias"`
	Services          []DetailGroupService          `json:"services" yaml:"services"`
	HostGroupServices []DetailGroupHostGroupService `json:"hostgroup_services" yaml:"hostgroup_services"`
}

//DetailResultGroup represents a service Group array
type DetailResultGroup struct {
	DetailGroups []DetailGroup `json:"result" yaml:"result"`
}

//DetailGroupService represents the caracteristics of a service
type DetailGroupService struct {
	HostID             string `json:"host id" yaml:"host id"`
	HostName           string `json:"host name" yaml:"host name"`
	ServiceID          string `json:"service id" yaml:"service id"`
	ServiceDescription string `json:"service description" yaml:"service description"`
}

//DetailResultGroupService represents a service array
type DetailResultGroupService struct {
	Services []DetailGroupService `json:"result" yaml:"result"`
}

//DetailGroupHostGroupService represents the caracteristics of a host group service
type DetailGroupHostGroupService struct {
	HostGroupID        string `json:"hostgroup id" yaml:"hostgroup id"`
	HostGroupName      string `json:"hostgroup name" yaml:"hostgroup name"`
	ServiceID          string `json:"service id" yaml:"service id"`
	ServiceDescription string `json:"service description" yaml:"service description"`
}

//DetailResultHostGroupService represents a host group service array
type DetailResultHostGroupService struct {
	Services []DetailGroupHostGroupService `json:"result" yaml:"result"`
}

//DetailGroupServer represents a server with informations
type DetailGroupServer struct {
	Server DetailGroupInformations `json:"server" yaml:"server"`
}

//DetailGroupInformations represents the informations of the server
type DetailGroupInformations struct {
	Name  string       `json:"name" yaml:"name"`
	Group *DetailGroup `json:"group" yaml:"group"`
}

//StringText permits to display the caracteristics of the service groups to text
func (s DetailGroupServer) StringText() string {
	var values string
	group := s.Server.Group
	if group != nil {
		elements := [][]string{{"0", "Group service:"}, {"1", "ID: " + (*group).ID}, {"1", "Name: " + (*group).Name + "\t" + "Alias: " + (*group).Alias}}
		if len((*group).Services) == 0 {
			elements = append(elements, []string{"1", "Services: []"})
		} else {
			elements = append(elements, []string{"1", "Services:"})
			for _, service := range (*group).Services {
				elements = append(elements, []string{"2", "Host: " + service.HostName + " (ID=" + service.HostID + ")"})
				elements = append(elements, []string{"2", "Service: " + service.ServiceDescription + " (ID=" + service.ServiceID + ")"})
			}
		}
		if len((*group).HostGroupServices) == 0 {
			elements = append(elements, []string{"1", "Host group services: []"})
		} else {
			elements = append(elements, []string{"1", "Host group services:"})
			for _, service := range (*group).HostGroupServices {
				elements = append(elements, []string{"2", "Host group: " + service.HostGroupName + " (ID=" + service.HostGroupID + ")"})
				elements = append(elements, []string{"2", "Host group service: " + service.ServiceDescription + " (ID=" + service.ServiceID + ")"})
			}
		}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "group: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the service groups to csv
func (s DetailGroupServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias\n"
	values += s.Server.Name + ","
	group := s.Server.Group
	if group != nil {
		values += "\"" + (*group).ID + "\"" + "," + "\"" + (*group).Name + "\"" + "," + "\"" + (*group).Alias + "\"" + "\n"
	} else {
		values += ",,\n"
	}

	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the service groups to json
func (s DetailGroupServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the service groups to yaml
func (s DetailGroupServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
