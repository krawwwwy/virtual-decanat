# PowerShell скрипт для сборки и запуска проекта Виртуальный деканат

param (
    [string]$Command = "help"
)

$dockerComposeFile = "decanat-dev-environment\docker-compose.yml"
$envFile = ".env"

# Функции
function Show-Help {
    Write-Host "Доступные команды:"
    Write-Host "  .\build.ps1 build   - собрать все сервисы"
    Write-Host "  .\build.ps1 run     - запустить все сервисы"
    Write-Host "  .\build.ps1 stop    - остановить все сервисы"
    Write-Host "  .\build.ps1 clean   - удалить все контейнеры и образы"
    Write-Host "  .\build.ps1 test    - запустить тесты"
    Write-Host "  .\build.ps1 demo    - демонстрация работы сервиса аутентификации"
    Write-Host "  .\build.ps1 help    - показать справку"
}

function Build-Services {
    if (-not (Test-Path $envFile)) {
        Write-Host "ПРЕДУПРЕЖДЕНИЕ: Файл $envFile не найден. Создайте его перед запуском."
        return
    }
    
    docker-compose -f $dockerComposeFile build
}

function Run-Services {
    if (-not (Test-Path $envFile)) {
        Write-Host "ПРЕДУПРЕЖДЕНИЕ: Файл $envFile не найден. Создайте его перед запуском."
        return
    }
    
    docker-compose -f $dockerComposeFile up -d
}

function Stop-Services {
    docker-compose -f $dockerComposeFile down
}

function Clean-Environment {
    docker-compose -f $dockerComposeFile down --rmi all --volumes --remove-orphans
}

function Run-Tests {
    Write-Host "Запуск тестов..."
    Set-Location -Path auth-service\backend
    go test ./...
    Set-Location -Path ../..
    # По мере добавления новых микросервисов, здесь будут добавляться команды для запуска их тестов
}

function Run-Demo {
    Write-Host "Запуск демонстрации сервиса аутентификации..."
    .\demo.ps1
}

# Выполнение команды
switch ($Command) {
    "build" { Build-Services }
    "run"   { Run-Services }
    "stop"  { Stop-Services }
    "clean" { Clean-Environment }
    "test"  { Run-Tests }
    "demo"  { Run-Demo }
    default { Show-Help }
} 