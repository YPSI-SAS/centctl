/*
MIT License

Copyright (c)  2020-2021 YPSI SAS


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

//ExportGroup represents the caracteristics of a group service
type ExportGroup struct {
	Name      string `json:"name" yaml:"name"`   //command Name
	Alias     string `json:"alias" yaml:"alias"` //command type
	Comment   string `json:"comment" yaml:"comment"`
	Activate  string `json:"activate" yaml:"activate"`
	GeoCoords string `json:"geo_coords" yaml:"geo_coords"`

	Services          []ExportGroupServices
	HostGroupServices []ExportGroupServicesHostGroupServices
}

//ExportResult represents a group service array send by the API
type ExportResult struct {
	GroupServices []ExportGroup `json:"result" yaml:"result"`
}

//ExportGroupServices represents the caracteristics of a service
type ExportGroupServices struct {
	HostID             string `json:"host id"`
	HostName           string `json:"host name"`
	ServiceID          string `json:"service id"`
	ServiceDescription string `json:"service description"`
}

//ExportGroupServicesHostGroupServices represents the caracteristics of a host group service
type ExportGroupServicesHostGroupServices struct {
	HostgroupID        string `json:"hostgroup id"`
	HostgroupName      string `json:"hostgroup name"`
	ServiceID          string `json:"service id"`
	ServiceDescription string `json:"service description"`
}

//ExportResultService represents a services array send by the API
type ExportResultService struct {
	GroupServices []ExportGroupServices `json:"result" yaml:"result"`
}

//ExportResultHostGroupServices represents a host group service array send by the API
type ExportResultHostGroupServices struct {
	HostGroupServices []ExportGroupServicesHostGroupServices `json:"result" yaml:"result"`
}
