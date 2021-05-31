FROM golang:alpine AS build

WORKDIR /build

COPY go.mod go.sum ./

COPY . .

RUN GOCGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main

FROM scratch

COPY --from=build /build/main /

CMD ["/main"]

EXPOSE 3000