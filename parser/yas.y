%{
package parser

import (
    "github.com/arikui1911/YAS/ast"
)

%}

%union {
    intVal int
    node ast.Node
    tok Token
}

%token<tok> 	IntLiteralToken

%token 	FloatLiteralToken
        StringLiteralToken
        IdentifierToken
        DotToken
        BangToken
        AddToken
        SubToken
        MulToken
        DivToken
        ModToken
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
            $$ = nil
        }
        | expr
        ;

expr : primary
     ;

primary : IntLiteralToken
        {
            $$ = ast.NewIntLiteral($1.Line(), $1.Column(), $1.Value())
        }
        ;

%%