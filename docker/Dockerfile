FROM golang:1.18

WORKDIR /src
COPY . .

RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o app -ldflags "-X 'main.version=${VERSION}'" ./cmd/bpzh-vk-bot.go

EXPOSE 4000

CMD ["./app"]
