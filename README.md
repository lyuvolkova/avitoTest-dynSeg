# Сервис динамического сегментирования пользователей
Тестовое задание для Avito стажировки.


# Запуск
Для запуска требуется установленные программы:
* Docker и docker compose
* Make
* Curl

Команда для запуска сервиса:
> docker compose up -d

Сервис будет доступен на порту 8080.

Описание методов можно найти в `swagger.json` файле.


# Тесты
Перед запуском тестов необходимо запустить сервис.

> make test

# Примеры запросов

Для IDE Goland можно использовать примеры из папки `examples`.
## Создание сегмента

``` 
curl -X 'POST' \
  'http://localhost:8080/segments' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "slug": "avito_discount_90"
}'
```
## Удаление сегмента
``` 
curl -X 'DELETE' \
  'http://localhost:8080/segments/avito_discount_70' \
  -H 'accept: application/json'
``` 
## Добавление пользователя в сегмент
``` 
curl -X 'PATCH' \
  'http://localhost:8080/users/89/segments' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "add_slug": [
    "avito_discount_70",
    "avito_discount_90"
  ],
  "delete_slug": [
    "avito_discount_30"
  ]
}'
```
## Получение активных сегментов пользователя
```
curl -X 'GET' \
  'http://localhost:8080/users/89/segments' \
  -H 'accept: application/json'
```

# Особенности работы сервиса

1. Если сегмент не существует, то при добавлении/удалении пользователя в сегмент, сервис не вернет ошибку
и не добавит/удалит его в этот сегмент.
2. Тест в начале и конце работы очищает сегменты и пользователей в них для чистого прогона.
3. Client для тестов сгенерирован с помощью go-swagger. В сгенерированном коде есть ошибки с полем `Error`.
   Если запускать команду `make swagger`, client будет перезаписан с ошибками и их необходимо исправить вручную.
