# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Install curl and git 
RUN apk add --no-cache curl git

COPY . .

RUN go mod tidy

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

RUN go build -o main main.go


# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrate ./migration

RUN chmod +x start.sh wait-for.sh migrate main

EXPOSE 8080

ENTRYPOINT [ "/app/start.sh" ]
CMD [ "/app/main" ]
