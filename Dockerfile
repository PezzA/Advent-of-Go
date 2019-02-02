# get the base image
FROM golang as builder

RUN apt-get install vim

# Build directory structure
WORKDIR /go/src/github.com/
RUN mkdir /pezza

WORKDIR /go/src/github.com/pezza
RUN mkdir /advent-of-code

WORKDIR /go/src/github.com/pezza/advent-of-code

# Get dependencies
RUN go get github.com/buger/goterm

COPY . /go/src/github.com/pezza/advent-of-code

# build
RUN go build -o advent-of-code .

CMD ["sh"]