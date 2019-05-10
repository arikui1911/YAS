package parser

import (
	"fmt"
	"github.com/arikui1911/YAS/ast"
	"io"
	"strings"
)

//go:generate goyacc -o yas_yacc.go yas.y

// ParseError represents parsing failure with some
// location of cause
type ParseError interface {
	error
	FileName() string
	Line() int
	Column() int
}

type parseError struct {
	message  string
	fileName string
	line     int
	column   int
}

func (e *parseError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.fileName, e.line, e.column, e.message)
}

func (e *parseError) FileName() string { return e.fileName }

func (e *parseError) Line() int { return e.line }

func (e *parseError) Column() int { return e.column }

type adaptor struct {
	lex       Lexer
	lastToken Token
	fileName  string
	result    ast.Node
}

type ejectionSeat struct {
	pilot error
}

func bailout(pilot error) {
	panic(ejectionSeat{pilot: pilot})
}

func doRecover(e interface{}) error {
	if e == nil {
		return nil
	}
	seat, ok := e.(ejectionSeat)
	if !ok {
		panic(e)
	}
	return seat.pilot
}

func (a *adaptor) Lex(lval *yySymType) int {
	tok, err := a.lex.Lex()
	if err != nil {
		bailout(err)
	}
	a.lastToken = tok
	lval.tok = tok
	return tok.Kind()
}

func (a *adaptor) Error(msg string) {
	line := 1
	column := 1
	if a.lastToken != nil {
		line = a.lastToken.Line()
		column = a.lastToken.Column()
	}
	bailout(&parseError{
		message:  msg,
		fileName: a.fileName,
		line:     line,
		column:   column,
	})
}

func doParse(l Lexer, fileName string) (retTree ast.Node, retErr error) {
	a := &adaptor{lex: l, fileName: fileName}
	defer func() {
		retTree = a.result
		retErr = doRecover(recover())
	}()
	yyParse(a)
	return
}

func ParseIO(src io.Reader, fileName string) (ast.Node, error) {
	return doParse(NewLexer(src, fileName), fileName)
}

func ParseString(src string, fileName string) (ast.Node, error) {
	return doParse(NewLexer(strings.NewReader(src), fileName), fileName)
}
