package scanner

import (
	"fmt"

	"github.com/Minnozz/gospl/token"
)

type ErrorHandler func(pos token.Position, msg string)

type Scanner struct {
	fileInfo     *token.FileInfo
	src          []byte
	errorHandler ErrorHandler

	ch     byte
	offset int

	ErrorCount int
}

func (s *Scanner) Init(fileInfo *token.FileInfo, src []byte, errorHandler ErrorHandler) {
	s.fileInfo = fileInfo
	s.src = src
	s.errorHandler = errorHandler

	s.ch = 0
	s.offset = -1 // Advance to 0 with first call to next()

	s.ErrorCount = 0

	// Read first byte
	s.next()
}

func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string) {
	s.skipWhitespace()

	// Record start of token
	pos = token.Pos(s.offset)

	// Determine token by looking at the first character
	switch ch := s.ch; {
	case s.offset >= len(s.src):
		tok = token.EOF
	case isAlpha(ch):
		lit = s.scanWord()
		tok, lit = token.LookupWord(lit)
	case isDigit(ch):
		tok = token.INTEGER
		lit = s.scanNumber()
	default:
		// First advance to the next character
		s.next()
		// Then look at the current(/previous) character
		switch ch {
		case '+':
			tok = token.PLUS
		case '-':
			tok = token.MINUS
		case '*':
			tok = token.MULTIPLY
		case '/':
			if s.ch == '/' || s.ch == '*' {
				tok = token.COMMENT
				lit = s.scanComment()
			} else {
				tok = token.DIVIDE
			}
		case '%':
			tok = token.MODULO
		case '&':
			tok = s.expect('&', token.AND)
		case '|':
			tok = s.expect('|', token.OR)
		case '=':
			tok = s.try('=', token.EQUALS, token.IS)
		case '<':
			tok = s.try('=', token.LESS_THAN_EQUALS, token.LESS_THAN)
		case '>':
			tok = s.try('=', token.GREATER_THAN_EQUALS, token.GREATER_THAN)
		case '!':
			tok = s.try('=', token.NOT_EQUALS, token.NOT)
		case ',':
			tok = token.COMMA
		case ';':
			tok = token.SEMICOLON
		case ':':
			tok = token.COLON
		case '(':
			tok = token.ROUND_BRACKET_OPEN
		case ')':
			tok = token.ROUND_BRACKET_CLOSE
		case '{':
			tok = token.CURLY_BRACKET_OPEN
		case '}':
			tok = token.CURLY_BRACKET_CLOSE
		case '[':
			tok = s.try(']', token.EMPTY_LIST, token.SQUARE_BRACKET_OPEN)
			if tok == token.EMPTY_LIST {
				lit = "[]"
			}
		case ']':
			tok = token.SQUARE_BRACKET_CLOSE
		default:
			s.error(s.offset, fmt.Sprintf("illegal character %+q", ch))
			tok, lit = token.INVALID, string(ch)
		}
	}

	return pos, tok, lit
}

func (s *Scanner) error(offset int, msg string) {
	if s.errorHandler != nil {
		s.errorHandler(s.fileInfo.Position(token.Pos(offset)), msg)
	}
	s.ErrorCount++
}

func (s *Scanner) next() {
	s.offset++
	if s.offset < len(s.src) {
		s.ch = s.src[s.offset]
		if s.ch == '\n' {
			s.fileInfo.AddLine(s.offset)
		}
	} else {
		s.ch = 0
	}
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) scanWord() string {
	start := s.offset
	for isWord(s.ch) {
		s.next()
	}
	return string(s.src[start:s.offset])
}

func (s *Scanner) scanNumber() string {
	start := s.offset
	for isDigit(s.ch) {
		s.next()
	}
	return string(s.src[start:s.offset])
}

func (s *Scanner) scanComment() string {
	// Initial '/' has been consumed; s.ch is the next character
	start := s.offset - 1

	if s.ch == '/' {
		// Line comment
		s.next()
		for s.ch != '\n' && s.ch != 0 {
			s.next()
		}
		goto ok
	}

	// Block comment
	s.next()
	for s.ch != 0 {
		ch := s.ch
		s.next()
		if ch == '*' && s.ch == '/' {
			s.next()
			goto ok
		}
	}
	s.error(start, "block comment not terminated")

ok:
	return string(s.src[start:s.offset])
}

func (s *Scanner) expect(ch byte, match token.Token) token.Token {
	if s.ch == ch {
		s.next()
		return match
	}
	s.error(s.offset, fmt.Sprintf("expected %+q to scan %v, got %+q", ch, match, s.ch))
	return token.INVALID
}

func (s *Scanner) try(ch byte, match, mismatch token.Token) token.Token {
	if s.ch == ch {
		s.next()
		return match
	}
	return mismatch
}
