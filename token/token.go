package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// トークンや文字が未知であること
	ILLEGAL = "ILLEGAL"
	// ファイル終端
	EOF = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミター
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
