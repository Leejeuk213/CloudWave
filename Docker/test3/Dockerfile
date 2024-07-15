# Alpine을 이용한 Dockerfile
ARG OS
FROM $OS
ENV BASE=$OS
WORKDIR /app
COPY src ./
RUN CGO_ENABLED=0 go build -o main
CMD ["/app/main"]


