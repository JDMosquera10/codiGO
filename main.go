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

	// Análisis sintáctico
	p := parser.New(tokens)
	ast, err := p.Parse()
	if err != nil {
		fmt.Printf("Error en análisis sintáctico: %v\n", err)
		os.Exit(1)
	}

	// Evaluación
	symbolTable := symbol.NewTable()
	eval := evaluator.New(symbolTable)
	
	fmt.Println("=== EJECUCION ===")
	err = eval.Evaluate(ast)
	if err != nil {
		// Si es un ReturnValue, ignorarlo (solo es relevante dentro de funciones)
		if !evaluator.IsReturnValue(err) {
			fmt.Printf("Error en ejecución: %v\n", err)
			os.Exit(1)
		}
	}
}

