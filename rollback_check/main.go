package main

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

var Analyzer = &analysis.Analyzer{
	Name: "rollbackcheck",
	Doc:  "check that db.Rollback is called in t.Run inside for loop with http.MethodUpdate and router.ServeHTTP",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if funcDecl, ok := n.(*ast.FuncDecl); ok {
				if !(strings.HasSuffix(funcDecl.Name.Name, "FailCase") || strings.HasSuffix(funcDecl.Name.Name, "Permission")) {
					ast.Inspect(funcDecl.Body, func(nn ast.Node) bool {
						// `for` 文を検出
						// rangeを使っているとRangeStmtになる
						if rangeStmt, ok := nn.(*ast.RangeStmt); ok {
							checkForLoop(pass, rangeStmt)
						}
						return true
					})
				}
			}
			return true
		})
	}
	return nil, nil
}

func checkForLoop(pass *analysis.Pass, forStmt *ast.RangeStmt) {
	// `for` 文の内部を探索
	ast.Inspect(forStmt.Body, func(n ast.Node) bool {
		// `t.Run` を検出
		if call, ok := n.(*ast.CallExpr); ok {
			if fun, ok := call.Fun.(*ast.SelectorExpr); ok && fun.Sel.Name == "Run" {
				if recv, ok := fun.X.(*ast.Ident); ok && recv.Name == "t" {
					// t.Run の中身を確認
					checkTRunBody(pass, call)
				}
			}
		}

		return true
	})
}

func checkTRunBody(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) < 2 {
		return
	}

	// t.Run の引数として渡される関数リテラルを取得
	funcLit, ok := call.Args[1].(*ast.FuncLit)
	if !ok {
		return
	}

	var foundMethod, foundServeHTTP, foundSavePoint, foundRollbackTo bool

	// t.Run 内部を探索
	ast.Inspect(funcLit.Body, func(n ast.Node) bool {
		switch expr := n.(type) {
		case *ast.CallExpr:
			// router.ServeHTTP や db.Rollback の検出
			if sel, ok := expr.Fun.(*ast.SelectorExpr); ok {
				switch sel.Sel.Name {
				case "ServeHTTP":
					// router.ServeHTTP を検出
					if recv, ok := sel.X.(*ast.Ident); ok && recv.Name == "router" {
						foundServeHTTP = true
					}
				case "SavePoint":
					// db.SavePoint を検出
					if recv, ok := sel.X.(*ast.Ident); ok && recv.Name == "db" {
						foundSavePoint = true
					}
				case "RollbackTo":
					// db.RollbackTo を検出
					if recv, ok := sel.X.(*ast.Ident); ok && recv.Name == "db" {
						foundRollbackTo = true
					}
				}
			}

		case *ast.SelectorExpr:
			switch expr.Sel.Name {
			case "MethodPost", "MethodPatch", "MethodUpdate", "MethodDelete":
				foundMethod = true
			}
		}
		return true
	})

	// 条件を満たしている場合に警告を出す
	if foundMethod && foundServeHTTP && !(foundSavePoint && foundRollbackTo) {
		pass.Reportf(call.Pos(), "テストケーススライスのt.Run内で 副作用のあるhttp.Method と router.ServeHTTP が使用されていますが、db.Rollback が呼び出されていません")
	}
}

func main() {
	unitchecker.Main(Analyzer)
}
