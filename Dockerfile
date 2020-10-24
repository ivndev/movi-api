FROM golang:1.15

LABEL Author="ivan cruz <ivancruz9818@gmail.com>" 

WORKDIR /go/src/app

RUN go get -d -v github.com/markbates/refresh && \ 
go install -v github.com/markbates/refresh

CMD ["refresh", "run"]

