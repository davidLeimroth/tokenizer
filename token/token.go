package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

func NewToken(t Type, literal string) *Token {
	return &Token{t, literal}
}

const (
	ILLEGAL Type = "ILLEGAL"
	EOF          = "EOF"

	WORD   = "WORD"
	NUMBER = "NUMBER"

	PERIOD               = "."
	COMMA                = ","
	SEMICOLON            = ";"
	COLON                = ":"
	EQUAL                = "="
	SLASH                = "/"
	DOLLAR               = "$"
	EURO              	 = "€"
	PERCENT              = "%"
	AMPERSAND            = "&"
	PIPE                 = "|"
	BACKSLASH            = "\\"
	OPEN_SQUARE_BRACKET  = "["
	CLOSE_SQUARE_BRACKET = "]"
	LESS_THAN			 = "<"
	GREATER_THAN		 = ">"
	OPEN_PARENTHESE      = "("
	CLOSE_PARENTHESE     = ")"
	OPEN_BRACE           = "{"
	CLOSE_BRACE          = "}"
	SPACE                = " "
	EXCLAMATION_MARK     = "!"
	QUESTION_MARK        = "?"
	QUOTATION_MARK		 = "\""
	APOSTROPHE			 = "'"
	GRAVE_ACCENT		 = "`"
	UNDERSCORE			 = "_"

	MINUS = "-"
	PLUS  = "+"
	HASHTAG  = "#"
	STAR  = "*"
)
