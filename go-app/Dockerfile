FROM golang:1.22 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o go-app

FROM gcr.io/distroless/static-debian12:latest-amd64

COPY --from=build /app/go-app .

EXPOSE 8080

CMD ["./go-app"]