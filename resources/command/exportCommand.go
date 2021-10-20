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

package command

//ExportCommand represents the caracteristics of a command
type ExportCommand struct {
	Name        string `json:"name" yaml:"name"` //command Name
	Type        string `json:"type" yaml:"type"` //command type
	Line        string `json:"line" yaml:"line"`
	Graph       string `json:"graph" yaml:"graph"`
	Example     string `json:"example" yaml:"example"`
	Comment     string `json:"comment" yaml:"comment"`
	Activate    string `json:"activate" yaml:"activate"`
	EnableShell string `json:"enable_shell" yaml:"enable_shell"`
}

//ExportResult represents a command array send by the API
type ExportResult struct {
	Commands []ExportCommand `json:"result" yaml:"result"`
}
