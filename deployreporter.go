package deployreporter

import (
	"github.com/sebasrp/deployreporter/internal/checkers"
)

type Deployment struct {
	ID          int
	Start       int64 // epoch
	End         int64
	Service     string
	Environment string
	Country     string
	Tool        string
}

func NewDeployment(annotation checkers.Annotation) Deployment {
	tags := checkers.GenerateMapFromTags(annotation.Tags)
	c := Deployment{
		ID:          annotation.ID,
		Start:       annotation.Created,
		End:         annotation.TimeEnd,
		Service:     tags["service"],
		Environment: tags["dh_env"],
		Country:     tags["location"],
		Tool:        tags["tool"],
	}
	return c
}

func GetDeployments(grafanaKey string) (ret []Deployment) {
	var deployments []Deployment
	annotations := checkers.GetDeploymentAnnotations(grafanaKey)
	for _, a := range annotations {
		deployments = append(deployments, NewDeployment(a))
	}
	return deployments
}
