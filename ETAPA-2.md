# Etapa 2

Por razones de seguridad nos han prohibido usar `ElephantSQL`, así que nos han autorizado crear una instancia Postgres en nuestro ambientes.

En esta etapa vamos a crear un servicio `postgres` en nuestro `docker-compose.yml`

## Actividades

Modifica el archivo `docker-compose.yml`, dejándolo así:

```
version: '3'
services:
  postgres:
    image: postgres:15.3
    restart: always
    environment:
      - POSTGRES_USER
      - POSTGRES_PASSWORD
      - POSTGRES_DB

  flyway:
    image: flyway/flyway:latest
    container_name: migration_flyway
    environment:
      - FLYWAY_USER=${POSTGRES_USER}
      - FLYWAY_PASSWORD=${POSTGRES_PASSWORD}
      - FLYWAY_URL=jdbc:postgresql://${POSTGRES_SERVER}:${POSTGRES_PORT}/${POSTGRES_DB}
      - FLYWAY_GROUP=true
    command: -locations=filesystem:/flyway/sql -connectRetries=60 migrate
    volumes:
      - ./sql_migrations:/flyway/sql
    depends_on:
      - postgres
```

Modifica el archivo .env, esta vez debes cambiar al menos la variable `POSTGRES_SERVER` que ahora debe tener el valor `postgres`.
De este modo tus migraciones se aplicarán en tu instancia local.

Ejecuta la migración igual que en la ETAPA-1:

```
	docker-compose up --build
```

Atención: si estabas ejecutando `docker-compose`` debes detenerlo con `ctrl-c`.

Fíjate en los logs y asegurate que se ejecuta correctamente la migración. 

Tip: elimina la instancia en ElephantSQL.


## Preguntas

¿Qué pasa si cambias el nombre del servicio de `postgres` a `db`? ¿Qué otros cambios tendrías que hacer?
