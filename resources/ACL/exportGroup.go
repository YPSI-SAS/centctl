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

package ACL

//ExportGroup represents the caracteristics of a ACL group
type ExportGroup struct {
	Name     string `json:"name" yaml:"name"`
	Alias    string `json:"alias" yaml:"alias"`
	Activate string `json:"activate" yaml:"activate"`

	Contact      []ExportGroupContact
	ContactGroup []ExportGroupContactGroup
	Menu         []ExportGroupMenu
	Action       []ExportGroupAction
	Resource     []ExportGroupResource
}

//ExportResult represents a ACL group array send by the API
type ExportResult struct {
	GroupACL []ExportGroup `json:"result" yaml:"result"`
}

//ExportGroupContact represents the caracteristics of a contact
type ExportGroupContact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultContact represents a contact array send by the API
type ExportResultContact struct {
	GroupContact []ExportGroupContact `json:"result" yaml:"result"`
}

//ExportGroupContactGroup represents the caracteristics of a contact group
type ExportGroupContactGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultContactGroup represents a contact group array send by the API
type ExportResultContactGroup struct {
	GroupContactGroup []ExportGroupContactGroup `json:"result" yaml:"result"`
}

//ExportGroupMenu represents the caracteristics of a menu
type ExportGroupMenu struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultMenu represents a menu array send by the API
type ExportResultMenu struct {
	GroupMenu []ExportGroupMenu `json:"result" yaml:"result"`
}

//ExportGroupAction represents the caracteristics of a action
type ExportGroupAction struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultAction represents a action array send by the API
type ExportResultAction struct {
	GroupAction []ExportGroupAction `json:"result" yaml:"result"`
}

//ExportGroupResource represents the caracteristics of a resource
type ExportGroupResource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultResource represents a resource array send by the API
type ExportResultResource struct {
	GroupResource []ExportGroupResource `json:"result" yaml:"result"`
}
