# Usamos una imagen base de Go
FROM golang:1.23-alpine

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos el código fuente al contenedor
COPY . .

# Instalamos las dependencias de Go
RUN go mod tidy

# Exponemos el puerto
EXPOSE 3000

# Ejecutamos la aplicación
CMD ["go", "run", "main.go"]
