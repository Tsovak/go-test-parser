package logger

import (
	"fmt"

	"github.com/tsovak/go-test-parser/models"
)

type ResultLogger struct {
	logger XunitResultFormatter
}

func (rl ResultLogger) Log(result models.TestResult) string {
	// output the package name for logical grouping
	builder := rl.logger.TestSuiteStarted(result)

	// always has to be a start
	builder += rl.logger.TestStart(result)

	switch result.Result {
	case models.Successful:
		builder += rl.logger.TestStdOut(result)
		break

	case models.Failed:
		builder += rl.logger.TestStdOut(result)
		builder += rl.logger.TestStdOut(result)
		builder += rl.logger.TestFailed(result)
		break

	case models.Ignored:
		builder += rl.logger.TestStdOut(result)
		builder += rl.logger.TestIgnored(result)
		break

	default:
		panic(fmt.Errorf("Unexpected Test State %q - this is a bug in the test runner", result.Result))
	}

	// always has to be a finish
	builder += rl.logger.TestFinished(result)

	// output the package name for logical grouping
	builder += rl.logger.TestSuiteFinished(result)

	return builder
}
