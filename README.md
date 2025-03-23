
# Zapping Test

## Stack de tecnologías

- Backend: Golang 1.23
- Frontend: Vue3 3.2.13
- Base de datos: PostgreSQL 13

## Observación

- Para manejar los valores de puertos, credenciales de base de datos, etc. sería ideal utilizar un archivo .env de secrets/configmap. Sin embargo, como este proyecto es solo a nivel local y esta dockerizado, no se implementa esta lógica por simplicidad de accesos a la hora probar, ya que se decide evitar el manejo de un archivo .env externo que se deba incluir de manera manual al proyecto.
- Las contraseñas se manejan con encriptado básico, ya que aunque se trate de una prueba local, se debe proteger la data sensible de los usuarios.

## Rutas backend

- /create-user -> Creación de cuentas de usuario backend
- /login -> Verificación de credenciales de cuenta registrada 
- /livestream -> Obtención de archivo m3u8 con 3 últimos segmentos actualizados
- /segments/ -> Obtención de segmentos a través de su nombre.ts

## Rutas frontend

- / -> Verificación de credenciales de cuenta registrada a través del navegador (login)
- /crear-cuenta -> Creación de cuenta de usuario nuevo 
- /player -> Reproductor de video en streaming (ruta protegida por token de autenticación)

## Comandos de ejecución

### Construir imágenes y levantar proyecto
```
docker-compose up --build
```

### Levatar proyecto ya construido
```
docker-compose up 
```

### Bajar proyecto
```
docker-compose down
```

### Bajar y eliminar proyecto
```
docker-compose down -v
```

## Para acceder a la aplicación se debe ingresar a la ruta:
```
http://localhost:3001
```

## Puertos del proyecto

- Frontend: http://localhost:3001
- Backend: http://localhost:3000
- Base de datos: http://localhost:5432
