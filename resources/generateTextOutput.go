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

package resources

import (
	"strconv"

	"github.com/pterm/pterm"
)

//Permits with pterm library to display a table list with custom aspect
func TableListWithHeader(table [][]string) string {
	values, _ := pterm.DefaultTable.WithHasHeader().WithData(table).WithHeaderStyle(pterm.NewStyle(pterm.FgMagenta, pterm.Bold, pterm.Underscore)).Srender()
	return values
}

//Permits with pterm library to display a table list with custom aspect
func TableList(table [][]string) string {
	values, _ := pterm.DefaultTable.WithData(table).Srender()
	return values
}

//Permits to generate a BulletListItem from list of elements
func GenerateListItems(elements [][]string, bulletStyle string) []pterm.BulletListItem {
	if bulletStyle == "" {
		bulletStyle = " "
	}
	items := []pterm.BulletListItem{}
	for _, elem := range elements {
		level, _ := strconv.Atoi(elem[0])
		items = append(items, pterm.BulletListItem{Level: level, Text: elem[1], Bullet: bulletStyle})
	}
	return items
}

//Permits with pterm library to display a bullet list with items
func BulletList(items []pterm.BulletListItem) string {
	values, _ := pterm.DefaultBulletList.WithItems(items).Srender()
	return values
}
