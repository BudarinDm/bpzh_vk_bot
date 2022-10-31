FROM golang:1.18

WORKDIR /src
COPY . .

RUN apk add --no-cache ca-certificates tzdata
RUN apk add gcc libc-dev musl-dev
RUN apk add git
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o app -ldflags "-X 'main.version=${VERSION}'" ./cmd/bpzh-vk-bot.go

EXPOSE 80

RUN apk add --no-cache ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /src/app .

CMD ["./app"]
