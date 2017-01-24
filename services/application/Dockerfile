FROM golang:1.7
RUN mkdir -p /go/src/github.com/beautiful-code/sal/services/application
WORKDIR /go/src/github.com/beautiful-code/sal/services/application
ADD . /go/src/github.com/beautiful-code/sal/services/application
# Copying the config file for now, but to use kube secrets.
RUN cp config.json.staging config.json
RUN go build main.go
EXPOSE 8080
CMD ["./main"]
