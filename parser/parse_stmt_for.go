package parser

import (
	"github.com/chai2010/ugo/ast"
	"github.com/chai2010/ugo/token"
)

func (p *parser) parseStmt_for(blk *ast.BlockStmt) {
	tokFor := p.r.MustAcceptToken(token.FOR)
	_ = tokFor
}
