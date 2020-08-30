package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
	_collector "github.com/tsovak/go-test-parser/collector"
	"github.com/tsovak/go-test-parser/parser"
)

type Executor struct {
	collector        *_collector.ResultCollector
	parser           parser.TestResultParser
	workingDirectory string
}

func NewExecutor() *Executor {
	collector := _collector.NewResultCollector(true)
	parser := parser.NewResultsParser(collector.CollectTestResult)

	return &Executor{
		collector: collector,
		parser:    parser,
	}
}

func (e *Executor) Execute(f *flag.FlagSet) error {
	args := f.Args()
	workingDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error retrieving current working directory: %+v", err)
	}
	e.workingDirectory = workingDirectory

	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	cmd.Dir = workingDirectory

	out, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error obtaining stdout: %+v", err)
	}

	errOut, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error obtaining stderr: %+v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting: %+v", err)
	}

	outScanner := bufio.NewScanner(out)
	go _collector.ReadFromScanner(e.parser, outScanner, true, false)

	errScanner := bufio.NewScanner(errOut)
	go _collector.ReadFromScanner(e.parser, errScanner, true, false)

	_ = cmd.Wait()
	var exitCode = 0
	if cmd.ProcessState != nil || cmd.ProcessState.ExitCode() > 0 {
		exitCode = cmd.ProcessState.ExitCode()
	}

	err = e.dump()
	if err != nil {
		return errors.Wrapf(err, "cannot dump test. Test exit code: %d", exitCode)
	}

	os.Exit(exitCode)
	return nil
}

func (e *Executor) dump() error {
	if len(e.collector.TestResult) == 0 {
		return nil
	}
	// save static assets to the package
	box := packr.NewBox("../../templates")
	tmplFile, err := box.FindString("report.tmpl.html")
	if err != nil {
		return err
	}
	tmpl := template.Must(template.New("tmpl").Parse(tmplFile))

	dir := fmt.Sprintf("%s/.reports/report_%s", e.workingDirectory, buildFileName())
	err = _collector.DumpStaticPageToDir(box, dir)
	if err != nil {
		return errors.Wrapf(err, "cannot dump static page")
	}
	f, err := _collector.CreateFile(dir + "/index.html")
	if err != nil {
		return errors.Wrapf(err, "error creating file")
	}
	err = tmpl.Execute(f, e.collector.TestResult)
	if err != nil {
		return err
	}
	fmt.Printf("report was generated in %s\n", dir)

	return nil
}

func buildFileName() string {
	return time.Now().Format("20060102_150405")
}
