FROM golang:1.22.5

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o receipt-processor

EXPOSE 8080

CMD [ "/app/receipt-processor" ]