package collector

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/gobuffalo/packr"
	"github.com/tsovak/go-test-parser/models"
	"github.com/tsovak/go-test-parser/parser"
)

// ResultCollector collects the test result and store inside
type ResultCollector struct {
	sync.Mutex
	TestResult             []models.TestResult
	reportSucceeded        bool
	loggedAtLeastOneResult bool
}

func NewResultCollector(reportSucceeded bool) *ResultCollector {
	return &ResultCollector{
		TestResult:      make([]models.TestResult, 0),
		reportSucceeded: reportSucceeded,
	}
}

func (l *ResultCollector) CollectTestResult(result models.TestResult) {
	l.loggedAtLeastOneResult = true
	if result.Result == models.Successful {
		if !l.reportSucceeded {
			return
		}
	}
	l.Lock()
	l.TestResult = append(l.TestResult, result)
	l.Unlock()
}

func (l *ResultCollector) HasAnyTest() bool {
	return l.loggedAtLeastOneResult
}

func ReadFromScanner(parser parser.TestResultParser, scanner *bufio.Scanner, logToOutput, verbose bool) {
	for scanner.Scan() {
		text := scanner.Text()
		if logToOutput {
			fmt.Println(text)
		}
		parser.ParseLine(text, verbose)
	}
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func DumpStaticPageToDir(box packr.Box, folder string) error {
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
		f, err := CreateFile(folder + string(os.PathSeparator) + fileName)
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
