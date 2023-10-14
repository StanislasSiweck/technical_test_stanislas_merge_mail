FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o merge-mail

EXPOSE 8181

CMD ["./merge-mail"]