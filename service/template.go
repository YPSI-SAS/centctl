package service

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Template represents the caracteristics of a service template
type Template struct {
	Description string `json:"description"` //Template Description
}

//ResultTemplate represents a service template array
type ResultTemplate struct {
	Templates []Template `json:"result"`
}

//TemplateServer represents a server with informations
type TemplateServer struct {
	Server TemplateInformations `json:"server"`
}

//TemplateInformations represents the informations of the server
type TemplateInformations struct {
	Name      string     `json:"name"`
	Templates []Template `json:"templates"`
}

//StringText permits to display the caracteristics of the service templates to text
func (s TemplateServer) StringText() string {
	var values string = "Service template list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Templates); i++ {
		values += s.Server.Templates[i].Description + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the service templates to csv
func (s TemplateServer) StringCSV() string {
	var values string = "Server,Description\n"
	for i := 0; i < len(s.Server.Templates); i++ {
		values += s.Server.Name + "," + s.Server.Templates[i].Description + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the service templates to json
func (s TemplateServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the service templates to yaml
func (s TemplateServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
