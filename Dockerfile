FROM golang:1.11

RUN go get github.com/pezza/advent-of-code

CMD ["sh"]