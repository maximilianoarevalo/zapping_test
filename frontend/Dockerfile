# Usamos una imagen base de Node.js
FROM node:16-alpine

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos de package.json e instalamos las dependencias
COPY package*.json ./
RUN npm install

# Copiamos el resto del código fuente
COPY . .

# Exponemos el puerto donde Vue correrá
EXPOSE 8081

# Ejecutamos el servidor de desarrollo de Vue
CMD ["npm", "run", "serve"]