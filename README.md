# Сервис динамического сегментирования пользователей
Тестовое задание для Avito стажировки


# Запуск

> docker compose up -d

Сервис будет доступен на порту 8080.

Описание методов можно найти в `swagger.json` файле.

# Примеры запросов
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
    "avito_discount_50"
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

