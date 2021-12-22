package parser

import (
	"fmt"

	"github.com/gramidt/mash-lang-for-codemash/ast"
	"github.com/gramidt/mash-lang-for-codemash/grammar"
	"github.com/gramidt/mash-lang-for-codemash/scanner"
)

type (
	parseFn       func() ast.Expr
	binaryParseFn func(ast.Expr) ast.Expr
)

type Parser struct {
	lexer  *scanner.Scanner
	errors []string

	tok     grammar.Token
	peekTok grammar.Token

	parseFunctions map[grammar.TokenType]parseFn
	binaryParseFns map[grammar.TokenType]binaryParseFn
}

func NewParser(lexer *scanner.Scanner) *Parser {
	p := &Parser{
		lexer:  lexer,
		errors: []string{},
	}

	p.parseFunctions = map[grammar.TokenType]parseFn{
		grammar.IDENT:  p.parseIdent,
		grammar.STRING: p.parseStringLit,
		grammar.TRUE:   p.parseBoolLit,
		grammar.FALSE:  p.parseBoolLit,
		grammar.LPAREN: p.parseGroupedExpr,
		grammar.FUN:    p.parseFunLit,
		grammar.IF:     p.parseIfSmt,
	}

	p.binaryParseFns = map[grammar.TokenType]binaryParseFn{
		grammar.ADD:    p.parseBinaryExpr,
		grammar.EQ:     p.parseBinaryExpr,
		grammar.LPAREN: p.parseCallExpr,
	}

	// Read the first two tokens, so tok and peekTok are set.
	p.nextTwo()

	return p
}

func (p *Parser) Parse() *ast.Root {
	root := &ast.Root{}
	root.Stmts = []ast.Stmt{}

	for !p.tokenIs(grammar.EOF) {
		stmt := p.parseStmt()
		root.Stmts = append(root.Stmts, stmt)
		p.next()
	}

	return root
}

func (p *Parser) Errors() []string {
	return p.errors
}

// ----------------------------------------------------------------------------
// Parsing support

func (p *Parser) next() {
	p.tok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *Parser) nextTwo() {
	p.next()
	p.next()
}

func (p *Parser) tokenIs(t grammar.TokenType) bool {
	return p.tok.Type == t
}

func (p *Parser) peekTokenIs(t grammar.TokenType) bool {
	return p.peekTok.Type == t
}

func (p *Parser) expectPeekTokenIs(t grammar.TokenType) bool {
	if p.peekTokenIs(t) {
		p.next()
		return true
	}

	msg := fmt.Sprintf("expected peek token to be %s, got %s instead", t, p.peekTok.Type)
	p.errors = append(p.errors, msg)
	return false
}

// ----------------------------------------------------------------------------
// Parsing statements and expressions

func (p *Parser) parseIdent() ast.Expr {
	return &ast.Ident{Token: p.tok, Value: p.tok.Lit}
}

func (p *Parser) parseStringLit() ast.Expr {
	return &ast.StringLit{Token: p.tok, Value: p.tok.Lit}
}

func (p *Parser) parseBoolLit() ast.Expr {
	return &ast.BoolLit{Token: p.tok, Value: p.tokenIs(grammar.TRUE)}
}

func (p *Parser) parseFunLit() ast.Expr {
	lit := &ast.FunLit{Token: p.tok}

	if !p.expectPeekTokenIs(grammar.LPAREN) {
		return nil
	}

	lit.Params = p.parseFunParams()

	if !p.expectPeekTokenIs(grammar.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStmt()

	return lit
}

func (p *Parser) parseFunParams() []*ast.Ident {
	idents := []*ast.Ident{}

	if p.peekTokenIs(grammar.RPAREN) {
		p.next()
		return idents
	}

	p.next()

	ident := &ast.Ident{Token: p.tok, Value: p.tok.Lit}
	idents = append(idents, ident)

	for p.peekTokenIs(grammar.COMMA) {
		p.nextTwo()
		ident := &ast.Ident{Token: p.tok, Value: p.tok.Lit}
		idents = append(idents, ident)
	}

	if !p.expectPeekTokenIs(grammar.RPAREN) {
		return nil
	}

	return idents
}

func (p *Parser) parseStmt() ast.Stmt {
	switch p.tok.Type {
	case grammar.VAR:
		return p.parseVarStmt()
	default:
		return p.parseExprStmt()
	}
}

func (p *Parser) parseBlockStmt() *ast.BlockStmt {
	block := &ast.BlockStmt{Token: p.tok}
	block.List = []ast.Stmt{}

	p.next()

	for !p.tokenIs(grammar.RBRACE) && !p.tokenIs(grammar.EOF) {
		stmt := p.parseStmt()
		block.List = append(block.List, stmt)
		p.next()
	}

	return block
}

func (p *Parser) parseVarStmt() *ast.VarStmt {
	stmt := &ast.VarStmt{Token: p.tok}

	if !p.expectPeekTokenIs(grammar.IDENT) {
		return nil
	}

	stmt.Name = &ast.Ident{Token: p.tok, Value: p.tok.Lit}

	if !p.expectPeekTokenIs(grammar.ASSIGN) {
		return nil
	}

	p.next()

	stmt.Value = p.parseExpr(grammar.LowestPrecedence)

	if p.peekTokenIs(grammar.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseIfSmt() ast.Expr {
	stmt := &ast.IfStmt{Token: p.tok}

	if !p.expectPeekTokenIs(grammar.LPAREN) {
		return nil
	}

	p.next()
	stmt.Cond = p.parseExpr(grammar.LowestPrecedence)

	if !p.expectPeekTokenIs(grammar.RPAREN) {
		return nil
	}

	if !p.expectPeekTokenIs(grammar.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStmt()

	if p.peekTokenIs(grammar.ELSE) {
		p.next()
		if !p.expectPeekTokenIs(grammar.LBRACE) {
			return nil
		}
		stmt.Else = p.parseBlockStmt()
	}

	return stmt
}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	stmt := &ast.ExprStmt{Token: p.tok}
	stmt.Expr = p.parseExpr(grammar.LowestPrecedence)

	if p.peekTokenIs(grammar.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseExpr(precedence int) ast.Expr {
	exprFun := p.parseFunctions[p.tok.Type]
	if exprFun == nil {
		msg := fmt.Sprintf("no parse function for %s found", p.tok.Type)
		p.errors = append(p.errors, msg)
		return nil
	}
	expr := exprFun()

	for !p.peekTokenIs(grammar.SEMICOLON) && precedence < p.peekTok.Precedence() {
		binaryExpr := p.binaryParseFns[p.peekTok.Type]
		if binaryExpr == nil {
			return expr
		}
		p.next()
		expr = binaryExpr(expr)
	}

	return expr
}

func (p *Parser) parseGroupedExpr() ast.Expr {
	p.next()

	expr := p.parseExpr(grammar.LowestPrecedence)

	if !p.expectPeekTokenIs(grammar.RPAREN) {
		return nil
	}

	return expr
}

func (p *Parser) parseBinaryExpr(left ast.Expr) ast.Expr {
	expr := &ast.BinaryExpr{
		Op:   p.tok,
		Left: left,
	}

	prec := p.tok.Precedence()
	p.next()
	expr.Right = p.parseExpr(prec)

	return expr
}

func (p *Parser) parseCallExpr(fun ast.Expr) ast.Expr {
	expr := &ast.CallExpr{Token: p.tok, Fun: fun}
	expr.Args = []ast.Expr{}

	if p.peekTokenIs(grammar.RPAREN) {
		p.next()
		return expr
	}

	p.next()
	expr.Args = append(expr.Args, p.parseExpr(grammar.LowestPrecedence))

	for p.peekTokenIs(grammar.COMMA) {
		p.nextTwo()
		expr.Args = append(expr.Args, p.parseExpr(grammar.LowestPrecedence))
	}

	if !p.expectPeekTokenIs(grammar.RPAREN) {
		return nil
	}

	return expr
}
