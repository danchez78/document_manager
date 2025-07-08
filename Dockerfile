FROM golang:1.24

WORKDIR /app

COPY . ./

RUN go build cmd/application/main.go

RUN chmod 700 ./entrypoint.sh
EXPOSE 50000
ENTRYPOINT ["./entrypoint.sh"]
