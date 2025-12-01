package main

import (
	"fmt"
	"os"
	"flux/lexer"
	"flux/parser"
	"flux/evaluator"
	"flux/symbol"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <archivo.flux>")
		os.Exit(1)
	}

	filename := os.Args[1]
	
	// Leer archivo fuente
	sourceBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v\n", err)
		os.Exit(1)
	}
	
	// Convertir a string UTF-8, removiendo BOM si existe
	source := string(sourceBytes)
	if len(sourceBytes) >= 3 && sourceBytes[0] == 0xEF && sourceBytes[1] == 0xBB && sourceBytes[2] == 0xBF {
		source = string(sourceBytes[3:])
	}

	// Análisis léxico
	l := lexer.New(string(source))
	tokens, err := l.Tokenize()
	if err != nil {
		fmt.Printf("Error en análisis léxico: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== ANÁLISIS LÉXICO ===")
	fmt.Printf("Total de tokens: %d\n\n", len(tokens))
	for i, token := range tokens {
		if i < 20 { // Mostrar primeros 20 tokens
			fmt.Printf("[%d] %s: %s (línea %d, columna %d)\n", i+1, token.Type, token.Value, token.Line, token.Column)
		}
	}
	if len(tokens) > 20 {
		fmt.Printf("... (%d tokens más)\n", len(tokens)-20)
	}
	fmt.Println()

	// Análisis sintáctico
	p := parser.New(tokens)
	ast, err := p.Parse()
	if err != nil {
		fmt.Printf("Error en análisis sintáctico: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== ÁRBOL SINTÁCTICO ===")
	ast.Print(0)
	fmt.Println()

	// Evaluación
	symbolTable := symbol.NewTable()
	eval := evaluator.New(symbolTable)
	
	fmt.Println("=== EJECUCIÓN ===")
	err = eval.Evaluate(ast)
	if err != nil {
		fmt.Printf("Error en ejecución: %v\n", err)
		os.Exit(1)
	}
}

