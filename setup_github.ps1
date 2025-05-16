# Скрипт для настройки GitHub репозитория для проекта Виртуальный деканат

# Имя репозитория
$repo_name = "virtual-decanat"

Write-Host "Настройка GitHub репозитория для проекта $repo_name..."

# Проверяем, инициализирован ли git
if (-not (Test-Path -Path ".git")) {
    Write-Host "Инициализация git репозитория..."
    git init
}

# Добавляем все файлы
Write-Host "Добавление всех файлов..."
git add .

# Создаем первый коммит
Write-Host "Создание первого коммита..."
git commit -m "Инициализация проекта Виртуальный деканат"

# Инструкции для пользователя по созданию репозитория на GitHub
Write-Host @"

Для загрузки проекта на GitHub выполните следующие шаги:

1. Создайте новый репозиторий на GitHub:
   - Откройте https://github.com/new
   - Введите имя репозитория: $repo_name
   - Не инициализируйте репозиторий README, .gitignore или лицензией
   - Нажмите 'Создать репозиторий'

2. Выполните следующие команды, чтобы связать локальный репозиторий с GitHub:
   git remote add origin https://github.com/YOUR_USERNAME/$repo_name.git
   git branch -M main
   git push -u origin main

Замените YOUR_USERNAME на ваше имя пользователя GitHub.
"@ 