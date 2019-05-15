package parser

import (
	"strings"
	"testing"
)

func testOneToken(t *testing.T, src string, kind int, val interface{}) {
	l := NewLexer(strings.NewReader(src), "(test)")
	tok, err := l.Lex()
	if err != nil {
		t.Errorf("%v: %v", tok, err)
		return
	}
	if tok.Kind() != kind {
		t.Errorf("%v: expect %v but %v", tok, kind, tok.Kind())
		return
	}
	if val == nil {
		return
	}
	if tok.Value() != val {
		t.Errorf("%v: expect %v but %v", tok, val, tok.Value())
	}
}

func TestLexer(t *testing.T) {
	dataSet := []struct {
		src  string
		kind int
		val  interface{}
	}{
		{"", 0, nil},
		{"123", IntLiteralToken, 123},
		{"0b101", IntLiteralToken, 5},
		{"0o123", IntLiteralToken, 83},
		{"0x12a", IntLiteralToken, 298},
		{"0.12", FloatLiteralToken, 0.12},
		{"1.23", FloatLiteralToken, 1.23},
		{"\"Hello\"", StringLiteralToken, "Hello"},
		{"hoge", IdentifierToken, "hoge"},

		{".", DotToken, nil},
		{":", ColonToken, nil},
		{";", SemicolonToken, nil},
		{",", CommaToken, nil},
		{"!", BangToken, nil},
		{"+", AddToken, nil},
		{"-", SubToken, nil},
		{"*", MulToken, nil},
		{"/", DivToken, nil},
		{"%", ModToken, nil},
		{"(", LParenToken, nil},
		{")", RParenToken, nil},
		{"{", LBraceToken, nil},
		{"}", RBraceToken, nil},
		{"[", LBracketToken, nil},
		{"]", RBracketToken, nil},
		{"=", AssignToken, nil},
		{"+=", AddAssignToken, nil},
		{"-=", SubAssignToken, nil},
		{"*=", MulAssignToken, nil},
		{"/=", DivAssignToken, nil},
		{"%=", ModAssignToken, nil},
		{":=", LetToken, nil},
		{"->", ArrowToken, nil},
		{"==", EqToken, nil},
		{"!=", NeToken, nil},
		{">", GtToken, nil},
		{"<", LtToken, nil},
		{">=", GeToken, nil},
		{"<= ", LeToken, nil},

		{"if", IfToken, nil},
		{"elsif", ElsifToken, nil},
		{"else", ElseToken, nil},
		{"while", WhileToken, nil},
		{"continue", ContinueToken, nil},
		{"break", BreakToken, nil},
		{"return", ReturnToken, nil},
		{"def", DefToken, nil},
		{"var", VarToken, nil},
	}

	for _, data := range dataSet {
		testOneToken(t, data.src, data.kind, data.val)
	}
}
