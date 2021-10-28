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
	"sort"
	"strconv"
	"strings"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//RealtimeService represents the caracteristics of a service
type RealtimeService struct {
	ServiceID      int                           `json:"id" yaml:"id"`     //Service ID
	Name           string                        `json:"name" yaml:"name"` //Service description
	RealtimeParent `json:"parent" yaml:"parent"` //Parent of service
	RealtimeStatus `json:"status" yaml:"status"` //State of the service
	Information    string                        `json:"information" yaml:"information"`     //Srevice Output
	Acknowledged   bool                          `json:"acknowledged" yaml:"acknowledged"`   //If the service is acknowledge or not
	ActiveCheck    bool                          `json:"active_checks" yaml:"active_checks"` //If the service is activate or not
}

type RealtimeParent struct {
	ID       int    `json:"id" yaml:"id" csv:"ParentID"`
	Name     string `json:"name" yaml:"name" csv:"ParentName"`
	Address  string `json:"fqdn" yaml:"fqdn" csv:"ParentAddress"`
	PollerID int    `json:"poller_id" yaml:"poller_id" csv:"ParentPollerID"` //Poller ID
}

type RealtimeStatus struct {
	Code         int    `json:"code" yaml:"code" csv:"StatusCode"`
	Name         string `json:"name" yaml:"name" csv:"StatusName"`
	SeverityCode int    `json:"severity_code" yaml:"severity_code" csv:"StatusSeverityCode"`
}

//RealtimeServer represents a server with informations
type RealtimeServer struct {
	Server RealtimeInformations `json:"server" yaml:"server"`
}

//RealtimeInformations represents the informations of the server
type RealtimeInformations struct {
	Name     string            `json:"name" yaml:"name"`
	Services []RealtimeService `json:"services" yaml:"services"`
}

type RealtimeResultBody struct {
	ListServices []RealtimeService `json:"result" yaml:"result"`
}

//StringText permits to display the caracteristics of the services to text
func (s RealtimeServer) StringText() string {
	sort.SliceStable(s.Server.Services, func(i, j int) bool {
		return strings.ToLower(s.Server.Services[i].Name) < strings.ToLower(s.Server.Services[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Parent ID", "Parent name", "Parent address", "PollerID", "Status code", "Status name", "Acknowledged", "ActiveCheck"})
	for i := 0; i < len(s.Server.Services); i++ {
		table = append(table, []string{strconv.Itoa(s.Server.Services[i].ServiceID), s.Server.Services[i].Name, strconv.Itoa(s.Server.Services[i].RealtimeParent.ID), s.Server.Services[i].RealtimeParent.Name, s.Server.Services[i].RealtimeParent.Address, strconv.Itoa(s.Server.Services[i].RealtimeParent.PollerID), strconv.Itoa(s.Server.Services[i].RealtimeStatus.Code), s.Server.Services[i].RealtimeStatus.Name, strconv.FormatBool(s.Server.Services[i].Acknowledged), strconv.FormatBool(s.Server.Services[i].ActiveCheck)})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the services to csv
func (s RealtimeServer) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.Services)
	return string(b)
}

//StringJSON permits to display the caracteristics of the services to json
func (s RealtimeServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the services to yaml
func (s RealtimeServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
