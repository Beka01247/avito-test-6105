# Тестовый бэкенд для Avito с использованием Docker

Этот проект представляет собой бэкенд API, написанный на Go с базой данных PostgreSQL. Проект контейнеризирован с использованием Docker и управляется с помощью Docker Compose.

## Структура проекта

```bash
/project-root
   /internal       # Внутренняя структура Go-приложения
      /db          # Логика подключения к базе данных
      /config      # Файлы конфигурации и переменные окружения
      /controllers # HTTP-контроллеры для API
      /routes      # Маршруты API
          routes.go
   /cmd            # Точка входа в приложение
      main.go
   Dockerfile      # Docker-конфигурация для сборки Go-приложения
   docker-compose.yml # Конфигурация Docker Compose для всего стека
   go.mod
   go.sum
```

Как запустить проект
Шаги:

1. Установлен Docker и Docker Compose.

2. Склонируйте репозиторий на локальную машину:

```bash
git clone https://github.com/yourusername/project.git
cd project
```

Создать .env файл в корне проекта с переменными окружения:

```bash
makefile
SERVER_ADDRESS=8080
POSTGRES_CONN=postgres://postgres:240219@localhost:5432/test_avito
POSTGRES_USERNAME=postgres
POSTGRES_PASSWORD=240219
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DATABASE=test_avito
```

Соберать и запустить контейнеры с помощью Docker Compose:

```bash
bash
docker-compose up --build
```

Приложение будет доступно по адресу: http://localhost:8080.

## Конфигурация Docker

В файле docker-compose.yml описаны два сервиса:

1. PostgreSQL:
   Используется образ postgres:latest.
   Порт: 5432.
   Данные сохраняются в volume database_dockerizing.
2. API:
   Сборка образа с помощью Dockerfile.
   Порт: 8080.
   Взаимодействие с базой данных через PostgreSQL.

## Команды

Сборка проекта:

```bash
docker-compose build
```

Запуск проекта:

```bash
docker-compose up
```

Остановка проекта:

```bash
docker-compose down
```

Переменные окружения
Все переменные окружения находятся в .env файле:

```bash
SERVER_ADDRESS — адрес сервера (по умолчанию 8080).
POSTGRES_CONN — строка подключения к базе данных.
POSTGRES_USERNAME — имя пользователя PostgreSQL.
POSTGRES_PASSWORD — пароль PostgreSQL.
POSTGRES_HOST — хост базы данных.
POSTGRES_PORT — порт базы данных.
POSTGRES_DATABASE — название базы данных.
```

## Деплой

Приложение задеплоено и доступно по адресу:
[https://cnrprod1725724653-team-77354-32538.avito2024.codenrock.com/](https://cnrprod1725724653-team-77354-32538.avito2024.codenrock.com/)
