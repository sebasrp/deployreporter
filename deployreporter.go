package deployreporter

import (
	"fmt"

	"github.com/sebasrp/deployreporter/internal/checkers"
)

func GetDeployments(grafanaKey string) {
	fmt.Printf("[get deployments]grafana key: %s", grafanaKey)
	checkers.GetAnnotations(grafanaKey)
}
