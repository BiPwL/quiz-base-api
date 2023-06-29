FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o apiserver ../cmd/apiserver/main.go

CMD [ "./apiserver" ]