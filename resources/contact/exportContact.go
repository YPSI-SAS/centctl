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

package contact

//ExportContact represents the caracteristics of a Contact
type ExportContact struct {
	Name      string `json:"name" yaml:"name"`
	Alias     string `json:"alias" yaml:"alias"`
	Email     string `json:"email" yaml:"email"`
	Pager     string `json:"pager" yaml:"pager"`
	GuiAccess string `json:"gui access" yaml:"gui access"`
	Admin     string `json:"admin" yaml:"admin"`
	Activate  string `json:"activate" yaml:"activate"`
}

//ExportResultContact represents a contact array send by the API
type ExportResultContact struct {
	Contacts []ExportContact `json:"result" yaml:"result"`
}
