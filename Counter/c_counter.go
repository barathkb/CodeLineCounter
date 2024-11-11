package Counter

import (
	"CodeLineCounter/DTO"
	"os"
	"regexp"
)

// CCounter implementation for C language
type CCounter struct{}

var cPattern = map[string]*regexp.Regexp{
	"comment":  regexp.MustCompile(`^\s*//|^\s*/\*|^\s*\*/|^\s*\*`),
	"import":   regexp.MustCompile(`^\s*#include\s`),
	"variable": regexp.MustCompile(`^\s*\w+\s+\w+\s*(=|;)`),
}

func (c *CCounter) CountLines(file *os.File) (DTO.Result, error) {
	return countLines(file, cPattern)
}
