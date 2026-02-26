// Package gormdeletedat provides a Go analysis pass that reports GORM model
// structs whose DeletedAt field is not typed as *gorm.DeletedAt.
package gormdeletedat

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const gormPkgPath = "gorm.io/gorm"
const deletedAtTypeName = "DeletedAt"

var Analyzer = &analysis.Analyzer{
	Name:     "gormdeletedat",
	Doc:      "reports GORM model structs where DeletedAt is not *gorm.DeletedAt",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		st := n.(*ast.StructType)
		for _, field := range st.Fields.List {
			for _, name := range field.Names {
				if name.Name != "DeletedAt" {
					continue
				}
				if isGormDeletedAtPtr(pass, field.Type) {
					continue
				}
				pass.Reportf(field.Pos(), "DeletedAt field should be *gorm.DeletedAt (from %q)", gormPkgPath)
			}
		}
	})

	return nil, nil
}

// isGormDeletedAtPtr returns true when expr resolves to *gorm.DeletedAt
// where gorm is the package gorm.io/gorm.
func isGormDeletedAtPtr(pass *analysis.Pass, expr ast.Expr) bool {
	typ := pass.TypesInfo.TypeOf(expr)
	if typ == nil {
		return false
	}

	ptr, ok := typ.(*types.Pointer)
	if !ok {
		return false
	}

	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}

	obj := named.Obj()
	return obj.Name() == deletedAtTypeName && obj.Pkg() != nil && obj.Pkg().Path() == gormPkgPath
}
