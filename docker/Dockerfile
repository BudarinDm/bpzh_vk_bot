FROM golang:1.18

WORKDIR /src


COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go build -o /vkbot

EXPOSE 80

CMD [ "/vkbot" ]