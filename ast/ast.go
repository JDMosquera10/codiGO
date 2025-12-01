package ast

import "fmt"

type Node interface {
	Print(indent int)
}

type Program struct {
	Statements []Statement
}

func (p *Program) Print(indent int) {
	fmt.Println("Programa")
	for _, stmt := range p.Statements {
		stmt.Print(indent + 1)
	}
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Declaraciones
type DeclareStatement struct {
	IsConst   bool
	Name      *Identifier
	Value     Expression
}

func (d *DeclareStatement) statementNode() {}
func (d *DeclareStatement) Print(indent int) {
	prefix := "definir"
	if d.IsConst {
		prefix = "constante"
	}
	printIndent(indent)
	fmt.Printf("%s %s\n", prefix, d.Name.Value)
	d.Value.Print(indent + 1)
}

// Asignaciones
type AssignStatement struct {
	Name  *Identifier
	Value Expression
}

func (a *AssignStatement) statementNode() {}
func (a *AssignStatement) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Asignación: %s\n", a.Name.Value)
	a.Value.Print(indent + 1)
}

// Condicionales
type IfStatement struct {
	Condition Expression
	Then      *BlockStatement
	Else      *BlockStatement
}

func (i *IfStatement) statementNode() {}
func (i *IfStatement) Print(indent int) {
	printIndent(indent)
	fmt.Println("Si")
	i.Condition.Print(indent + 1)
	printIndent(indent)
	fmt.Println("Entonces")
	i.Then.Print(indent + 1)
	if i.Else != nil {
		printIndent(indent)
		fmt.Println("Sino")
		i.Else.Print(indent + 1)
	}
}

// Bucles
type WhileStatement struct {
	Condition Expression
	Body      *BlockStatement
}

func (w *WhileStatement) statementNode() {}
func (w *WhileStatement) Print(indent int) {
	printIndent(indent)
	fmt.Println("Mientras")
	w.Condition.Print(indent + 1)
	w.Body.Print(indent + 1)
}

type RepeatStatement struct {
	Variable *Identifier
	From     Expression
	To       Expression
	Body     *BlockStatement
}

func (r *RepeatStatement) statementNode() {}
func (r *RepeatStatement) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Repetir %s desde\n", r.Variable.Value)
	r.From.Print(indent + 1)
	printIndent(indent)
	fmt.Println("hasta")
	r.To.Print(indent + 1)
	r.Body.Print(indent + 1)
}

// Funciones
type FunctionStatement struct {
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionStatement) statementNode() {}
func (f *FunctionStatement) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Función: %s\n", f.Name.Value)
	f.Body.Print(indent + 1)
}

// Mostrar
type ShowStatement struct {
	Value Expression
}

func (s *ShowStatement) statementNode() {}
func (s *ShowStatement) Print(indent int) {
	printIndent(indent)
	fmt.Println("Mostrar")
	s.Value.Print(indent + 1)
}

// Retorno
type ReturnStatement struct {
	Value Expression
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) Print(indent int) {
	printIndent(indent)
	fmt.Println("Retornar")
	if r.Value != nil {
		r.Value.Print(indent + 1)
	}
}

// Bloques
type BlockStatement struct {
	Statements []Statement
}

func (b *BlockStatement) statementNode() {}
func (b *BlockStatement) Print(indent int) {
	for _, stmt := range b.Statements {
		stmt.Print(indent)
	}
}

// Expresiones
type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Identificador: %s\n", i.Value)
}

type IntegerLiteral struct {
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Entero: %d\n", i.Value)
}

type FloatLiteral struct {
	Value float64
}

func (f *FloatLiteral) expressionNode() {}
func (f *FloatLiteral) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Decimal: %f\n", f.Value)
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) expressionNode() {}
func (s *StringLiteral) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Cadena: %s\n", s.Value)
}

type BooleanLiteral struct {
	Value bool
}

func (b *BooleanLiteral) expressionNode() {}
func (b *BooleanLiteral) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Booleano: %v\n", b.Value)
}

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Operador: %s\n", i.Operator)
	i.Left.Print(indent + 1)
	i.Right.Print(indent + 1)
}

type PrefixExpression struct {
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) Print(indent int) {
	printIndent(indent)
	fmt.Printf("Operador: %s\n", p.Operator)
	p.Right.Print(indent + 1)
}

type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) expressionNode() {}
func (c *CallExpression) Print(indent int) {
	printIndent(indent)
	fmt.Println("Llamada función")
	c.Function.Print(indent + 1)
	for _, arg := range c.Arguments {
		arg.Print(indent + 1)
	}
}

func printIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
}

