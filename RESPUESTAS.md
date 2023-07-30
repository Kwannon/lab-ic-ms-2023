# Respuestas

Indica tu nombre a continuación: 

Por cada etapa agrega una sección abajo y escribe las respuestas a las preguntas de cada etapa.

## ETAPA 1

* ¿Cuál es la diferencia entre los archivos con el verbo Create con los archivos con el verbo Add?
  
    Los archivos Create tienen los comandos SQL para crear las tablas, mientras los Add tienen las sentencias para insertar datos en las tablas recien creadas

* ¿Cómo se llama el servicio que se declara en el archivo docker-compose.yml?
  
    flyway

* ¿Cuál es el comando que se ejecuta en el servicio declarado?
  
    migrate

## ETAPA 2

* ¿Qué pasa si cambias el nombre del servicio de postgres a db? ¿Qué otros cambios tendrías que hacer?

    Si se cambia el nombre del servicio tambien hay que actualizar en ´docker-compose´ el valor para "depends_on" y en el archivo .env el valor de "POSTGRES_SERVER"

## ETAPA 3

  **Revisa el archivo movies-api/Dockerfile.**

  * ¿Qué te llama la atención?

    El archivo Dockerfile contiene las instrucciones para levantar la app contenida en "movies-api"

  **Revisa el archivo docker-compose.yml.**

  * ¿Cómo se relacionan el archivo docker-compose.yml y el archivo movies-api/Dockerfile?

    Desde el archivo docker-compose se invoca y levanta el contenedor y la API contenida en movies-api 

  * ¿Qué crees que hace el atributo context debajo de build (está en la linea 6 del archivo docker-compose.yml)?

    Indica el directorio donde puede encontrar los recursos a que se hace referencia para levantar movies-api


## Etapa 4

  * Compara los archivos `Dockerfile` de `movies-api` y `movies-front`.

    Son muy distintos. Mientras el del front depende de NPM para levantar la app basicamente, el otro contruye e instala el backend instruccion por instruccion.

  * Compara el atributo `build` del servicio `movies-api` con el de `movies-front`. 
¿Cuál es la diferencia? 
¿Qué pasa si los dejas iguales?

   Son diferentes maneras de hacer referencia a una ubicación, si se dejan igual sigue funcionando

