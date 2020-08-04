package models

import (
	"fmt"
	"time"
)

type Action string

var (
	Run    Action = "run"
	Pause  Action = "pause"
	Cont   Action = "cont"
	Pass   Action = "pass"
	Bench  Action = "bench"
	Fail   Action = "fail"
	Output Action = "output"
	Skip   Action = "skip"
)

type GoTestEvent struct {
	Time    time.Time `json:"Time"`
	Action  Action    `json:"Action"`
	Package string    `json:"Package"`
	Test    string    `json:"Test"`
	Elapsed float64   `json:"Elapsed"`
	Output  string    `json:"Output"`
}

func (t GoTestEvent) CacheKey() string {
	return fmt.Sprintf("%s-%s", t.Package, t.Test)
}

func (t GoTestEvent) Duration() int64 {
	return int64(t.Elapsed * 1000)
}

type Result string

var (
	Failed     Result = "Failed"
	Ignored    Result = "Ignored"
	Pending    Result = "Pending"
	Successful Result = "Successful"
)

type TestResult struct {
	Package  string
	TestName string
	Result   Result
	Duration float64
	StdOut   string
	StdErr   string
}

type Tests struct {
	Packages []Package
}
type Package struct {
	Name    []string
	Results []TestResult
}
