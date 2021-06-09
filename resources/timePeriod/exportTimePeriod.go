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

package timePeriod

//ExportTimePeriod represents the caracteristics of a timePeriod
type ExportTimePeriod struct {
	Name      string `json:"name" yaml:"name"`
	Alias     string `json:"alias" yaml:"alias"`
	Sunday    string `json:"sunday" yaml:"sunday"`
	Monday    string `json:"monday" yaml:"monday"`
	Tuesday   string `json:"tuesday" yaml:"tuesday"`
	Wednesday string `json:"wednesday" yaml:"wednesday"`
	Thursday  string `json:"thursday" yaml:"thursday"`
	Friday    string `json:"friday" yaml:"friday"`
	Saturday  string `json:"saturday" yaml:"saturday"`

	Exceptions []ExportTimePeriodException
}

//ExportResultTimePeriod represents a timePeriod array send by the API
type ExportResultTimePeriod struct {
	TimePeriods []ExportTimePeriod `json:"result" yaml:"result"`
}

//ExportTimePeriodException represents the caracteristics of a exception
type ExportTimePeriodException struct {
	Days      string `json:"days"`
	Timerange string `json:"timerange"`
}

//ExportResultTimePeriodExecption represents a exception array send by the API
type ExportResultTimePeriodExecption struct {
	Exceptions []ExportTimePeriodException `json:"result" yaml:"result"`
}
