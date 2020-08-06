package cmd

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tsovak/go-test-parser/parser"
)

const ApplicationShortDescription string = "go-test-parser is a simple CLI tool for generating the UI report from go test output"

var (
	reportDir        string
	httpServeAddress int32
	includeSucceeded bool
	startHttp        bool
	verbose          bool
)

func parseInputParams(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.Int32Var(&httpServeAddress, "http", 8080, "HTTP address to serve")
	flags.BoolVarP(&includeSucceeded, "includeSucceeded", "s", false, "Include the successful test results or not")
	flags.BoolVarP(&startHttp, "web", "w", false, "Start only web serer for report displaying")
	flags.StringVarP(&reportDir, "output", "o", "./report", "The report output directory")
	flags.BoolVarP(&verbose, "verbose", "v", false, "Print verbose output")
	flags.BoolP("help", "h", false, "Help for application")
}

func GetWebAppCommand() *cobra.Command {
	retCmd := &cobra.Command{
		Use:     "./go-test-parser <go test -json -v output>",
		Short:   ApplicationShortDescription,
		Long:    "",
		Example: "./go-test-parser ./test.log -v \n go test -json -v ./... | go-test-parser -o ./report_directory",
		RunE: func(_ *cobra.Command, args []string) error {

			var scanner *bufio.Scanner
			if len(args) > 0 {
				filePath := args[0]
				ok := isFileExists(filePath)
				if !ok {
					return errors.New("file does not exists")
				}

				file, err := os.Open(filePath)
				if err != nil {
					return err
				}
				scanner = bufio.NewScanner(file)
				defer file.Close()
			} else {
				scanner = bufio.NewScanner(os.Stdin)
			}

			collector := NewResultCollector(includeSucceeded)
			parser := parser.NewResultsParser(collector.collectTestResult)

			// save static assets to the package
			box := packr.NewBox("../templates")
			tmplFile, err := box.FindString("report.tmpl.html")
			if err != nil {
				return err
			}
			tmpl := template.Must(template.New("tmpl").Parse(tmplFile))
			ReadFromScanner(parser, scanner, verbose)

			if startHttp {

				http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
					err := tmpl.Execute(w, collector.testResult)
					if err != nil {
						panic(err)
					}
				})

				http.Handle("/", http.StripPrefix("/", http.FileServer(box)))
				return http.ListenAndServe(fmt.Sprintf(":%d", httpServeAddress), nil)
			} else {
				err = dumpStaticPageToDir(box, reportDir)
				if err != nil {
					return errors.Wrapf(err, "cannot dump static page")
				}
				f, err := createFile(reportDir + "/index.html")
				if err != nil {
					return errors.Wrapf(err, "error creating file")
				}
				err = tmpl.Execute(f, collector.testResult)
				if err != nil {
					return err
				}

			}

			// fail if test weren't found
			if !collector.hasAnyTest() {
				return errors.New("no tests were found/logged!")
			}

			return nil
		},
	}
	parseInputParams(retCmd)
	return retCmd
}

func Execute() error {
	return GetWebAppCommand().Execute()
}
