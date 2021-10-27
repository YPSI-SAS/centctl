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
	"sort"
	"strings"

	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

/*
* Est encore essentiel pour la recherche du nom du poller dans la fonction NamePollerHost de requestAPIv1 appelée lors de l'ajout d'un service
 */

//RealtimeHost represents the caracteristics of a host
type RealtimeHost struct {
	ID           string `json:"id" yaml:"id"`                       //Host ID
	Name         string `json:"name" yaml:"name"`                   //Host name
	Alias        string `json:"alias" yaml:"alias"`                 //Host alias
	Address      string `json:"address" yaml:"address"`             //Host address
	State        string `json:"state" yaml:"state"`                 //State of the host
	Acknowledged string `json:"acknowledged" yaml:"acknowledged"`   //If the host is acknowledge or not
	Activate     string `json:"active_checks" yaml:"active_checks"` //If the host is activate or not
	PollerName   string `json:"instance_name" yaml:"instance_name"` //Poller name of the host
}

//RealtimeServer represents a server with informations
type RealtimeServer struct {
	Server RealtimeInformations `json:"server" yaml:"server"`
}

//RealtimeInformations represents the informations of the server
type RealtimeInformations struct {
	Name  string         `json:"name" yaml:"name"`
	Hosts []RealtimeHost `json:"hosts" yaml:"hosts"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s RealtimeServer) StringText() string {
	sort.SliceStable(s.Server.Hosts, func(i, j int) bool {
		return strings.ToLower(s.Server.Hosts[i].Name) < strings.ToLower(s.Server.Hosts[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Alias", "IP address", "State", "Acknowledged", "Activate", "Poller name"})
	for i := 0; i < len(s.Server.Hosts); i++ {
		table = append(table, []string{s.Server.Hosts[i].ID, s.Server.Hosts[i].Name, s.Server.Hosts[i].Alias, s.Server.Hosts[i].Address, GetState(s.Server.Hosts[i].State), GetAcknowledgment(s.Server.Hosts[i].Acknowledged), s.Server.Hosts[i].Activate, s.Server.Hosts[i].PollerName})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s RealtimeServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,IPAddress,State,Acknowledged,Activate,PollerName\n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + s.Server.Hosts[i].ID + "\"" + "," + "\"" + s.Server.Hosts[i].Name + "\"" + "," + "\"" + s.Server.Hosts[i].Alias + "\"" + "," + "\"" + s.Server.Hosts[i].Address + "\"" + "," + "\"" + GetState(s.Server.Hosts[i].State) + "\"" + "," + "\"" + GetAcknowledgment(s.Server.Hosts[i].Acknowledged) + "\"" + "," + "\"" + s.Server.Hosts[i].Activate + "\"" + "," + "\"" + s.Server.Hosts[i].PollerName + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s RealtimeServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s RealtimeServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}

//GetState permits to obtain the value of the state
func GetState(stateValue string) string {
	state := ""
	switch stateValue {
	case "0":
		state = "UP"
	case "1":
		state = "DOWN"
	case "2":
		state = "UNREACHABLE"
	case "4":
		state = "PENDING"
	}
	return state
}

//GetAcknowledgment permits to obtain the value of the acknowledgement
func GetAcknowledgment(acknowledgeValue string) string {
	acknowledge := ""
	switch acknowledgeValue {
	case "0":
		acknowledge = "no"
	case "1":
		acknowledge = "yes"
	}
	return acknowledge
}
