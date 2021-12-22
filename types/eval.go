package types

import (
	"github.com/gramidt/mash-lang-for-codemash/ast"
	"github.com/gramidt/mash-lang-for-codemash/grammar"
)

func Eval(node ast.Node, env *Env) Object {
	switch node := node.(type) {
	case *ast.Root:
		return evalRoot(node, env)

	case *ast.ExprStmt:
		return Eval(node.Expr, env)

	case *ast.BlockStmt:
		return evalBlockStmt(node, env)

	case *ast.IfStmt:
		return evalIfStmt(node, env)

	case *ast.Ident:
		return evalIdent(node, env)

	case *ast.VarStmt:
		return evalVarStmt(node, env)

	case *ast.BoolLit:
		return evalBoolLit(node)

	case *ast.StringLit:
		return evalStringLit(node)

	case *ast.FunLit:
		return evalFunLit(node, env)

	case *ast.CallExpr:
		return evalCallExpr(node, env)

	case *ast.BinaryExpr:
		return evalBinaryExpr(node, env)
	}

	return nil
}

func evalRoot(root *ast.Root, env *Env) Object {
	var result Object

	for _, stmt := range root.Stmts {
		result = Eval(stmt, env)

		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		}
	}

	return result
}

func evalBlockStmt(block *ast.BlockStmt, env *Env) Object {
	var result Object

	for _, stmt := range block.List {
		result = Eval(stmt, env)

		if result != nil {
			rt := result.Type()

			if rt == RETURN_VALUE_OBJ {
				return result
			}
		}
	}

	return result
}

func evalIdent(node *ast.Ident, env *Env) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("invalid identifier: " + node.Value)
}

func evalVarStmt(node *ast.VarStmt, env *Env) Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	env.Set(node.Name.Value, val)
	return nil
}

func evalIfStmt(node *ast.IfStmt, env *Env) Object {
	cond := Eval(node.Cond, env)
	if isError(cond) {
		return cond
	}

	if cond.IsTruthy() {
		return Eval(node.Body, env)
	} else if node.Else != nil {
		return Eval(node.Else, env)
	}

	return NULL
}

func evalBoolLit(node *ast.BoolLit) *Bool {
	if node.Value {
		return TRUE
	}
	return FALSE
}

func evalStringLit(node *ast.StringLit) Object {
	return &String{Value: node.Value}
}

func evalFunLit(node *ast.FunLit, env *Env) Object {
	return &Fun{Params: node.Params, Body: node.Body, Env: env}
}

func evalExprs(exprs []ast.Expr, env *Env) []Object {
	var result []Object

	for _, expr := range exprs {
		e := Eval(expr, env)
		if isError(e) {
			return []Object{e}
		}
		result = append(result, e)
	}

	return result
}

func evalCallExpr(node *ast.CallExpr, env *Env) Object {
	fun := Eval(node.Fun, env)
	if isError(fun) {
		return fun
	}

	args := evalExprs(node.Args, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	switch f := fun.(type) {
	case *Fun:
		env := NewEnclosedEnv(f.Env)
		for paramIdx, param := range f.Params {
			env.Set(param.Value, args[paramIdx])
		}

		evaluated := Eval(f.Body, env)
		if returnVal, ok := evaluated.(*ReturnValue); ok {
			return returnVal.Value
		}
		return evaluated

	case *Builtin:
		return f.Fun(args...)

	default:
		return newError("invalid function: %s", f.Type().String())
	}
}

func evalBinaryExpr(node *ast.BinaryExpr, env *Env) Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}

	switch {
	case left.Type() == BOOL_OBJ && right.Type() == BOOL_OBJ:
		leftVal := left.(*Bool)
		rightVal := right.(*Bool)
		return evalBoolBinaryExpr(node.Op.Type, leftVal, rightVal)
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		leftVal := left.(*String)
		rightVal := right.(*String)
		return evalStringBinaryExpr(node.Op.Type, leftVal, rightVal)
	default:
		return newError("invalid operation: %s %s %s", left.Type().String(), node.Op.Lit, right.Type().String())
	}
}

func evalBoolBinaryExpr(opTokType grammar.TokenType, left, right *Bool) Object {
	switch opTokType {
	case grammar.EQ:
		if left.Value == right.Value {
			return TRUE
		}
		return FALSE
	}

	return newError("unknown operator for Bool: %s", opTokType)
}

func evalStringBinaryExpr(opTokType grammar.TokenType, left, right *String) Object {
	switch opTokType {
	case grammar.ADD:
		return &String{Value: left.Value + right.Value}
	case grammar.EQ:
		if left.Value == right.Value {
			return TRUE
		}
		return FALSE
	}

	return newError("unknown operator for String: %s", opTokType)
}
