# Etapa de compilación
FROM golang:1.23 AS builder

WORKDIR /app

# Copia go.mod y go.sum primero para aprovechar el cache de Docker
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Compila el binario del servidor
RUN go build -o server_bin ./server

# Imagen final
FROM debian:bookworm-slim

WORKDIR /app

# Copia el binario desde la etapa de build
COPY --from=builder /app/server_bin .

# Expone el puerto gRPC
EXPOSE 50052

# Comando para ejecutar el servidor
CMD ["./server_bin"]
