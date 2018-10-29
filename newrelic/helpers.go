package newrelic

import (
	"fmt"
	"strconv"
	"strings"
)

func parseIDs(serializedID string, count int) ([]int, error) {
	rawIDs := strings.SplitN(serializedID, ":", count)
	if len(rawIDs) != count {
		return []int{}, fmt.Errorf("Unable to parse ID %v", serializedID)
	}

	ids := make([]int, count)

	for i, rawID := range rawIDs {
		id, err := strconv.ParseInt(rawID, 10, 32)
		if err != nil {
			return ids, err
		}

		ids[i] = int(id)
	}

	return ids, nil
}

func serializeIDs(ids []int) string {
	idStrings := make([]string, len(ids))

	for i, id := range ids {
		idStrings[i] = strconv.Itoa(id)
	}

	return strings.Join(idStrings, ":")
}

func sortedTerms(terms []map[string]interface{}) []map[string]interface{} {
	returnTerms := make([]map[string]interface{}, 0)

	var warningCondition, criticalCondition map[string]interface{}

	for _, src := range terms {
		if src["priority"] == "warning" {
			warningCondition = src
		} else if src["priority"] == "critical" {
			criticalCondition = src
		}
	}

	if warningCondition["priority"] == "warning" {
		returnTerms = append(returnTerms, warningCondition)
	}
	if criticalCondition["priority"] == "critical" {
		returnTerms = append(returnTerms, criticalCondition)
	}

	return returnTerms
}
