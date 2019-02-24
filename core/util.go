package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ptflp/adrenalinBank/models"
	"net/http"
	"sort"
	"time"
)

func Operations(oc models.User) []models.Operation {
	client := &http.Client{}
	data := []byte(`{"productId": ` + oc.ProductId + `}`)
	ff := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "https://diybank.aeb-it.ru/api/operation", ff)
	authType := "Bearer"
	authToken := oc.Token
	req.Header.Add("Authorization", authType+" "+authToken)
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
			fmt.Println(err)
			fmt.Println(index)
		}
		operations[index].Timestamp = t.Unix()
	}

	return operations
}

func SortOperationsAsc(o []models.Operation) (operations []models.Operation) {
	sort.Slice(operations, func(i, j int) bool {
		return operations[i].Timestamp < operations[j].Timestamp
	})
	return operations
}
