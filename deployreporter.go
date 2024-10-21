package deployreporter

import (
	"fmt"
	"regexp"

	"github.com/sebasrp/deployreporter/internal"
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
	Tribe       string
	Squad       string
	Tier        string
	Source      string
}

func NewDeployment(annotation checkers.Annotation) Deployment {
	tags := checkers.GenerateMapFromTags(annotation.Tags)
	service := tags["service"]
	emailRegex := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
	email := emailRegex.FindString(annotation.Text)
	tribe, squad, tier, err := internal.GetOrgFromServiceName(service)
	if err != nil {
		fmt.Printf("Error retrieving org information: %v", err)
	}

	c := Deployment{
		ID:          annotation.ID,
		Start:       annotation.Created,
		End:         annotation.TimeEnd,
		Operator:    email,
		Service:     service,
		Environment: tags["dh_env"],
		Country:     tags["location"],
		Tribe:       tribe,
		Squad:       squad,
		Tier:        tier,
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
