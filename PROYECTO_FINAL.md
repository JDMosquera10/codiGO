# PROYECTO FINAL: ANÁLISIS LÉXICO Y SINTÁCTICO
## Diseño de un Mini-Lenguaje Basado en Go

---

## ÍNDICE

1. [Análisis del Lenguaje Base (Go)](#a-análisis-del-lenguaje-base-go)
2. [Creación del Nuevo Lenguaje: "Flux"](#b-creación-del-nuevo-lenguaje-flux)
3. [Tabla Léxica Completa](#c-tabla-léxica-completa)
4. [Gramática Formal (BNF/EBNF)](#d-gramática-formal-bnfebnf)
5. [Programa de Ejemplo Completo](#e-programa-de-ejemplo-completo)
6. [Análisis Léxico del Ejemplo](#f-análisis-léxico-del-ejemplo)
7. [Árbol Sintáctico (AST)](#g-árbol-sintáctico-ast)
8. [Simulación de Ejecución](#h-simulación-de-ejecución)
9. [Implementación del Analizador](#i-implementación-del-analizador-en-go)

---

# A. ANÁLISIS DEL LENGUAJE BASE (GO)

## 1. Palabras Reservadas

Go tiene 25 palabras reservadas que definen la estructura del lenguaje:

| Palabra Reservada | Categoría | Descripción | Ejemplo |
|-------------------|-----------|-------------|---------|
| `package` | Declaración | Define el paquete al que pertenece el archivo | `package main` |
| `import` | Declaración | Importa paquetes externos | `import "fmt"` |
| `func` | Declaración | Define una función | `func main() {}` |
| `var` | Declaración | Declara una variable | `var x int` |
| `const` | Declaración | Declara una constante | `const PI = 3.14` |
| `type` | Declaración | Define un nuevo tipo | `type User struct {}` |
| `struct` | Tipo | Define una estructura | `type Point struct { x, y int }` |
| `interface` | Tipo | Define una interfaz | `type Writer interface {}` |
| `map` | Tipo | Tipo de datos mapa | `var m map[string]int` |
| `chan` | Tipo | Canal para concurrencia | `var c chan int` |
| `if` | Control | Condicional | `if x > 0 {}` |
| `else` | Control | Rama alternativa | `if x > 0 {} else {}` |
| `switch` | Control | Estructura de selección múltiple | `switch x {}` |
| `case` | Control | Caso en switch | `case 1:` |
| `default` | Control | Caso por defecto | `default:` |
| `for` | Control | Bucle iterativo | `for i := 0; i < 10; i++ {}` |
| `range` | Control | Iteración sobre colecciones | `for k, v := range m {}` |
| `break` | Control | Salir de bucle | `break` |
| `continue` | Control | Continuar iteración | `continue` |
| `return` | Control | Retornar valor | `return 42` |
| `go` | Concurrencia | Lanzar goroutine | `go f()` |
| `select` | Concurrencia | Selección de canales | `select {}` |
| `defer` | Control | Ejecutar al finalizar función | `defer f.Close()` |
| `goto` | Control | Salto incondicional | `goto label` |
| `fallthrough` | Control | Continuar en switch | `fallthrough` |

### Explicación Detallada de Categorías:

**Declaración:** `package`, `import`, `func`, `var`, `const`, `type` - Definen la estructura del programa y sus componentes.

**Tipo:** `struct`, `interface`, `map`, `chan` - Definen tipos de datos compuestos.

**Control:** `if`, `else`, `switch`, `case`, `for`, `break`, `continue`, `return`, `defer`, `goto`, `fallthrough` - Controlan el flujo de ejecución.

**Concurrencia:** `go`, `select` - Gestionan la ejecución concurrente mediante goroutines y canales.

## 2. Operadores

### 2.1 Operadores Aritméticos

| Operador | Nombre | Descripción | Ejemplo | Precedencia |
|----------|--------|-------------|---------|-------------|
| `+` | Suma | Suma dos operandos | `x + y` | 5 |
| `-` | Resta | Resta dos operandos | `x - y` | 5 |
| `*` | Multiplicación | Multiplica dos operandos | `x * y` | 4 |
| `/` | División | Divide dos operandos | `x / y` | 4 |
| `%` | Módulo | Resto de la división | `x % y` | 4 |
| `++` | Incremento | Incrementa en 1 (sufijo) | `x++` | 2 |
| `--` | Decremento | Decrementa en 1 (sufijo) | `x--` | 2 |

**Ejemplo:**
```go
x := 10
y := 3
suma := x + y        // 13
producto := x * y    // 30
resto := x % y       // 1
x++                  // x ahora es 11
```

### 2.2 Operadores Relacionales

| Operador | Nombre | Descripción | Ejemplo | Precedencia |
|----------|--------|-------------|---------|-------------|
| `==` | Igualdad | Compara igualdad | `x == y` | 7 |
| `!=` | Desigualdad | Compara desigualdad | `x != y` | 7 |
| `<` | Menor que | Compara menor | `x < y` | 6 |
| `<=` | Menor o igual | Compara menor o igual | `x <= y` | 6 |
| `>` | Mayor que | Compara mayor | `x > y` | 6 |
| `>=` | Mayor o igual | Compara mayor o igual | `x >= y` | 6 |

**Ejemplo:**
```go
x := 5
y := 10
esMenor := x < y     // true
esIgual := x == y    // false
```

### 2.3 Operadores Lógicos

| Operador | Nombre | Descripción | Ejemplo | Precedencia |
|----------|--------|-------------|---------|-------------|
| `&&` | AND lógico | Conjunción lógica | `x && y` | 3 |
| `\|\|` | OR lógico | Disyunción lógica | `x \|\| y` | 2 |
| `!` | NOT lógico | Negación lógica | `!x` | 1 |

**Ejemplo:**
```go
x := true
y := false
resultado := x && y  // false
resultado2 := x || y // true
negado := !x         // false
```

### 2.4 Operadores de Asignación

| Operador | Nombre | Descripción | Ejemplo |
|----------|--------|-------------|---------|
| `=` | Asignación simple | Asigna valor | `x = 5` |
| `:=` | Declaración corta | Declara y asigna | `x := 5` |
| `+=` | Suma y asigna | `x += y` equivale a `x = x + y` | `x += 3` |
| `-=` | Resta y asigna | `x -= y` equivale a `x = x - y` | `x -= 3` |
| `*=` | Multiplica y asigna | `x *= y` equivale a `x = x * y` | `x *= 3` |
| `/=` | Divide y asigna | `x /= y` equivale a `x = x / y` | `x /= 3` |
| `%=` | Módulo y asigna | `x %= y` equivale a `x = x % y` | `x %= 3` |
| `&=` | AND bit y asigna | Operación bit a bit | `x &= y` |
| `\|=` | OR bit y asigna | Operación bit a bit | `x \|= y` |
| `^=` | XOR bit y asigna | Operación bit a bit | `x ^= y` |
| `<<=` | Desplazamiento izquierda y asigna | `x <<= y` | `x <<= 2` |
| `>>=` | Desplazamiento derecha y asigna | `x >>= y` | `x >>= 2` |

**Ejemplo:**
```go
x := 10
x += 5   // x = 15
x *= 2   // x = 30
x /= 3   // x = 10
```

## 3. Delimitadores

| Delimitador | Nombre | Uso | Ejemplo |
|-------------|--------|-----|---------|
| `(` `)` | Paréntesis | Agrupación, llamadas de función | `f(x)` |
| `{` `}` | Llaves | Bloques de código, estructuras | `{ x := 5 }` |
| `[` `]` | Corchetes | Arrays, slices, índices | `arr[0]` |
| `;` | Punto y coma | Terminador de sentencia (opcional) | `x := 5;` |
| `,` | Coma | Separador de elementos | `x, y := 1, 2` |
| `.` | Punto | Acceso a miembros, paquetes | `fmt.Println()` |
| `:` | Dos puntos | Etiquetas, casos switch | `case 1:` |
| `...` | Ellipsis | Variadic, unpacking | `func f(...int) {}` |

**Ejemplo:**
```go
func calcular(x, y int) int {
    return x + y
}

arr := []int{1, 2, 3}
valor := arr[0]
```

## 4. Tipos de Datos Básicos

| Tipo | Descripción | Tamaño | Rango | Ejemplo |
|------|-------------|--------|-------|---------|
| `bool` | Booleano | 1 byte | `true`, `false` | `var b bool = true` |
| `int` | Entero con signo | 32 o 64 bits | Plataforma dependiente | `var x int = 42` |
| `int8` | Entero 8 bits | 8 bits | -128 a 127 | `var x int8 = 100` |
| `int16` | Entero 16 bits | 16 bits | -32768 a 32767 | `var x int16 = 1000` |
| `int32` | Entero 32 bits | 32 bits | -2³¹ a 2³¹-1 | `var x int32 = 100000` |
| `int64` | Entero 64 bits | 64 bits | -2⁶³ a 2⁶³-1 | `var x int64 = 1000000` |
| `uint` | Entero sin signo | 32 o 64 bits | 0 a 2ⁿ-1 | `var x uint = 42` |
| `uint8` | Entero sin signo 8 bits | 8 bits | 0 a 255 | `var x uint8 = 200` |
| `uint16` | Entero sin signo 16 bits | 16 bits | 0 a 65535 | `var x uint16 = 50000` |
| `uint32` | Entero sin signo 32 bits | 32 bits | 0 a 2³²-1 | `var x uint32 = 4000000` |
| `uint64` | Entero sin signo 64 bits | 64 bits | 0 a 2⁶⁴-1 | `var x uint64 = 10000000` |
| `uintptr` | Puntero sin signo | Plataforma dependiente | Dirección de memoria | `var p uintptr` |
| `float32` | Punto flotante 32 bits | 32 bits | ±3.4e±38 | `var f float32 = 3.14` |
| `float64` | Punto flotante 64 bits | 64 bits | ±1.7e±308 | `var f float64 = 3.14159` |
| `complex64` | Complejo 64 bits | 64 bits | float32 real + float32 imag | `var c complex64 = 1+2i` |
| `complex128` | Complejo 128 bits | 128 bits | float64 real + float64 imag | `var c complex128 = 1+2i` |
| `byte` | Alias de uint8 | 8 bits | 0 a 255 | `var b byte = 'A'` |
| `rune` | Alias de int32 (Unicode) | 32 bits | Puntos de código Unicode | `var r rune = '中'` |
| `string` | Cadena de caracteres | Variable | Secuencia de bytes | `var s string = "Hola"` |

**Ejemplo:**
```go
var entero int = 42
var decimal float64 = 3.14159
var texto string = "Hola Mundo"
var booleano bool = true
var caracter rune = 'A'
```

## 5. Estructuras de Control

### 5.1 Condicionales

**if-else:**
```go
if condición {
    // código si verdadero
} else {
    // código si falso
}

// if-else if-else
if x > 10 {
    fmt.Println("Mayor que 10")
} else if x > 5 {
    fmt.Println("Mayor que 5")
} else {
    fmt.Println("Menor o igual a 5")
}
```

**switch:**
```go
switch variable {
case valor1:
    // código
case valor2:
    // código
default:
    // código por defecto
}

// switch sin expresión
switch {
case x > 10:
    fmt.Println("Mayor que 10")
case x > 5:
    fmt.Println("Mayor que 5")
}
```

### 5.2 Bucles

**for tradicional:**
```go
for inicialización; condición; incremento {
    // código
}

for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

**for como while:**
```go
for condición {
    // código
}

x := 0
for x < 10 {
    fmt.Println(x)
    x++
}
```

**for infinito:**
```go
for {
    // código (requiere break para salir)
    if condición {
        break
    }
}
```

**for-range:**
```go
// Arrays/Slices
for índice, valor := range slice {
    fmt.Println(índice, valor)
}

// Maps
for clave, valor := range mapa {
    fmt.Println(clave, valor)
}

// Strings
for índice, carácter := range "Hola" {
    fmt.Println(índice, carácter)
}
```

### 5.3 Control de Flujo

**break:** Sale del bucle más interno
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Sale cuando i == 5
    }
}
```

**continue:** Salta a la siguiente iteración
```go
for i := 0; i < 10; i++ {
    if i % 2 == 0 {
        continue  // Salta números pares
    }
    fmt.Println(i)
}
```

**return:** Retorna de una función
```go
func suma(a, b int) int {
    return a + b
}
```

**defer:** Ejecuta código al finalizar la función
```go
func ejemplo() {
    defer fmt.Println("Esto se ejecuta al final")
    fmt.Println("Esto se ejecuta primero")
}
```

---

# B. CREACIÓN DEL NUEVO LENGUAJE: "FLUX"

## 1. Identificación y Paradigma

**Nombre del Lenguaje:** **Flux**

**Justificación del Nombre:** 
"Flux" representa el flujo constante de datos y ejecución, inspirado en el concepto de flujo de información en sistemas reactivos. El nombre evoca movimiento, transformación y dinamismo, características que el lenguaje busca promover mediante una sintaxis declarativa y expresiva.

**Paradigma:** 
- **Multi-paradigma** con énfasis en:
  - **Programación Imperativa:** Estructuras de control tradicionales
  - **Programación Funcional:** Expresiones puras, funciones como valores
  - **Programación Declarativa:** Sintaxis orientada a expresar "qué" en lugar de "cómo"

**Filosofía de Diseño:**
Flux busca ser un lenguaje expresivo y legible, donde la sintaxis refleja la intención del programador de manera natural. Se inspira en lenguajes como Python (legibilidad) y Go (simplicidad), pero con una identidad propia que prioriza la claridad sobre la concisión.

## 2. Especificación Inicial (Sintaxis y Reglas)

### 2.1 Palabras Reservadas Nuevas

| Flux | Go | Descripción |
|------|-----|-------------|
| `definir` | `var` | Declara una variable |
| `constante` | `const` | Declara una constante |
| `función` | `func` | Define una función |
| `si` | `if` | Condicional |
| `entonces` | (implícito) | Rama verdadera del condicional |
| `sino` | `else` | Rama falsa del condicional |
| `mientras` | `for` | Bucle condicional |
| `repetir` | `for` | Bucle con contador |
| `desde` | `range` | Iteración sobre colecciones |
| `mostrar` | `fmt.Println` | Imprime en consola |
| `retornar` | `return` | Retorna valor de función |
| `salir` | `break` | Sale de un bucle |
| `continuar` | `continue` | Continúa siguiente iteración |
| `verdadero` | `true` | Valor booleano verdadero |
| `falso` | `false` | Valor booleano falso |
| `nulo` | `nil` | Valor nulo |
| `tipo` | `type` | Define un nuevo tipo |
| `estructura` | `struct` | Define una estructura |
| `fin` | `}` | Cierra un bloque |
| `hacer` | `{` | Abre un bloque (opcional) |

**Total: 19 palabras reservadas**

### 2.2 Operadores Nuevos

| Operador Flux | Operador Go | Descripción |
|---------------|-------------|-------------|
| `→` | `=` | Asignación |
| `↔` | `==` | Igualdad |
| `≠` | `!=` | Desigualdad |
| `≤` | `<=` | Menor o igual |
| `≥` | `>=` | Mayor o igual |
| `∧` | `&&` | AND lógico |
| `∨` | `\|\|` | OR lógico |
| `¬` | `!` | NOT lógico |
| `⊕` | `^` | XOR bit a bit |
| `·` | `.` | Acceso a miembros |

### 2.3 Estructuras del Lenguaje

#### Declaraciones
```flux
definir x → 10
definir nombre → "Flux"
constante PI → 3.14159
```

#### Asignaciones
```flux
x → 20
y → x + 5
```

#### Condicionales
```flux
si x > 5 entonces
    mostrar("Mayor que 5")
sino
    mostrar("Menor o igual")
fin
```

#### Ciclos
```flux
// Bucle condicional
mientras x > 0 hacer
    mostrar(x)
    x → x - 1
fin

// Bucle con contador
repetir i desde 0 hasta 10 hacer
    mostrar(i)
fin

// Iteración sobre colección
repetir elemento desde lista hacer
    mostrar(elemento)
fin
```

#### Funciones
```flux
función suma(a, b) hacer
    retornar a + b
fin

función saludar(nombre) hacer
    mostrar("Hola, " + nombre)
fin
```

#### Bloques
```flux
// Bloque explícito con 'hacer'
si condición hacer
    // código
fin

// Bloque implícito (sin 'hacer')
si condición entonces
    // código
fin
```

## 3. Gramática Formal (BNF/EBNF)

### 3.1 Gramática BNF Completa

```
<programa>          ::= <sentencia> | <sentencia> <programa>

<sentencia>         ::= <declaracion> | <asignacion> | <expresion> | 
                        <condicional> | <ciclo> | <funcion> | <mostrar> | 
                        <retorno> | <salida> | <continuacion>

<declaracion>       ::= "definir" <identificador> "→" <expresion> |
                        "constante" <identificador> "→" <expresion>

<asignacion>        ::= <identificador> "→" <expresion>

<condicional>       ::= "si" <expresion> "entonces" <bloque> |
                        "si" <expresion> "entonces" <bloque> "sino" <bloque> |
                        "si" <expresion> "hacer" <bloque> "fin" |
                        "si" <expresion> "hacer" <bloque> "sino" <bloque> "fin"

<ciclo>             ::= <mientras> | <repetir>

<mientras>          ::= "mientras" <expresion> "hacer" <bloque> "fin"

<repetir>           ::= "repetir" <identificador> "desde" <expresion> "hasta" <expresion> "hacer" <bloque> "fin" |
                        "repetir" <identificador> "desde" <coleccion> "hacer" <bloque> "fin"

<funcion>           ::= "función" <identificador> "(" <parametros>? ")" "hacer" <bloque> "fin"

<parametros>        ::= <identificador> | <identificador> "," <parametros>

<bloque>            ::= <sentencia> | <sentencia> <bloque>

<expresion>         ::= <expresion_logica> | <expresion_aritmetica> | 
                        <expresion_relacional> | <literal> | <identificador> |
                        <llamada_funcion> | "(" <expresion> ")"

<expresion_logica>  ::= <expresion> "∧" <expresion> |
                        <expresion> "∨" <expresion> |
                        "¬" <expresion>

<expresion_relacional> ::= <expresion> "↔" <expresion> |
                            <expresion> "≠" <expresion> |
                            <expresion> "<" <expresion> |
                            <expresion> ">" <expresion> |
                            <expresion> "≤" <expresion> |
                            <expresion> "≥" <expresion>

<expresion_aritmetica> ::= <expresion> "+" <expresion> |
                            <expresion> "-" <expresion> |
                            <expresion> "*" <expresion> |
                            <expresion> "/" <expresion> |
                            <expresion> "%" <expresion>

<llamada_funcion>   ::= <identificador> "(" <argumentos>? ")"

<argumentos>        ::= <expresion> | <expresion> "," <argumentos>

<mostrar>           ::= "mostrar" "(" <expresion> ")"

<retorno>           ::= "retornar" <expresion>?

<salida>            ::= "salir"

<continuacion>      ::= "continuar"

<literal>           ::= <numero> | <cadena> | <booleano> | "nulo"

<numero>            ::= <entero> | <decimal>

<entero>            ::= [0-9]+

<decimal>           ::= [0-9]+ "." [0-9]+

<cadena>            ::= '"' <caracteres> '"' | "'" <caracteres> "'"

<booleano>          ::= "verdadero" | "falso"

<identificador>     ::= [a-zA-Z_][a-zA-Z0-9_]*

<coleccion>         ::= <identificador>
```

### 3.2 Gramática EBNF (Extended BNF)

```
programa           = { sentencia }

sentencia          = declaracion | asignacion | condicional | ciclo | 
                     funcion | mostrar | retorno | salida | continuacion

declaracion        = ("definir" | "constante") identificador "→" expresion

asignacion         = identificador "→" expresion

condicional        = "si" expresion ("entonces" | "hacer") bloque 
                     ["sino" bloque] "fin"

ciclo              = mientras | repetir

mientras           = "mientras" expresion "hacer" bloque "fin"

repetir            = "repetir" identificador "desde" 
                     (expresion "hasta" expresion | coleccion) 
                     "hacer" bloque "fin"

funcion            = "función" identificador "(" [parametros] ")" 
                     "hacer" bloque "fin"

parametros         = identificador { "," identificador }

bloque             = { sentencia }

expresion          = expresion_logica | expresion_aritmetica | 
                     expresion_relacional | literal | identificador |
                     llamada_funcion | "(" expresion ")"

expresion_logica   = expresion ("∧" | "∨") expresion | "¬" expresion

expresion_relacional = expresion ("↔" | "≠" | "<" | ">" | "≤" | "≥") expresion

expresion_aritmetica = expresion ("+" | "-" | "*" | "/" | "%") expresion

llamada_funcion    = identificador "(" [argumentos] ")"

argumentos         = expresion { "," expresion }

mostrar            = "mostrar" "(" expresion ")"

retorno            = "retornar" [expresion]

salida             = "salir"

continuacion       = "continuar"

literal            = numero | cadena | booleano | "nulo"

numero             = entero | decimal

entero             = digitos+

decimal            = digitos+ "." digitos+

cadena             = ('"' caracteres '"') | ("'" caracteres "'")

booleano           = "verdadero" | "falso"

identificador      = letra { letra | digito | "_" }

coleccion          = identificador

letra              = [a-zA-Z]

digito             = [0-9]

caracteres         = { cualquier_caracter_excepto_comillas }
```

### 3.3 Precedencia de Operadores

| Precedencia | Operadores | Asociatividad |
|-------------|------------|---------------|
| 1 (más alta) | `¬` (NOT) | Derecha a izquierda |
| 2 | `*`, `/`, `%` | Izquierda a derecha |
| 3 | `+`, `-` | Izquierda a derecha |
| 4 | `<`, `>`, `≤`, `≥` | Izquierda a derecha |
| 5 | `↔`, `≠` | Izquierda a derecha |
| 6 | `∧` (AND) | Izquierda a derecha |
| 7 (más baja) | `∨` (OR) | Izquierda a derecha |

---

# C. TABLA LÉXICA COMPLETA

| Token | Categoría | Patrón (RegEx) | Descripción |
|-------|-----------|----------------|-------------|
| `definir` | PALABRA_RESERVADA | `^definir$` | Declara una variable |
| `constante` | PALABRA_RESERVADA | `^constante$` | Declara una constante |
| `función` | PALABRA_RESERVADA | `^función$` | Define una función |
| `si` | PALABRA_RESERVADA | `^si$` | Inicia condicional |
| `entonces` | PALABRA_RESERVADA | `^entonces$` | Rama verdadera |
| `sino` | PALABRA_RESERVADA | `^sino$` | Rama falsa |
| `mientras` | PALABRA_RESERVADA | `^mientras$` | Bucle condicional |
| `repetir` | PALABRA_RESERVADA | `^repetir$` | Bucle con contador |
| `desde` | PALABRA_RESERVADA | `^desde$` | Inicio de rango |
| `hasta` | PALABRA_RESERVADA | `^hasta$` | Fin de rango |
| `hacer` | PALABRA_RESERVADA | `^hacer$` | Inicia bloque |
| `fin` | PALABRA_RESERVADA | `^fin$` | Termina bloque |
| `mostrar` | PALABRA_RESERVADA | `^mostrar$` | Función de impresión |
| `retornar` | PALABRA_RESERVADA | `^retornar$` | Retorna valor |
| `salir` | PALABRA_RESERVADA | `^salir$` | Sale de bucle |
| `continuar` | PALABRA_RESERVADA | `^continuar$` | Continúa iteración |
| `verdadero` | PALABRA_RESERVADA | `^verdadero$` | Valor booleano true |
| `falso` | PALABRA_RESERVADA | `^falso$` | Valor booleano false |
| `nulo` | PALABRA_RESERVADA | `^nulo$` | Valor nulo |
| `→` | OPERADOR_ASIGNACION | `^→$` | Operador de asignación |
| `↔` | OPERADOR_RELACIONAL | `^↔$` | Operador de igualdad |
| `≠` | OPERADOR_RELACIONAL | `^≠$` | Operador de desigualdad |
| `≤` | OPERADOR_RELACIONAL | `^≤$` | Menor o igual |
| `≥` | OPERADOR_RELACIONAL | `^≥$` | Mayor o igual |
| `∧` | OPERADOR_LOGICO | `^∧$` | AND lógico |
| `∨` | OPERADOR_LOGICO | `^∨$` | OR lógico |
| `¬` | OPERADOR_LOGICO | `^¬$` | NOT lógico |
| `+` | OPERADOR_ARITMETICO | `^\+$` | Suma |
| `-` | OPERADOR_ARITMETICO | `^\-$` | Resta |
| `*` | OPERADOR_ARITMETICO | `^\*$` | Multiplicación |
| `/` | OPERADOR_ARITMETICO | `^\/$` | División |
| `%` | OPERADOR_ARITMETICO | `^\%$` | Módulo |
| `<` | OPERADOR_RELACIONAL | `^<$` | Menor que |
| `>` | OPERADOR_RELACIONAL | `^>$` | Mayor que |
| `(` | DELIMITADOR | `^\($` | Paréntesis izquierdo |
| `)` | DELIMITADOR | `^\)$` | Paréntesis derecho |
| `{` | DELIMITADOR | `^\{$` | Llave izquierda |
| `}` | DELIMITADOR | `^\}$` | Llave derecha |
| `,` | DELIMITADOR | `^,$` | Coma |
| `·` | DELIMITADOR | `^·$` | Acceso a miembros |
| `ENTERO` | LITERAL | `^\d+$` | Número entero |
| `DECIMAL` | LITERAL | `^\d+\.\d+$` | Número decimal |
| `CADENA` | LITERAL | `^"[^"]*"$` \| `^'[^']*'$` | Cadena de caracteres |
| `IDENTIFICADOR` | IDENTIFICADOR | `^[a-zA-Z_][a-zA-Z0-9_]*$` | Nombre de variable/función |
| `COMENTARIO_LINEA` | COMENTARIO | `^//.*$` | Comentario de línea |
| `COMENTARIO_BLOQUE` | COMENTARIO | `^/\*[\s\S]*?\*/$` | Comentario de bloque |
| `ESPACIO` | ESPACIO | `^\s+$` | Espacios en blanco |
| `NUEVA_LINEA` | ESPACIO | `^\n$` | Salto de línea |

**Total: 45 tokens**

---

# D. GRAMÁTICA FORMAL (BNF/EBNF)

(Véase sección B.3 para la gramática completa)

---

# E. PROGRAMA DE EJEMPLO COMPLETA

## Programa en Flux

```flux
// Programa de ejemplo: Calculadora de números primos
definir numero → 20
definir contador → 2
definir esPrimo → verdadero

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

si verificarPrimo(numero) entonces
    mostrar("El número " + numero + " es primo")
sino
    mostrar("El número " + numero + " no es primo")
fin

// Mostrar todos los primos hasta el número
mostrar("Números primos hasta " + numero + ":")
repetir i desde 2 hasta numero hacer
    si verificarPrimo(i) entonces
        mostrar(i)
    fin
fin
```

## Traducción al Lenguaje Base (Go)

```go
package main

import "fmt"

func main() {
    numero := 20
    contador := 2
    esPrimo := true

    verificarPrimo := func(n int) bool {
        if n <= 1 {
            return false
        }
        
        for i := 2; i < n; i++ {
            if n%i == 0 {
                return false
            }
        }
        
        return true
    }

    fmt.Println("Verificando si", numero, "es primo...")

    if verificarPrimo(numero) {
        fmt.Println("El número", numero, "es primo")
    } else {
        fmt.Println("El número", numero, "no es primo")
    }

    fmt.Println("Números primos hasta", numero, ":")
    for i := 2; i <= numero; i++ {
        if verificarPrimo(i) {
            fmt.Println(i)
        }
    }
}
```

## Output Esperado

```
Verificando si 20 es primo...
El número 20 no es primo
Números primos hasta 20:
2
3
5
7
11
13
17
19
```

---

# F. ANÁLISIS LÉXICO DEL EJEMPLO

Analizando el programa de ejemplo línea por línea:

## Línea 1: `definir numero → 20`

| Token | Tipo | Valor Semántico |
|-------|------|-----------------|
| `definir` | PALABRA_RESERVADA | Declaración de variable |
| `numero` | IDENTIFICADOR | Nombre de variable |
| `→` | OPERADOR_ASIGNACION | Asignación |
| `20` | ENTERO | Valor numérico 20 |

## Línea 2: `definir contador → 2`

| Token | Tipo | Valor Semántico |
|-------|------|-----------------|
| `definir` | PALABRA_RESERVADA | Declaración de variable |
| `contador` | IDENTIFICADOR | Nombre de variable |
| `→` | OPERADOR_ASIGNACION | Asignación |
| `2` | ENTERO | Valor numérico 2 |

## Línea 3: `definir esPrimo → verdadero`

| Token | Tipo | Valor Semántico |
|-------|------|-----------------|
| `definir` | PALABRA_RESERVADA | Declaración de variable |
| `esPrimo` | IDENTIFICADOR | Nombre de variable |
| `→` | OPERADOR_ASIGNACION | Asignación |
| `verdadero` | PALABRA_RESERVADA | Valor booleano true |

## Línea 5: `función verificarPrimo(n) hacer`

| Token | Tipo | Valor Semántico |
|-------|------|-----------------|
| `función` | PALABRA_RESERVADA | Definición de función |
| `verificarPrimo` | IDENTIFICADOR | Nombre de función |
| `(` | DELIMITADOR | Inicio de parámetros |
| `n` | IDENTIFICADOR | Parámetro |
| `)` | DELIMITADOR | Fin de parámetros |
| `hacer` | PALABRA_RESERVADA | Inicio de bloque |

## Línea 6: `si n ≤ 1 entonces`

| Token | Tipo | Valor Semántico |
|-------|------|-----------------|
| `si` | PALABRA_RESERVADA | Condicional |
| `n` | IDENTIFICADOR | Variable |
| `≤` | OPERADOR_RELACIONAL | Menor o igual |
| `1` | ENTERO | Valor numérico 1 |
| `entonces` | PALABRA_RESERVADA | Rama verdadera |

## Línea 7: `retornar falso`

| Token | Tipo | Valor Semántico |
|-------|------|-----------------|
| `retornar` | PALABRA_RESERVADA | Retorno |
| `falso` | PALABRA_RESERVADA | Valor booleano false |

## Resumen de Tokens del Programa Completo

| Token | Frecuencia | Categoría |
|-------|------------|-----------|
| `definir` | 3 | PALABRA_RESERVADA |
| `función` | 1 | PALABRA_RESERVADA |
| `si` | 4 | PALABRA_RESERVADA |
| `entonces` | 4 | PALABRA_RESERVADA |
| `sino` | 1 | PALABRA_RESERVADA |
| `retornar` | 3 | PALABRA_RESERVADA |
| `repetir` | 2 | PALABRA_RESERVADA |
| `desde` | 2 | PALABRA_RESERVADA |
| `hasta` | 2 | PALABRA_RESERVADA |
| `hacer` | 3 | PALABRA_RESERVADA |
| `fin` | 6 | PALABRA_RESERVADA |
| `mostrar` | 4 | PALABRA_RESERVADA |
| `→` | 3 | OPERADOR_ASIGNACION |
| `≤` | 1 | OPERADOR_RELACIONAL |
| `↔` | 1 | OPERADOR_RELACIONAL |
| `%` | 1 | OPERADOR_ARITMETICO |
| `-` | 2 | OPERADOR_ARITMETICO |
| `+` | 2 | OPERADOR_ARITMETICO |
| `(` | 6 | DELIMITADOR |
| `)` | 6 | DELIMITADOR |
| `numero`, `contador`, `esPrimo`, `verificarPrimo`, `n`, `i` | 15 | IDENTIFICADOR |
| `20`, `2`, `1`, `0` | 6 | ENTERO |
| `verdadero`, `falso` | 2 | BOOLEANO |
| `"Verificando si "`, `" es primo..."`, etc. | 4 | CADENA |

**Total de tokens: ~75 tokens**

---

# G. ÁRBOL SINTÁCTICO (AST)

## Representación del AST del Programa Completo

```
Programa
│
├── Declaración (numero → 20)
│   ├── Identificador: numero
│   └── Literal: 20
│
├── Declaración (contador → 2)
│   ├── Identificador: contador
│   └── Literal: 2
│
├── Declaración (esPrimo → verdadero)
│   ├── Identificador: esPrimo
│   └── Literal: verdadero
│
├── Función (verificarPrimo)
│   ├── Nombre: verificarPrimo
│   ├── Parámetros
│   │   └── n
│   └── Cuerpo
│       └── Condicional (si)
│           ├── Condición
│           │   └── Expresión Relacional
│           │       ├── Identificador: n
│           │       ├── Operador: ≤
│           │       └── Literal: 1
│           ├── Rama Verdadera (entonces)
│           │   └── Retorno
│           │       └── Literal: falso
│           └── Rama Falsa (sino)
│               └── Bloque
│                   ├── Repetir
│                   │   ├── Variable: i
│                   │   ├── Desde: 2
│                   │   ├── Hasta: Expresión Aritmética
│                   │   │   ├── Identificador: n
│                   │   │   ├── Operador: -
│                   │   │   └── Literal: 1
│                   │   └── Cuerpo
│                   │       └── Condicional (si)
│                   │           ├── Condición
│                   │           │   └── Expresión Relacional
│                   │           │       ├── Expresión Aritmética
│                   │           │       │   ├── Identificador: n
│                   │           │       │   ├── Operador: %
│                   │           │       │   └── Identificador: i
│                   │           │       ├── Operador: ↔
│                   │           │       └── Literal: 0
│                   │           └── Rama Verdadera
│                   │               └── Retorno
│                   │                   └── Literal: falso
│                   └── Retorno
│                       └── Literal: verdadero
│
├── Mostrar
│   └── Expresión
│       └── Concatenación de Cadenas
│           ├── "Verificando si "
│           ├── Identificador: numero
│           └── " es primo..."
│
├── Condicional (si)
│   ├── Condición
│   │   └── Llamada Función
│   │       ├── Nombre: verificarPrimo
│   │       └── Argumentos
│   │           └── Identificador: numero
│   ├── Rama Verdadera (entonces)
│   │   └── Mostrar
│   │       └── Expresión
│   │           └── Concatenación de Cadenas
│   │               ├── "El número "
│   │               ├── Identificador: numero
│   │               └── " es primo"
│   └── Rama Falsa (sino)
│       └── Mostrar
│           └── Expresión
│               └── Concatenación de Cadenas
│                   ├── "El número "
│                   ├── Identificador: numero
│                   └── " no es primo"
│
└── Repetir
    ├── Variable: i
    ├── Desde: 2
    ├── Hasta: Identificador: numero
    └── Cuerpo
        └── Condicional (si)
            ├── Condición
            │   └── Llamada Función
            │       ├── Nombre: verificarPrimo
            │       └── Argumentos
            │           └── Identificador: i
            └── Rama Verdadera
                └── Mostrar
                    └── Identificador: i
```

## Representación Textual Simplificada

```
Programa
├── Declaraciones
│   ├── numero = 20
│   ├── contador = 2
│   └── esPrimo = verdadero
│
├── Función: verificarPrimo(n)
│   └── Si n ≤ 1 entonces retornar falso
│   └── Repetir i desde 2 hasta n-1
│       └── Si n % i ↔ 0 entonces retornar falso
│   └── Retornar verdadero
│
├── Mostrar("Verificando si " + numero + " es primo...")
│
├── Si verificarPrimo(numero) entonces
│   ├── Verdadero: Mostrar("El número " + numero + " es primo")
│   └── Falso: Mostrar("El número " + numero + " no es primo")
│
└── Repetir i desde 2 hasta numero
    └── Si verificarPrimo(i) entonces Mostrar(i)
```

---

# H. SIMULACIÓN DE EJECUCIÓN

## Paso 1: Lectura del Archivo

El intérprete lee el archivo fuente `programa.flux` y lo carga en memoria como una secuencia de caracteres:

```
definir numero → 20
definir contador → 2
...
```

## Paso 2: Análisis Léxico (Tokenización)

El lexer procesa el código fuente y genera una secuencia de tokens:

```
[PALABRA_RESERVADA: definir] [IDENTIFICADOR: numero] [OPERADOR: →] [ENTERO: 20]
[PALABRA_RESERVADA: definir] [IDENTIFICADOR: contador] [OPERADOR: →] [ENTERO: 2]
[PALABRA_RESERVADA: definir] [IDENTIFICADOR: esPrimo] [OPERADOR: →] [BOOLEANO: verdadero]
[PALABRA_RESERVADA: función] [IDENTIFICADOR: verificarPrimo] [DELIMITADOR: (] [IDENTIFICADOR: n] [DELIMITADOR: )] ...
```

**Estado del Lexer:**
- Posición: 0
- Línea actual: 1
- Columna actual: 1
- Buffer de tokens: []

## Paso 3: Análisis Sintáctico (Parsing)

El parser construye el AST a partir de los tokens:

1. **Reconoce declaraciones:**
   - `numero → 20` → Nodo Declaración
   - `contador → 2` → Nodo Declaración
   - `esPrimo → verdadero` → Nodo Declaración

2. **Reconoce función:**
   - `función verificarPrimo(n)` → Nodo Función
   - Construye cuerpo de función recursivamente

3. **Reconoce sentencias:**
   - `mostrar(...)` → Nodo Mostrar
   - `si ... entonces ... sino ...` → Nodo Condicional
   - `repetir ... desde ... hasta ...` → Nodo Repetir

**Estado del Parser:**
- AST construido completamente
- Tabla de símbolos inicializada
- Listo para evaluación

## Paso 4: Construcción de Tabla de Símbolos

| Símbolo | Tipo | Valor | Alcance |
|---------|------|-------|---------|
| `numero` | Variable | 20 | Global |
| `contador` | Variable | 2 | Global |
| `esPrimo` | Variable | true | Global |
| `verificarPrimo` | Función | (función) | Global |
| `n` | Parámetro | - | Local (verificarPrimo) |
| `i` | Variable | - | Local (repetir) |

## Paso 5: Evaluación (Ejecución)

### 5.1 Ejecución de Declaraciones

```
Ejecutar: definir numero → 20
  → Tabla de símbolos: numero = 20

Ejecutar: definir contador → 2
  → Tabla de símbolos: contador = 2

Ejecutar: definir esPrimo → verdadero
  → Tabla de símbolos: esPrimo = true
```

### 5.2 Definición de Función

```
Ejecutar: función verificarPrimo(n) hacer ...
  → Función registrada en tabla de símbolos
  → Cuerpo de función almacenado (no ejecutado aún)
```

### 5.3 Ejecución de Mostrar

```
Ejecutar: mostrar("Verificando si " + numero + " es primo...")
  → Evaluar expresión: "Verificando si " + 20 + " es primo..."
  → Resultado: "Verificando si 20 es primo..."
  → Output: Verificando si 20 es primo...
```

### 5.4 Ejecución de Condicional

```
Ejecutar: si verificarPrimo(numero) entonces ...
  → Evaluar condición: verificarPrimo(20)
    → Llamar función verificarPrimo con n = 20
      → Evaluar: si 20 ≤ 1 entonces → false
      → Repetir i desde 2 hasta 19:
          → i = 2: Evaluar 20 % 2 ↔ 0 → true
          → Retornar falso
    → Resultado: false
  → Condición es falsa → Ejecutar rama sino
  → Ejecutar: mostrar("El número " + numero + " no es primo")
  → Output: El número 20 no es primo
```

### 5.5 Ejecución de Repetir

```
Ejecutar: repetir i desde 2 hasta numero hacer ...
  → i = 2:
    → Evaluar: si verificarPrimo(2) entonces
      → Llamar verificarPrimo(2)
        → 2 ≤ 1 → false
        → Repetir i desde 2 hasta 1 → no ejecuta (2 > 1)
        → Retornar verdadero
      → Mostrar(2)
      → Output: 2
  → i = 3:
    → Evaluar: si verificarPrimo(3) entonces
      → Llamar verificarPrimo(3)
        → 3 ≤ 1 → false
        → Repetir i desde 2 hasta 2 → no ejecuta (2 >= 2)
        → Retornar verdadero
      → Mostrar(3)
      → Output: 3
  → ... (continúa hasta i = 20)
```

## Paso 6: Output Final

```
Verificando si 20 es primo...
El número 20 no es primo
Números primos hasta 20:
2
3
5
7
11
13
17
19
```

## Diagrama de Flujo de Ejecución

```
INICIO
  │
  ├─→ [Lexer] Tokenización
  │     │
  │     └─→ Lista de Tokens
  │
  ├─→ [Parser] Construcción AST
  │     │
  │     └─→ Árbol Sintáctico
  │
  ├─→ [Tabla de Símbolos] Registro
  │     │
  │     └─→ Entradas de Variables/Funciones
  │
  ├─→ [Evaluador] Ejecución
  │     │
  │     ├─→ Declaraciones → Tabla de Símbolos
  │     ├─→ Funciones → Registro
  │     ├─→ Mostrar → Output
  │     ├─→ Condicionales → Evaluación de condición
  │     └─→ Repetir → Iteración
  │
  └─→ FIN
```

---

# I. IMPLEMENTACIÓN DEL ANALIZADOR EN GO

## Arquitectura del Sistema

El sistema está compuesto por los siguientes módulos:

1. **Lexer (Analizador Léxico):** Convierte código fuente en tokens
2. **Parser (Analizador Sintáctico):** Construye el AST a partir de tokens
3. **Evaluador:** Ejecuta el programa representado por el AST
4. **Tabla de Símbolos:** Gestiona variables y funciones

## Estructura de Archivos

```
flux/
├── main.go              # Punto de entrada
├── lexer/
│   └── lexer.go         # Analizador léxico
├── parser/
│   └── parser.go        # Analizador sintáctico
├── ast/
│   └── ast.go           # Definiciones del AST
├── evaluator/
│   └── evaluator.go     # Evaluador/e intérprete
└── symbol/
    └── symbol.go        # Tabla de símbolos
```

## Implementación Completa

(Véase archivos de código fuente adjuntos)

---

# CONCLUSIÓN

Este proyecto demuestra un dominio completo de:

1. **Análisis Léxico:** Diseño de tokens, patrones regex, categorización
2. **Análisis Sintáctico:** Gramáticas formales (BNF/EBNF), construcción de AST
3. **Diseño de Lenguajes:** Creación de un lenguaje nuevo con sintaxis coherente
4. **Implementación Práctica:** Código ejecutable en Go

El lenguaje **Flux** representa una síntesis entre legibilidad y expresividad, manteniendo la simplicidad de Go pero con una identidad propia que prioriza la claridad semántica.

---

**Autor:** [Tu Nombre]  
**Fecha:** [Fecha Actual]  
**Proyecto:** Análisis Léxico y Sintáctico - Lenguaje Flux

