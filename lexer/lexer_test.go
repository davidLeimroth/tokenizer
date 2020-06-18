package lexer

import (
	"search/tokenizer/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := ".,;(){} -+ hello world:!?=/$%&|\\[]<>\"'`_#*€"
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.PERIOD, "."},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.OPEN_PARENTHESE, "("},
		{token.CLOSE_PARENTHESE, ")"},
		{token.OPEN_BRACE, "{"},
		{token.CLOSE_BRACE, "}"},
		{token.SPACE, " "},
		{token.MINUS, "-"},
		{token.PLUS, "+"},
		{token.SPACE, " "},
		{token.WORD, "hello"},
		{token.SPACE, " "},
		{token.WORD, "world"},
		{token.COLON, ":"},
		{token.EXCLAMATION_MARK, "!"},
		{token.QUESTION_MARK, "?"},
		{token.EQUAL, "="},
		{token.SLASH, "/"},
		{token.DOLLAR, "$"},
		{token.PERCENT, "%"},
		{token.AMPERSAND, "&"},
		{token.PIPE, "|"},
		{token.BACKSLASH, "\\"},
		{token.OPEN_SQUARE_BRACKET, "["},
		{token.CLOSE_SQUARE_BRACKET, "]"},
		{token.LESS_THAN, "<"},
		{token.GREATER_THAN, ">"},
		{token.QUOTATION_MARK, "\""},
		{token.APOSTROPHE, "'"},
		{token.GRAVE_ACCENT, "`"},
		{token.UNDERSCORE, "_"},
		{token.HASHTAG, "#"},
		{token.STAR, "*"},
		{token.EURO, "€"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWords(t *testing.T) {
	input := `he wo`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.WORD, "he"},
		{token.SPACE, " "},
		{token.WORD, "wo"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWordsGerman(t *testing.T) {
	input := `Fußball törö überfluss tätärätä`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.WORD, "Fußball"},
		{token.SPACE, " "},
		{token.WORD, "törö"},
		{token.SPACE, " "},
		{token.WORD, "überfluss"},
		{token.SPACE, " "},
		{token.WORD, "tätärätä"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestInteger(t *testing.T) {
	input := `12345 4234`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.NUMBER, "12345"},
		{token.SPACE, " "},
		{token.NUMBER, "4234"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestFloat(t *testing.T) {
	input := `12.345 4.234`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.NUMBER, "12.345"},
		{token.SPACE, " "},
		{token.NUMBER, "4.234"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestFloatComma(t *testing.T) {
	input := `12,345 4,234`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.NUMBER, "12,345"},
		{token.SPACE, " "},
		{token.NUMBER, "4,234"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestIntegerWithEndingPeriodAndComma(t *testing.T) {
	input := `12345. 4234,`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.NUMBER, "12345"},
		{token.PERIOD, "."},
		{token.SPACE, " "},
		{token.NUMBER, "4234"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestFloatWithEndingPeriodAndComma(t *testing.T) {
	input := `12.345. 4.234,`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.NUMBER, "12.345"},
		{token.PERIOD, "."},
		{token.SPACE, " "},
		{token.NUMBER, "4.234"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestFloatCommaWithEndingPeriodAndComma(t *testing.T) {
	input := `12,345. 4,234,`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.NUMBER, "12,345"},
		{token.PERIOD, "."},
		{token.SPACE, " "},
		{token.NUMBER, "4,234"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestSentences(t *testing.T) {
	input := `Hello World. This is a test.`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.WORD, "Hello"},
		{token.SPACE, " "},
		{token.WORD, "World"},
		{token.PERIOD, "."},
		{token.SPACE, " "},
		{token.WORD, "This"},
		{token.SPACE, " "},
		{token.WORD, "is"},
		{token.SPACE, " "},
		{token.WORD, "a"},
		{token.SPACE, " "},
		{token.WORD, "test"},
		{token.PERIOD, "."},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, test := range tests {
		tok := l.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%v]: wrong type: expected '%v' - got '%v'", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%v]: wrong literal: expected '%v' - got '%v'", i, test.expectedLiteral, tok.Literal)
		}
	}
}
