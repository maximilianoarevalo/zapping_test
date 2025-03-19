# Usamos una imagen base de Go
FROM golang:1.21-alpine

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos el código fuente al contenedor
COPY . .

# Instalamos las dependencias de Go
RUN go mod tidy

# Exponemos el puerto
EXPOSE ${GO_PORT}

# Ejecutamos la aplicación
CMD ["go", "run", "main.go"]
