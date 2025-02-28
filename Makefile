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
# Теперь при вызове make run мы запустим наш сервер

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
# oapi-codegen - обращается к нашей утилите oapi
# -config openapi/.openapi - за конфиг берем файл .openapi из папки openapi
# -inclede-tags tasks - Генерируем описанные ручки под тегом tasks из файла openapi.yaml
# -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go 
# генерируем все описанное по пути interanl/web/tasks. Этот путь вам нужно создать
# Создать в папке internal папку web и в ней папку tasks