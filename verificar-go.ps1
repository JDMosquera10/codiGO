# Script de Verificación de Instalación de Go
# Ejecutar con: .\verificar-go.ps1

Write-Host "=== Verificacion de Instalacion de Go ===" -ForegroundColor Cyan
Write-Host ""

# Verificar si go está en el PATH
$goPath = (Get-Command go -ErrorAction SilentlyContinue)

if ($goPath) {
    Write-Host "[OK] Go esta instalado y disponible en el PATH" -ForegroundColor Green
    Write-Host "Ubicación: $($goPath.Source)" -ForegroundColor Gray
    Write-Host ""
    
    # Verificar versión
    Write-Host "Version instalada:" -ForegroundColor Yellow
    go version
    Write-Host ""
    
    # Verificar variables de entorno
    Write-Host "Variables de entorno Go:" -ForegroundColor Yellow
    go env GOPATH
    go env GOROOT
    Write-Host ""
    
    Write-Host "[OK] Go esta correctamente instalado. Puedes ejecutar el proyecto." -ForegroundColor Green
} else {
    Write-Host "[ERROR] Go NO esta instalado o no esta en el PATH" -ForegroundColor Red
    Write-Host ""
    
    # Verificar si existe en ubicación por defecto
    $defaultPath = "C:\Program Files\Go\bin\go.exe"
    if (Test-Path $defaultPath) {
        Write-Host "[ADVERTENCIA] Go parece estar instalado en: $defaultPath" -ForegroundColor Yellow
        Write-Host "   Pero no esta en el PATH del sistema." -ForegroundColor Yellow
        Write-Host ""
        Write-Host "Solución:" -ForegroundColor Cyan
        Write-Host "1. Agrega 'C:\Program Files\Go\bin' al PATH del sistema" -ForegroundColor White
        Write-Host "2. O ejecuta temporalmente:" -ForegroundColor White
        Write-Host "   `$env:PATH += ';C:\Program Files\Go\bin'" -ForegroundColor Gray
        Write-Host "   go version" -ForegroundColor Gray
    } else {
        Write-Host "[INFO] Go no esta instalado." -ForegroundColor Yellow
        Write-Host ""
        Write-Host "Pasos para instalar:" -ForegroundColor Cyan
        Write-Host "1. Descarga Go desde: https://go.dev/dl/" -ForegroundColor White
        Write-Host "2. Ejecuta el instalador .msi" -ForegroundColor White
        Write-Host "3. Asegúrate de marcar 'Add to PATH'" -ForegroundColor White
        Write-Host "4. Cierra y vuelve a abrir PowerShell" -ForegroundColor White
        Write-Host ""
        Write-Host "[INFO] Consulta INSTALACION_GO.md para instrucciones detalladas" -ForegroundColor Cyan
    }
    
    Write-Host ""
    Write-Host "Verificando ubicaciones comunes..." -ForegroundColor Yellow
    
    $commonPaths = @(
        "C:\Program Files\Go\bin\go.exe",
        "C:\Program Files (x86)\Go\bin\go.exe",
        "$env:USERPROFILE\go\bin\go.exe"
    )
    
    $found = $false
    foreach ($path in $commonPaths) {
        if (Test-Path $path) {
            Write-Host "  [OK] Encontrado en: $path" -ForegroundColor Green
            $found = $true
        }
    }
    
    if (-not $found) {
        Write-Host "  [ERROR] No se encontro Go en ubicaciones comunes" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "=== Fin de Verificacion ===" -ForegroundColor Cyan

