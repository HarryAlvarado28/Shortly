# Usar una imagen de Go ligera
FROM golang:1.21-alpine

# Establecer directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar archivos del proyecto al contenedor
COPY . .

# Descargar dependencias y compilar el binario
RUN go mod tidy
RUN go build -o shortly

# Exponer el puerto de la app
EXPOSE 8080

# Ejecutar el binario al iniciar el contenedor
CMD ["./shortly"]
