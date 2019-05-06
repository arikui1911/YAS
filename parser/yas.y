%{
package parser

%}

%union {
    intVal int
}

%token 	IntLiteralToken
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

%%

program :
        ;

%%