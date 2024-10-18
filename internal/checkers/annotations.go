package checkers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Annotation struct {
	ID           int      `json:"id"`
	AlertID      int      `json:"alertId"`
	AlertName    string   `json:"alertName"`
	DashboardID  int      `json:"dashboardId"`
	DashboardUID any      `json:"dashboardUID"`
	PanelID      int      `json:"panelId"`
	UserID       int      `json:"userId"`
	NewState     string   `json:"newState"`
	PrevState    string   `json:"prevState"`
	Created      int64    `json:"created"`
	Updated      int64    `json:"updated"`
	Time         int64    `json:"time"`
	TimeEnd      int64    `json:"timeEnd"`
	Text         string   `json:"text"`
	Tags         []string `json:"tags"`
	Login        string   `json:"login"`
	Email        string   `json:"email"`
	AvatarURL    string   `json:"avatarUrl"`
	Data         struct {
	} `json:"data"`
}

func GetAnnotations(grafanaKey string) {
	// see https://grafana.com/docs/grafana/latest/developers/http_api/annotations/

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://deliveryhero.grafana.net/api/annotations", nil)
	fmt.Printf("grafana key: %s", grafanaKey)

	req.Header.Set("Authorization", "Bearer "+grafanaKey)
	res, _ := client.Do(req)
	if res.StatusCode == http.StatusOK {
		var annotations []Annotation
		err := json.NewDecoder(res.Body).Decode(&annotations)
		if err != nil {
			fmt.Printf("error retrieving annotations. %v", err)
		}
		for _, a := range annotations {
			fmt.Printf("%+v\n", a)
		}
	}
}
