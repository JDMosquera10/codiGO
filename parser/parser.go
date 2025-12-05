package parser

import (
	"fmt"
	"flux/ast"
	"flux/lexer"
	"strconv"
	"strings"
)

type Parser struct {
	tokens       []lexer.Token
	position     int
	currentToken lexer.Token
	errors       []string
}

func New(tokens []lexer.Token) *Parser {
	p := &Parser{
		tokens: tokens,
		position: 0,
		errors: []string{},
	}
	
	if len(tokens) > 0 {
		p.currentToken = tokens[0]
	} else {
		p.currentToken = lexer.Token{Type: lexer.TOKEN_EOF}
	}
	
	return p
}

func (p *Parser) nextToken() {
	p.position++
	if p.position < len(p.tokens) {
		p.currentToken = p.tokens[p.position]
	} else {
		p.currentToken = lexer.Token{Type: lexer.TOKEN_EOF}
	}
}

func (p *Parser) peekToken() lexer.Token {
	if p.position+1 < len(p.tokens) {
		return p.tokens[p.position+1]
	}
	return lexer.Token{Type: lexer.TOKEN_EOF}
}

func (p *Parser) Parse() (*ast.Program, error) {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}
	
	for p.currentToken.Type != lexer.TOKEN_EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		} else {
			// Si no se pudo parsear la sentencia, avanzar para evitar bucle infinito
			if p.currentToken.Type != lexer.TOKEN_EOF {
				p.nextToken()
			}
		}
	}
	
	if len(p.errors) > 0 {
		return nil, fmt.Errorf("errores de parsing: %v", p.errors)
	}
	
	return program, nil
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case lexer.TOKEN_DEFINIR, lexer.TOKEN_CONSTANTE:
		return p.parseDeclareStatement()
	case lexer.TOKEN_SI:
		return p.parseIfStatement()
	case lexer.TOKEN_MIENTRAS:
		return p.parseWhileStatement()
	case lexer.TOKEN_REPETIR:
		return p.parseRepeatStatement()
	case lexer.TOKEN_FUNCION:
		return p.parseFunctionStatement()
	case lexer.TOKEN_MOSTRAR:
		return p.parseShowStatement()
	case lexer.TOKEN_RETORNAR:
		return p.parseReturnStatement()
	case lexer.TOKEN_IDENTIFICADOR:
		// Podría ser una llamada a función
		if p.peekToken().Type == lexer.TOKEN_PARENTESIS_IZQ {
			expr := p.parseCallExpression()
			if expr != nil {
				return &ast.ExpressionStatement{Expression: expr}
			}
		}
		// Si no es una llamada, intentar parsear como expresión
		expr := p.parseExpression(0)
		if expr != nil {
			return &ast.ExpressionStatement{Expression: expr}
		}
		// Si no se pudo parsear como expresión, podría ser un error
		p.errors = append(p.errors, fmt.Sprintf("línea %d, columna %d: no se pudo parsear el identificador '%s' como una sentencia válida", 
			p.currentToken.Line, p.currentToken.Column, p.currentToken.Value))
		return nil
	default:
		// Intentar parsear como expresión (para casos como llamadas a función)
		expr := p.parseExpression(0)
		if expr != nil {
			return &ast.ExpressionStatement{Expression: expr}
		}
		// Si no se pudo parsear, es un error
		if p.currentToken.Type != lexer.TOKEN_EOF {
			p.errors = append(p.errors, fmt.Sprintf("línea %d, columna %d: token inesperado '%s' (tipo: %s)", 
				p.currentToken.Line, p.currentToken.Column, p.currentToken.Value, p.currentToken.Type))
		}
		return nil
	}
}

func (p *Parser) parseDeclareStatement() *ast.DeclareStatement {
	stmt := &ast.DeclareStatement{
		IsConst: p.currentToken.Type == lexer.TOKEN_CONSTANTE,
	}
	
	p.nextToken()
	
	if p.currentToken.Type != lexer.TOKEN_IDENTIFICADOR {
		p.errors = append(p.errors, fmt.Sprintf("línea %d, columna %d: se esperaba un identificador después de 'definir', pero se encontró '%s' (tipo: %s)", 
			p.currentToken.Line, p.currentToken.Column, p.currentToken.Value, p.currentToken.Type))
		return nil
	}
	
	stmt.Name = &ast.Identifier{Value: p.currentToken.Value}
	p.nextToken()
	
	if p.currentToken.Type != lexer.TOKEN_ASIGNACION {
		expected := "="
		found := p.currentToken.Value
		if found == "" {
			found = string(p.currentToken.Type)
		}
		line := p.currentToken.Line
		column := p.currentToken.Column
		if line == 0 {
			line = 1 // Valor por defecto si no hay información de línea
		}
		if column == 0 {
			column = 1 // Valor por defecto si no hay información de columna
		}
		p.errors = append(p.errors, fmt.Sprintf("línea %d, columna %d: se esperaba '%s' después del identificador '%s', pero se encontró '%s' (tipo: %s)", 
			line, column, expected, stmt.Name.Value, found, p.currentToken.Type))
		return nil
	}
	
	p.nextToken()
	stmt.Value = p.parseExpression(0)
	
	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{}
	
	p.nextToken()
	stmt.Condition = p.parseExpression(0)
	
	if p.currentToken.Type == lexer.TOKEN_ENTONCES || p.currentToken.Type == lexer.TOKEN_HACER {
		p.nextToken()
	}
	
	stmt.Then = p.parseBlockStatement()
	
	if p.currentToken.Type == lexer.TOKEN_SINO {
		p.nextToken()
		stmt.Else = p.parseBlockStatement()
	}
	
	if p.currentToken.Type == lexer.TOKEN_FIN {
		p.nextToken()
	}
	
	return stmt
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{}
	
	p.nextToken()
	stmt.Condition = p.parseExpression(0)
	
	if p.currentToken.Type == lexer.TOKEN_HACER {
		p.nextToken()
	}
	
	stmt.Body = p.parseBlockStatement()
	
	if p.currentToken.Type == lexer.TOKEN_FIN {
		p.nextToken()
	}
	
	return stmt
}

func (p *Parser) parseRepeatStatement() *ast.RepeatStatement {
	stmt := &ast.RepeatStatement{}
	
	p.nextToken()
	stmt.Variable = &ast.Identifier{Value: p.currentToken.Value}
	p.nextToken()
	
	if p.currentToken.Type != lexer.TOKEN_DESDE {
		p.errors = append(p.errors, "se esperaba 'desde'")
		return nil
	}
	
	p.nextToken()
	stmt.From = p.parseExpression(0)
	
	if p.currentToken.Type != lexer.TOKEN_HASTA {
		p.errors = append(p.errors, fmt.Sprintf("se esperaba 'hasta', pero se encontró '%s' (tipo: %s)", p.currentToken.Value, p.currentToken.Type))
		return nil
	}
	
	p.nextToken()
	stmt.To = p.parseExpression(0)
	
	if p.currentToken.Type == lexer.TOKEN_HACER {
		p.nextToken()
	}
	
	stmt.Body = p.parseBlockStatement()
	
	if p.currentToken.Type == lexer.TOKEN_FIN {
		p.nextToken()
	}
	
	return stmt
}

func (p *Parser) parseFunctionStatement() *ast.FunctionStatement {
	stmt := &ast.FunctionStatement{}
	
	p.nextToken()
	stmt.Name = &ast.Identifier{Value: p.currentToken.Value}
	p.nextToken()
	
	if p.currentToken.Type == lexer.TOKEN_PARENTESIS_IZQ {
		p.nextToken()
		stmt.Parameters = p.parseParameters()
		if p.currentToken.Type == lexer.TOKEN_PARENTESIS_DER {
			p.nextToken()
		}
	}
	
	if p.currentToken.Type == lexer.TOKEN_HACER {
		p.nextToken()
	}
	
	stmt.Body = p.parseBlockStatement()
	
	if p.currentToken.Type == lexer.TOKEN_FIN {
		p.nextToken()
	}
	
	return stmt
}

func (p *Parser) parseParameters() []*ast.Identifier {
	var params []*ast.Identifier
	
	if p.currentToken.Type == lexer.TOKEN_IDENTIFICADOR {
		params = append(params, &ast.Identifier{Value: p.currentToken.Value})
		p.nextToken()
		
		for p.currentToken.Type == lexer.TOKEN_COMA {
			p.nextToken()
			if p.currentToken.Type == lexer.TOKEN_IDENTIFICADOR {
				params = append(params, &ast.Identifier{Value: p.currentToken.Value})
				p.nextToken()
			}
		}
	}
	
	return params
}

func (p *Parser) parseShowStatement() *ast.ShowStatement {
	stmt := &ast.ShowStatement{}
	
	p.nextToken()
	if p.currentToken.Type == lexer.TOKEN_PARENTESIS_IZQ {
		p.nextToken()
		stmt.Value = p.parseExpression(0)
		if p.currentToken.Type == lexer.TOKEN_PARENTESIS_DER {
			p.nextToken()
		}
	} else {
		stmt.Value = p.parseExpression(0)
	}
	
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{}
	
	p.nextToken()
	if p.currentToken.Type != lexer.TOKEN_FIN {
		stmt.Value = p.parseExpression(0)
	}
	
	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Statements: []ast.Statement{},
	}
	
	for p.currentToken.Type != lexer.TOKEN_FIN && p.currentToken.Type != lexer.TOKEN_SINO && p.currentToken.Type != lexer.TOKEN_EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		} else {
			// Si no se pudo parsear, avanzar para evitar bucle infinito
			if p.currentToken.Type != lexer.TOKEN_FIN && p.currentToken.Type != lexer.TOKEN_SINO && p.currentToken.Type != lexer.TOKEN_EOF {
				p.nextToken()
			}
		}
	}
	
	return block
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	left := p.parsePrefixExpression()
	if left == nil {
		return nil
	}
	
	for p.currentToken.Type != lexer.TOKEN_EOF && 
		p.currentToken.Type != lexer.TOKEN_PARENTESIS_DER &&
		p.currentToken.Type != lexer.TOKEN_FIN &&
		p.currentToken.Type != lexer.TOKEN_SINO &&
		p.currentToken.Type != lexer.TOKEN_ENTONCES &&
		p.currentToken.Type != lexer.TOKEN_HACER &&
		p.currentToken.Type != lexer.TOKEN_COMA &&
		precedence < p.currentPrecedence() {
		// No avanzamos aquí porque currentToken ya es el operador infijo
		left = p.parseInfixExpression(left)
		if left == nil {
			return nil
		}
	}
	
	return left
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	switch p.currentToken.Type {
	case lexer.TOKEN_NOT:
		expr := &ast.PrefixExpression{
			Operator: p.currentToken.Value,
		}
		p.nextToken()
		expr.Right = p.parseExpression(1)
		return expr
	case lexer.TOKEN_IDENTIFICADOR:
		// Verificar si es una llamada a función
		if p.peekToken().Type == lexer.TOKEN_PARENTESIS_IZQ {
			return p.parseCallExpression()
		}
		ident := &ast.Identifier{Value: p.currentToken.Value}
		p.nextToken()
		return ident
	case lexer.TOKEN_ENTERO:
		val, _ := strconv.ParseInt(p.currentToken.Value, 10, 64)
		lit := &ast.IntegerLiteral{Value: val}
		p.nextToken()
		return lit
	case lexer.TOKEN_DECIMAL:
		val, _ := strconv.ParseFloat(p.currentToken.Value, 64)
		lit := &ast.FloatLiteral{Value: val}
		p.nextToken()
		return lit
	case lexer.TOKEN_CADENA:
		val := strings.Trim(p.currentToken.Value, "\"'")
		lit := &ast.StringLiteral{Value: val}
		p.nextToken()
		return lit
	case lexer.TOKEN_VERDADERO:
		lit := &ast.BooleanLiteral{Value: true}
		p.nextToken()
		return lit
	case lexer.TOKEN_FALSO:
		lit := &ast.BooleanLiteral{Value: false}
		p.nextToken()
		return lit
	case lexer.TOKEN_PARENTESIS_IZQ:
		p.nextToken()
		expr := p.parseExpression(0)
		if p.currentToken.Type == lexer.TOKEN_PARENTESIS_DER {
			p.nextToken()
		}
		return expr
	default:
		return nil
	}
}

func (p *Parser) parseCallExpression() ast.Expression {
	expr := &ast.CallExpression{
		Function: &ast.Identifier{Value: p.currentToken.Value},
	}
	p.nextToken() // Consumir el identificador
	
	if p.currentToken.Type != lexer.TOKEN_PARENTESIS_IZQ {
		return nil
	}
	
	p.nextToken() // Consumir el paréntesis izquierdo
	
	// Parsear argumentos
	expr.Arguments = []ast.Expression{}
	if p.currentToken.Type != lexer.TOKEN_PARENTESIS_DER {
		expr.Arguments = append(expr.Arguments, p.parseExpression(0))
		
		for p.currentToken.Type == lexer.TOKEN_COMA {
			p.nextToken()
			expr.Arguments = append(expr.Arguments, p.parseExpression(0))
		}
	}
	
	if p.currentToken.Type == lexer.TOKEN_PARENTESIS_DER {
		p.nextToken()
	}
	
	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Left:     left,
		Operator: p.currentToken.Value,
	}
	
	precedence := p.currentPrecedence()
	p.nextToken()
	// Para operadores de izquierda a derecha, usar precedence - 1 para que operadores de mayor precedencia
	// se agrupen primero. Esto asegura que n % i == 0 se parsea como (n % i) == 0
	expr.Right = p.parseExpression(precedence - 1)
	
	return expr
}

func (p *Parser) currentPrecedence() int {
	switch p.currentToken.Type {
	case lexer.TOKEN_SUMA, lexer.TOKEN_RESTA:
		return 3
	case lexer.TOKEN_MULTIPLICAR, lexer.TOKEN_DIVIDIR, lexer.TOKEN_MODULO:
		return 4
	case lexer.TOKEN_MENOR, lexer.TOKEN_MAYOR, lexer.TOKEN_MENOR_IGUAL, lexer.TOKEN_MAYOR_IGUAL:
		return 5
	case lexer.TOKEN_IGUAL, lexer.TOKEN_DIFERENTE:
		return 6
	case lexer.TOKEN_AND:
		return 7
	case lexer.TOKEN_OR:
		return 8
	default:
		return 0
	}
}

func (p *Parser) peekPrecedence() int {
	peek := p.peekToken()
	switch peek.Type {
	case lexer.TOKEN_SUMA, lexer.TOKEN_RESTA:
		return 3
	case lexer.TOKEN_MULTIPLICAR, lexer.TOKEN_DIVIDIR, lexer.TOKEN_MODULO:
		return 4
	case lexer.TOKEN_MENOR, lexer.TOKEN_MAYOR, lexer.TOKEN_MENOR_IGUAL, lexer.TOKEN_MAYOR_IGUAL:
		return 5
	case lexer.TOKEN_IGUAL, lexer.TOKEN_DIFERENTE:
		return 6
	case lexer.TOKEN_AND:
		return 7
	case lexer.TOKEN_OR:
		return 8
	default:
		return 0
	}
}

