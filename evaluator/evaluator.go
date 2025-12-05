package evaluator

import (
	"fmt"
	"flux/ast"
	"flux/symbol"
)

type Evaluator struct {
	symbolTable *symbol.Table
	parentScope *symbol.Table // Para manejar scopes anidados
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
	case *ast.FunctionStatement:
		return e.evaluateFunctionStatement(n)
	case *ast.BlockStatement:
		return e.evaluateBlockStatement(n)
	case *ast.ExpressionStatement:
		// Evaluar la expresión pero ignorar el resultado (para llamadas a función sin asignación)
		e.evaluateExpression(n.Expression)
		return nil
	default:
		return fmt.Errorf("tipo de nodo no soportado: %T", node)
	}
}

func (e *Evaluator) evaluateProgram(program *ast.Program) error {
	for _, stmt := range program.Statements {
		if err := e.Evaluate(stmt); err != nil {
			// Si es un ReturnValue, ignorarlo (solo es relevante dentro de funciones)
			if IsReturnValue(err) {
				continue
			}
			// Para otros errores, retornarlos
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
		if stmt.Then != nil {
			err := e.Evaluate(stmt.Then)
			// Si es un ReturnValue, propagarlo
			if _, ok := err.(*ReturnValue); ok {
				return err
			}
			return err
		}
	} else if stmt.Else != nil {
		err := e.Evaluate(stmt.Else)
		// Si es un ReturnValue, propagarlo
		if _, ok := err.(*ReturnValue); ok {
			return err
		}
		return err
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
			// Si es un ReturnValue, propagarlo
			if IsReturnValue(err) {
				return err
			}
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
		e.symbolTable.Set(stmt.Variable.Value, int64(i))
		if err := e.Evaluate(stmt.Body); err != nil {
			// Si es un ReturnValue, propagarlo (para retornos dentro de bucles en funciones)
			if IsReturnValue(err) {
				return err
			}
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

// ReturnValue es un error especial para indicar que se debe retornar de una función
type ReturnValue struct {
	Value interface{}
}

func (r *ReturnValue) Error() string {
	return "return"
}

// IsReturnValue verifica si un error es un ReturnValue
func IsReturnValue(err error) bool {
	_, ok := err.(*ReturnValue)
	return ok
}

func (e *Evaluator) evaluateReturnStatement(stmt *ast.ReturnStatement) error {
	var value interface{}
	if stmt.Value != nil {
		value = e.evaluateExpression(stmt.Value)
	}
	return &ReturnValue{Value: value}
}

// Tipo para representar funciones
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
}

func (e *Evaluator) evaluateFunctionStatement(stmt *ast.FunctionStatement) error {
	// Registrar la función en la tabla de símbolos
	fn := &Function{
		Parameters: stmt.Parameters,
		Body:       stmt.Body,
	}
	e.symbolTable.Set(stmt.Name.Value, fn)
	return nil
}

func (e *Evaluator) evaluateBlockStatement(stmt *ast.BlockStatement) error {
	for _, s := range stmt.Statements {
		if err := e.Evaluate(s); err != nil {
			// Si es un ReturnValue, propagarlo
			if IsReturnValue(err) {
				return err
			}
			// Para otros errores, retornarlos también
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
		// Buscar en el scope actual (que buscará recursivamente en los padres)
		val, ok := e.symbolTable.Get(ex.Value)
		if ok {
			return val
		}
		return nil
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
	// Obtener el nombre de la función
	var fnName string
	if ident, ok := expr.Function.(*ast.Identifier); ok {
		fnName = ident.Value
	} else {
		return nil
	}
	
	// Buscar la función en la tabla de símbolos
	// Primero buscar en el scope actual
	val, ok := e.symbolTable.Get(fnName)
	
	// Si no está en el scope actual, buscar en el scope padre (variables globales)
	if !ok && e.parentScope != nil {
		val, ok = e.parentScope.Get(fnName)
	}
	
	// Si aún no está, buscar recursivamente en los scopes padres
	if !ok {
		parent := e.parentScope
		for parent != nil {
			val, ok = parent.Get(fnName)
			if ok {
				break
			}
			// En una implementación completa, tendríamos un campo parentScope en Table
			// Por ahora, solo buscamos en el scope padre inmediato
			break
		}
	}
	
	if !ok {
		// Función no encontrada - esto podría ser un error
		return nil
	}
	
	fn, ok := val.(*Function)
	if !ok {
		return nil
	}
	
	// Evaluar los argumentos
	args := make([]interface{}, len(expr.Arguments))
	for i, arg := range expr.Arguments {
		args[i] = e.evaluateExpression(arg)
	}
	
	// Guardar el scope actual
	oldTable := e.symbolTable
	oldParent := e.parentScope
	
	// Crear un nuevo scope para los parámetros con el scope anterior como padre
	newTable := symbol.NewTableWithParent(oldTable)
	e.symbolTable = newTable
	e.parentScope = oldTable // El scope anterior es el padre
	
	// Asignar los argumentos a los parámetros
	for i, param := range fn.Parameters {
		if i < len(args) {
			newTable.Set(param.Value, args[i])
		}
	}
	
	// Ejecutar el cuerpo de la función
	var result interface{}
	
	// Ejecutar todas las sentencias del cuerpo
	for _, stmt := range fn.Body.Statements {
		err := e.Evaluate(stmt)
		if err != nil {
			// Si es un ReturnValue, capturar el valor y terminar
			if retVal, ok := err.(*ReturnValue); ok {
				result = retVal.Value
				break
			}
			// Si es otro error, continuar (en producción se manejaría mejor)
		}
	}
	
	// Restaurar el scope anterior
	e.symbolTable = oldTable
	e.parentScope = oldParent
	
	return result
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

