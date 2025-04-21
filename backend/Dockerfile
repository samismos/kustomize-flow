FROM golang

WORKDIR /app

COPY  . /app

RUN go mod tidy

EXPOSE 8080

ENTRYPOINT ["go", "run", "cmd/main.go"]