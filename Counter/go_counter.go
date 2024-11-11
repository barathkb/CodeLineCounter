package Counter

import (
	"CodeLineCounter/DTO"
	"os"
	"regexp"
)

// GoCounter implementation for Go language
type GoCounter struct{}

var goPattern = map[string]*regexp.Regexp{
	"comment":  regexp.MustCompile(`^\s*//|^\s*/\*|^\s*\*/|^\s*\*`),
	"import":   regexp.MustCompile(`^\s*import\s`),
	"variable": regexp.MustCompile(`^\s*(var|const)\s`),
}

func (c *GoCounter) CountLines(file *os.File) (DTO.Result, error) {
	return countLines(file, goPattern)
}
