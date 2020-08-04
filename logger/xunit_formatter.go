package logger

import (
	"fmt"
	"strings"

	"github.com/tsovak/go-test-parser/models"
)

type XunitResultFormatter struct{}

func (tl XunitResultFormatter) TestSuiteStarted(result models.TestResult) string {
	return fmt.Sprintf("<testsuite failures=\"0\" "+
		"package=\"%s\" "+
		"tests=\"1\" "+
		"time=\"%f\">",
		result.Package, result.Duration)
}

func (tl XunitResultFormatter) TestStart(result models.TestResult) string {
	return fmt.Sprintf("<testcase classname=\"%s\" name=\"%s\" time=\"%f\">", result.Package, result.TestName, result.Duration)
}

func (tl XunitResultFormatter) TestIgnored(result models.TestResult) string {
	return ""
}

func (tl XunitResultFormatter) TestFinished(result models.TestResult) string {
	return fmt.Sprintf("</testcase>")
}

func (tl XunitResultFormatter) TestStdErr(result models.TestResult) string {
	err := strings.TrimSuffix(result.StdErr, "\n")
	if len(err) > 0 {
		return fmt.Sprintf("<syserr>\n %s</syserr>\n", err)
	}
	return ""
}

func (tl XunitResultFormatter) TestStdOut(result models.TestResult) string {
	out := strings.TrimSuffix(result.StdOut, "\n")
	if len(out) > 0 {
		return fmt.Sprintf("<sysout>\n %s</sysout>\n", out)
	}

	return ""
}

func (tl XunitResultFormatter) TestFailed(result models.TestResult) string {
	return ""
}

func (tl XunitResultFormatter) TestSuiteFinished(result models.TestResult) string {
	return fmt.Sprintf("</testsuite>\n")
}
