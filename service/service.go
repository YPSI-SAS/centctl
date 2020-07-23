package service

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Service represents the caracteristics of a service
type Service struct {
	ServiceID   string `json:"id"`          //Service ID
	Description string `json:"description"` //Service description
	HostID      string `json:"host id"`     //Host ID of the service
	HostName    string `json:"host name"`   //Host name of the service
	Activate    string `json:"activate"`    //If the service is activate or not
}

//Result represents a poller array
type Result struct {
	Services []Service `json:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name     string    `json:"name"`
	Services []Service `json:"services"`
}

//StringText permits to display the caracteristics of the services to text
func (s Server) StringText() string {
	var values string = "Service list for server" + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Services); i++ {
		values += "ID: " + s.Server.Services[i].ServiceID + "\t"
		values += "Description: " + s.Server.Services[i].Description + "\t"
		values += "Host ID: " + s.Server.Services[i].HostID + "\t"
		values += "Host name: " + s.Server.Services[i].HostName + "\t"
		values += "Activate: " + s.Server.Services[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the services to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Description,HostID,HostName,Activate\n"
	for i := 0; i < len(s.Server.Services); i++ {
		values += s.Server.Name + "," + s.Server.Services[i].ServiceID + "," + s.Server.Services[i].Description + "," + s.Server.Services[i].HostID + "," + s.Server.Services[i].HostName + "," + s.Server.Services[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the services to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the services to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
