package LineCounter

import (
	"CodeLineCounter/Counter"
	"CodeLineCounter/DTO"
	"errors"
	"os"
	"path/filepath"
)

// analyzeFolder processes each file in the folder and aggregates results
func analyzeFolder(folderPath string) (DTO.Summary, error) {
	var summary DTO.Summary
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(path)
			language, supported := fileExtensions[ext]
			if !supported {
				err := errors.New("file not supported")
				return err // Skip unsupported file types
			}

			counter, err := Counter.NewCounter(language)
			if err != nil {
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			fileResult, err := counter.CountLines(file)
			if err != nil {
				return err
			}
			summary.Files = append(summary.Files, fileResult)

			// Aggregate results into the overall summary
			summary.Overall.TotalLines += fileResult.TotalLines
			summary.Overall.BlankLines += fileResult.BlankLines
			summary.Overall.CommentLines += fileResult.CommentLines
			summary.Overall.CodeLines += fileResult.CodeLines
			summary.Overall.ImportLines += fileResult.ImportLines
			summary.Overall.VariableLines += fileResult.VariableLines
		}
		return nil
	})
	return summary, err
}
