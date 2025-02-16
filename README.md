## Requisitos

- Go 1.16 o superior
- PostgreSQL
- MongoDB
- MySQL
- SQLite

## Instalación

1. Clona el repositorio:

    ```sh
    git clone https://github.com/tu-usuario/masivo.git
    cd masivo
    ```

2. Instala las dependencias:

    ```sh
    go mod tidy
    ```

## Uso

1. Configura las conexiones a las bases de datos en los archivos correspondientes en el directorio [repositories](http://_vscodecontentref_/3).

2. Ejecuta el programa:

    ```sh
    go run main.go <tipo_repositorio>
    ```

    Donde `<tipo_repositorio>` puede ser `sqlite`, `placebo`, `pgx`, `mongo` o `mysql`.

3. Elige el repositorio que deseas utilizar pasando el tipo de repositorio como argumento al ejecutar el programa:

    ```sh
    go run main.go sqlite
    ```

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o envía un pull request.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.