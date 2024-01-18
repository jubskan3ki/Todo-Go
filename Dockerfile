FROM golang:1.21.6

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app

COPY ./app .

RUN go mod download \
    && go mod verify \
    && go build -o /build/buildedApp main.go

WORKDIR /build
ENTRYPOINT [ "./buildedApp" ]
    
