package Counter

import (
	"regexp"
	"strings"
)

// Common classifyLine function to be used by all counters
func classifyLine(line string, patterns map[string]*regexp.Regexp) LineType {
	line = strings.TrimSpace(line)
	if line == "" {
		return Blank
	}
	if patterns["comment"].MatchString(line) {
		return Comment
	}
	if patterns["import"].MatchString(line) {
		return Import
	}
	if patterns["variable"].MatchString(line) {
		return Variable
	}
	
	return Code
}
