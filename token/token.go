package token

type Token int

const (
	// Special
	INVALID Token = iota
	EOF
	COMMENT

	// Literals
	IDENTIFIER // Void
	INTEGER    // 12345

	// Operators and delimiters
	PLUS     // +
	MINUS    // -
	MULTIPLY // *
	DIVIDE   // /
	MODULO   // %

	AND // &&
	OR  // ||

	EQUALS       // ==
	LESS_THAN    // <
	GREATER_THAN // >
	IS           // =
	NOT          // !

	NOT_EQUALS          // !=
	LESS_THAN_EQUALS    // <=
	GREATER_THAN_EQUALS // >=

	COMMA     // ,
	SEMICOLON // ;
	COLON     // :

	ROUND_BRACKET_OPEN   // (
	ROUND_BRACKET_CLOSE  // )
	CURLY_BRACKET_OPEN   // {
	CURLY_BRACKET_CLOSE  // }
	SQUARE_BRACKET_OPEN  // [
	SQUARE_BRACKET_CLOSE // ]

	// Keywords
	IF     // if
	ELSE   // else
	WHILE  // while
	RETURN // return
	TRUE   // True
	FALSE  // False
)

//go:generate stringer -type=Token

// Words that are keywords
var keywords = map[string]Token{
	"if":     IF,
	"else":   ELSE,
	"while":  WHILE,
	"return": RETURN,
	"True":   TRUE,
	"False":  FALSE,
}

// LookupWord returns the Token and literal for a scanned word
// (the corresponding keyword Token if it is a keyword, or else IDENTIFIER)
func LookupWord(word string) (Token, string) {
	if tok, ok := keywords[word]; ok {
		return tok, ""
	} else {
		return IDENTIFIER, word
	}
}
