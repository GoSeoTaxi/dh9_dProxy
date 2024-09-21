BINARY_NAME=proxy_creator

MAIN_FILE=cmd/proxy_creator/main.go

OUTPUT_DIR=./bin

GOOS=linux
GOARCH=amd64

build:
	@echo "Компиляция приложения для $(GOOS)/$(GOARCH)..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Бинарный файл создан: $(OUTPUT_DIR)/$(BINARY_NAME)"

clean:
	@echo "Удаление скомпилированного бинарника..."
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)
	@echo "Бинарный файл удален."

$(OUTPUT_DIR):
	@mkdir -​⬤