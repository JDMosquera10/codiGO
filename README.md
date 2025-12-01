# Proyecto Final: Análisis Léxico y Sintáctico
## Lenguaje Flux - Basado en Go

Este proyecto implementa un analizador léxico y sintáctico completo para el lenguaje **Flux**, un mini-lenguaje diseñado como parte del proyecto final de análisis léxico y sintáctico.

## Estructura del Proyecto

```
codiGO/
├── PROYECTO_FINAL.md      # Documento completo del proyecto
├── DEV.md                  # Especificaciones del experto PLT
├── PROMT.md                # Requerimientos del proyecto
├── main.go                 # Punto de entrada
├── go.mod                  # Módulo Go
├── ejemplo.flux            # Programa de ejemplo en Flux
├── lexer/
│   └── lexer.go           # Analizador léxico
├── parser/
│   └── parser.go          # Analizador sintáctico
├── ast/
│   └── ast.go             # Definiciones del AST
├── evaluator/
│   └── evaluator.go       # Evaluador/interprete
└── symbol/
    └── symbol.go          # Tabla de símbolos
```

## Características del Lenguaje Flux

- **Sintaxis expresiva:** Palabras clave en español con operadores Unicode
- **Tipado dinámico:** Inferencia de tipos automática
- **Estructuras de control:** Condicionales, bucles, funciones
- **Operadores Unicode:** →, ↔, ≠, ≤, ≥, ∧, ∨, ¬

## Compilación y Ejecución

### Requisitos
- Go 1.21 o superior

### Compilar
```bash
go build -o flux main.go
```

### Ejecutar
```bash
go run main.go ejemplo.flux
```

## Ejemplo de Código Flux

```flux
definir numero → 20

función verificarPrimo(n) hacer
    si n ≤ 1 entonces
        retornar falso
    fin
    
    repetir i desde 2 hasta n - 1 hacer
        si n % i ↔ 0 entonces
            retornar falso
        fin
    fin
    
    retornar verdadero
fin

mostrar("Verificando si " + numero + " es primo...")
```

## Documentación Completa

Ver `PROYECTO_FINAL.md` para la documentación completa del proyecto, incluyendo:
- Análisis completo del lenguaje base (Go)
- Diseño del nuevo lenguaje Flux
- Tabla léxica completa
- Gramática formal (BNF/EBNF)
- Análisis léxico y sintáctico
- Árbol sintáctico (AST)
- Simulación de ejecución

## Autor

Proyecto desarrollado como parte del curso de Análisis Léxico y Sintáctico.

