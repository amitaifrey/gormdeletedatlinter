package gormdeletedat_test

import (
	"testing"

	"github.com/amitaifrey/gormdeletedat"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, gormdeletedat.Analyzer, "example")
}
