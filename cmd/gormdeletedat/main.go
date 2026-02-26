// Command gormdeletedat runs the gormdeletedat analyzer.
package main

import (
	"github.com/amitaifrey/gormdeletedat"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gormdeletedat.Analyzer)
}
