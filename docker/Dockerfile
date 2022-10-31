FROM golang:1.18

WORKDIR /src


COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go build -o /docker-gs-ping

EXPOSE 80

CMD [ "/docker-gs-ping" ]