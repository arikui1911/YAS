package yas

import (
	"fmt"
	"github.com/arikui1911/YAS/ast"
)

type Value interface {
	InspectType() string
	opBinPlus(other Value) (Value, error)
	opBinMinus(other Value) (Value, error)
	opBinMul(other Value) (Value, error)
	opBinDiv(other Value) (Value, error)
	opBinMod(other Value) (Value, error)
}

type RuntimeError interface {
	error
	FileName() string
	Line() int
	Column() int
}

type runtimeError struct {
	fileName string
	line int
	column int
	message string
}

func (r *runtimeError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", r.fileName, r.line, r.column, r.message)
}

func (r *runtimeError) FileName() string { return r.fileName }

func (r *runtimeError) Line() int { return r.line }

func (r *runtimeError) Column() int { return r.column }

type Runtime interface {
	Evaluate(tree ast.Node) (Value, error)
}

type environment struct {
}

type runtime struct {
	stack []Value
	sp int
	toplevel *environment
}

func NewRuntime() *runtime {
	return &runtime{
		stack: make([]Value, 32),
		sp: 0,
		toplevel: &environment{},
	}
}

func (r *runtime) isStackEmpty() bool {
	return r.sp == 0
}

func (r *runtime) checkStackCapa() {
	if r.sp >= len(r.stack) {
		new := make([]Value, len(r.stack)*2)
		copy(new, r.stack)
		r.stack = new
	}
}

func (r *runtime) pushStack(v Value) {
	r.checkStackCapa()
	r.sp++
	r.stack[r.sp-1] = v
}

func (r *runtime) popStack() Value {
	r.sp--
	return r.stack[r.sp]
}

func (r *runtime) Evaluate(tree ast.Node) (Value, error) {
	err := r.eval(tree, r.toplevel)
	if err != nil {
		return nil, err
	}
	if r.isStackEmpty() {
		return nil, nil
	}
	return r.popStack(), nil
}

func (r *runtime) eval(node ast.Node, env *environment) (err error) {
	switch n := node.(type) {
	case *ast.Root:
		err = r.eval(n.TopLevel, env)
		if err != nil {
			if e, ok := err.(*runtimeError); ok {
				e.fileName = n.FileName
			}
		}
	case *ast.AdditionExpression:
		err = r.opBinary(n, env, opBinPlus)
	case *ast.SubtractionExpression:
		err = r.opBinary(n, env, opBinMinus)
	case *ast.MultiplicationExpression:
		err = r.opBinary(n, env, opBinMul)
	case *ast.DivisionExpression:
		err = r.opBinary(n, env, opBinDiv)
	case *ast.ModuloExpression:
		err = r.opBinary(n, env, opBinMod)
	case *ast.IntLiteral:
		r.pushStack(IntValue(n.Value()))
	case *ast.FloatLiteral:
		r.pushStack(FloatValue(n.Value()))
	case *ast.StringLiteral:
		r.pushStack(StringValue(n.Value()))
	}
	if err != nil {
		if _, ok := err.(*runtimeError); !ok {
			err = &runtimeError {
				line: node.Line(),
				column: node.Column(),
				message: err.Error(),
			}
		}
	}
	return
}

func opBinPlus(left Value, right Value) (Value, error) {
	return left.opBinPlus(right)
}

func opBinMinus(left Value, right Value) (Value, error) {
	return left.opBinMinus(right)
}

func opBinMul(left Value, right Value) (Value, error) {
	return left.opBinMul(right)
}

func opBinDiv(left Value, right Value) (Value, error) {
	return left.opBinDiv(right)
}

func opBinMod(left Value, right Value) (Value, error) {
	return left.opBinMod(right)
}

func (r *runtime) opBinary(exp ast.BinaryExpression, env *environment, op func(left Value, right Value) (Value, error)) error {
	err := r.eval(exp.Left(), env)
	if err != nil {
		return err
	}
	err = r.eval(exp.Right(), env)
	if err != nil {
		return err
	}
	right := r.popStack()
	left := r.popStack()
	result, err := op(left, right)
	if err != nil {
		return err
	}
	r.pushStack(result)
	return nil
}