/*MIT License

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
package request

import (
	"centctl/colorMessage"
	"centctl/debug"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//List of available map in the Centreon MAP
type availableMaps struct {
	ID     int    `json:"id"`
	Label  string `json:"label"`
	ViewId int    `json:"viewId"`
}

//ConnectionMAP permits the connection to the centreon MAP for get the studio-token and the server's version
func ConnectionMAP(debugV bool) {
	colorRed := colorMessage.GetColorRed()
	if os.Getenv("URLMAP") == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You have not entered a MAP URL in a config file")
		os.Exit(1)
	}

	requestBody, _ := json.Marshal(map[string]string{
		"login":    os.Getenv("LOGIN"),
		"password": os.Getenv("PASSWORD"),
	})
	urlCentreon := os.Getenv("URLMAP") + "/api/beta/authentication"
	client := NewClientMAP(urlCentreon)
	err := client.GetSessionToken(requestBody, debugV)
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	urlCentreon = os.Getenv("URLMAP") + "/api/beta/authentication/serverVersion"
	client = NewClientMAP(urlCentreon)
	err = client.GetServerVersion(debugV)
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

//MapsExist permits to check if the map already exists
func MapsExist(labelMap string, debugV bool, debugNameCmd string) (int, int) {
	colorRed := colorMessage.GetColorRed()

	urlCentreon := os.Getenv("URLMAP") + "/api/beta/maps"
	client := NewClientMAP(urlCentreon)
	statusCode, body, err := client.Get()
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if debugV {
		debug.Show(debugNameCmd, "", urlCentreon, statusCode, body)
	}

	//Permits to recover the maps contain into the response body
	maps := []availableMaps{}
	json.Unmarshal(body, &maps)

	//Check the label of the map and return his id and his viewID if exists
	for _, elem := range maps {
		if strings.ToLower(elem.Label) == strings.ToLower(labelMap) {
			return elem.ViewId, elem.ID
		}
	}

	return -1, -1
}

//CreateMaps permits to create a Map in the Centreon MAP
func CreateMaps(labelMap string, debugV bool, debugNameCmd string) int {
	colorRed := colorMessage.GetColorRed()

	requestBody, _ := json.Marshal(map[string]string{
		"label": labelMap,
	})

	urlCentreon := os.Getenv("URLMAP") + "/api/beta/maps"
	client := NewClientMAP(urlCentreon)
	statusCode, body, err := client.Post(requestBody)
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if debugV {
		debug.Show(debugNameCmd, "", urlCentreon, statusCode, body)
	}

	//Permits to recover the maps contain into the response body
	mapCreated := availableMaps{}
	json.Unmarshal(body, &mapCreated)

	return mapCreated.ViewId
}

//DeleteMap permits to delete one map
func DeleteMap(mapID int, debugV bool, debugNameCmd string) {
	colorRed := colorMessage.GetColorRed()

	urlCentreon := os.Getenv("URLMAP") + "/api/beta/maps/" + strconv.Itoa(mapID)
	client := NewClientMAP(urlCentreon)
	statusCode, body, err := client.Delete()
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if debugV {
		debug.Show(debugNameCmd, "", urlCentreon, statusCode, body)
	}

}

//PutResource permits to put a resource in a view
func PutResource(viewID int, elementID int, debugV bool, debugNameCmd string) {
	colorRed := colorMessage.GetColorRed()
	urlCentreon := os.Getenv("URLMAP") + "/api/beta/views/" + strconv.Itoa(viewID) + "/elements/" + strconv.Itoa(elementID)
	client := NewClientMAP(urlCentreon)
	statusCode, body, err := client.Put()
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if debugV {
		debug.Show(debugNameCmd, "", urlCentreon, statusCode, body)
	}
}

//PutLink permits to put a link in a view
func PutLink(viewID int, linkID int, debugV bool, debugNameCmd string) {
	colorRed := colorMessage.GetColorRed()
	urlCentreon := os.Getenv("URLMAP") + "/api/beta/views/" + strconv.Itoa(viewID) + "/links/" + strconv.Itoa(linkID)
	client := NewClientMAP(urlCentreon)
	statusCode, body, err := client.Put()
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if debugV {
		debug.Show(debugNameCmd, "", urlCentreon, statusCode, body)
	}
}
