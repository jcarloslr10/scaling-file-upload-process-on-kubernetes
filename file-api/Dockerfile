FROM golang:1.19.2-alpine

WORKDIR /app

COPY . .

RUN go install -v .

RUN go build -o .

EXPOSE 8080

CMD [ "./file-api" ]