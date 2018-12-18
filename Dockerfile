FROM golang:alpine

ADD ./ /go

RUN go build -o prober.out ./main.go

CMD [ "/go/prober.out" ]
