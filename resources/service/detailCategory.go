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
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailCategory represents the caracteristics of a service Category
type DetailCategory struct {
	ID               string                          `json:"id" yaml:"id"`
	Name             string                          `json:"name" yaml:"name"` //Category name
	Alias            string                          `json:"alias" yaml:"alias"`
	Level            string                          `json:"level" yaml:"level"`
	Services         []DetailCategoryService         `json:"services" yaml:"services"`
	ServiceTemplates []DetailCategoryServiceTemplate `json:"service_templates" yaml:"service_templates"`
}

//DetailResultCategory represents a service Category array
type DetailResultCategory struct {
	Categories []DetailCategory `json:"result" yaml:"result"`
}

//DetailCategoryService represents the caracteristics of a service
type DetailCategoryService struct {
	HostID             string `json:"host id" yaml:"host id"`
	HostName           string `json:"host name" yaml:"host name"`
	ServiceID          string `json:"service id" yaml:"service id"`
	ServiceDescription string `json:"service description" yaml:"service description"`
}

//DetailResultCategoryService represents a service array
type DetailResultCategoryService struct {
	Services []DetailCategoryService `json:"result" yaml:"result"`
}

//DetailCategoryServiceTemplate represents the caracteristics of a service
type DetailCategoryServiceTemplate struct {
	TemplateID                 string `json:"template id" yaml:"template id"`
	ServiceTemplateDescription string `json:"service template description" yaml:"service template description"`
}

//DetailResultCategoryServiceTemplate represents a service array
type DetailResultCategoryServiceTemplate struct {
	ServiceTemplates []DetailCategoryServiceTemplate `json:"result" yaml:"result"`
}

//DetailCategoryServer represents a server with informations
type DetailCategoryServer struct {
	Server DetailCategoryInformations `json:"server" yaml:"server"`
}

//DetailCategoryInformations represents the informations of the server
type DetailCategoryInformations struct {
	Name     string          `json:"name" yaml:"name"`
	Category *DetailCategory `json:"category" yaml:"category"`
}

//StringText permits to display the caracteristics of the service categories to text
func (s DetailCategoryServer) StringText() string {
	var values string = "Service categories list for server " + s.Server.Name + ": \n"

	category := s.Server.Category
	if category != nil {
		values += (*category).ID + "\t"
		values += (*category).Name + "\t"
		values += (*category).Alias + "\t"
		values += (*category).Level + "\n"
	} else {
		values += "category: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the service category to csv
func (s DetailCategoryServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Level\n"
	values += s.Server.Name + ","
	category := s.Server.Category
	if category != nil {
		values += (*category).ID + "," + (*category).Name + "," + (*category).Alias + "," + (*category).Level + "\n"

	} else {
		values += ",,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the service category to json
func (s DetailCategoryServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the service category to yaml
func (s DetailCategoryServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
