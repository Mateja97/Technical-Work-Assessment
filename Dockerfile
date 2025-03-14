FROM golang:1.23.6
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /alluvial-task


EXPOSE 8080

CMD ["/alluvial-task"]