# Script para agregar Go al PATH permanentemente
# Ejecutar como Administrador: .\SOLUCION_PATH.ps1

Write-Host "=== Agregando Go al PATH del Sistema ===" -ForegroundColor Cyan
Write-Host ""

$goPath = "C:\Program Files\Go\bin"

# Verificar si Go existe
if (-not (Test-Path $goPath)) {
    Write-Host "[ERROR] Go no se encuentra en: $goPath" -ForegroundColor Red
    Write-Host "Por favor, instala Go primero desde: https://go.dev/dl/" -ForegroundColor Yellow
    exit 1
}

Write-Host "[OK] Go encontrado en: $goPath" -ForegroundColor Green

# Obtener PATH actual del sistema
$currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")

# Verificar si ya está en el PATH
if ($currentPath -like "*$goPath*") {
    Write-Host "[INFO] Go ya esta en el PATH del sistema" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Para aplicar los cambios:" -ForegroundColor Cyan
    Write-Host "1. Cierra esta ventana de PowerShell" -ForegroundColor White
    Write-Host "2. Abre una nueva ventana de PowerShell" -ForegroundColor White
    Write-Host "3. Ejecuta: go version" -ForegroundColor White
    exit 0
}

# Agregar al PATH
try {
    Write-Host "Agregando Go al PATH del sistema..." -ForegroundColor Yellow
    [Environment]::SetEnvironmentVariable(
        "Path",
        $currentPath + ";$goPath",
        "Machine"
    )
    
    Write-Host "[OK] Go agregado al PATH exitosamente!" -ForegroundColor Green
    Write-Host ""
    Write-Host "Para aplicar los cambios:" -ForegroundColor Cyan
    Write-Host "1. Cierra esta ventana de PowerShell" -ForegroundColor White
    Write-Host "2. Abre una nueva ventana de PowerShell" -ForegroundColor White
    Write-Host "3. Ejecuta: go version" -ForegroundColor White
    Write-Host ""
    Write-Host "O ejecuta temporalmente en esta sesion:" -ForegroundColor Cyan
    Write-Host "`$env:PATH += ';$goPath'" -ForegroundColor Gray
    Write-Host "go version" -ForegroundColor Gray
    
} catch {
    Write-Host "[ERROR] No se pudo modificar el PATH. Asegurate de ejecutar como Administrador." -ForegroundColor Red
    Write-Host "Error: $_" -ForegroundColor Red
    exit 1
}

