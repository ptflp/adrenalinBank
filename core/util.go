package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

func chetoTam () {

	client := &http.Client{}
	data := []byte(`{"productId":"2268927"}`)
	ff := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "https://diybank.aeb-it.ru/api/operation", ff)
	authType := "Bearer"
	authToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBYUk1QRzJIU0NNMElRQlY1Mlk0WEVIQjhDSlVWRE5ZTExNRUIwSlIiLCJ0eXAiOiJKV1QifQ.eyJzdWIiOiIzIiwibmFtZSI6ImxvZ2luMSIsInRva2VuX3VzYWdlIjoiYWNjZXNzX3Rva2VuIiwianRpIjoiMWEyYjE0YjMtOTViMC00MjQwLTg4OTYtMGE2NGE4OWRkZjVjIiwiYXVkIjoiQWViSXQuRElZQmFuayIsIm5iZiI6MTU1MDg5OTc5NiwiZXhwIjoxNTUxNTA0NTk2LCJpYXQiOjE1NTA4OTk3OTYsImlzcyI6Imh0dHA6Ly9kaXliYW5rLmFlYi1pdC5ydS8ifQ.nzPux_UlisWbxxtalV257Y35zLAOnheHokQxJprg0dUkvQzLQTzrGUsjAk3BpoNuwJuL6aJcD5e3rq9Zbi8Khn2o_ZYjrz2jyWehJJ3tPo7IOzWdVhCp5mGz6ducPjRRUTIcGcv3jCeP3T8KTJAv_W4I1xUkS3PcJ1DXWVrbJwRWbEsMpnZS6MiSOluWZ-AV_SBDX0dXZr5y0OnexHNxg9DnHTczHug5xaUlJNQKp_H_S0usGbroHc8bwnSb0NnS_K03WaisSILPw8fUDBqDr8S60tmZodRGqmR6OarbswpgjGmUwmYWP5BsELJFDQI-uWEy0Mg8pHwarcPRLIJJrg"
	req.Header.Add("Authorization", authType + " " + authToken)
	req.Header.Add("accept", `application/json`)
	req.Header.Add("Content-Type", `application/json-patch+json`)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("response Status:", err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	bytes := []byte(newStr)
	var operations []models.Operation
	json.Unmarshal(bytes, &operations)
	for index, element := range operations {
		layout := "2006-01-02T15:04:05"
		t, err := time.Parse(layout, element.Date)
		if err != nil {
			fmt.Println(err )
			fmt.Println(index)
		}
		operations[index].Timestamp = t.Unix()
	}
	sort.Slice(operations, func(i, j int) bool {
		return operations[i].Timestamp < operations[j].Timestamp
	})
}