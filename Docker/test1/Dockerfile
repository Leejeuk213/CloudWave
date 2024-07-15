

# Debian (bullseye)을 이용한 Dockerfile
# FROM golang:1.19-bullseye
# WORKDIR /app
# COPY src ./
# RUN CGO_ENABLED=0 go build -o main
# CMD ["/app/main"]

# Alpine을 이용한 Dockerfile
# FROM golang:1.19-alpine
# WORKDIR /app
# COPY src ./
# RUN CGO_ENABLED=0 go build -o main
# CMD ["/app/main"]

#Scratch를 이용한 Dockerfile
FROM golang:1.19-alpine as build
WORKDIR /app
COPY src ./
RUN CGO_ENABLED=0 go build -o main
FROM scratch as release
COPY --from=build /app/main /app/
WORKDIR /app
CMD ["/app/main"]