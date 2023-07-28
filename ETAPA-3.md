# Etapa 3

Al fin nuestro primer micro servicio.

Corresponde a una API REST que despliega las películas.

La api tiene los siguientes endpoints:

    - `GET /movies` : lista todas las películas disponibles en la base de datos
    - `GET /movies/:id` : obtiene una película basado en su identificador

Este micro servicio está escrito en Go y se encuentra en la carpeta `movies-api`.

## Requisitos

Haber completado la Etapa 2.

### Preguntas

Revisa el archivo `movies-api/Dockerfile`.

¿Qué te llama la atención?

Revisa el archivo `docker-compose.yml`.

¿Cómo se relacionan el archivo `docker-compose.yml` y el archivo `movies-api/Dockerfile`?

¿Qué crees que hace el atributo `context` debajo de `build` (está en la linea 6 del archivo `docker-compose.yml`)?

Recuerda registrar tus respuestas en el archivo `RESPUESTAS.md`

## Actividades

Modifica tu archivo .env agregando dos variables:

    - `BIND_IP`: asignale el valor `movies-api`
    - `BIND_PORT`: asignale el valor `8000` o el valor que te indicó el profesor si estás ejecutando esto en clases.

Con esto defines la ip y el puerto que usa el microservicio.

Intenta reiniciar la aplicación:

```
	docker-compose up --build
```

Si te fijas vas a obtener un error en `movies_api_container` indicando que no puede inicializar la conexión a la base de datos.
Esto ocurre porque falta definir algunas variables de ambiente en el archivo `docker-compose.yml`, corríge el error agregando las variables que faltan en el arreglo `environment`

Reinicia la aplicación ejecutando: 

```
	docker-compose up --build -d
```

Fíjate que agregamos el parámetro `-d` eso libera la consola y deja los contenedores corriendo en background.

Para detener los containers debes hacer `docker-compose down`.

Ahora navega a la dirección `http://localhost:8000/movies`, deberías poder ver la lista de películas.

Tip: te recomiendo usar la herramienta [curl](https://curl.se) de este modo:

    curl -v http://localhost:8000/movies


Nota: si estás corriendo esto en el laboratorio usa la URL que te indique el profesor.

Si todo está ok haz un commit con la corrección y el archivo `RESPUESTAS.md`.

No olvides hacer un push de los cambios a tu fork.


## Opcional

Intenta cambiar el puerto a 8080.

¿Qué pasa si en vez `movies-api` usas `localhost` para la variabla `BIND_IP`?


