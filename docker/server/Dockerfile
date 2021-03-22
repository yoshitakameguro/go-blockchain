FROM golang:1.16.0-alpine
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
RUN go get github.com/pilu/fresh
CMD ["fresh"]
COPY . .
