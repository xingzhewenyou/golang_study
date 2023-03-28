package lexer

package lexer

// TokenType represents the type of a token.
type TokenType string

// Token represents a token in the input.
type Token struct {
	Type    TokenType // The type of the token.
	Literal string    // The literal value of the token.
}

// TokenType constants.
const (
	ILLEGAL = "ILLEGAL" // Illegal token.
	EOF     = "EOF"     // End-of-file token.

	// Identifiers and literals.
	IDENT = "IDENT" // Identifier token.
	INT   = "INT"   // Integer literal token.

	// Operators.
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ  = "=="
	NEQ = "!="

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
)

// keywords maps identifiers to their corresponding TokenType.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// lookupIdent returns the corresponding TokenType for the given identifier.
func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
