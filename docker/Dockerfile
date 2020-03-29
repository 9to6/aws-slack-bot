FROM golang:1.13.5-alpine as builder
WORKDIR /aws-slack-bot
COPY . .

WORKDIR /aws-slack-bot

RUN go build -mod=vendor cmd/aws-slack-bot/main.go

FROM alpine:3

COPY --from=builder /aws-slack-bot/main /bin/aws-slack-bot

CMD [""]
ENTRYPOINT ["/bin/aws-slack-bot"]
