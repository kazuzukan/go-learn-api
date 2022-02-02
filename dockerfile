FROM golang:alpine

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN go get ./
RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main