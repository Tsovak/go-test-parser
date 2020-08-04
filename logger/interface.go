package logger

import (
	"github.com/tsovak/go-test-parser/models"
)

type TestResultLogger interface {
	Log(result models.TestResult) string
}

type TestResultFormatter interface {
	TestSuiteStarted(result models.TestResult) string

	TestStart(result models.TestResult) string

	TestIgnored(result models.TestResult) string

	TestFinished(result models.TestResult) string

	TestStdErr(result models.TestResult) string

	TestStdOut(result models.TestResult) string

	TestFailed(result models.TestResult) string

	TestSuiteFinished(result models.TestResult) string
}
