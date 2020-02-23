# agentApi

Servicio encargado de guardar los datos del request recibido, en un archivo en formato json.

#### - Herramientas y Librerias Globales :

El DockerFile contiene la instalacíon librerias o herramientas globales para el proyecto.

Librerias Globales:

- Mux: Permite el manejo de peticiones Http.
  https://github.com/gorilla/mux

#### - Docker :

- Ejecutar container: `docker-compose up --build`

#### - Estructura de carpetas :
Listado con la separación de namespace (y directorios) :

- **app**: Carpeta que contiene los archivos con la lógica del servicio.
  - **main**: Archivo en el cual esta configurada la ruta y por el cual se recibe la request.
  - **serverinfo**: Lógica de parseo y guardado de los datos recibidos a través de la request, también posee la estructura de los datos a guardar.
