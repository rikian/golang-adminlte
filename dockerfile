FROM golang:1.18-alpine

WORKDIR /Users/Rikian/Desktop/go-lte

COPY . .

RUN go mod tidy
RUN go build -o ./adminlte-gin .

EXPOSE 9093

CMD ["./adminlte-gin"]