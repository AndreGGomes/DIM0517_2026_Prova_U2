FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod .
COPY src ./src
RUN CGO_ENABLED=0 go build -o app ./src

FROM alpine:3.21
WORKDIR /app
COPY --from=build /app/app .
EXPOSE 8080
CMD ["./app"]
