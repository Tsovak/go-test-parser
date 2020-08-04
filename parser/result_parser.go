package parser

import (
	"log"

	"github.com/tsovak/go-test-parser/models"
)

type TestResultCompleted func(result models.TestResult)

type TestResultParser struct {
	testCompleted TestResultCompleted
	pending       map[string]models.TestResult
}

func NewResultsParser(completed TestResultCompleted) TestResultParser {
	return TestResultParser{
		testCompleted: completed,
		pending:       map[string]models.TestResult{},
	}
}

func (trp *TestResultParser) ParseLine(input string, verbose bool) {
	parsed, err := ParseLine(input)
	if err != nil {
		if verbose {
			log.Printf("[ERROR] %+v", err)
		}
		return
	}

	// e.g. go mod || a package
	if parsed == nil || parsed.Test == "" {
		return
	}

	// do we already have a pending item for this line?
	cacheKey := parsed.CacheKey()
	existing, ok := trp.pending[cacheKey]
	if !ok {
		existing = models.TestResult{
			Package:  parsed.Package,
			TestName: parsed.Test,
			Result:   models.Pending,
			Duration: 0,
			StdOut:   "",
			StdErr:   "",
		}
		trp.pending[cacheKey] = existing
	}

	// nothing of value to parse out at this time
	if parsed.Action == models.Cont ||
		parsed.Action == models.Pause ||
		parsed.Action == models.Run {
		return
	}

	if parsed.Action == models.Output {
		existing.StdOut += parsed.Output
	}

	if parsed.Action == models.Pass {
		existing.Result = models.Successful
		existing.Duration = parsed.Elapsed
	}

	if parsed.Action == models.Skip {
		existing.Result = models.Ignored
		existing.Duration = parsed.Elapsed
	}

	if parsed.Action == models.Fail {
		existing.Result = models.Failed
		existing.Duration = parsed.Elapsed
	}

	// store it for the next invocation
	trp.pending[cacheKey] = existing

	if existing.Result == models.Failed ||
		existing.Result == models.Ignored ||
		existing.Result == models.Successful {
		delete(trp.pending, cacheKey)
		trp.testCompleted(existing)
	}
}
