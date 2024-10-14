FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go build -o /bin/verse ./src

FROM ubuntu:24.10
COPY --from=build /bin/verse /bin/verse
ENTRYPOINT ["/bin/verse", "-f", "/etc/verse/config.yml"]

