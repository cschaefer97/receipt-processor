FROM golang:1.22.5

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8080
run build -o bin .

ENTRYPOINT ["/app/bin"]