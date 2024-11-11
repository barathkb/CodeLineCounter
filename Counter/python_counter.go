package Counter

import (
	"CodeLineCounter/DTO"
	"os"
	"regexp"
)

// PythonCounter implementation for Python language
type PythonCounter struct{}

var pythonPattern = map[string]*regexp.Regexp{
	"comment":  regexp.MustCompile(`^\s*#`),
	"import":   regexp.MustCompile(`^\s*import\s|^\s*from\s`),
	"variable": regexp.MustCompile(`^\s*\w+\s*=`),
}

func (c *PythonCounter) CountLines(file *os.File) (DTO.Result, error) {
	return countLines(file, pythonPattern)
}
