package lexer

import (
	"fmt"
	"search/tokenizer/token"
)

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input: []rune(input),
	}
	lexer.readRune()
	return lexer
}

type Lexer struct {
	input        []rune
	position     int
	nextPosition int
	currentRune  rune
}

func (l *Lexer) String() string {
	return fmt.Sprintf("%v@%v", string(l.currentRune), l.position)
}

func (l *Lexer) NextToken() *token.Token {
	var tok *token.Token

	switch l.currentRune {
	case '.':
		tok = token.NewToken(token.PERIOD, string(l.currentRune))
	case ',':
		tok = token.NewToken(token.COMMA, string(l.currentRune))
	case ';':
		tok = token.NewToken(token.SEMICOLON, string(l.currentRune))
	case ':':
		tok = token.NewToken(token.COLON, string(l.currentRune))
	case '(':
		tok = token.NewToken(token.OPEN_PARENTHESE, string(l.currentRune))
	case ')':
		tok = token.NewToken(token.CLOSE_PARENTHESE, string(l.currentRune))
	case '{':
		tok = token.NewToken(token.OPEN_BRACE, string(l.currentRune))
	case '}':
		tok = token.NewToken(token.CLOSE_BRACE, string(l.currentRune))
	case ' ':
		tok = token.NewToken(token.SPACE, string(l.currentRune))
	case '!':
		tok = token.NewToken(token.EXCLAMATION_MARK, string(l.currentRune))
	case '?':
		tok = token.NewToken(token.QUESTION_MARK, string(l.currentRune))
	case '-':
		tok = token.NewToken(token.MINUS, string(l.currentRune))
	case '+':
		tok = token.NewToken(token.PLUS, string(l.currentRune))
	case '=':
		tok = token.NewToken(token.EQUAL, string(l.currentRune))
	case '/':
		tok = token.NewToken(token.SLASH, string(l.currentRune))
	case '$':
		tok = token.NewToken(token.DOLLAR, string(l.currentRune))
	case '%':
		tok = token.NewToken(token.PERCENT, string(l.currentRune))
	case '&':
		tok = token.NewToken(token.AMPERSAND, string(l.currentRune))
	case '|':
		tok = token.NewToken(token.PIPE, string(l.currentRune))
	case '\\':
		tok = token.NewToken(token.BACKSLASH, string(l.currentRune))
	case '[':
		tok = token.NewToken(token.OPEN_SQUARE_BRACKET, string(l.currentRune))
	case ']':
		tok = token.NewToken(token.CLOSE_SQUARE_BRACKET, string(l.currentRune))
	case '<':
		tok = token.NewToken(token.LESS_THAN, string(l.currentRune))
	case '>':
		tok = token.NewToken(token.GREATER_THAN, string(l.currentRune))
	case '"':
		tok = token.NewToken(token.QUOTATION_MARK, string(l.currentRune))
	case '\'':
		tok = token.NewToken(token.APOSTROPHE, string(l.currentRune))
	case '`':
		tok = token.NewToken(token.GRAVE_ACCENT, string(l.currentRune))
	case '_':
		tok = token.NewToken(token.UNDERSCORE, string(l.currentRune))
	case '#':
		tok = token.NewToken(token.HASHTAG, string(l.currentRune))
	case '*':
		tok = token.NewToken(token.STAR, string(l.currentRune))
	case '€':
		tok = token.NewToken(token.EURO, string(l.currentRune))
	case 0:
		tok = token.NewToken(token.EOF, "")
	default:
		if isLetter(l.currentRune) {
			return token.NewToken(token.WORD, string(l.readWord()))
		}

		if isDigit(l.currentRune, l.peekRune()) {
			return token.NewToken(token.NUMBER, string(l.readNumber()))
		}
		tok = token.NewToken(token.ILLEGAL, string(l.currentRune))
	}

	l.readRune()
	return tok
}

func (l *Lexer) readWord() []rune {
	pos := l.position
	for isLetter(l.currentRune) {
		l.readRune()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() []rune {
	pos := l.position
	for isDigit(l.currentRune, l.peekRune()) {
		l.readRune()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readRune() {
	if l.nextPosition >= len(l.input) {
		l.currentRune = 0
	} else {
		l.currentRune = l.input[l.nextPosition]
	}

	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) peekRune() rune {
	if l.nextPosition >= len(l.input) {
		return 0
	}
	return l.input[l.nextPosition]
}
