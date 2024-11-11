package DTO

// Result struct holds the line count results for an individual file
type Result struct {
	FileName      string `json:"file_name"`
	TotalLines    int    `json:"total_lines"`
	BlankLines    int    `json:"blank_lines"`
	CommentLines  int    `json:"comment_lines"`
	CodeLines     int    `json:"code_lines"`
	ImportLines   int    `json:"import_lines"`
	VariableLines int    `json:"variable_lines"`
}

// Summary struct holds both the overall summary and individual file results
type Summary struct {
	Overall Result   `json:"overall"`
	Files   []Result `json:"files"`
}
