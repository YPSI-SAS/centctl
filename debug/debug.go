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
	if statusCode != 200 {
		logger.Error("centctl " + nameCmd + " - statusCode: " + strconv.Itoa(statusCode))
		logger.Error("centctl " + nameCmd + " - Body response: " + string(body))
	} else {
		logger.Info("centctl " + nameCmd + " - statusCode: " + strconv.Itoa(statusCode))
		logger.Info("centctl " + nameCmd + " - Body response: " + string(body))
	}
	fmt.Println()
}
