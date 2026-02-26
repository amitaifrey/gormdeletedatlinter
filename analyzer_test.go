package gormdeletedat_test

import (
	"testing"

	gormdeletedat "github.com/amitaifrey/gormdeletedatlinter"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, gormdeletedat.Analyzer, "example")
}
