# Solución al Error: "go : El término 'go' no se reconoce"

## Problema Resuelto ✅

**Diagnóstico:** Go está instalado en `C:\Program Files\Go\bin\go.exe` pero no está en el PATH del sistema.

## Solución Temporal (Sesión Actual)

Para esta sesión de PowerShell, ejecuta:

```powershell
$env:PATH += ';C:\Program Files\Go\bin'
```

Luego verifica:
```powershell
go version
```

Deberías ver: `go version go1.25.4 windows/amd64`

## Solución Permanente (Recomendada)

Para que Go funcione en todas las sesiones futuras:

### Método 1: Variables de Entorno del Sistema (GUI)

1. Presiona `Win + R`
2. Escribe: `sysdm.cpl` y presiona Enter
3. Ve a la pestaña **"Opciones avanzadas"**
4. Click en **"Variables de entorno"**
5. En **"Variables del sistema"**, busca **"Path"** y haz clic en **"Editar"**
6. Click en **"Nuevo"**
7. Agrega: `C:\Program Files\Go\bin`
8. Click en **"Aceptar"** en todos los diálogos
9. **Cierra y vuelve a abrir PowerShell**

### Método 2: PowerShell (Administrador)

Abre PowerShell como **Administrador** y ejecuta:

```powershell
[Environment]::SetEnvironmentVariable(
    "Path",
    [Environment]::GetEnvironmentVariable("Path", "Machine") + ";C:\Program Files\Go\bin",
    "Machine"
)
```

Luego cierra y vuelve a abrir PowerShell.

### Método 3: Verificación Rápida

Ejecuta el script de verificación:
```powershell
.\verificar-go.ps1
```

## Ejecutar el Proyecto

Después de agregar Go al PATH:

```powershell
# Navegar al directorio
cd C:\Users\DESAR_JUNIORII\Desktop\workU\codiGO

# Ejecutar el proyecto
go run main.go ejemplo.flux
```

## Notas Importantes

- **Siempre cierra y vuelve a abrir PowerShell** después de modificar el PATH
- El cambio solo afecta nuevas sesiones de PowerShell
- Si usas CMD, también necesitarás cerrarlo y abrirlo de nuevo

## Verificar que Funciona

```powershell
# Verificar versión
go version

# Verificar ubicación
where.exe go

# Debería mostrar: C:\Program Files\Go\bin\go.exe
```

---

**Estado:** ✅ Go está instalado correctamente, solo necesita estar en el PATH

