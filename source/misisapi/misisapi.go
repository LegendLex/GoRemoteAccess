package misisapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Returns the group ID parsed from the local JSON and confirmation of existence by name
func GetGroups(groupName string) (bool, int) {
	var (
		groups           []misisGroup
		groupID          int
		groupIsConfirmed = false
	)

	data, err := os.ReadFile("misisapi/groups.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &groups)
	for _, group := range groups {
		if group.Name == groupName {
			groupID = group.ID
			groupIsConfirmed = true
			break
		}
	}
	return groupIsConfirmed, groupID
}

// Returns a newSchedule structure containing a weekly schedule that includes the day specified as a parameter.
func GetSchedule(groupID string, day string) misisSchedule {

	url := "https://lk.misis.ru/method/schedule.get"
	payload := strings.NewReader("{\"filial\":880,\"group\":\"" + groupID + "\",\"room\":null,\"teacher\":null,\"start_date\":\"" + day + "\",\"end_date\":null}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "ru,en;q=0.9,cy;q=0.8")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("Referer", "https://edu.misis.ru/")
	req.Header.Add("Referrer-Policy", "strict-origin-when-cross-origin")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var newSchedule misisSchedule
	json.Unmarshal(body, &newSchedule)
	return newSchedule
}
