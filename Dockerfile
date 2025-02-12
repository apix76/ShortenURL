FROM golang:1.22-alpine

WORKDIR /ReductionAPI
COPY . .

RUN go build ./

CMD ./ShortenURL