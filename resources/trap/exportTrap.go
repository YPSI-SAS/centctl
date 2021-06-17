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

package trap

//ExportTrap represents the caracteristics of a trap
type ExportTrap struct {
	Name         string `json:"name" yaml:"name"`
	Oid          string `json:"oid" yaml:"oid"`
	Manufacturer string `json:"manufacturer" yaml:"manufacturer"`

	Matchings []ExportTrapMatching
}

//ExportResultTrap represents a trap array send by the API
type ExportResultTrap struct {
	Traps []ExportTrap `json:"result" yaml:"result"`
}

//ExportTrapMatching represents the caracteristics of a matching
type ExportTrapMatching struct {
	String string `json:"string"`
	Regexp string `json:"regexp"`
	Status string `json:"status"`
	Order  string `json:"order"`
}

//ExportResultTrapMatching represents a matching array send by the API
type ExportResultTrapMatching struct {
	Matchings []ExportTrapMatching `json:"result" yaml:"result"`
}
