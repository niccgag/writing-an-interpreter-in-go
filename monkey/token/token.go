package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

var keywords = map[string]TokenType {
  "fn": FUNCTION,
  "let": LET,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers + literals
  IDENT = "INDENT" // add, forbar, x, y, ...
  INT = "INT" // 1343456

  // Operators
  ASSIGN = "="
  PLUS = "+"

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)
