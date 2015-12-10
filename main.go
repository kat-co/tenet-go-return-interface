package main

import (
	"go/ast"

	"github.com/lingo-reviews/tenets/go/dev/server"
	"github.com/lingo-reviews/tenets/go/dev/tenet"
)

func main() {
	t := &interfaceReturned{}
	t.setup()

	server.Serve(t)
}

type interfaceReturned struct {
	tenet.Base
}

func (t *interfaceReturned) setup() {
	t.SetInfo(tenet.Info{
		Name:     "returning interfaces",
		Usage:    "return structs, not interfaces",
		Language: "Go",
		Description: `
A Go convention is to return types, and take in interfaces.
`[1:],
	})

	issueName := t.RegisterIssue(
		"issue-name",
		tenet.AddComment("usually it's idiomatic to return concrete types"),
	)

	t.SmellNode(func(rvw tenet.Review, returnStmt *ast.FuncType) error {
		if returnStmt.Results == nil {
			return nil
		}

		for _, r := range returnStmt.Results.List {
			if ident, ok := r.Type.(*ast.Ident); ok {
				if IsInterface(ident) {
					rvw.RaiseNodeIssue(issueName, returnStmt)
				}
			}
		}
		return nil
	})
}

func IsInterface(ident *ast.Ident) bool {
	if ident.Obj != nil {
		switch ident.Obj.Decl.(type) {
		case *ast.TypeSpec:
			return true
		}
	}
	return false
}
