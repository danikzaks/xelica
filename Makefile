# Название проекта
PROJECT_NAME := xelica

# Директории для бинарных файлов
BIN_DIR := ./bin

# Названия бинарных файлов
BIN_NAME := $(PROJECT_NAME)

# Путь к Go модулям
GO := go
GOFMT := gofmt
GOTEST := go test
BUILD_FLAGS := -ldflags="-s -w"

# Путь к основной директории с исходным кодом
MAIN_DIR := ./cmd

# Определения целей
.PHONY: all build run fmt test clean

# Основная цель
all: build

# Сборка приложения
build:
	@echo "Сборка приложения..."
	$(GO) build $(BUILD_FLAGS) -o $(BIN_DIR)/$(BIN_NAME) $(MAIN_DIR)

# Запуск приложения
run: build
	@echo "Запуск приложения..."
	$(BIN_DIR)/$(BIN_NAME)

# Форматирование кода
fmt:
	@echo "Форматирование кода..."
	$(GOFMT) -w .

# Запуск тестов
test:
	@echo "Запуск тестов..."
	$(GOTEST) ./...

# Очистка временных файлов
clean:
	@echo "Удаление бинарных файлов..."
	rm -rf $(BIN_DIR)/$(BIN_NAME)
