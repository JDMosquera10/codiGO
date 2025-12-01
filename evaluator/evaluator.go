package evaluator

import (
	"fmt"
	"flux/ast"
	"flux/symbol"
)

type Evaluator struct {
	symbolTable *symbol.Table
}

func New(symbolTable *symbol.Table) *Evaluator {
	return &Evaluator{
		symbolTable: symbolTable,
	}
}

func (e *Evaluator) Evaluate(node ast.Node) error {
	switch n := node.(type) {
	case *ast.Program:
		return e.evaluateProgram(n)
	case *ast.DeclareStatement:
		return e.evaluateDeclareStatement(n)
	case *ast.AssignStatement:
		return e.evaluateAssignStatement(n)
	case *ast.IfStatement:
		return e.evaluateIfStatement(n)
	case *ast.WhileStatement:
		return e.evaluateWhileStatement(n)
	case *ast.RepeatStatement:
		return e.evaluateRepeatStatement(n)
	case *ast.ShowStatement:
		return e.evaluateShowStatement(n)
	case *ast.ReturnStatement:
		return e.evaluateReturnStatement(n)
	case *ast.BlockStatement:
		return e.evaluateBlockStatement(n)
	default:
		return fmt.Errorf("tipo de nodo no soportado: %T", node)
	}
}

func (e *Evaluator) evaluateProgram(program *ast.Program) error {
	for _, stmt := range program.Statements {
		if err := e.Evaluate(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (e *Evaluator) evaluateDeclareStatement(stmt *ast.DeclareStatement) error {
	value := e.evaluateExpression(stmt.Value)
	if value == nil {
		return fmt.Errorf("no se pudo evaluar el valor de la declaración")
	}
	
	if stmt.IsConst {
		e.symbolTable.SetConst(stmt.Name.Value, value)
	} else {
		e.symbolTable.Set(stmt.Name.Value, value)
	}
	
	return nil
}

func (e *Evaluator) evaluateAssignStatement(stmt *ast.AssignStatement) error {
	value := e.evaluateExpression(stmt.Value)
	if value == nil {
		return fmt.Errorf("no se pudo evaluar el valor de la asignación")
	}
	
	e.symbolTable.Set(stmt.Name.Value, value)
	return nil
}

func (e *Evaluator) evaluateIfStatement(stmt *ast.IfStatement) error {
	condition := e.evaluateExpression(stmt.Condition)
	
	if isTruthy(condition) {
		return e.Evaluate(stmt.Then)
	} else if stmt.Else != nil {
		return e.Evaluate(stmt.Else)
	}
	
	return nil
}

func (e *Evaluator) evaluateWhileStatement(stmt *ast.WhileStatement) error {
	for {
		condition := e.evaluateExpression(stmt.Condition)
		if !isTruthy(condition) {
			break
		}
		
		if err := e.Evaluate(stmt.Body); err != nil {
			return err
		}
	}
	
	return nil
}

func (e *Evaluator) evaluateRepeatStatement(stmt *ast.RepeatStatement) error {
	from := e.evaluateExpression(stmt.From)
	to := e.evaluateExpression(stmt.To)
	
	fromVal := getIntValue(from)
	toVal := getIntValue(to)
	
	if fromVal == nil || toVal == nil {
		return fmt.Errorf("los valores de 'desde' y 'hasta' deben ser enteros")
	}
	
	for i := *fromVal; i <= *toVal; i++ {
		e.symbolTable.Set(stmt.Variable.Value, *fromVal)
		if err := e.Evaluate(stmt.Body); err != nil {
			return err
		}
	}
	
	return nil
}

func (e *Evaluator) evaluateShowStatement(stmt *ast.ShowStatement) error {
	value := e.evaluateExpression(stmt.Value)
	fmt.Println(value)
	return nil
}

func (e *Evaluator) evaluateReturnStatement(stmt *ast.ReturnStatement) error {
	// En una implementación completa, esto retornaría el valor
	// Por ahora solo lo evaluamos
	if stmt.Value != nil {
		e.evaluateExpression(stmt.Value)
	}
	return nil
}

func (e *Evaluator) evaluateBlockStatement(stmt *ast.BlockStatement) error {
	for _, s := range stmt.Statements {
		if err := e.Evaluate(s); err != nil {
			return err
		}
	}
	return nil
}

func (e *Evaluator) evaluateExpression(expr ast.Expression) interface{} {
	switch ex := expr.(type) {
	case *ast.IntegerLiteral:
		return ex.Value
	case *ast.FloatLiteral:
		return ex.Value
	case *ast.StringLiteral:
		return ex.Value
	case *ast.BooleanLiteral:
		return ex.Value
	case *ast.Identifier:
		val, ok := e.symbolTable.Get(ex.Value)
		if !ok {
			return nil
		}
		return val
	case *ast.InfixExpression:
		return e.evaluateInfixExpression(ex)
	case *ast.PrefixExpression:
		return e.evaluatePrefixExpression(ex)
	case *ast.CallExpression:
		return e.evaluateCallExpression(ex)
	default:
		return nil
	}
}

func (e *Evaluator) evaluateInfixExpression(expr *ast.InfixExpression) interface{} {
	left := e.evaluateExpression(expr.Left)
	right := e.evaluateExpression(expr.Right)
	
	switch expr.Operator {
	case "+":
		return add(left, right)
	case "-":
		return subtract(left, right)
	case "*":
		return multiply(left, right)
	case "/":
		return divide(left, right)
	case "%":
		return modulo(left, right)
	case "↔", "==":
		return left == right
	case "≠", "!=":
		return left != right
	case "<":
		return lessThan(left, right)
	case ">":
		return greaterThan(left, right)
	case "≤", "<=":
		return lessOrEqual(left, right)
	case "≥", ">=":
		return greaterOrEqual(left, right)
	case "∧", "&&":
		return isTruthy(left) && isTruthy(right)
	case "∨", "||":
		return isTruthy(left) || isTruthy(right)
	default:
		return nil
	}
}

func (e *Evaluator) evaluatePrefixExpression(expr *ast.PrefixExpression) interface{} {
	right := e.evaluateExpression(expr.Right)
	
	switch expr.Operator {
	case "¬", "!":
		return !isTruthy(right)
	case "-":
		if val, ok := right.(int64); ok {
			return -val
		}
		if val, ok := right.(float64); ok {
			return -val
		}
	}
	
	return nil
}

func (e *Evaluator) evaluateCallExpression(expr *ast.CallExpression) interface{} {
	// Implementación simplificada
	// En producción, buscaría la función en la tabla de símbolos
	return nil
}

// Funciones auxiliares
func isTruthy(obj interface{}) bool {
	switch obj {
	case nil:
		return false
	case false:
		return false
	case 0:
		return false
	case 0.0:
		return false
	case "":
		return false
	default:
		return true
	}
}

func getIntValue(val interface{}) *int64 {
	if v, ok := val.(int64); ok {
		return &v
	}
	return nil
}

func add(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int64:
		if r, ok := right.(int64); ok {
			return l + r
		}
		if r, ok := right.(float64); ok {
			return float64(l) + r
		}
	case float64:
		if r, ok := right.(float64); ok {
			return l + r
		}
		if r, ok := right.(int64); ok {
			return l + float64(r)
		}
	case string:
		return l + fmt.Sprintf("%v", right)
	}
	return nil
}

func subtract(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int64:
		if r, ok := right.(int64); ok {
			return l - r
		}
		if r, ok := right.(float64); ok {
			return float64(l) - r
		}
	case float64:
		if r, ok := right.(float64); ok {
			return l - r
		}
		if r, ok := right.(int64); ok {
			return l - float64(r)
		}
	}
	return nil
}

func multiply(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int64:
		if r, ok := right.(int64); ok {
			return l * r
		}
		if r, ok := right.(float64); ok {
			return float64(l) * r
		}
	case float64:
		if r, ok := right.(float64); ok {
			return l * r
		}
		if r, ok := right.(int64); ok {
			return l * float64(r)
		}
	}
	return nil
}

func divide(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int64:
		if r, ok := right.(int64); ok {
			if r == 0 {
				return nil
			}
			return float64(l) / float64(r)
		}
		if r, ok := right.(float64); ok {
			if r == 0 {
				return nil
			}
			return float64(l) / r
		}
	case float64:
		if r, ok := right.(float64); ok {
			if r == 0 {
				return nil
			}
			return l / r
		}
		if r, ok := right.(int64); ok {
			if r == 0 {
				return nil
			}
			return l / float64(r)
		}
	}
	return nil
}

func modulo(left, right interface{}) interface{} {
	if l, ok := left.(int64); ok {
		if r, ok := right.(int64); ok {
			if r == 0 {
				return nil
			}
			return l % r
		}
	}
	return nil
}

func lessThan(left, right interface{}) bool {
	switch l := left.(type) {
	case int64:
		if r, ok := right.(int64); ok {
			return l < r
		}
		if r, ok := right.(float64); ok {
			return float64(l) < r
		}
	case float64:
		if r, ok := right.(float64); ok {
			return l < r
		}
		if r, ok := right.(int64); ok {
			return l < float64(r)
		}
	}
	return false
}

func greaterThan(left, right interface{}) bool {
	return lessThan(right, left)
}

func lessOrEqual(left, right interface{}) bool {
	return !greaterThan(left, right)
}

func greaterOrEqual(left, right interface{}) bool {
	return !lessThan(left, right)
}

