package LineCounter

import (
	"CodeLineCounter/DTO"
)

// LineCounterResponse response struct
type LineCounterResponse struct {
	Body    DTO.Summary `json:"body,omitempty"`
	Message string      `json:"message,omitempty"`
}
