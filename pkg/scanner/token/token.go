package token

const (
	Int = iota
	Add
	Sub
	Div
	Mul
	Lparen
	Rparen
	Begin
	End
	Assign
	Semi
	Dot
	Identifier
	EOF
)

// Token contains the token type and the lexeme
type Token struct {
	Type   int
	Lexeme string
}
