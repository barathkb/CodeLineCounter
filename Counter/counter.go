package Counter

import (
	"CodeLineCounter/DTO"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Counter interface that defines methods for counting lines
type Counter interface {
	CountLines(file *os.File) (DTO.Result, error)
}

// Factory function to create the appropriate counter based on file extension
func NewCounter(language string) (Counter, error) {
	switch language {
	case "go":
		return &GoCounter{}, nil
	case "c":
		return &CCounter{}, nil
	case "python":
		return &PythonCounter{}, nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
}

func countLines(file *os.File, patterns map[string]*regexp.Regexp) (DTO.Result, error) {
	var res DTO.Result
	res.FileName = filepath.Base(file.Name())
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		res.TotalLines++
		lineType := classifyLine(line, patterns)

		switch lineType {
		case Blank:
			res.BlankLines++
		case Comment:
			res.CommentLines++
		case Import:
			res.ImportLines++
		case Variable:
			res.VariableLines++
		case Code:
			res.CodeLines++
		}
	}

	return res, scanner.Err()
}
