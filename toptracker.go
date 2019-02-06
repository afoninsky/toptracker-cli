package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const topTrackerHost = "https://tracker-api.toptal.com"
const timeout = 5 * time.Second

func toJSON(body io.ReadCloser, output interface{}) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, &output)
}

func get(url string, jsonRes interface{}) error {
	client := &http.Client{Timeout: timeout}
	host := fmt.Sprintf("%s%s", topTrackerHost, url)
	res, err := client.Get(host)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return toJSON(res.Body, jsonRes)
}

func post(url string, payload map[string]string, jsonRes interface{}) error {
	client := &http.Client{Timeout: timeout}
	jsonPayload, _ := json.Marshal(payload)

	host := fmt.Sprintf("%s%s", topTrackerHost, url)
	res, err := client.Post(host, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return toJSON(res.Body, jsonRes)
}

func fetchSeconds(user userInfo, projects projectsInfo, from, till string) (int, error) {

	// generate query url
	query := url.Values{}
	// add own projects
	for _, project := range projects.Projects {
		query.Add("project_ids[]", strconv.Itoa(project.ID))
	}

	// add own ids in each project
	for _, profile := range user.Profiles {
		query.Add("worker_ids[]", strconv.Itoa(profile.ID))
	}
	query.Add("start_date", from)
	query.Add("end_date", till)
	query.Add("type", "workers")
	query.Add("access_token", user.Token)

	// collect all found seconds
	requestURL := fmt.Sprintf("/reports/work_summary?%s", query.Encode())
	var data workSummaryInfo
	if err := get(requestURL, &data); err != nil {
		return 0, err
	}

	// calculate total amount of seconds
	secondsTotal := 0
	for _, projects := range data.WorkSummary.WorkersProjects {
		for _, project := range projects.Projects {
			secondsTotal = secondsTotal + project.Seconds
		}
	}
	return secondsTotal, nil
}

func getTopTrackerInfo(email, password string) (topTrackerInfo, error) {
	info := topTrackerInfo{}
	creds := map[string]string{"email": email, "password": password}
	var user userInfo

	// get token and user id
	if err := post("/sessions", creds, &user); err != nil {
		return info, err
	}
	if user.Token == "" || user.User.ID == 0 {
		return info, errors.New("Login failed: bad credentials?")
	}

	// get available project ids
	var projects projectsInfo
	if err := get(fmt.Sprintf("/projects?access_token=%s", user.Token), &projects); err != nil {
		return info, err
	}

	now := time.Now()
	firstday := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	secondsPerMonth, err := fetchSeconds(user, projects, firstday.Format("2006-01-02"), now.Format("2006-01-02"))
	if err != nil {
		return info, err
	}
	secondsPerDay, err := fetchSeconds(user, projects, now.Format("2006-01-02"), now.Format("2006-01-02"))
	if err != nil {
		return info, err
	}

	info.SecondsPerMonth = secondsPerMonth
	info.SecondsPerDay = secondsPerDay
	return info, nil
}
