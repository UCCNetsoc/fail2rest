FROM golang:1.13 as dev

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /fail2rest

RUN go get github.com/go-task/task/v2/cmd/task \
    github.com/go-delve/delve/cmd/dlv

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN go build -a -ldflags '-linkmode external -extldflags "-static"' -o /go/bin/fail2rest

RUN go mod vendor

CMD [ "go", "run", "main.go" ]

FROM scratch

COPY --from=dev /go/bin/fail2rest /fail2rest

CMD [ "/fail2rest" ]