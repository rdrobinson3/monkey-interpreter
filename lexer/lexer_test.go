package lexer

import (
	"testing"

	"github.com/rdrobinson3/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{}
}
