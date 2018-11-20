FROM golang:1.8

RUN apt-get update \
	&& apt-get install -y python-pip \
    && pip install --upgrade pip==10.0.1 supervisor==3.3.4 setuptools==39.1.0

COPY . $GOPATH/src/github.com/vlam1/acceptance_test_example
WORKDIR $GOPATH/src/github.com/vlam1/acceptance_test_example

CMD ["go", "run", "cmd/myapp/main.go"]

