package ast

import "github.com/gramidt/mash-lang-for-codemash/grammar"

// Nodes implement the Node interface
type Node interface {
	TokenLit() string
}

// Statement nodes implement the Stmt interface
type Stmt interface {
	Node
	stmtNode()
}

// Expression nodes impelemtn the Expr interface
type Expr interface {
	Node
	exprNode()
}

// Root node for the AST
type Root struct {
	Stmts []Stmt
}

func (r *Root) exprNode()        {}
func (r *Root) TokenLit() string { return r.Stmts[0].TokenLit() }

// An Ident node represents an identifier
type Ident struct {
	Token grammar.Token
	Value string
}

func (i *Ident) exprNode()        {}
func (i *Ident) TokenLit() string { return i.Token.Lit }

// An AssignStmt represents a
type AssignStmt struct {
	Token grammar.Token // the grammar.VAR token
	Name  *Ident
	Value Expr
}

func (vs *AssignStmt) stmtNode()        {}
func (vs *AssignStmt) TokenLit() string { return vs.Token.Lit }

// An ExprStmt represents a stand-alone expression
type ExprStmt struct {
	Token grammar.Token
	Expr  Expr
}

func (es *ExprStmt) stmtNode()        {}
func (es *ExprStmt) TokenLit() string { return es.Token.Lit }

// A BlockStmt node represents a braced statement list
type BlockStmt struct {
	Token grammar.Token
	List  []Stmt
}

func (bs *BlockStmt) stmtNode()        {}
func (bs *BlockStmt) TokenLit() string { return bs.Token.Lit }

// A VarStmt node represents a variable statement
type VarStmt struct {
	Token grammar.Token
	Name  *Ident
	Value Expr
}

func (vs *VarStmt) stmtNode()        {}
func (vs *VarStmt) TokenLit() string { return vs.Token.Lit }

// An IfStmt node represents an if statement
type IfStmt struct {
	Token grammar.Token
	Cond  Expr
	Body  *BlockStmt
	Else  *BlockStmt
}

func (is *IfStmt) exprNode()        {}
func (is *IfStmt) TokenLit() string { return is.Token.Lit }

// An BinaryExpr node represents a binary expression
type BinaryExpr struct {
	Op    grammar.Token // operator
	Left  Expr          // left operand
	Right Expr          // right operand
}

func (be *BinaryExpr) exprNode()        {}
func (be *BinaryExpr) TokenLit() string { return be.Op.Lit }

// An FunLit represents a function literal
type FunLit struct {
	Token  grammar.Token
	Params []*Ident
	Body   *BlockStmt
}

func (fl *FunLit) exprNode()        {}
func (fl *FunLit) TokenLit() string { return fl.Token.Lit }

// An CallExpr node represents an expression followed by an argument list
type CallExpr struct {
	Token grammar.Token
	Fun   Expr
	Args  []Expr
}

func (ce *CallExpr) exprNode()        {}
func (ce *CallExpr) TokenLit() string { return ce.Token.Lit }

// An StringLit node represents a string literal
type StringLit struct {
	Token grammar.Token
	Value string
}

func (sl *StringLit) exprNode()        {}
func (sl *StringLit) TokenLit() string { return sl.Token.Lit }

// An BoolLit node represents a boolean literal
type BoolLit struct {
	Token grammar.Token
	Value bool
}

func (bl *BoolLit) exprNode()        {}
func (bl *BoolLit) TokenLit() string { return bl.Token.Lit }
