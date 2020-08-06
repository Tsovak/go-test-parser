package cmd

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/gobuffalo/packr"
	"github.com/tsovak/go-test-parser/models"
	"github.com/tsovak/go-test-parser/parser"
)

type resultCollector struct {
	sync.Mutex
	testResult             []models.TestResult
	reportSucceeded        bool
	loggedAtLeastOneResult bool
}

func NewResultCollector(reportSucceeded bool) *resultCollector {
	return &resultCollector{
		testResult:      make([]models.TestResult, 0),
		reportSucceeded: reportSucceeded,
	}
}

func (l *resultCollector) collectTestResult(result models.TestResult) {
	l.loggedAtLeastOneResult = true
	if result.Result == models.Successful {
		if !l.reportSucceeded {
			return
		}
	}
	l.Lock()
	l.testResult = append(l.testResult, result)
	l.Unlock()
}

func (l *resultCollector) hasAnyTest() bool {
	return l.loggedAtLeastOneResult
}

func ReadFromScanner(parser parser.TestResultParser, scanner *bufio.Scanner, verbose bool) {
	for scanner.Scan() {
		text := scanner.Text()
		parser.ParseLine(text, verbose)
	}
}

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func dumpStaticPageToDir(box packr.Box, folder string) error {
	file, err := os.Stat(folder)
	if !os.IsNotExist(err) {
		if !file.IsDir() {
			return errors.New("is not a directory")
		}
		err := os.RemoveAll(folder)
		if err != nil {
			return err
		}
	}

	if err := os.MkdirAll(filepath.Dir(folder), 0770); err != nil {
		return err
	}

	fileNames := box.List()

	for _, fileName := range fileNames {
		bytes, err := box.Find(fileName)
		if err != nil {
			return err
		}
		f, err := createFile(folder + string(os.PathSeparator) + fileName)
		if err != nil {
			return err
		}
		_, err = f.Write(bytes)
		if err != nil {
			return err
		}
	}

	return nil
}
