package checkers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
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

func GetDeploymentAnnotations(from string, to string, limit int, grafanaKey string) (ret []Annotation) {
	// see https://grafana.com/docs/grafana/latest/developers/http_api/annotations/
	var annotations []Annotation

	client := &http.Client{}
	requestURL := fmt.Sprintf("http://deliveryhero.grafana.net/api/annotations?from=%s&to=%s&limit=%d", from, to, limit)
	req, _ := http.NewRequest(http.MethodGet, requestURL, nil)

	req.Header.Set("Authorization", "Bearer "+grafanaKey)
	res, _ := client.Do(req)
	if res.StatusCode == http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&annotations)
		if err != nil {
			fmt.Printf("error retrieving annotations. %v", err)
		}
	}
	var mortyAnnotations []Annotation
	for _, annotation := range annotations {
		if slices.Contains(annotation.Tags, "tool:morty") {
			mortyAnnotations = append(mortyAnnotations, annotation)
		}
	}
	return mortyAnnotations
}

func GenerateMapFromTags(tags []string) (ret map[string]string) {
	result := make(map[string]string)

	for _, item := range tags {
		key, value, err := extractTagInfo(item)
		if err != nil {
			fmt.Printf("error extracting tag: %v", err)
		}
		result[key] = value
	}
	return result
}

func extractTagInfo(tag string) (string, string, error) {
	parts := strings.SplitN(tag, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("input string must contain exactly one semicolon. tag: %v", tag)
	}
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	return key, value, nil
}
