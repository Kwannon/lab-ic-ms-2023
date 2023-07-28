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
