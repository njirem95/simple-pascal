package token

const (
	Int    = iota
	Add    = iota
	Sub    = iota
	Div    = iota
	Mul    = iota
	Lparen = iota
	Rparen = iota
	EOF    = iota
)

// Token contains the token type and the lexeme
type Token struct {
	Type   int
	Lexeme string
}
