## Cài đặt & kiểm tra Docker

docker --version
docker ps

## Viết Dockerfile

## Build image

docker build -t gin-backend:latest .

## Run container từ image

docker run -p 8080:8080 gin-backend

# Chạy container ở background

docker run -d -p 8080:8080 --name my-app gin-backend

## Kiểm tra container hiện dang chay

docker ps -a

## Nếu muốn dừng

docker stop gin-app

# Xem logs

docker logs

## xóa container cũ

docker rm gin-app

# Xem danh sách images

docker images

# Xóa image

docker rmi gin-backend

## Chạy với .env file

docker run -p 8080:8080 --env-file .env --name gin-app gin-backend

## Nếu muốn giữ container cũ mà tạo thêm cái mới

docker run -p 8080:8080 --env-file .env --name gin-app-v2 gin-backend

#### ---------------------------------------------------------

# docker-gin-crud

## Run locally (Go)

go run main.go
Open http://localhost:8080/users

## Run with Docker

docker build -t gin-backend .
docker run -p 8080:8080 --env-file .env gin-backend

## Run with docker-compose (dev)

docker compose up

## Test

GET /users
POST /users { "name": "Bob" }

###### ---------------------------------------------------------------------------------------------------------

## Dừng container → docker stop gin-app

## Xoá container → docker rm gin-app

## Dừng tất cả → docker stop $(docker ps -q)
