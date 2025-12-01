# Instrucciones de Compilación y Ejecución

## Requisitos Previos

- **Go 1.21 o superior** instalado
- Terminal/Consola de comandos
- Editor de texto (opcional, para modificar código)

### ⚠️ Si obtienes el error: "go : El término 'go' no se reconoce..."

**Go no está instalado.** Consulta el archivo `INSTALACION_GO.md` para instrucciones detalladas de instalación en Windows.

## Pasos para Ejecutar el Proyecto

### 1. Verificar Instalación de Go

```bash
go version
```

Debería mostrar algo como: `go version go1.21.x windows/amd64`

### 2. Navegar al Directorio del Proyecto

```bash
cd C:\Users\DESAR_JUNIORII\Desktop\workU\codiGO
```

### 3. Compilar el Proyecto

```bash
go build -o flux.exe main.go
```

O simplemente ejecutar sin compilar:

```bash
go run main.go ejemplo.flux
```

### 4. Ejecutar el Analizador

```bash
flux.exe ejemplo.flux
```

O si compiló con otro nombre:

```bash
go run main.go ejemplo.flux
```

## Estructura de Salida Esperada

El programa mostrará:

1. **ANÁLISIS LÉXICO**
   - Lista de tokens encontrados
   - Tipo y valor de cada token
   - Posición (línea, columna)

2. **ÁRBOL SINTÁCTICO**
   - Representación del AST
   - Estructura jerárquica del programa

3. **EJECUCIÓN**
   - Output del programa Flux
   - Resultados de las operaciones

## Crear Nuevos Programas Flux

1. Cree un archivo con extensión `.flux`
2. Escriba código en el lenguaje Flux
3. Ejecute: `go run main.go su_archivo.flux`

## Ejemplo de Programa Flux Mínimo

Cree un archivo `test.flux`:

```flux
definir x → 10
mostrar("El valor es: " + x)
```

Ejecute:

```bash
go run main.go test.flux
```

## Solución de Problemas

### Error: "cannot find package"
```bash
go mod tidy
```

### Error: "syntax error"
- Verifique que el archivo `.flux` tenga sintaxis correcta
- Revise los tokens Unicode (→, ↔, etc.)

### Error: "file not found"
- Verifique la ruta del archivo
- Use rutas relativas desde el directorio del proyecto

## Notas Importantes

1. **Caracteres Unicode:** El lenguaje Flux usa caracteres Unicode especiales (→, ↔, etc.). Asegúrese de que su editor y terminal los soporten.

2. **Encoding:** Los archivos `.flux` deben estar en UTF-8.

3. **Líneas de Comentarios:** Use `//` para comentarios de una línea.

## Documentación Adicional

- `PROYECTO_FINAL.md` - Documentación completa del proyecto
- `README.md` - Guía general del proyecto
- `RESUMEN_EJECUTIVO.md` - Resumen de entregables

