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

package downtime

import (
	"centctl/request"
	"centctl/resources/contact"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type timezoneHost struct {
	Timezone string `json:"timezone"`
}

func getTimezoneHost(hostId int, debugV bool) string {
	//Recovery the  output of a service
	urlCentreon := "/monitoring/hosts/" + strconv.Itoa(hostId)
	_, body := request.GeneriqueCommandV2Get(urlCentreon, "downtime get timezone", debugV)

	//Permits to recover the array result
	var timezone timezoneHost
	json.Unmarshal(body, &timezone)

	return timezone.Timezone
}

//VerifyStartDayAndHour permits to verify the syntax of the day and hour start
func VerifyStartDayAndHour(startDay string, startHour string) error {
	matched, err := regexp.MatchString(`^(\d{4}[-](0[1-9]|1[012])[-](0[1-9]|[12][0-9]|3[01]))$`, startDay)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("The start day must be on the form : YYYY-MM-DD")
	}

	matched, err = regexp.MatchString(`^((0[0-9]|1[0-9]|2[0-3])[:]([0-5][0-9]|[6][0]))$`, startHour)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("The start hour must be on the form : HH-MM")
	}
	return nil
}

//GetAuthorId permits to find the ID of the person login
func GetAuthorId(debugV bool) (int, error) {
	login := os.Getenv("LOGIN")
	err, body := request.GeneriqueCommandV1Post("show", "contact", login, "get author_id", debugV, false, "")
	if err != nil {
		return -1, err
	}

	//Permits to recover the contacts contain into the response body
	contacts := contact.DetailResult{}
	json.Unmarshal(body, &contacts)

	//Permits to find the good contact in the array
	var ContactFind contact.DetailContact
	for _, v := range contacts.DetailContacts {
		if strings.ToLower(v.Alias) == strings.ToLower(login) {
			ContactFind = v
		}
	}
	if ContactFind.Alias != "" {
		id, _ := strconv.Atoi(ContactFind.ID)
		return id, nil
	}

	return -1, fmt.Errorf("Contact " + login + " not find")
}

//GetEndDowntime permits to get start and end downtime in time type
func GetEndDowntime(startDay string, startHour string, duration int, timezone string) (time.Time, time.Time) {
	daySplit := strings.Split(startDay, "-")
	hourSplit := strings.Split(startHour, ":")
	var year, _ = strconv.Atoi(daySplit[0])
	var month, _ = strconv.Atoi(daySplit[1])
	var day, _ = strconv.Atoi(daySplit[2])
	var hour, _ = strconv.Atoi(hourSplit[0])
	var minute, _ = strconv.Atoi(hourSplit[1])

	durationFinal, _ := time.ParseDuration(strconv.Itoa(duration) + "s")

	var timeStart time.Time
	if timezone != "" {
		timezoneFinal, err := time.LoadLocation(timezone)
		if err != nil {
			timezoneFinal = time.Local
		}
		timeStart = time.Date(year, time.Month(month), day, hour, minute, 0, 0, timezoneFinal)
	} else {
		timeStart = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
	}
	timeEnd := timeStart.Add(durationFinal)

	timeStart.Format(time.RFC3339)
	timeEnd.Format(time.RFC3339)

	return timeStart, timeEnd
}
