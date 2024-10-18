package checkers

import (
	"fmt"
	"io"
	"net/http"
)

func GetAnnotations(grafanaKey string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://deliveryhero.grafana.net/api/annotations", nil)
	fmt.Printf("grafana key: %s", grafanaKey)

	req.Header.Set("Authorization", "Bearer "+grafanaKey)
	res, _ := client.Do(req)
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("error retrieving annotations. %v", err)
		}
		bodyString := string(bodyBytes)
		fmt.Printf("annotations: %s", bodyString)
	}
}
