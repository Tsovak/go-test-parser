package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/tsovak/go-test-parser/models"
)

var githubTestResultRegexp = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{7}Z ")

func ParseLine(input string) (*models.GoTestEvent, error) {
	// e.g. "go: downloading github.com/....."
	if strings.HasPrefix(input, "go:") {
		return nil, nil
	}

	input = githubTestResultRegexp.ReplaceAllString(input, "")

	var output models.GoTestEvent

	if err := json.Unmarshal([]byte(input), &output); err != nil {
		return nil, fmt.Errorf("error deserializing: %q", input)
	}

	return &output, nil
}
