package operations

import (
	"strings"
)

type Operation struct {
	Time     string
	Units    string
	Report   string
	District string
	Location string
}

var extraUnits = []string{"ILS Lagedienst", "Kreisbrandinspektion", "THW", "UG-ÖEL"}

func ProcessUnits(units string) string {
	return strings.Join(parseUnits(units), ", ")
}

func parseUnits(source string) []string {
	var matchedUnits []string
	var matchedExtraUnits []string

	// Extract extraUnits from unit string
	for _, match := range extraUnits {
		if strings.Contains(source, match) {
			source = strings.Replace(source, match, "", -1)
			matchedExtraUnits = append(matchedExtraUnits, match)
		}
	}

	t := strings.Split(source, "FF ")
	for cnt, unit := range t {
		if cnt != 0 {
			tmp := "FF " + unit
			matchedUnits = append(matchedUnits, tmp)
		}
	}

	matchedUnits = append(matchedUnits, matchedExtraUnits...)

	return matchedUnits
}
