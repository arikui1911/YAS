package parser

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"io"
	"unicode"
)

const (
	EOFToken = iota
	IntLiteralToken
	FloatLiteralToken
	StringLiteralToken
	IdentifierToken
)

// Token are results of `Lexer.Lex`.
// Entity of `Value` depends on `Kind`.
type Token interface {
	Kind() int
	Value() interface{}
	Line() int
	Column() int
}

type token struct {
	kind int
	value interface{}
	line int
	column int
}

func (t *token) Kind() int { return t.kind }

func (t *token) Value() interface{} { return t.value }

func (t *token) Line() int { return t.line }

func (t *token) Column() int { return t.column }

func (t *token) String() string {
	return fmt.Sprintf("%T{%v %v %v %v}", t, t.kind, t.value, t.line, t.column)
}

// Lexer is YAS lexical analizer.
type Lexer interface {
	Lex() (Token, error)
}

type lexer struct {
	src *bufio.Reader
	fileName string
	line int
	column int

	pushbacked rune
	isPushbacked bool
	lastLine int
	lastColumn int
}

// NewLexer creates a new instance of YAS lexer and return it.
// `fileName` is used for some error message.
func NewLexer(src io.Reader, fileName string) Lexer {
	return &lexer{
		src: bufio.NewReader(src),
		fileName: fileName,
		line: 1,
		column: 0,
		isPushbacked: false,
	}
}

func (l *lexer) errorf(t Token, format string, args ...interface{}) error {
	return fmt.Errorf("%s:%d:%d: " + format, append([](interface{}){l.fileName, t.Line(), t.Column()}, args...))
}

type lexingState uint

const (
	lexingInitial lexingState = iota
	lexingComment
	lexingZeroBegun
	lexingBinary
	lexingOctet
	lexingHex
	lexingNumber
	lexingFloat
	lexingFloatEPartOrSign
	lexingFloatEPart
	lexingDoubleQuoted
	lexingDoubleQuotedEscapeSequence
	lexingIdentifier
	lexingOperator
)

func (l *lexer) Lex() (Token, error) {
	buf := []rune{}
	addc := func(c rune){
		buf = append(buf, c)
	}
	state := lexingInitial
	t := token{kind: EOFToken}

LEXING:
	for {
		c, err := l.getc()
		if err == io.EOF {
			break LEXING
		}
		if err != nil {
			return nil, err
		}

		switch state {
		case lexingInitial:
			switch {
			case unicode.IsSpace(c):
				continue
			case c == '#':
				state = lexingComment
				continue
			}
			t.line = l.line
			t.column = l.column
			switch {
			case c == '0':
				state = lexingZeroBegun
			case unicode.IsDigit(c):
				addc(c)
				state = lexingNumber
			case c == '"':
				state = lexingDoubleQuoted
			case c == '_' || unicode.IsLetter(c):
				addc(c)
				state = lexingIdentifier
			default:
				addc(c)
				if !isPrefixOfAnyOperator(string(buf)) {
					return nil, l.errorf(&t, "invalid character: %c", c)
				}
				state = lexingOperator
			}
		case lexingComment:
			if c == '\n' {
				state = lexingInitial
			}
		case lexingZeroBegun:
			switch c {
			case '.':
				addc('0')
				addc(c)
				state = lexingFloat
			case 'b':
				state = lexingBinary
			case 'o':
				state = lexingOctet
			case 'x':
				state = lexingHex
			default:
				l.ungetc(c)
				t.kind = IntLiteralToken
				t.value = 0
				break LEXING
			}
		case lexingBinary:
			switch c {
			case '0', '1':
				addc(c)
			default:
				l.ungetc(c)
				v, err := strconv.ParseInt(string(buf), 2, 64)
				if err != nil {
					return nil, err
				}
				t.kind = IntLiteralToken
				t.value = int(v)
				break LEXING
			}
		case lexingOctet:
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7':
				addc(c)
			default:
				l.ungetc(c)
				v, err := strconv.ParseInt(string(buf), 8, 64)
				if err != nil {
					return nil, err
				}
				t.kind = IntLiteralToken
				t.value = int(v)
				break LEXING
			}
		case lexingHex:
			switch c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F':
				addc(c)
			default:
				l.ungetc(c)
				v, err := strconv.ParseInt(string(buf), 16, 64)
				if err != nil {
					return nil, err
				}
				t.kind = IntLiteralToken
				t.value = int(v)
				break LEXING
			}
		case lexingNumber:
			switch {
			case unicode.IsDigit(c):
				addc(c)
			case c == '.':
				addc(c)
				state = lexingFloat
			default:
				l.ungetc(c)
				v, err := strconv.ParseInt(string(buf), 10, 64)
				if err != nil {
					return nil, err
				}
				t.kind = IntLiteralToken
				t.value = int(v)
				break LEXING
			}
		case lexingFloat:
			switch {
			case unicode.IsDigit(c):
				addc(c)
			case c == 'e' || c == 'E':
				addc(c)
				state = lexingFloatEPartOrSign
			default:
				l.ungetc(c)
				x, err := strconv.ParseFloat(string(buf), 64)
				if err != nil {
					return nil, err
				}
				t.kind = FloatLiteralToken
				t.value = x
				break LEXING
			}
		case lexingFloatEPartOrSign:
			switch {
			case unicode.IsDigit(c):
				addc(c)
			case c == '+' || c == '-':
				addc(c)
			default:
				l.ungetc(c)
			}
			state = lexingFloatEPart
		case lexingFloatEPart:
			if unicode.IsDigit(c) {
				addc(c)
			} else {
				l.ungetc(c)
				x, err := strconv.ParseFloat(string(buf), 64)
				if err != nil {
					return nil, err
				}
				t.kind = FloatLiteralToken
				t.value = x
				break LEXING
			}
		case lexingDoubleQuoted:
			switch c {
			case '"':
				state = lexingInitial
				t.kind = StringLiteralToken
				t.value = string(buf)
				break LEXING
			case '\\':
				state = lexingDoubleQuotedEscapeSequence
			default:
				addc(c)
			}
		case lexingDoubleQuotedEscapeSequence:
			switch c {
			case 'n':
				addc('\n')
			case 't':
				addc('\t')
			default:
				addc(c)
			}
			state = lexingDoubleQuoted
		case lexingIdentifier:
			if c == '_' || unicode.IsLetter(c) || unicode.IsDigit(c) {
				addc(c)
			} else {
				l.ungetc(c)
				s := string(buf)
				t.kind = IdentifierToken
				if k, ok := keywords[s]; ok {
					t.kind = k
				}
				t.value = s
				break LEXING
			}
		case lexingOperator:
			addc(c)
			if !isPrefixOfAnyOperator(string(buf)) {
				l.ungetc(c)
				s := string(buf[:len(buf)-1])
				kind, ok := operators[s]
				if !ok {
					panic("must not happen")
				}
				t.kind = kind
				t.value = s
				break LEXING
			}
		default:
			panic("must not happen")
		}
	}

	switch state {
	case lexingDoubleQuoted, lexingDoubleQuotedEscapeSequence:
		return nil, l.errorf(&t, "unterminated string literal")
	}

	return &t, nil
}

var keywords = map[string]int{}

var operators = map[string]int{}

func isPrefixOfAnyOperator(prefix string) bool {
	for op := range operators {
		if strings.HasPrefix(op, prefix) {
			return true
		}
	}
	return false
}

func (l *lexer) getc() (c rune, err error) {
	if l.isPushbacked {
		l.isPushbacked = false
		c = l.pushbacked
	} else {
		c, _, err = l.src.ReadRune()
		if err != nil {
			return
		}
	}

	l.lastLine = l.line
	l.lastColumn = l.column
	l.column++
	if c == '\n' {
		l.line++
		l.column = 0
	}
	return
}

func (l *lexer) ungetc(c rune) {
	l.pushbacked = c
	l.isPushbacked = true
	l.line = l.lastLine
	l.column = l.lastColumn
}