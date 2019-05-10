%{
package parser

import (
    "github.com/arikui1911/YAS/ast"
)

%}

%union {
    node ast.Node
    tok Token
}

%token<tok> IntLiteralToken
            FloatLiteralToken
            StringLiteralToken
            IdentifierToken
            DotToken
            BangToken
            AddToken
            SubToken
            MulToken
            DivToken
            ModToken
            LParenToken
            RParenToken
            LBraceToken
            RBraceToken
            LBracketToken
            RBracketToken
            AssignToken
            AddAssignToken
            SubAssignToken
            MulAssignToken
            DivAssignToken
            ModAssignToken
            LetToken
            ArrowToken
            EqToken
            NeToken
            GtToken
            LtToken
            GeToken
            LeToken
            IfToken
            ElsifToken
            ElseToken
            WhileToken
            ContinueToken
            BreakToken
            ReturnToken
            DefToken
            VarToken

%type<node> program expr primary

%%

program :
        {
            yylex.(*adaptor).result = nil
            $$ = nil
        }
        | expr
        {
            yylex.(*adaptor).result = $1
            $$ = $1
        }
        ;

expr : primary
     ;

primary : LParenToken expr RParenToken
        {
            $$ = $2
        }
        | IntLiteralToken
        {
            $$ = ast.NewIntLiteral($1.Line(), $1.Column(), $1.Value().(int))
        }
        | FloatLiteralToken
        {
            $$ = ast.NewFloatLiteral($1.Line(), $1.Column(), $1.Value().(float64))
        }
        | StringLiteralToken
        {
            $$ = ast.NewStringLiteral($1.Line(), $1.Column(), $1.Value().(string))
        }
        ;

%%
