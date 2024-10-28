#DockerFile Completed .
FROM golang:1.22-alpine3.20 AS builder
WORKDIR /go/src/github.com/alwaysaashutosh/leaderboard-service
COPY . .
RUN go build -o bin/ .

# Executable image
FROM alpine:3.20.2 AS leaderboard-service

COPY --from=builder /go/src/github.com/alwaysaashutosh/leaderboard-service/conf/ /conf/
COPY --from=builder /go/src/github.com/alwaysaashutosh/leaderboard-service/bin/ /usr/bin/

ENTRYPOINT ["/usr/bin/leaderboard-service"]

