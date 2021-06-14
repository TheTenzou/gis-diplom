# Дипломный проект геоинформационные системы

## Технологии

 - Golang, Kotlin, Java, JavaScript.
 - Gin-gonic, Spring(Spring-Security, Spring-Data, Spring-web), React, Leaflet.
 - Redis, Mongodb, Postgres.

## Архитектура

![Link](https://github.com/TheTenzou/gis-diplom/images/architecture.png?raw=true)

Приложение состоит из следующих сервисов:
 - Управление аккаунтами пользователей осушествляется сервисом аунтефикации, написаном на Go с использованием Gin-gonic. Mongo используется для храниения данных пользователя. Redis - для refresh токенов.
 - ТСОДД сервис CRUD приложение для данных связанных с ТСОДД, написанов на Kotlin с использованием Spring. Postgres используется для хранения данных. Оптимальный план обслуживания ТСОДД составляется припомоши Optaplanner.
 - УДС сервис CRUD приложение для данных связанных с модернизацией УДС, написанов на Java с использованием Spring. Postgres используется для хранения данных. Оптимальный план модернизации УДС составляется припомоши Optaplanner.
 - фронт написан на React с Leaflet для отображения карты.

## Требования

- [Docker](https://www.docker.com/)
- Docker-compose

В случае windows лучше запускать из под wsl

## Запуск
```sh
docker-compose up
```

Приложение запускается на 8080 порте
