FROM golang:1.18.1
RUN mkdir app
WORKDIR /app
COPY ./ ./
RUN go build -o main ./cmd
CMD ["./main"]