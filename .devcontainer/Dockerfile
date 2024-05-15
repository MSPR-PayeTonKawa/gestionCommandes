FROM golang:1.22-bullseye

WORKDIR /app

COPY go.* ./

RUN go mod download && go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air"]
