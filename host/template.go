package host

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Template represents the caracteristics of a host template
type Template struct {
	Name string `json:"name"` //Template Name
	ID   string `json:"id"`   //Template ID
}

//ResultTemplate represents a host template array
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

//StringText permits to display the caracteristics of the host templates to text
func (s TemplateServer) StringText() string {
	var values string = "Host template list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Templates); i++ {
		values += s.Server.Templates[i].Name + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the host ResultTemplate to csv
func (s TemplateServer) StringCSV() string {
	var values string = "Server,Name\n"
	for i := 0; i < len(s.Server.Templates); i++ {
		values += s.Server.Name + "," + s.Server.Templates[i].Name + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the host ResultTemplate to json
func (s TemplateServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the host ResultTemplate to yaml
func (s TemplateServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
