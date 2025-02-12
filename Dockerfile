FROM golang:1.22-alpine

WORKDIR /ShortenURL
COPY . .

RUN go build ./

CMD ./ShortenURL