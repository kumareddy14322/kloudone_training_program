ROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/kumareddy14322/kloudone_training_program/grpc3
WORKDIR /go/src/github.com/kumareddy14322/kloudone_training_program/grpc3
RUN go mod download
COPY . /go/src/ggithub.com/kumareddy14322/kloudone_training_program/grpc3
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/grpc3 github.com/kumareddy14322/kloudone_training_program/grpc3

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/kumareddy14322/kloudone_training_program/grpc3/build/grpc3 /usr/bin/grpc3
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/grpc3"]