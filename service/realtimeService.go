/*
MIT License

Copyright (c) 2020 YPSI SAS
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
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//RealtimeService represents the caracteristics of a service
type RealtimeService struct {
	ServiceID    string `json:"service_id"`    //Service ID
	Description  string `json:"description"`   //Service description
	HostID       string `json:"host_id"`       //Host ID of the service
	HostName     string `json:"host_name"`     //Host name of the service
	State        string `json:"state"`         //State of the service
	Output       string `json:"output"`        //Srevice Output
	Acknowledged string `json:"acknowledged"`  //If the service is acknowledge or not
	Activate     string `json:"active_checks"` //If the service is activate or not
}

//RealtimeServer represents a server with informations
type RealtimeServer struct {
	Server RealtimeInformations `json:"server"`
}

//RealtimeInformations represents the informations of the server
type RealtimeInformations struct {
	Name     string            `json:"name"`
	Services []RealtimeService `json:"services"`
}

//StringText permits to display the caracteristics of the services to text
func (s RealtimeServer) StringText() string {
	var values string = "Service list for server=" + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Services); i++ {
		values += "ID: " + s.Server.Services[i].ServiceID + "\t"
		values += "Description: " + s.Server.Services[i].Description + "\t"
		values += "Host ID: " + s.Server.Services[i].HostID + "\t"
		values += "Host name: " + s.Server.Services[i].HostName + "\t"
		values += "State: " + GetState(s.Server.Services[i].State) + "\t"
		values += "Output: " + s.Server.Services[i].Output + "\t"
		values += "Acknowledged: " + GetAcknowledgment(s.Server.Services[i].Acknowledged) + "\t"
		values += "Activate: " + s.Server.Services[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the services to csv
func (s RealtimeServer) StringCSV() string {
	var values string = "Server,ID,Description,HostID,HostName,State,Output,Acknowledged,Activate\n"
	for i := 0; i < len(s.Server.Services); i++ {
		values += s.Server.Name + "," + s.Server.Services[i].ServiceID + "," + s.Server.Services[i].Description + "," + s.Server.Services[i].HostID + "," + s.Server.Services[i].HostName + "," + GetState(s.Server.Services[i].State) + "," + s.Server.Services[i].Output + "," + GetAcknowledgment(s.Server.Services[i].Acknowledged) + "," + s.Server.Services[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
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

//GetState permits to obtain the state of the service
func GetState(stateValue string) string {
	state := ""
	switch stateValue {
	case "0":
		state = "OK"
	case "1":
		state = "Warning"
	case "2":
		state = "Critical"
	case "3":
		state = "Unknow"
	case "4":
		state = "Pending"
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
