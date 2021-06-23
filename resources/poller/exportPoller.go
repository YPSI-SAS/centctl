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

package poller

//ExportPoller represents the caracteristics of a Poller
type ExportPoller struct {
	Name             string `json:"name" yaml:"name"`
	Localhost        string `json:"localhost" yaml:"localhost"`
	IPAddress        string `json:"ip address" yaml:"ip address"`
	Activate         string `json:"activate" yaml:"activate"`
	Status           string `json:"status" yaml:"status"`
	EngineRestartCmd string `json:"engine restart command" yaml:"engine restart command"`
	EngineReloadCmd  string `json:"engine reload command" yaml:"engine reload command"`
	BorkerReloadCmd  string `json:"broker reload command" yaml:"broker reload command"`
	Bin              string `json:"bin" yaml:"bin"`
	StatsBin         string `json:"stats bin" yaml:"stats bin"`
	SSHPort          string `json:"ssh port" yaml:"ssh port"`
	GorgonePorotocol string `json:"gorgone protocol" yaml:"gorgone protocol"`
	GorgonePort      string `json:"gorgone port" yaml:"gorgone port"`
}

//ExportResultPoller represents a Poller array send by the API
type ExportResultPoller struct {
	Pollers []ExportPoller `json:"result" yaml:"result"`
}
