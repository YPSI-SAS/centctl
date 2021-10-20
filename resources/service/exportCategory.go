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

//ExportCategory represents the caracteristics of a category service
type ExportCategory struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`

	Services         []ExportCategoryService
	ServiceTemplates []ExportCategoryServiceTemplate
}

//ExportResultCategory represents a category service array send by the API
type ExportResultCategory struct {
	CategoryServices []ExportCategory `json:"result" yaml:"result"`
}

//ExportCategoryService represents the caracteristics of a service member
type ExportCategoryService struct {
	HostID             string `json:"host id"`
	HostName           string `json:"host name"`
	ServiceID          string `json:"service id"`
	ServiceDescription string `json:"service description"`
}

//ExportResultCategoryService represents a service member array send by the API
type ExportResultCategoryService struct {
	CategoryService []ExportCategoryService `json:"result" yaml:"result"`
}

//ExportCategoryServiceTemplate represents the caracteristics of a service template member
type ExportCategoryServiceTemplate struct {
	TemplateID                 string `json:"template id"`
	ServiceTemplateDescription string `json:"service template description"`
}

//ExportResultCategoryServiceTemplate represents a service template member array send by the API
type ExportResultCategoryServiceTemplate struct {
	CategoryServiceTemplate []ExportCategoryServiceTemplate `json:"result" yaml:"result"`
}
