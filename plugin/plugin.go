package plugin

import (
	"github.com/amitaifrey/gormdeletedat"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

// gormdeletedatPlugin implements the LinterPlugin interface.
type gormdeletedatPlugin struct{}

func (p *gormdeletedatPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{gormdeletedat.Analyzer}, nil
}

func (p *gormdeletedatPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func init() {
	register.Plugin("gormdeletedat", New)
}

// New is the entrypoint expected by golangci-lint Go Plugin System.
func New(conf any) (register.LinterPlugin, error) {
	return &gormdeletedatPlugin{}, nil
}
