package token

// TODO write godoc for the constants below
const (
	Int = iota
	Add = iota
	Sub = iota
	Div = iota
	Mul = iota
	EOF = iota
)

// Token contains the type of the token and the lexeme.
type Token struct {
	Type   int
	Lexeme string
}
