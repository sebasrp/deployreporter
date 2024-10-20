package deployreporter

import (
	"regexp"

	"github.com/sebasrp/deployreporter/internal/checkers"
)

type Deployment struct {
	ID          int
	Start       int64
	End         int64
	Operator    string
	Service     string
	Environment string
	Country     string
	Source      string
}

func NewDeployment(annotation checkers.Annotation) Deployment {
	tags := checkers.GenerateMapFromTags(annotation.Tags)
	emailRegex := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
	email := emailRegex.FindString(annotation.Text)

	c := Deployment{
		ID:          annotation.ID,
		Start:       annotation.Created,
		End:         annotation.TimeEnd,
		Operator:    email,
		Service:     tags["service"],
		Environment: tags["dh_env"],
		Country:     tags["location"],
		Source:      tags["tool"],
	}
	return c
}

func GetDeployments(from string, to string, limit int, grafanaKey string) (ret []Deployment) {
	var deployments []Deployment
	annotations := checkers.GetDeploymentAnnotations(from, to, limit, grafanaKey)
	for _, a := range annotations {
		deployments = append(deployments, NewDeployment(a))
	}
	return deployments
}
