package parser

import (
	"go/ast"
	"go/parser"
)

func main() {
	expr, err := parser.ParseExpr(`v + 1`)
	if err != nil {
		panic(err)
	}
	ast.Print(nil, expr)
}

// go run parser.go -v
//      0  *ast.BinaryExpr {
//      1  .  X: *ast.Ident {
//      2  .  .  NamePos: 1
//      3  .  .  Name: "v"
//      4  .  }
//      5  .  OpPos: 3
//      6  .  Op: +
//      7  .  Y: *ast.BasicLit {
//      8  .  .  ValuePos: 5
//      9  .  .  Kind: INT
//     10  .  .  Value: "1"
//     11  .  }
//     12  }
