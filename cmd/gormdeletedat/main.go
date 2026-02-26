// Command gormdeletedat runs the gormdeletedat analyzer.
package main

import (
	gormdeletedat "github.com/amitaifrey/gormdeletedatlinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gormdeletedat.Analyzer)
}
