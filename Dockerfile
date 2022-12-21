FROM golang:1.19-alpine

WORKDIR /app

CMD ["tail", "-f", "/dev/null"]