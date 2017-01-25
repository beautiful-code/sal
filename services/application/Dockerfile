FROM golang:1.7

ENV APP_PATH /go/src/github.com/beautiful-code/sal/services/application

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

# Install glide
RUN curl https://glide.sh/get | sh

# Copy glide.yaml and glide.lock
ADD glide* $APP_PATH/

# Run glide install
RUN glide install

# Copy the application source
ADD . $APP_PATH

# Copying the config file for now, but to use kube secrets.
RUN cp config.json.staging config.json

# Go Build
RUN go build main.go && chmod +x main

EXPOSE 8090

CMD ["./main"]
