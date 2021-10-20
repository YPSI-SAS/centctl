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

package LDAP

//ExportLDAP represents the caracteristics of a LDAP
type ExportLDAP struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Status      string `json:"status" yaml:"status"`

	Servers []ExportLDAPServer
}

//ExportResultLDAP represents a LDAP array send by the API
type ExportResultLDAP struct {
	LDAPs []ExportLDAP `json:"result" yaml:"result"`
}

//ExportLDAPServer represents the caracteristics of a Server
type ExportLDAPServer struct {
	Address string `json:"address"`
	Port    string `json:"port"`
	SSL     string `json:"ssl"`
	TLS     string `json:"tls"`
	Order   string `json:"order"`
}

//ExportResultLDAPServer represents a Server array send by the API
type ExportResultLDAPServer struct {
	Server []ExportLDAPServer `json:"result" yaml:"result"`
}
