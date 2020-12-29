FROM golang:1.15
ENV GO111MODULE=on
WORKDIR /usr/app
COPY ./main ./main

EXPOSE 8080

CMD ["./main"]
