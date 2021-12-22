package types

import (
	"fmt"

	"github.com/gramidt/mash-lang-for-codemash/ast"
)

type ObjType int

const (
	NULL_OBJ ObjType = iota
	ERROR_OBJ
	BOOL_OBJ
	STRING_OBJ
	FUN_OBJ
	BUILTIN_OBJ
	RETURN_VALUE_OBJ
)

var (
	NULL  = &Null{}
	TRUE  = &Bool{Value: true}
	FALSE = &Bool{Value: false}

	objTypes = map[ObjType]string{
		NULL_OBJ:         "NULL",
		ERROR_OBJ:        "ERROR",
		BOOL_OBJ:         "BOOL",
		STRING_OBJ:       "STRING",
		FUN_OBJ:          "FUNCTION",
		BUILTIN_OBJ:      "BUILTIN",
		RETURN_VALUE_OBJ: "RETURN_VALUE",
	}
)

func (ot ObjType) String() string {
	if t, ok := objTypes[ot]; ok {
		return t
	}
	return ""
}

type Object interface {
	Type() ObjType
	Inspect() string
	IsTruthy() bool
}

type Bool struct {
	Value bool
}

func (b *Bool) Type() ObjType   { return BOOL_OBJ }
func (b *Bool) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Bool) IsTruthy() bool  { return b.Value }

type String struct {
	Value string
}

func (s *String) Type() ObjType   { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }
func (b *String) IsTruthy() bool  { return true }

type Fun struct {
	Params []*ast.Ident
	Body   *ast.BlockStmt
	Env    *Env
}

func (f *Fun) Type() ObjType { return FUN_OBJ }
func (f *Fun) Inspect() string {
	return "<function>"
}
func (b *Fun) IsTruthy() bool { return true }

type BuiltinFun func(args ...Object) Object

type Builtin struct {
	Fun BuiltinFun
}

func (b *Builtin) Type() ObjType   { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string { return "<function>" }
func (b *Builtin) IsTruthy() bool  { return true }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjType   { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
func (b *ReturnValue) IsTruthy() bool   { return true }

type Null struct{}

func (n *Null) Type() ObjType   { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }
func (b *Null) IsTruthy() bool  { return false }

type Error struct {
	Msg string
}

func (e *Error) Type() ObjType   { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Msg }
func (b *Error) IsTruthy() bool  { return true }

func newError(format string, a ...interface{}) *Error {
	return &Error{Msg: fmt.Sprintf(format, a...)}
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}
