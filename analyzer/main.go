// https://docs.google.com/presentation/d/1I4pHnzV2dFOMbRcpA-XD0TaLcX6PBKpls6WxGHoMjOg/edit#slide=id.g6298a0e67590aa1e_6
// 自作の静的解析

package main

import (
	"go/ast"

	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/unitchecker"
)

var Analyzer = &analysis.Analyzer{
	Name:     "simple",
	Doc:      "simple is simple analyzer.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{new(ast.Ident)}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name == "gopher" {
				pass.Reportf(n.Pos(), "identifier is gopher!")
			}
		}
	})
	return nil, nil
}

func main() { unitchecker.Main(Analyzer) }

// go build .
// go vet -vettool=analyzer
