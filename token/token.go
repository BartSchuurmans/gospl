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

	EMPTY_LIST // []

	// Keywords
	IF     // if
	ELSE   // else
	WHILE  // while
	RETURN // return
)

//go:generate stringer -type=Token

// Words that are keywords
var keywords = map[string]Token{
	"if":     IF,
	"else":   ELSE,
	"while":  WHILE,
	"return": RETURN,
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

var printStrings = map[Token]string{
	PLUS:     "+",
	MINUS:    "-",
	MULTIPLY: "*",
	DIVIDE:   "/",
	MODULO:   "%",

	AND: "&&",
	OR:  "||",

	EQUALS:       "==",
	LESS_THAN:    "<",
	GREATER_THAN: ">",
	IS:           "=",
	NOT:          "!",

	NOT_EQUALS:          "!=",
	LESS_THAN_EQUALS:    "<=",
	GREATER_THAN_EQUALS: ">=",

	COMMA:     ",",
	SEMICOLON: ";",
	COLON:     ":",

	ROUND_BRACKET_OPEN:   "(",
	ROUND_BRACKET_CLOSE:  ")",
	CURLY_BRACKET_OPEN:   "{",
	CURLY_BRACKET_CLOSE:  "}",
	SQUARE_BRACKET_OPEN:  "[",
	SQUARE_BRACKET_CLOSE: "]",

	EMPTY_LIST: "[]",

	IF:     "if",
	ELSE:   "else",
	WHILE:  "while",
	RETURN: "return",
}

func (t Token) Print() string {
	s, ok := printStrings[t]
	if !ok {
		panic("Token " + t.String() + " is not printable")
	}
	return s
}
