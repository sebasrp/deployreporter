package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Mapping struct {
	Service string
	Tribe   string
	Squad   string
	Tier    string
}

var _orgMappings = make(map[string]Mapping)

func GetOrgFromServiceName(serviceName string) (tribe string, squad string, tier string, err error) {
	if len(_orgMappings) == 0 {
		_orgMappings = GetOrgMappings("dh.csv") // TODO: we should really extract this
	}

	scvMapping, exists := _orgMappings[serviceName]
	if !exists {
		return "", "", "", fmt.Errorf("service mapping does not exist for service %v", serviceName)
	}

	return scvMapping.Tribe, scvMapping.Squad, scvMapping.Tier, nil
}

func GetOrgMappings(filename string) (svcMappings map[string]Mapping) {
	result := make(map[string]Mapping)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error while reading the file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	mappings, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}
	for _, entry := range mappings {
		svcName := entry[0]
		svcTribe := entry[1]
		svcSquad := entry[2]
		svcTier := entry[3]
		result[entry[0]] = Mapping{Service: svcName, Tribe: svcTribe, Squad: svcSquad, Tier: svcTier}
	}
	return result
}
