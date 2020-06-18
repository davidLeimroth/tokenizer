package tokenizer

import (
	"search/tokenizer/lexer"
	"search/tokenizer/token"
	"strings"
)

func WordTokenizer(input string) []string {
	words := make([]string, 0)
	l := lexer.NewLexer(input)

	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		if tok.Type == token.WORD || tok.Type == token.NUMBER {
			words = append(words, tok.Literal)
		}
	}
	return words
}

func NWordTokenizer(n int) func (input string) []string {
	return func(input string) []string {
		out := WordTokenizer(input)
		length := len(out)
		for i := 0; i <= length - n; i++ {
			out = append(out, strings.Join(out[i:i+n], " "))
		}
		return out
	}
}