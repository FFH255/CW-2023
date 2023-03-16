FROM golang:alpine3.17 as build-env

RUN mkdir /app
WORKDIR /app

COPY go.mod . 
COPY go.sum .


RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/main
FROM scratch 
COPY --from=build-env /go/bin/main /bin/main
COPY --from=build-env /app/.env /
ENTRYPOINT ["/bin/main"]