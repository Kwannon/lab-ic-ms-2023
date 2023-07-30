# Etapa 4

Tenemos frontend \o/.

El equipo front ha agregado una aplicación que permite ver y mantener películas y directores.

Se trata de una aplicación escrita en REACT y requiere Node.js para operar.

Esta aplicación interopera con el microservicio en Go definido en la etapa 3.

El archivo docker-compose ahora incorpora un nuevo servicio llamado: `movies-front`.

## Actividades

Deten el servicio:

```
    docker-compose down 
```

Configura la variable `REACT_PORT` en tu archivo `.env` con un valor distinto al usado por `movies-api`, por ejemplo: 8001.

Reinicia el servicio:

```
    docker-compose up -d 
```

Fijate que hemos eliminado el parametro --build para no tener que reconstruir innecesariamente todas las imágenes

Verifica que el servidor rest está funcionando:

```
    curl -v http://localhost:8000/movies
```

(Usa el puerto que definiste en la etapa 3)

Verifica que el nuevo endpoint `directors` funciona

```
    curl -v http://localhost:8000/directors
```

Nota: Si no tienes `curl` en tu sistema puedes acceder a las urls en tu browser.

Ahora intenta acceder a la aplicación a través de tu navegador en la dirección `http://localhost:REACT_PORT` (o la que te indique tu profesor).

Este intento va a fallar porque debes configurar el puerto de salida con los atributos `expose` y `ports`.

Para solucionar esto fíjate en cómo lo hacemos con `movies-api`, revisa los atributos `expose` y `ports`  definidos bajo el servicio `movies-api` (lineas 24 a 27 del archivo `docker-compose.yml`).

Reinicia el servicio con `docker-compose up` y trata de acceder otra vez a la aplicación web con el browser.

Podras ver que a pesar de que los endpoints funcionan y retornan los valores esperados, el front-end no muestra datos.

La razón es que falta configurar una variable en el servicio `movies-front`, la variable perdida está definida en el archivo `movies-front/src/const.js`.

El valor que debes dar a esa variable es `http://localhost:BIND_PORT`, donde `BIND_PORT` es el valor que has definido para el servidor `movies-api`. Configura esta variable debajo de la variable `PORT` y reinicia el servicio.

Si has configurado todo correctamente vas a ver los valores de las películas y los directores en la aplicación.

Intenta agregar directores y películas en tu aplicación.

Nota: para obtener el valor de las variables de ambiente en node usamos `process.env.VARIABLE`.

## Preguntas

Compara los archivos `Dockerfile` de `movies-api` y `movies-front`. 

Compara el atributo `build` del servicio `movies-api` con el de `movies-front`. 
¿Cuál es la diferencia? 
¿Qué pasa si los dejas iguales?

