package main

import "fmt"

type TokenType uint16

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	LBRACKET
	RBRACKET
	INT
)

type Token interface {
	Type() TokenType
	Lexeme() string
}

type Match struct {
	tokenType TokenType
	lexeme    string
}

func (m *Match) Type() TokenType {
	return m.tokenType
}

func (m *Match) Lexeme() string {
	return m.lexeme
}

/*type IntegerConstant struct {
	token Token
	value uint64
}

func (i *IntegerConstant) Type() TokenType {
	return i.token.Type()
}

func (i *IntegerConstant) Lexeme() string {
	return i.token.Lexeme()
}

func (i *IntegerConstant) Value() uint64 {
	return i.value
}*/

// **** 下面两段代码相当于上面的4段 ****
type IntegerConstant struct {
	Token
	value uint64
}

func (i *IntegerConstant) Value() uint64 {
	return i.value
}

func main() {
	t := IntegerConstant{
		Token: &Match{
			tokenType: KEYWORD,
			lexeme:    "abcde",
		},
		value: 555,
	}
	fmt.Println(t.Type(), t.Lexeme(), t.Value())
	x := Token(t)
	fmt.Println(x.Type(), x.Lexeme())
}
