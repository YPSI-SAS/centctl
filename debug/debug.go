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

package debug

import (
	"fmt"

	"os"
	"strconv"

	"github.com/withmandala/go-log"
)

//Show permits to show the debug informations
func Show(nameCmd string, requestBody string, urlCentreon string, statusCode int, body []byte) {
	logger := log.New(os.Stdout).WithColor()
	logger.Info("centctl authentification - successful authentification, token generate ")
	if requestBody != "" {
		logger.Info("centctl " + nameCmd + " - Body request: " + requestBody)
	}
	logger.Info("centctl " + nameCmd + " - URL: " + urlCentreon)
	if statusCode != 200 && statusCode != 201 && statusCode != 204 {
		logger.Error("centctl " + nameCmd + " - statusCode: " + strconv.Itoa(statusCode))
		logger.Error("centctl " + nameCmd + " - Body response: " + string(body))
	} else {
		logger.Info("centctl " + nameCmd + " - statusCode: " + strconv.Itoa(statusCode))
		logger.Info("centctl " + nameCmd + " - Body response: " + string(body))
	}
	fmt.Println()
}
