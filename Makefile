# Makefile для проекта Виртуальный деканат

# Переменные
DOCKER_COMPOSE_FILE = decanat-dev-environment/docker-compose.yml
ENV_FILE = .env

# Цели
.PHONY: build run stop clean test help

help:
	@echo "Доступные команды:"
	@echo "  make build     - собрать все сервисы"
	@echo "  make run       - запустить все сервисы"
	@echo "  make stop      - остановить все сервисы"
	@echo "  make clean     - удалить все контейнеры и образы"
	@echo "  make test      - запустить тесты"
	@echo "  make help      - показать справку"

build:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "ПРЕДУПРЕЖДЕНИЕ: Файл $(ENV_FILE) не найден. Создайте его перед запуском."; \
		exit 1; \
	fi
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

run:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "ПРЕДУПРЕЖДЕНИЕ: Файл $(ENV_FILE) не найден. Создайте его перед запуском."; \
		exit 1; \
	fi
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

clean:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi all --volumes --remove-orphans

test:
	@echo "Запуск тестов..."
	cd auth-service/backend && go test ./...
	# По мере добавления новых микросервисов, здесь будут добавляться команды для запуска их тестов

demo: 
	@echo "Запуск демонстрации сервиса аутентификации..."
	chmod +x demo.sh
	./demo.sh 