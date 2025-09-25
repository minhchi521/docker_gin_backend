# Sử dụng base image golang trên Alpine Linux cho build stage 
FROM golang:1.24.4-alpine AS builder

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy file dependencies trước để tận dụng Docker layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Sử dụng Alpine làm base image cho production (nhẹ hơn)
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# --- final stage ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .

# Thông báo container sử dụng port 8080
EXPOSE 8080

# Lệnh mặc định khi container start
CMD ["./server"]
