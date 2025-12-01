# Guía de Instalación de Go en Windows

## Problema Detectado
El error `go : El término 'go' no se reconoce...` indica que Go no está instalado o no está en el PATH del sistema.

## Solución: Instalar Go

### Opción 1: Instalación Manual (Recomendada)

1. **Descargar Go**
   - Visita: https://go.dev/dl/
   - Descarga el instalador para Windows (archivo `.msi`)
   - Ejemplo: `go1.21.x.windows-amd64.msi`

2. **Ejecutar el Instalador**
   - Haz doble clic en el archivo `.msi` descargado
   - Sigue el asistente de instalación
   - **Importante:** Asegúrate de que la opción "Add to PATH" esté marcada
   - La instalación por defecto es en `C:\Program Files\Go`

3. **Verificar la Instalación**
   - Abre una **nueva** ventana de PowerShell (cierra y vuelve a abrir)
   - Ejecuta:
     ```powershell
     go version
     ```
   - Deberías ver algo como: `go version go1.21.x windows/amd64`

### Opción 2: Instalación con Chocolatey (Si lo tienes instalado)

```powershell
choco install golang
```

### Opción 3: Instalación con Winget (Windows 10/11)

```powershell
winget install GoLang.Go
```

## Verificar PATH

Si Go está instalado pero aún no funciona:

1. **Verificar que Go esté en el PATH:**
   ```powershell
   $env:PATH -split ';' | Select-String -Pattern "Go"
   ```

2. **Si no aparece, agregar manualmente:**
   - Presiona `Win + R`
   - Escribe: `sysdm.cpl` y presiona Enter
   - Ve a la pestaña "Opciones avanzadas"
   - Click en "Variables de entorno"
   - En "Variables del sistema", busca "Path" y haz clic en "Editar"
   - Agrega: `C:\Program Files\Go\bin`
   - Acepta todos los diálogos
   - **Cierra y vuelve a abrir PowerShell**

## Verificar Instalación Completa

Después de instalar, ejecuta estos comandos:

```powershell
# Verificar versión
go version

# Verificar variables de entorno
go env GOPATH
go env GOROOT

# Verificar que go está en el PATH
where.exe go
```

## Después de Instalar Go

Una vez que Go esté instalado correctamente:

1. **Navega al directorio del proyecto:**
   ```powershell
   cd C:\Users\DESAR_JUNIORII\Desktop\workU\codiGO
   ```

2. **Inicializa el módulo Go (si es necesario):**
   ```powershell
   go mod init flux
   go mod tidy
   ```

3. **Ejecuta el proyecto:**
   ```powershell
   go run main.go ejemplo.flux
   ```

## Solución de Problemas

### Error: "go: cannot find module"
```powershell
go mod tidy
```

### Error: "go: command not found" después de instalar
- Cierra y vuelve a abrir PowerShell
- Verifica que Go esté en el PATH
- Reinicia el sistema si es necesario

### Verificar instalación manualmente
```powershell
# Verificar que el ejecutable existe
Test-Path "C:\Program Files\Go\bin\go.exe"

# Si existe pero no funciona, agregar al PATH temporalmente
$env:PATH += ";C:\Program Files\Go\bin"
```

## Enlaces Útiles

- **Descarga oficial:** https://go.dev/dl/
- **Documentación:** https://go.dev/doc/
- **Guía de instalación:** https://go.dev/doc/install

## Nota Importante

Después de instalar Go, **siempre cierra y vuelve a abrir PowerShell** para que los cambios en el PATH surtan efecto.

