# Название Микросервиса

> [!NOTE]
> В процессе разработки

## Описание

Краткое описание микросервиса и его назначение.

## Стек разработки

Список используемых технологии

## Установка

- Установка
- развертывание
- `make` команды

Пример: 

```bash
git clone https://gitlab.toledo24.ru/web/ms_layout.git
```

Изменить внутри `Makefile` и `.scripts/deploy.sh` переменную `PROJECT_NAME` на названия микросервиса. Если в проекте используется `docker` настройте его. После этого запустить `make init`

Убрать из комментария в `.gitignore` файл `.env`.

Запустить `make remove-readme` - удаление ненужных README.md

Основные команды в `Makefile`:

```
Usage:
  help             print this help message
  init             used to initialize the Go project, tidy, docker, migration, build and deploy
  deploy           executing the deployment command
  fast-start       quick launch of ms_layout
  start            build start of $(PROJECT_NAME)
  migration-up     start the migration stage with the database
  migration-down   down the migration with the database
  build            build a project
  lint             format and golangci-lint the project
  test             start test
  start-server     start systemctl server
  stop-server      stop systemctl server
  remove-readme    delete all files README.md
```

## Использование

Примеры запросов к API и описание конечных точек:

- GET /api/endpoint - Описание запроса
- POST /api/endpoint - Описание запроса
