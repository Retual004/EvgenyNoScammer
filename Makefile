# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:77788999@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go 

generate:
	@echo "Генерация кода из спецификации OpenAPI..."
	oapi-codegen -config openapi/.openapi.tasks -include-tags tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi.users -include-tags users openapi/openapi.yaml > ./internal/web/users/api.gen.go


git:
	git add .
	git commit -m "$(commit)"
	git push

lint:
	golangci-lint run --out-format=colored-line-number

# oapi-codegen - обращается к нашей утилите oapi
# -config openapi/.openapi - за конфиг берем файл .openapi из папки openapi
# -inclede-tags tasks - Генерируем описанные ручки под тегом tasks из файла openapi.yaml
# -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go 
# генерируем все описанное по пути interanl/web/tasks. Этот путь вам нужно создать
# Создать в папке internal папку web и в ней папку tasks