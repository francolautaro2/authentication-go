FROM golang:1.19

WORKDIR /app

COPY . . 

RUN make build

EXPOSE 3030

CMD ["./auth-go"]
