package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type TokenType string

const (
	// Palabras reservadas
	TOKEN_DEFINIR   TokenType = "DEFINIR"
	TOKEN_CONSTANTE TokenType = "CONSTANTE"
	TOKEN_FUNCION   TokenType = "FUNCION"
	TOKEN_SI        TokenType = "SI"
	TOKEN_ENTONCES  TokenType = "ENTONCES"
	TOKEN_SINO      TokenType = "SINO"
	TOKEN_MIENTRAS  TokenType = "MIENTRAS"
	TOKEN_REPETIR   TokenType = "REPETIR"
	TOKEN_DESDE     TokenType = "DESDE"
	TOKEN_HASTA     TokenType = "HASTA"
	TOKEN_HACER     TokenType = "HACER"
	TOKEN_FIN       TokenType = "FIN"
	TOKEN_MOSTRAR   TokenType = "MOSTRAR"
	TOKEN_RETORNAR  TokenType = "RETORNAR"
	TOKEN_SALIR     TokenType = "SALIR"
	TOKEN_CONTINUAR TokenType = "CONTINUAR"
	TOKEN_VERDADERO TokenType = "VERDADERO"
	TOKEN_FALSO     TokenType = "FALSO"
	TOKEN_NULO      TokenType = "NULO"

	// Operadores
	TOKEN_ASIGNACION  TokenType = "ASIGNACION"  // =
	TOKEN_IGUAL       TokenType = "IGUAL"       // ==
	TOKEN_DIFERENTE   TokenType = "DIFERENTE"   // !=
	TOKEN_MENOR_IGUAL TokenType = "MENOR_IGUAL" // <=
	TOKEN_MAYOR_IGUAL TokenType = "MAYOR_IGUAL" // >=
	TOKEN_AND         TokenType = "AND"         // &&
	TOKEN_OR          TokenType = "OR"          // ||
	TOKEN_NOT         TokenType = "NOT"         // !
	TOKEN_SUMA        TokenType = "SUMA"        // +
	TOKEN_RESTA       TokenType = "RESTA"       // -
	TOKEN_MULTIPLICAR TokenType = "MULTIPLICAR" // *
	TOKEN_DIVIDIR     TokenType = "DIVIDIR"     // /
	TOKEN_MODULO      TokenType = "MODULO"      // %
	TOKEN_MENOR       TokenType = "MENOR"       // <
	TOKEN_MAYOR       TokenType = "MAYOR"       // >

	// Delimitadores
	TOKEN_PARENTESIS_IZQ TokenType = "PARENTESIS_IZQ" // (
	TOKEN_PARENTESIS_DER TokenType = "PARENTESIS_DER" // )
	TOKEN_LLAVE_IZQ      TokenType = "LLAVE_IZQ"      // {
	TOKEN_LLAVE_DER      TokenType = "LLAVE_DER"      // }
	TOKEN_COMA           TokenType = "COMA"           // ,
	TOKEN_PUNTO          TokenType = "PUNTO"          // ·

	// Literales
	TOKEN_ENTERO   TokenType = "ENTERO"
	TOKEN_DECIMAL  TokenType = "DECIMAL"
	TOKEN_CADENA   TokenType = "CADENA"
	TOKEN_BOOLEANO TokenType = "BOOLEANO"

	// Identificadores
	TOKEN_IDENTIFICADOR TokenType = "IDENTIFICADOR"

	// Especiales
	TOKEN_EOF     TokenType = "EOF"
	TOKEN_ILLEGAL TokenType = "ILLEGAL"
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
	line         int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 1,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		l.position = l.readPosition
		return
	}

	r, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
	if r == utf8.RuneError && size == 0 {
		l.ch = 0
		l.position = l.readPosition
		return
	}

	l.ch = r
	l.position = l.readPosition
	l.readPosition += size

	if l.ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return r
}

func (l *Lexer) skipWhitespace() {
	for {
		// Saltar espacios en blanco
		for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
			l.readChar()
		}

		// Saltar comentarios de línea
		if l.ch == '/' && l.peekChar() == '/' {
			for l.ch != '\n' && l.ch != 0 {
				l.readChar()
			}
			if l.ch == '\n' {
				l.readChar()
			}
			continue
		}

		// Saltar comentarios de bloque
		if l.ch == '/' && l.peekChar() == '*' {
			l.readChar() // consume '/'
			l.readChar() // consume '*'
			for {
				if l.ch == 0 {
					return
				}
				if l.ch == '*' && l.peekChar() == '/' {
					l.readChar() // consume '*'
					l.readChar() // consume '/'
					break
				}
				l.readChar()
			}
			continue
		}

		break
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() (string, error) {
	quote := l.ch
	l.readChar()
	position := l.position

	for l.ch != quote && l.ch != 0 {
		if l.ch == '\\' {
			l.readChar()
		}
		l.readChar()
	}

	if l.ch == 0 {
		return "", fmt.Errorf("cadena no cerrada en línea %d", l.line)
	}

	l.readChar()
	return l.input[position-1 : l.position], nil
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	tok.Line = l.line
	tok.Column = l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TOKEN_IGUAL, Value: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: TOKEN_ASIGNACION, Value: "="}
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TOKEN_DIFERENTE, Value: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: TOKEN_NOT, Value: "!"}
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TOKEN_MENOR_IGUAL, Value: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: TOKEN_MENOR, Value: "<"}
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TOKEN_MAYOR_IGUAL, Value: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: TOKEN_MAYOR, Value: ">"}
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TOKEN_AND, Value: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: TOKEN_ILLEGAL, Value: string(l.ch)}
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TOKEN_OR, Value: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: TOKEN_ILLEGAL, Value: string(l.ch)}
		}
	case '+':
		tok = Token{Type: TOKEN_SUMA, Value: "+"}
	case '-':
		tok = Token{Type: TOKEN_RESTA, Value: "-"}
	case '*':
		tok = Token{Type: TOKEN_MULTIPLICAR, Value: "*"}
	case '/':
		tok = Token{Type: TOKEN_DIVIDIR, Value: "/"}
	case '%':
		tok = Token{Type: TOKEN_MODULO, Value: "%"}
	case '(':
		tok = Token{Type: TOKEN_PARENTESIS_IZQ, Value: "("}
	case ')':
		tok = Token{Type: TOKEN_PARENTESIS_DER, Value: ")"}
	case '{':
		tok = Token{Type: TOKEN_LLAVE_IZQ, Value: "{"}
	case '}':
		tok = Token{Type: TOKEN_LLAVE_DER, Value: "}"}
	case ',':
		tok = Token{Type: TOKEN_COMA, Value: ","}
	case '·':
		tok = Token{Type: TOKEN_PUNTO, Value: "·"}
	case '"', '\'':
		str, err := l.readString()
		if err != nil {
			tok = Token{Type: TOKEN_ILLEGAL, Value: err.Error()}
			return tok
		}
		tok = Token{Type: TOKEN_CADENA, Value: str}
		return tok
	case 0:
		tok = Token{Type: TOKEN_EOF, Value: ""}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tokType := lookupIdent(ident)
			if tokType == TOKEN_ILLEGAL {
				// Es un error de palabra clave mal escrita
				suggestion := getKeywordSuggestion(ident)
				tok = Token{
					Type:   TOKEN_ILLEGAL,
					Value:  fmt.Sprintf("palabra clave incorrecta '%s' (¿quisiste decir '%s'?)", ident, suggestion),
					Line:   l.line,
					Column: l.column,
				}
				return tok
			}
			tok.Type = tokType
			tok.Value = ident
			return tok
		} else if isDigit(l.ch) {
			num := l.readNumber()
			if strings.Contains(num, ".") {
				tok.Type = TOKEN_DECIMAL
			} else {
				tok.Type = TOKEN_ENTERO
			}
			tok.Value = num
			return tok
		} else if l.ch == ':' {
			// Ignorar dos puntos (puede ser parte de comentarios o sintaxis)
			l.readChar()
			return l.NextToken()
		} else if l.ch > 127 {
			// Caracteres Unicode no reconocidos (probablemente en comentarios o cadenas)
			// Intentar leer como parte de una cadena o ignorar
			l.readChar()
			return l.NextToken()
		} else if l.ch != 0 {
			tok = Token{Type: TOKEN_ILLEGAL, Value: string(l.ch)}
		} else {
			tok = Token{Type: TOKEN_EOF, Value: ""}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) Tokenize() ([]Token, error) {
	var tokens []Token

	for {
		tok := l.NextToken()
		if tok.Type == TOKEN_ILLEGAL {
			// El mensaje de error ya está formateado en el token
			return nil, fmt.Errorf("línea %d, columna %d: %s", tok.Line, tok.Column, tok.Value)
		}
		tokens = append(tokens, tok)
		if tok.Type == TOKEN_EOF {
			break
		}
	}

	return tokens, nil
}

func lookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"definir":    TOKEN_DEFINIR,
		"constante":  TOKEN_CONSTANTE,
		"función":    TOKEN_FUNCION,
		"funcion":    TOKEN_FUNCION,
		"si":         TOKEN_SI,
		"entonces":   TOKEN_ENTONCES,
		"sino":       TOKEN_SINO,
		"mientras":   TOKEN_MIENTRAS,
		"repetir":    TOKEN_REPETIR,
		"desde":      TOKEN_DESDE,
		"hasta":      TOKEN_HASTA,
		"hacer":      TOKEN_HACER,
		"fin":        TOKEN_FIN,
		"mostrar":    TOKEN_MOSTRAR,
		"retornar":   TOKEN_RETORNAR,
		"salir":      TOKEN_SALIR,
		"continuar":  TOKEN_CONTINUAR,
		"verdadero":  TOKEN_VERDADERO,
		"falso":      TOKEN_FALSO,
		"true":       TOKEN_VERDADERO,  // Soporte para inglés
		"false":      TOKEN_FALSO,      // Soporte para inglés
		"nulo":       TOKEN_NULO,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	// Detectar errores comunes de escritura en palabras clave
	commonMistakes := map[string]string{
		"defenir":   "definir",
		"definr":    "definir",
		"defnir":    "definir",
		"defini":    "definir",
		"definor":   "definir",
		"constnte":  "constante",
		"constatne": "constante",
		"funcio":    "función",
		"funcion":   "función",
		"entoces":   "entonces",
		"entonces":  "entonces",
		"mientas":   "mientras",
		"repetr":    "repetir",
		"repeti":    "repetir",
		"mostar":    "mostrar",
		"mostrr":    "mostrar",
		"retornr":   "retornar",
		"retorna":   "retornar",
	}

	if _, ok := commonMistakes[strings.ToLower(ident)]; ok {
		// Retornar un token especial que indique un error de palabra clave
		return TOKEN_ILLEGAL
	}

	return TOKEN_IDENTIFICADOR
}

func getKeywordSuggestion(ident string) string {
	suggestions := map[string]string{
		"defenir":   "definir",
		"definr":    "definir",
		"defnir":    "definir",
		"defini":    "definir",
		"constnte":  "constante",
		"constatne": "constante",
		"funcio":    "función",
		"entoces":   "entonces",
		"mientas":   "mientras",
		"repetr":    "repetir",
		"repeti":    "repetir",
		"mostar":    "mostrar",
		"mostrr":    "mostrar",
		"retornr":   "retornar",
		"retorna":   "retornar",
	}

	if suggestion, ok := suggestions[strings.ToLower(ident)]; ok {
		return suggestion
	}
	return "definir" // Por defecto
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' ||
		(ch >= 0x00C0 && ch <= 0x024F) || // Latin Extended-A y B (incluye acentos)
		(ch >= 0x1E00 && ch <= 0x1EFF) // Latin Extended Additional
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
