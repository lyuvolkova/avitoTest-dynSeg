FROM golang:1.21-alpine as build
WORKDIR /build
COPY . /build
RUN go build -o /build/service cmd/service/main.go

FROM alpine
WORKDIR /app
COPY --from=build /build/service /app
CMD ["/app/service"]
